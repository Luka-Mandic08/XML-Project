package repository

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"rating_service/domain/model"
)

type GuestAccommodationGraphStore struct {
	driver neo4j.DriverWithContext
}

func NewGuestAccommodationGraphStore(driver neo4j.DriverWithContext) GuestAccommodationGraphStore {
	return GuestAccommodationGraphStore{
		driver: driver,
	}
}

// CheckConnection Check if connection is established
func (store *GuestAccommodationGraphStore) CheckConnection() {
	ctx := context.Background()
	err := store.driver.VerifyConnectivity(ctx)
	if err != nil {
		return
	}
	// Print Neo4J server address
}

// CloseDriverConnection Disconnect from database
func (store *GuestAccommodationGraphStore) CloseDriverConnection(ctx context.Context) {
	store.driver.Close(ctx)
}

func (store *GuestAccommodationGraphStore) CreateGuestNode(guestId string) error {
	ctx := context.Background()
	session := store.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (g:Guest) SET g.guestId = $guestId RETURN g.guestId + ', from node ' + id(g)",
				map[string]any{"guestId": guestId})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		return err
	}
	return nil
}

func (store *GuestAccommodationGraphStore) CreateAccommodationNode(accommodationId string) error {
	ctx := context.Background()
	session := store.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (a:Accommodation) SET a.accommodationId = accommodationId RETURN a.accommodationId + ', from node ' + id(a)",
				map[string]any{"accommodationId": accommodationId})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
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
	ctx := context.Background()
	session := store.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	_, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (g:Guest), (a:Accommodation) WHERE g.guestId = $guestId AND a.accommodationId = $accommodationId CREATE (g)-[r:RATED {score:$score, date:$date}]->(a) RETURN type(r)",
				map[string]any{"guestId": accommodationRating.GuestId, "accommodationId": accommodationRating.AccommodationId, "score": accommodationRating.Score, "date": accommodationRating.Date})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		return err
	}
	return nil
}
