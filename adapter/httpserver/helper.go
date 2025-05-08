package httpserver

import (
	"context"

	"github.com/loan-service/infra/way"
)

type RouterHelper struct{}

func (r RouterHelper) GetParam(ctx context.Context, param string) string {
	return way.Param(ctx, param)
}

type HelperInterface interface {
	GetParam(ctx context.Context, param string) string
}
