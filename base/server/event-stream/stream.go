package event_stream

import (
	"github.com/cilloparch/boilerplates/base/app"
	"github.com/cilloparch/boilerplates/base/config"
	"github.com/cilloparch/cillop/events"
	"github.com/cilloparch/cillop/server"
)

type srv struct {
	config config.App
	app    app.Application
	engine events.Engine
}

type Config struct {
	Config config.App
	App    app.Application
	Engine events.Engine
}

func New(cfg Config) server.Server {
	return &srv{
		config: cfg.Config,
		app:    cfg.App,
		engine: cfg.Engine,
	}
}

func (s *srv) Listen() error {
	return s.engine.Subscribe(s.config.EventStream.TopicProductCreate, s.CreateProduct)
}
