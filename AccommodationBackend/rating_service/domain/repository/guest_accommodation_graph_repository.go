package repository

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"rating_service/domain/model"
)

type GuestAccommodationGraphStore struct {
	driver neo4j.Driver
}

func NewGuestAccommodationGraphStore(driver neo4j.Driver) *GuestAccommodationGraphStore {
	return &GuestAccommodationGraphStore{
		driver: driver,
	}
}

func (store *GuestAccommodationGraphStore) CheckConnection() {
	err := store.driver.VerifyConnectivity()
	if err != nil {
		println("Error while checking connection: ", err.Error())
		return
	}
	println("Neo4J server address: ", store.driver.Target().Host)
}

func (store *GuestAccommodationGraphStore) CreateOrUpdateGuestAccommodationConnection(accommodationRating *model.AccommodationRating) error {
	session := store.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			// Check if guest node exists
			guestResult, err := transaction.Run(
				"MATCH (g:Guest {guestId: $guestId}) RETURN g",
				map[string]interface{}{"guestId": accommodationRating.GuestId})
			if err != nil {
				return nil, err
			}

			// Create or update guest node
			if !guestResult.Next() {
				_, err := transaction.Run(
					"CREATE (g:Guest {guestId: $guestId})",
					map[string]interface{}{"guestId": accommodationRating.GuestId})
				if err != nil {
					return nil, err
				}
			}

			// Check if accommodation node exists
			accommodationResult, err := transaction.Run(
				"MATCH (a:Accommodation {accommodationId: $accommodationId}) RETURN a",
				map[string]interface{}{"accommodationId": accommodationRating.AccommodationId})
			if err != nil {
				return nil, err
			}

			// Create or update accommodation node
			if !accommodationResult.Next() {
				_, err := transaction.Run(
					"CREATE (a:Accommodation {accommodationId: $accommodationId})",
					map[string]interface{}{"accommodationId": accommodationRating.AccommodationId})
				if err != nil {
					return nil, err
				}
			}

			// Check if a relationship between guest and accommodation exists
			relationshipResult, err := transaction.Run(
				"MATCH (g:Guest {guestId: $guestId})-[r:RATED {ratingId: $ratingId}]->(a:Accommodation {accommodationId: $accommodationId}) RETURN r",
				map[string]interface{}{"guestId": accommodationRating.GuestId, "accommodationId": accommodationRating.AccommodationId, "ratingId": accommodationRating.Id.Hex()})
			if err != nil {
				return nil, err
			}

			// Create or update the relationship
			if relationshipResult.Next() {
				_, err := transaction.Run(
					"MATCH (g:Guest {guestId: $guestId})-[r:RATED]->(a:Accommodation {accommodationId: $accommodationId}) SET r.score = $score, r.date = $date, r.ratingId = $ratingId",
					map[string]interface{}{"guestId": accommodationRating.GuestId, "accommodationId": accommodationRating.AccommodationId, "score": accommodationRating.Score, "date": accommodationRating.Date, "ratingId": accommodationRating.Id.Hex()})
				if err != nil {
					return nil, err
				}
			} else {
				// Create the relationship if it doesn't exist
				_, err := transaction.Run(
					"MATCH (g:Guest), (a:Accommodation) WHERE g.guestId = $guestId AND a.accommodationId = $accommodationId CREATE (g)-[r:RATED {score:$score, date:$date, ratingId:$ratingId}]->(a) RETURN r AS rel",
					map[string]interface{}{"guestId": accommodationRating.GuestId, "accommodationId": accommodationRating.AccommodationId, "score": accommodationRating.Score, "date": accommodationRating.Date, "ratingId": accommodationRating.Id.Hex()})
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		})

	if err != nil {
		return err
	}

	return nil
}

func (store *GuestAccommodationGraphStore) RecommendAccommodationsForGuest(guestId string) ([]string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	query := `
	MATCH (guest:Guest {guestId: $guestId})-[r:RATED]->(accommodation:Accommodation)
	WITH guest, accommodation, r.score AS userRating

	MATCH (otherGuest:Guest)-[otherR:RATED]->(accommodation)
	WHERE otherGuest <> guest AND (otherR.score = userRating + 1 OR otherR.score = userRating - 1)
	WITH otherGuest
	
	// Find accommodations rated by otherGuests with a score of 4 or more
	MATCH (otherGuest)-[r:RATED]->(goodAccommodations:Accommodation)
	WHERE r.score >= 4
	WITH COLLECT(goodAccommodations) AS collectedGoodAccommodations
	
	// Find accommodations rated by any guest with a score of 2 or less
	OPTIONAL MATCH (badAccommodations:Accommodation)<-[r:RATED]-(:Guest)
	WHERE r.score <= 2
	WITH collectedGoodAccommodations, COLLECT(badAccommodations) AS collectedBadAccommodations
	
	// Filter out accommodations from collectedGoodAccommodations that match collectedBadAccommodations
	WITH collectedGoodAccommodations, collectedBadAccommodations
	UNWIND collectedGoodAccommodations AS goodAccommodation
	WITH goodAccommodation, collectedBadAccommodations
	WHERE NOT goodAccommodation IN collectedBadAccommodations
	
	WITH DISTINCT(goodAccommodation) AS distinctGoodAccommodations

	RETURN distinctGoodAccommodations.accommodationId
	`
	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		queryResult, err := tx.Run(query, map[string]interface{}{
			"guestId": guestId,
		})
		if err != nil {
			return nil, err
		}

		recommendations := make([]string, 0)
		for queryResult.Next() {
			record := queryResult.Record()
			recommendation, ok := record.Get("distinctGoodAccommodations.accommodationId")
			if ok {
				recommendations = append(recommendations, recommendation.(string))
			}
		}

		return recommendations, nil
	})

	if err != nil {
		return nil, err
	}

	return result.([]string), nil
}

func (store *GuestAccommodationGraphStore) DeleteGuestAndConnection(guestID, accommodationID, ratingId string) error {
	session := store.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.WriteTransaction(
		func(transaction neo4j.Transaction) (interface{}, error) {
			// Delete the relationship between the guest and accommodation
			_, err := transaction.Run(
				"MATCH (g:Guest {guestId: $guestId})-[r:RATED {ratingId:$ratingId}]->(a:Accommodation {accommodationId: $accommodationId}) DELETE r",
				map[string]interface{}{"guestId": guestID, "accommodationId": accommodationID, "ratingId": ratingId})
			if err != nil {
				return nil, err
			}

			// Check if the guest node is no longer connected to any accommodation
			checkResult, err := transaction.Run(
				"MATCH (g:Guest {guestId: $guestId}) WHERE NOT (g)-[:RATED]->() DELETE g",
				map[string]interface{}{"guestId": guestID})
			if err != nil {
				return nil, err
			}

			// If the guest node is no longer connected to any accommodation, delete it
			if checkResult.Next() {
				_, err := transaction.Run(
					"MATCH (g:Guest {guestId: $guestId}) DELETE g",
					map[string]interface{}{"guestId": guestID})
				if err != nil {
					return nil, err
				}
			}

			return nil, nil
		})

	if err != nil {
		return err
	}

	return nil
}
