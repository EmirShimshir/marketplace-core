package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
	"time"
)

type OrderService struct {
	orderRepo port.IOrderRepository
	userRepo  port.IUserRepository
	cartRepo  port.ICartRepository
	shopRepo  port.IShopRepository
}

func NewOrderService(orderRepo port.IOrderRepository,
	userRepo port.IUserRepository,
	cartRepo port.ICartRepository,
	shopRepo port.IShopRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
		userRepo:  userRepo,
		cartRepo:  cartRepo,
		shopRepo:  shopRepo,
	}
}

func (o *OrderService) GetOrderCustomerByCustomerID(ctx context.Context, customerID domain.ID) ([]domain.OrderCustomer, error) {
	return o.orderRepo.GetOrderCustomerByCustomerID(ctx, customerID)
}

func (o *OrderService) GetOrderCustomerByID(ctx context.Context, ID domain.ID) (domain.OrderCustomer, error) {
	return o.orderRepo.GetOrderCustomerByID(ctx, ID)
}

func (o *OrderService) getCartItemsByShopID(ctx context.Context, cart domain.Cart) (map[domain.ID][]domain.CartItem, error) {
	cartItemsByShopID := make(map[domain.ID][]domain.CartItem)
	for _, cartItem := range cart.Items {
		shopItem, err := o.shopRepo.GetShopItemByProductID(ctx, cartItem.ProductID)
		if err != nil {
			return nil, err
		}
		cartItemsByShopID[shopItem.ShopID] = append(cartItemsByShopID[shopItem.ShopID], cartItem)
		if err != nil {
			return nil, err
		}
	}
	return cartItemsByShopID, nil
}

func (o *OrderService) buildOrderCustomer(param port.CreateOrderCustomerParam, totalPrice int64,
	cartItemsByShopID map[domain.ID][]domain.CartItem) domain.OrderCustomer {
	orderCustomer := domain.OrderCustomer{
		ID:         domain.NewID(),
		CustomerID: param.CustomerID,
		Address:    param.Address,
		CreatedAt:  time.Now(),
		TotalPrice: totalPrice,
		Payed:      false,
		OrderShops: make([]domain.OrderShop, 0, len(cartItemsByShopID)),
	}
	for shopID, cartItems := range cartItemsByShopID {
		orderShop := domain.OrderShop{
			ID:              domain.NewID(),
			ShopID:          shopID,
			OrderCustomerID: orderCustomer.ID,
			Status:          domain.OrderShopStatusStart,
			OrderShopItems:  make([]domain.OrderShopItem, len(cartItems)),
		}
		for i, cartItem := range cartItems {
			orderShop.OrderShopItems[i] = domain.OrderShopItem{
				ID:          domain.NewID(),
				OrderShopID: orderShop.ID,
				ProductID:   cartItem.ProductID,
				Quantity:    cartItem.Quantity,
			}
		}
		orderCustomer.OrderShops = append(orderCustomer.OrderShops, orderShop)
	}
	return orderCustomer
}

func (o *OrderService) CreateOrderCustomer(ctx context.Context, param port.CreateOrderCustomerParam) (domain.OrderCustomer, error) {
	if param.Address == "" {
		return domain.OrderCustomer{}, domain.ErrAddress
	}
	customer, err := o.userRepo.GetByID(ctx, param.CustomerID)
	if err != nil {
		return domain.OrderCustomer{}, err
	}

	cart, err := o.cartRepo.GetCartByID(ctx, customer.CartID)
	if err != nil {
		return domain.OrderCustomer{}, err
	}

	if len(cart.Items) < 1 {
		return domain.OrderCustomer{}, domain.ErrEmptyCart
	}

	cartItemsByShopID, err := o.getCartItemsByShopID(ctx, cart)
	if err != nil {
		return domain.OrderCustomer{}, err
	}
	orderCustomer := o.buildOrderCustomer(param, cart.Price, cartItemsByShopID)

	return o.orderRepo.CreateOrderCustomer(ctx, orderCustomer)
}

func (o *OrderService) GetOrderShopByID(ctx context.Context, orderShopID domain.ID) (domain.OrderShop, error) {
	return o.orderRepo.GetOrderShopByID(ctx, orderShopID)
}

func (o *OrderService) GetOrderShopByShopID(ctx context.Context, shopID domain.ID) ([]domain.OrderShop, error) {
	return o.orderRepo.GetOrderShopByShopID(ctx, shopID)
}

func (o *OrderService) UpdateOrderShop(ctx context.Context, orderShopID domain.ID, param port.UpdateOrderShopParam) (domain.OrderShop, error) {
	orderShop, err := o.orderRepo.GetOrderShopByID(ctx, orderShopID)
	if err != nil {
		return domain.OrderShop{}, err
	}

	if param.Status != nil {
		orderShop.Status = *param.Status
	}

	return o.orderRepo.UpdateOrderShop(ctx, orderShop)
}
