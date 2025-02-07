package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/thiagodutra/fake-cart/models"
	"github.com/thiagodutra/fake-cart/services"
)

// AddCartHandler is the main handler for cart-related requests.
func AddCartHandler(w http.ResponseWriter, r *http.Request, cartService *services.CartService) {
	switch r.Method {
	case "POST":
		postCartHandler(w, r, cartService)
	default:
		http.Error(w, "Only POST method is allowed for now", http.StatusMethodNotAllowed)
	}
}

func postCartHandler(w http.ResponseWriter, r *http.Request, cartService *services.CartService) {
	var cart models.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	updatedCart, err := cartService.UpsertCart(ctx, cart)
	if err != nil {
		http.Error(w, "Failed to process cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(updatedCart); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
