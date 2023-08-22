package service

import (
	"github.com/cilloparch/boilerplates/base/app"
	"github.com/cilloparch/boilerplates/base/app/command"
	"github.com/cilloparch/boilerplates/base/app/query"
	"github.com/cilloparch/boilerplates/base/config"
	"github.com/cilloparch/boilerplates/base/domains/product"
	"github.com/cilloparch/cillop/db/mongodb"
)

type Config struct {
	App config.App
	Db  *mongodb.DB
}

func NewApp(config Config) app.Application {
	productRepo := product.NewRepository(config.Db.GetCollection(config.App.ProductMongo.Collection))

	return app.Application{
		Commands: app.Commands{
			ProductCreate: command.NewProductCreateHandler(productRepo),
		},
		Queries: app.Queries{
			ProductGet: query.NewProductGetHandler(productRepo),
		},
	}
}
