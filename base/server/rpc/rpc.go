package rpc

import (
	"github.com/cilloparch/boilerplates/base/app"
	"github.com/cilloparch/boilerplates/base/config"
	"github.com/cilloparch/boilerplates/base/server/rpc/routes"
	"github.com/cilloparch/cillop/helpers/rpc"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"google.golang.org/grpc"
)

type srv struct {
	routes.ProductServiceServer
	config config.App
	app    app.Application
	valid  validation.Validator
	i18n   i18np.I18n
}

type Config struct {
	Config config.App
	App    app.Application
	Valid  validation.Validator
	I18n   i18np.I18n
}

func New(cfg Config) server.Server {
	return &srv{
		config: cfg.Config,
		app:    cfg.App,
		valid:  cfg.Valid,
		i18n:   cfg.I18n,
	}
}

func (s *srv) Listen() error {
	rpc.RunServer(s.config.Rpc.Port, func(server *grpc.Server) {
		routes.RegisterProductServiceServer(server, s)
	})
	return nil
}
