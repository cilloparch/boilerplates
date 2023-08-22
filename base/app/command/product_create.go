package command

import (
	"context"

	"github.com/cilloparch/boilerplates/base/domains/product"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type ProductCreateCmd struct {
	Name string `json:"name"`
}

type ProductCreateRes struct {
	ID string `json:"id" validate:"required"`
}

type ProductCreateHandler cqrs.Handler[ProductCreateCmd, *ProductCreateRes]

type productCreateHandler struct {
	repo product.Repository
}

func NewProductCreateHandler(repo product.Repository) ProductCreateHandler {
	return &productCreateHandler{repo: repo}
}

func (h *productCreateHandler) Handle(ctx context.Context, cmd ProductCreateCmd) (*ProductCreateRes, *i18np.Error) {
	id, err := h.repo.Create(ctx, cmd.Name)
	if err != nil {
		return nil, err
	}
	return &ProductCreateRes{ID: id}, nil
}
