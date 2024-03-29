package query

import (
	"context"

	"github.com/cilloparch/boilerplates/base/domains/product"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type ProductGetQuery struct {
	ID string `param:"id" json:"id" validate:"required"`
}

type ProductGetRes struct {
	Name string `json:"name"`
}

type ProductGetHandler cqrs.Handler[ProductGetQuery, *ProductGetRes]

type productGetHandler struct {
	repo product.Repository
}

func NewProductGetHandler(repo product.Repository) ProductGetHandler {
	return &productGetHandler{repo: repo}
}

func (h *productGetHandler) Handle(ctx context.Context, query ProductGetQuery) (*ProductGetRes, *i18np.Error) {
	prod, err := h.repo.Get(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return &ProductGetRes{Name: prod.Name}, nil
}
