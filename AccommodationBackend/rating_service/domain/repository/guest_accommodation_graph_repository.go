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

func (store *GuestAccommodationGraphStore) RecommendAccommodationsForUser(guestId string) ([]string, error) {
	session := store.driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	query := `
	MATCH (user:Guest)-[r:RATED]->(accommodation:Accommodation)
	WITH user, accommodation, r.score AS userRating

	MATCH (otherUser:Guest)-[otherR:RATED]->(accommodation)
	WHERE otherUser <> user
	WITH user, otherUser, COLLECT(DISTINCT accommodation) AS sharedAccommodations,
	     COLLECT(DISTINCT otherR.score) AS otherRatings, userRating

	WITH user, otherUser, sharedAccommodations, otherRatings, userRating,
	     algo.similarity.pearson(userRating, otherRatings) AS similarity

	MATCH (otherUser)-[r:RATED]->(recommended:Accommodation)
	WHERE NOT recommended IN sharedAccommodations
	WITH user, recommended, COLLECT(DISTINCT otherUser) AS similarUsers,
	     COLLECT(DISTINCT r.score) AS recommendationScores

	WITH user, recommended,
	     REDUCE(s = 0, score IN recommendationScores | s + score) AS totalScore,
	     SIZE(similarUsers) AS numSimilarUsers

	MATCH (recommended)<-[r:RATED]-(similarUser:Guest)
	WHERE r.score < 3 AND r.date >= datetime() - duration('P3M')
	WITH user, recommended, totalScore, numSimilarUsers,
	     COUNT(DISTINCT similarUser) AS lowRatingCount

	WHERE lowRatingCount <= 5

	RETURN recommended, totalScore / numSimilarUsers AS recommendationScore
	ORDER BY recommendationScore DESC
	LIMIT 10
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
