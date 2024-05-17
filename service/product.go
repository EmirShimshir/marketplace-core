package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
)

type ProductService struct {
	repo    port.IProductRepository
	storage port.IObjectStorage
}

func NewProductService(repo port.IProductRepository, storage port.IObjectStorage) *ProductService {
	return &ProductService{
		repo:    repo,
		storage: storage,
	}
}

func (p *ProductService) Get(ctx context.Context, limit, offset int64) ([]domain.Product, error) {
	return p.repo.Get(ctx, limit, offset)
}

func (p *ProductService) GetByID(ctx context.Context, productID domain.ID) (domain.Product, error) {
	return p.repo.GetByID(ctx, productID)
}

func (p *ProductService) Create(ctx context.Context, param port.CreateProductParam) (domain.Product, error) {
	productID := domain.NewID()
	url, err := p.storage.SaveFile(ctx, domain.File{
		Name:   productID.String() + ".png",
		Path:   "product",
		Reader: param.PhotoReader,
	})
	if err != nil {
		return domain.Product{}, err
	}

	return p.repo.Create(ctx, domain.Product{
		ID:          productID,
		Name:        param.Name,
		Description: param.Description,
		Price:       param.Price,
		Category:    param.Category,
		PhotoUrl:    url.String(),
	})
}

func (p *ProductService) Update(ctx context.Context, productID domain.ID, param port.UpdateProductParam) (domain.Product, error) {
	product, err := p.repo.GetByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}

	if param.Name.Valid {
		product.Name = param.Name.String
	}
	if param.Description.Valid {
		product.Description = param.Description.String
	}
	if param.Price.Valid {
		product.Price = param.Price.Int64
	}
	if param.Category != nil {
		product.Category = *param.Category
	}
	if param.PhotoReader != nil {
		url, err := p.storage.SaveFile(ctx, domain.File{
			Name:   productID.String() + ".png",
			Path:   "product",
			Reader: *param.PhotoReader,
		})
		if err != nil {
			return domain.Product{}, err
		}
		product.PhotoUrl = url.String()
	}

	return p.repo.Update(ctx, product)
}

func (p *ProductService) Delete(ctx context.Context, productID domain.ID) error {
	return p.repo.Delete(ctx, productID)
}
