package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/thiagodutra/fake-cart/models"
	"github.com/thiagodutra/fake-cart/repositories"
	"sync"
)

type CartServiceInterface interface {
	GetCartById(ctx context.Context, cartId string) (*models.Cart, error)
	UpsertCart(ctx context.Context, cart *models.Cart) (*models.Cart, error)
	DeleteCart(ctx context.Context, cartId string) error
	SubmitCheckout(ctx context.Context, cart *models.Cart) (*models.Cart, error)
}

type CartService struct {
	cartRepo repositories.CartRepositoryInterface
}

var (
	serviceInstance *CartService
	serviceOnce     sync.Once
)

func NewCartService(cartRepo repositories.CartRepositoryInterface) *CartService {
	serviceOnce.Do(func() {
		serviceInstance = &CartService{
			cartRepo: cartRepo,
		}
	})
	return serviceInstance
}

func (service *CartService) GetCartById(ctx context.Context, cartId string) (*models.Cart, error) {
	return service.cartRepo.GetById(ctx, cartId)
}

func (service *CartService) UpsertCart(ctx context.Context, cart models.Cart) (*models.Cart, error) {
	if cart.ID == "" {
		cart.ID = uuid.New().String()
	}

	return service.cartRepo.Upsert(ctx, cart)
}

func (service *CartService) DeleteCart(ctx context.Context, cartId string) error {
	return service.cartRepo.Delete(ctx, cartId)
}

func (service *CartService) SubmitCheckout(cart *models.Cart) (*models.Cart, error) {
	//TODO create client to submit to checkout
	//TODO do some logic here, call validation webhooks
	//

	return nil, nil
}
