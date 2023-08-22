package http

import (
	"github.com/cilloparch/boilerplates/base/app/command"
	"github.com/cilloparch/boilerplates/base/app/query"
	"github.com/cilloparch/cillop/result"
	"github.com/gofiber/fiber/v2"
)

func (s *srv) ProductCreate(ctx *fiber.Ctx) error {
	cmd := &command.ProductCreateCmd{}
	s.parseBody(ctx, cmd)
	res, err := s.app.Commands.ProductCreate.Handle(ctx.UserContext(), *cmd)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessDetail(Messages.Ok, res)
}

func (s *srv) ProductGet(ctx *fiber.Ctx) error {
	q := &query.ProductGetQuery{}
	s.parseParams(ctx, q)
	res, err := s.app.Queries.ProductGet.Handle(ctx.UserContext(), *q)
	if err != nil {
		return result.Error(err.Error())
	}
	return result.SuccessDetail(Messages.Ok, res)
}
