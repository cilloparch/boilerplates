package http

import (
	"github.com/cilloparch/boilerplates/base/app"
	"github.com/cilloparch/boilerplates/base/config"
	"github.com/cilloparch/cillop/helpers/http"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/server"
	"github.com/cilloparch/cillop/validation"
	"github.com/gofiber/fiber/v2"
)

type srv struct {
	config config.App
	app    app.Application
	valid  validation.Validator
	i18n   *i18np.I18n
}

type Config struct {
	Config config.App
	App    app.Application
	Valid  validation.Validator
	I18n   *i18np.I18n
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
	http.RunServer(http.Config{
		Host:        s.config.Http.Host,
		Port:        s.config.Http.Port,
		I18n:        s.i18n,
		AcceptLangs: []string{},
		CreateHandler: func(router fiber.Router) fiber.Router {
			router.Post("/", s.ProductCreate)
			router.Get("/:id", s.ProductGet)
			return router
		},
	})
	return nil
}

func (s *srv) parseBody(ctx *fiber.Ctx, dto interface{}) {
	http.ParseBody(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) parseParams(ctx *fiber.Ctx, dto interface{}) {
	http.ParseParams(ctx, s.valid, *s.i18n, dto)
}

func (s *srv) parseQuery(ctx *fiber.Ctx, dto interface{}) {
	http.ParseQuery(ctx, s.valid, *s.i18n, dto)
}
