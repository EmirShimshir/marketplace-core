package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
	"github.com/guregu/null"
)

type CartService struct {
	cartRepo    port.ICartRepository
	shopRepo    port.IShopRepository
	productRepo port.IProductRepository
}

func NewCartService(cartRepo port.ICartRepository, shopRepo port.IShopRepository, productRepo port.IProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		shopRepo:    shopRepo,
		productRepo: productRepo,
	}
}

func (c *CartService) GetCartByID(ctx context.Context, cartID domain.ID) (domain.Cart, error) {
	cart, err := c.cartRepo.GetCartByID(ctx, cartID)
	if err != nil {
		return domain.Cart{}, err
	}

	var totalPrice int64 = 0
	for _, cartItem := range cart.Items {
		shopItem, err := c.shopRepo.GetShopItemByProductID(ctx, cartItem.ProductID)
		if err != nil {
			return domain.Cart{}, err
		}
		if cartItem.Quantity > shopItem.Quantity {
			cartItem, err = c.UpdateCartItem(ctx, cartItem.ID, port.UpdateCartItemParam{
				Quantity: null.IntFrom(shopItem.Quantity),
			})
			if err != nil {
				return domain.Cart{}, err
			}
		}
		product, err := c.productRepo.GetByID(ctx, cartItem.ProductID)
		if err != nil {
			return domain.Cart{}, err
		}
		totalPrice += product.Price * cartItem.Quantity
	}

	cart.Price = totalPrice
	return c.cartRepo.UpdateCart(ctx, cart)
}

func (c *CartService) ClearCart(ctx context.Context, cartID domain.ID) error {
	cart, err := c.cartRepo.GetCartByID(ctx, cartID)
	if err != nil {
		return err
	}

	for _, item := range cart.Items {
		err = c.DeleteCartItem(ctx, item.ID)
		if err != nil {
			return err
		}
	}

	cart.Price = 0
	_, err = c.cartRepo.UpdateCart(ctx, cart)

	return err
}

func (c *CartService) GetCartItemByID(ctx context.Context, cartItemID domain.ID) (domain.CartItem, error) {
	return c.cartRepo.GetCartItemByID(ctx, cartItemID)
}

func (c *CartService) CreateCartItem(ctx context.Context, param port.CreateCartItemParam) (domain.CartItem, error) {
	shopItem, err := c.shopRepo.GetShopItemByProductID(ctx, param.ProductID)
	if err != nil {
		return domain.CartItem{}, err
	}

	if param.Quantity < 1 || param.Quantity > shopItem.Quantity {
		return domain.CartItem{}, domain.ErrQuantityItems
	}

	return c.cartRepo.CreateCartItem(ctx, domain.CartItem{
		ID:        domain.NewID(),
		CartID:    param.CartID,
		ProductID: param.ProductID,
		Quantity:  param.Quantity,
	})
}

func (c *CartService) UpdateCartItem(ctx context.Context, cartItemID domain.ID, param port.UpdateCartItemParam) (domain.CartItem, error) {
	cartItem, err := c.cartRepo.GetCartItemByID(ctx, cartItemID)
	if err != nil {
		return domain.CartItem{}, err
	}

	if param.Quantity.Valid {
		cartItem.Quantity = param.Quantity.Int64
	}

	if cartItem.Quantity < 1 {
		return domain.CartItem{}, c.DeleteCartItem(ctx, cartItemID)
	}

	shopItem, err := c.shopRepo.GetShopItemByProductID(ctx, cartItem.ProductID)
	if err != nil {
		return domain.CartItem{}, err
	}

	if cartItem.Quantity > shopItem.Quantity {
		return domain.CartItem{}, domain.ErrQuantityItems
	}

	return c.cartRepo.UpdateCartItem(ctx, cartItem)
}

func (c *CartService) DeleteCartItem(ctx context.Context, cartItemID domain.ID) error {
	return c.cartRepo.DeleteCartItem(ctx, cartItemID)
}