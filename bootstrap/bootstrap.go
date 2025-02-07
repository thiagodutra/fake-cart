package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/thiagodutra/fake-cart/config"
	"github.com/thiagodutra/fake-cart/handlers"
	"github.com/thiagodutra/fake-cart/repositories"
	"github.com/thiagodutra/fake-cart/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppComponents struct {
	DBClient    *mongo.Client
	CartService *services.CartService
}

func InitializeComponents() (*AppComponents, error) {
	// Initialize database
	log.Printf("Bootstrapping server ...")
	log.Printf("Connecting to MongoDB...")

	maxConnections := uint64(20)
	client, err := config.ConnectDatabase("mongodb://localhost:27017/carts", maxConnections)

	if err != nil {
		return nil, fmt.Errorf("... error connecting to database: %w", err)
	}
	log.Printf("Connected to the database ... ")

	// Initialize repository
	log.Printf("initializing repositories ... ")
	cartRepo := repositories.NewCartRepository(client)

	log.Printf("initializing services ... ")
	// Initialize service
	cartService := services.NewCartService(cartRepo)

	log.Printf("Server initialized ... ")
	return &AppComponents{
		DBClient:    client,
		CartService: cartService,
	}, nil
}

func SetupRoutes(cartService *services.CartService) {
	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddCartHandler(w, r, cartService)
	})
}

func Shutdown(appComponents *AppComponents) {
	if err := appComponents.DBClient.Disconnect(context.Background()); err != nil {
		fmt.Printf("Failed to disconnect from database: %v\n", err)
	}
}
