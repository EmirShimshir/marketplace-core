package port

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/guregu/null"
)

type CreateShopParam struct {
	Name        string
	Description string
	Requisites  string
	Email       string
}

type UpdateShopParam struct {
	Name        null.String
	Description null.String
	Requisites  null.String
	Email       null.String
}

type CreateShopItemParam struct {
	ShopID       domain.ID
	ProductParam CreateProductParam
	Quantity     int64
}

type UpdateShopItemParam struct {
	Quantity null.Int
}

type IShopService interface {
	GetShops(ctx context.Context, limit, offset int64) ([]domain.Shop, error)
	GetShopByID(ctx context.Context, shopID domain.ID) (domain.Shop, error)
	GetShopBySellerID(ctx context.Context, sellerID domain.ID) ([]domain.Shop, error)
	CreateShop(ctx context.Context, sellerID domain.ID, param CreateShopParam) (domain.Shop, error)
	UpdateShop(ctx context.Context, shopID domain.ID, param UpdateShopParam) (domain.Shop, error)
	DeleteShop(ctx context.Context, shopID domain.ID) error
	GetShopItems(ctx context.Context, limit, offset int64) ([]domain.ShopItem, error)
	GetShopItemByID(ctx context.Context, shopItemID domain.ID) (domain.ShopItem, error)
	GetShopItemByProductID(ctx context.Context, productID domain.ID) (domain.ShopItem, error)
	CreateShopItem(ctx context.Context, param CreateShopItemParam) (domain.ShopItem, error)
	UpdateShopItem(ctx context.Context, shopItemID domain.ID, param UpdateShopItemParam) (domain.ShopItem, error)
	DeleteShopItem(ctx context.Context, shopItemID domain.ID) error
}
