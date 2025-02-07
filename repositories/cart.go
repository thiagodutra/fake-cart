package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/thiagodutra/fake-cart/logger"
	"github.com/thiagodutra/fake-cart/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var log = logger.NewLogger()
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
	log.Log(logger.INFO, "Repository initialized", instance, nil)
	return instance
}

func (repo *CartRepository) GetById(ctx context.Context, cartId string) (*models.Cart, error) {
	var cart models.Cart
	err := repo.dbCollection.FindOne(ctx, bson.M{"_id": cartId}).Decode(&cart)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil, fmt.Errorf("cart %s not found", cartId)
		}
		log.Log(logger.ERROR, "An error occurred trying to find cart by id %s", instance, err)
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
		log.Log(logger.ERROR, "An error occurred trying to insert/update cart by id %s", instance, err)
		return nil, err
	}
	return repo.GetById(ctx, cart.ID)
}

func (repo *CartRepository) Delete(ctx context.Context, cartId string) error {
	_, err := repo.dbCollection.DeleteOne(ctx, bson.M{"id": cartId})
	if err != nil {
		log.Log(logger.ERROR, "An error occurred trying to delete cart by id %s", instance, err)
		return err
	}
	return nil
}
