package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/thiagodutra/fake-cart/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var cartCollection = "carts"
var database = "carts"

type CartRepositoryInterface interface {
	GetById(ctx context.Context, id string) (*models.Cart, error)
	Upsert(ctx context.Context, cart models.Cart) (*models.Cart, error)
	Delete(ctx context.Context, id string) error
}

type CartRepository struct {
	dbClient     *mongo.Client
	dbCollection *mongo.Collection
}

// This makes the instance singleton
var (
	instance *CartRepository
	once     sync.Once
)

func NewCartRepository(dbClient *mongo.Client) *CartRepository {
	once.Do(func() {
		instance = &CartRepository{
			dbClient:     dbClient,
			dbCollection: dbClient.Database(database).Collection(cartCollection),
		}
	})
	return instance
}

func (repo *CartRepository) GetById(ctx context.Context, cartId string) (*models.Cart, error) {
	var cart models.Cart
	err := repo.dbCollection.FindOne(ctx, bson.M{"_id": cartId}).Decode(&cart)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, fmt.Errorf("cart %s not found", cartId)
		}
		return nil, err
	}
	return &cart, nil
}

func (repo *CartRepository) Upsert(ctx context.Context, cart models.Cart) (*models.Cart, error) {
	filter := bson.M{"id": cart.ID}
	update := bson.M{"$set": cart}
	opts := options.Update().SetUpsert(true)
	_, err := repo.dbCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}
	return repo.GetById(ctx, cart.ID)
}

func (repo *CartRepository) Delete(ctx context.Context, cartId string) error {
	_, err := repo.dbCollection.DeleteOne(ctx, bson.M{"id": cartId})
	if err != nil {
		return err
	}
	return nil
}
