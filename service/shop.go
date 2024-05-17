package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
)

type ShopService struct {
	repo    port.IShopRepository
	storage port.IObjectStorage
}

func NewShopService(repo port.IShopRepository, storage port.IObjectStorage) *ShopService {
	return &ShopService{
		repo:    repo,
		storage: storage,
	}
}

func (s *ShopService) GetShops(ctx context.Context, limit, offset int64) ([]domain.Shop, error) {
	return s.repo.GetShops(ctx, limit, offset)
}

func (s *ShopService) GetShopByID(ctx context.Context, shopID domain.ID) (domain.Shop, error) {
	return s.repo.GetShopByID(ctx, shopID)
}

func (s *ShopService) GetShopBySellerID(ctx context.Context, sellerID domain.ID) ([]domain.Shop, error) {
	return s.repo.GetShopBySellerID(ctx, sellerID)
}

func (s *ShopService) CreateShop(ctx context.Context, sellerID domain.ID, param port.CreateShopParam) (domain.Shop, error) {
	if param.Name == "" {
		return domain.Shop{}, domain.ErrName
	}
	if param.Description == "" {
		return domain.Shop{}, domain.ErrDescription
	}
	if param.Requisites == "" {
		return domain.Shop{}, domain.ErrRequisites
	}

	return s.repo.CreateShop(ctx, domain.Shop{
		ID:          domain.NewID(),
		SellerID:    sellerID,
		Name:        param.Name,
		Description: param.Description,
		Requisites:  param.Requisites,
		Email:       param.Email,
		Items:       make([]domain.ShopItem, 0),
	})
}

func (s *ShopService) UpdateShop(ctx context.Context, shopID domain.ID, param port.UpdateShopParam) (domain.Shop, error) {
	shop, err := s.GetShopByID(ctx, shopID)
	if err != nil {
		return domain.Shop{}, err
	}

	if param.Name.Valid {
		shop.Name = param.Name.String
	}
	if param.Description.Valid {
		shop.Description = param.Description.String
	}
	if param.Requisites.Valid {
		shop.Requisites = param.Requisites.String
	}
	if param.Email.Valid {
		shop.Email = param.Email.String
	}

	if shop.Name == "" {
		return domain.Shop{}, domain.ErrName
	}
	if shop.Description == "" {
		return domain.Shop{}, domain.ErrDescription
	}
	if shop.Requisites == "" {
		return domain.Shop{}, domain.ErrRequisites
	}

	return s.repo.UpdateShop(ctx, shop)
}

func (s *ShopService) DeleteShop(ctx context.Context, shopID domain.ID) error {
	return s.repo.DeleteShop(ctx, shopID)
}

func (s *ShopService) GetShopItemByID(ctx context.Context, shopItemID domain.ID) (domain.ShopItem, error) {
	return s.repo.GetShopItemByID(ctx, shopItemID)
}

func (s *ShopService) GetShopItems(ctx context.Context, limit, offset int64) ([]domain.ShopItem, error) {
	return s.repo.GetShopItems(ctx, limit, offset)
}

func (s *ShopService) GetShopItemByProductID(ctx context.Context, productID domain.ID) (domain.ShopItem, error) {
	return s.repo.GetShopItemByProductID(ctx, productID)
}

func (s *ShopService) CreateShopItem(ctx context.Context, param port.CreateShopItemParam) (domain.ShopItem, error) {
	productID := domain.NewID()
	url, err := s.storage.SaveFile(ctx, domain.File{
		Name:   productID.String() + ".png",
		Path:   "product",
		Reader: param.ProductParam.PhotoReader,
	})
	if err != nil {
		return domain.ShopItem{}, err
	}

	if param.ProductParam.Name == "" {
		return domain.ShopItem{}, domain.ErrName
	}
	if param.ProductParam.Description == "" {
		return domain.ShopItem{}, domain.ErrDescription
	}
	if param.Quantity < 1 {
		return domain.ShopItem{}, domain.ErrQuantityItems
	}
	if param.ProductParam.Price < 1 {
		return domain.ShopItem{}, domain.ErrPrice
	}

	product := domain.Product{
		ID:          productID,
		Name:        param.ProductParam.Name,
		Description: param.ProductParam.Description,
		Price:       param.ProductParam.Price,
		Category:    param.ProductParam.Category,
		PhotoUrl:    url.String(),
	}

	shopItem := domain.ShopItem{
		ID:        domain.NewID(),
		ShopID:    param.ShopID,
		ProductID: product.ID,
		Quantity:  param.Quantity,
	}

	return s.repo.CreateShopItem(ctx, shopItem, product)
}

func (s *ShopService) UpdateShopItem(ctx context.Context, shopItemID domain.ID, param port.UpdateShopItemParam) (domain.ShopItem, error) {
	shopItem, err := s.repo.GetShopItemByID(ctx, shopItemID)
	if err != nil {
		return domain.ShopItem{}, err
	}

	if param.Quantity.Valid {
		if param.Quantity.Int64 < 0 {
			return domain.ShopItem{}, domain.ErrQuantityItems
		}
		shopItem.Quantity = param.Quantity.Int64
	}

	return s.repo.UpdateShopItem(ctx, shopItem)
}

func (s *ShopService) DeleteShopItem(ctx context.Context, shopItemID domain.ID) error {
	return s.repo.DeleteShopItem(ctx, shopItemID)

}
