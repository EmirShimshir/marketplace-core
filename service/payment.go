package service

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
	"github.com/EmirShimshir/marketplace-core/port"
	"net/url"
)

type PaymentService struct {
	gateway   port.IPaymentGateway
	orderRepo port.IOrderRepository
}

func NewPaymentService(gateway port.IPaymentGateway, orderRepo port.IOrderRepository) *PaymentService {
	return &PaymentService{
		gateway:   gateway,
		orderRepo: orderRepo,
	}
}

func (p *PaymentService) GetOrderPaymentUrl(ctx context.Context, orderID domain.ID) (url.URL, error) {
	orderCustomer, err := p.orderRepo.GetOrderCustomerByID(ctx, orderID)
	if err != nil {
		return url.URL{}, err
	}

	if orderCustomer.Payed {
		return url.URL{}, domain.ErrOrderAlreadyPayed
	}

	return p.gateway.GetPaymentUrl(ctx, domain.PaymentPayload{
		OrderID: orderCustomer.ID,
		PaySum:  orderCustomer.TotalPrice,
	})
}

func (p *PaymentService) ProcessOrderPayment(ctx context.Context, key string) error {
	payload, err := p.gateway.ProcessPayment(ctx, key)
	if err != nil {
		return err
	}

	orderCustomer, err := p.orderRepo.GetOrderCustomerByID(ctx, payload.OrderID)
	if err != nil {
		return err
	}

	if orderCustomer.TotalPrice < payload.PaySum {
		return domain.ErrInvalidPaymentSum
	}

	return p.orderRepo.UpdatePaymentStatus(ctx, orderCustomer.ID)
}
