package port

import (
	"context"
	"github.com/EmirShimshir/marketplace-core/domain"
)

type IObjectStorage interface {
	SaveFile(ctx context.Context, file domain.File) (domain.Url, error)
}
