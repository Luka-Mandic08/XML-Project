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

func (store *GuestAccommodationGraphStore) CreateGuestNode(guestNode *model.GuestNode) error {
	ctx := context.Background()
	session := store.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	savedPerson, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (p:Person) SET p.born = $born, p.name = $name RETURN p.name + ', from node ' + id(p)",
				map[string]any{"born": person.Born, "name": person.Name})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		pr.logger.Println("Error inserting Person:", err)
		return err
	}
	pr.logger.Println(savedPerson.(string))
	return nil
}

func (pr *PersonRepo) CreateConnectionBetweenPersons() error {
	ctx := context.Background()
	session := pr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// ExecuteWrite for write transactions (Create/Update/Delete)
	savedPerson, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (a:Person), (b:Person) WHERE a.name = $nameOne AND b.name = $nameTwo CREATE (a)-[r:IS_FRIEND]->(b) RETURN type(r)",
				map[string]any{"nameOne": "Luka", "nameTwo": "Pera"})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		pr.logger.Println("Error inserting Person:", err)
		return err
	}
	pr.logger.Println(savedPerson.(string))
	return nil
}
