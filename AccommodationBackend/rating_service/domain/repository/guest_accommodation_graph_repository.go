package repository

import (
	"fmt"
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

/*
	func (store *GuestAccommodationGraphStore) CreateGuestNode(newGuestId string) error {
		session := store.driver.NewSession(neo4j.SessionConfig{})
		defer session.Close()

		// ExecuteWrite for write transactions (Create/Update/Delete)
		_, err := session.WriteTransaction(
			func(transaction neo4j.Transaction) (interface{}, error) {
				result, err := transaction.Run(
					"CREATE (g:Guest {guestId: $newGuestId})",
					map[string]interface{}{"newGuestId": newGuestId})
				if err != nil {
					return nil, err
				}

				if result.Next() {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
		if err != nil {
			return err
		}
		return nil
	}

	func (store *GuestAccommodationGraphStore) CreateAccommodationNode(newAccommodationId string) error {
		session := store.driver.NewSession(neo4j.SessionConfig{})
		defer session.Close()

		// ExecuteWrite for write transactions (Create/Update/Delete)
		_, err := session.WriteTransaction(
			func(transaction neo4j.Transaction) (interface{}, error) {
				result, err := transaction.Run(
					"CREATE (a:Accommodation {accommodationId: $newAccommodationId})",
					map[string]interface{}{"$newAccommodationId": newAccommodationId})
				if err != nil {
					return nil, err
				}

				if result.Next() {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
		if err != nil {
			return err
		}
		return nil
	}

	func (store *GuestAccommodationGraphStore) CreateConnectionBetweenGuestAndAccommodation(accommodationRating *model.AccommodationRating) error {
		session := store.driver.NewSession(neo4j.SessionConfig{})
		defer session.Close()

		// ExecuteWrite for write transactions (Create/Update/Delete)
		_, err := session.WriteTransaction(
			func(transaction neo4j.Transaction) (any, error) {
				result, err := transaction.Run(
					"MATCH (g:Guest), (a:Accommodation) WHERE g.guestId = $guestId AND a.accommodationId = $accommodationId CREATE (g)-[r:RATED {score:$score, date:$date}]->(a) RETURN r AS rel",
					map[string]any{"guestId": accommodationRating.GuestId, "accommodationId": accommodationRating.AccommodationId, "score": accommodationRating.Score, "date": accommodationRating.Date})
				if err != nil {
					return nil, err
				}

				if result.Next() {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
		if err != nil {
			return err
		}
		return nil
	}
*/
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

			// Create connection between guest and accommodation
			_, err = transaction.Run(
				"MATCH (g:Guest), (a:Accommodation) WHERE g.guestId = $guestId AND a.accommodationId = $accommodationId CREATE (g)-[r:RATED {score:$score, date:$date}]->(a) RETURN r AS rel",
				map[string]interface{}{"guestId": accommodationRating.GuestId, "accommodationId": accommodationRating.AccommodationId, "score": accommodationRating.Score, "date": accommodationRating.Date})
			if err != nil {
				return nil, err
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
	MATCH (guest:Guest {guestId: "64ef06e805cdcb24e0f2180f"})-[r:RATED]->(accommodation:Accommodation)
	WITH guest, accommodation, r.score AS userRating
	
	MATCH (otherGuest:Guest)-[otherR:RATED]->(accommodation)
	WHERE otherGuest <> guest AND (otherR.score = userRating + 1 OR otherR.score = userRating - 1)
	WITH otherGuest
	
	// Find accommodations rated by otherGuests with a score of 4 or more
	MATCH (otherGuest)-[r:RATED]->(goodAccommodations:Accommodation)
	WHERE r.score >= 4
	WITH COLLECT(goodAccommodations) AS collectedGoodAccommodations
	
	// Find accommodations rated by any guest with a score of 2 or less
	MATCH (badAccommodations:Accommodation)<-[r:RATED]-(:Guest)
	WHERE r.score <= 2
	WITH collectedGoodAccommodations, COLLECT(badAccommodations) AS collectedBadAccommodations
	
	// Filter out accommodations from collectedGoodAccommodations that match collectedBadAccommodations
	WITH collectedGoodAccommodations, collectedBadAccommodations
	UNWIND collectedGoodAccommodations AS goodAccommodation
	WITH goodAccommodation, collectedBadAccommodations
	WHERE NONE (badAccommodation IN collectedBadAccommodations WHERE badAccommodation = goodAccommodation)
	LIMIT 10
	RETURN goodAccommodation
	

	RETURN recommended, totalScore / numSimilarUsers AS recommendationScore
	ORDER BY recommendationScore DESC
	`

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		queryResult, err := tx.Run(query, map[string]interface{}{
			"guestId": guestId,
		})
		if err != nil {
			return nil, err
		}

		var recommendations []string
		for queryResult.Next() {
			record := queryResult.Record()
			recommendation := fmt.Sprintf("Recommended Accommodation: %s, Score: %f",
				record.Values[0].(string), record.Values[1].(float64))
			recommendations = append(recommendations, recommendation)
		}

		return recommendations, nil
	})

	if err != nil {
		return nil, err
	}

	return result.([]string), nil
}