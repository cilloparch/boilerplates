package app

import (
	"github.com/cilloparch/boilerplates/base/app/command"
	"github.com/cilloparch/boilerplates/base/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	ProductCreate command.ProductCreateHandler
}

type Queries struct {
	ProductGet query.ProductGetHandler
}
