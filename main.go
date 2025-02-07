package main

import (
	"fmt"
	"net/http"

	"github.com/thiagodutra/fake-cart/bootstrap"
)

func main() {
	appComponents, err := bootstrap.InitializeComponents()
	if err != nil {
		fmt.Printf("Error initializing components: %v\n", err)
		return
	}
	defer bootstrap.Shutdown(appComponents)

	bootstrap.SetupRoutes(appComponents.CartService)

	fmt.Println("Server is running on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
