package main

import (
	"github.com/cilloparch/boilerplates/base/config"
	event_stream "github.com/cilloparch/boilerplates/base/server/event-stream"
	"github.com/cilloparch/boilerplates/base/server/http"
	"github.com/cilloparch/boilerplates/base/server/rpc"
	"github.com/cilloparch/boilerplates/base/service"
	"github.com/cilloparch/cillop/db/mongodb"
	"github.com/cilloparch/cillop/env"
	"github.com/cilloparch/cillop/events/nats"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
)

func main() {
	cnf := config.App{}
	env.Load(&cnf)
	i18n := i18np.New(cnf.I18n.Fallback)
	i18n.Load(cnf.I18n.Dir, cnf.I18n.Locales...)
	valid := validation.New(i18n)
	valid.ConnectCustom()
	valid.RegisterTagName()
	mongo := loadMongo(cnf)
	eventEngine := nats.New(nats.Config{
		Url:     cnf.Nats.Url,
		Streams: cnf.Nats.Streams,
	})
	app := service.NewApp(service.Config{
		App: cnf,
		Db:  mongo,
	})
	http := http.New(http.Config{
		Config: cnf,
		App:    app,
		Valid:  *valid,
		I18n:   i18n,
	})
	rpc := rpc.New(rpc.Config{
		Config: cnf,
		App:    app,
		Valid:  *valid,
		I18n:   *i18n,
	})
	stream := event_stream.New(event_stream.Config{
		Config: cnf,
		App:    app,
		Engine: eventEngine,
	})
	go rpc.Listen()
	go stream.Listen()
	http.Listen()
}

func loadMongo(cnf config.App) *mongodb.DB {
	uri := mongodb.CalcMongoUri(mongodb.UriParams{
		Host:  cnf.ProductMongo.Host,
		Port:  cnf.ProductMongo.Port,
		User:  cnf.ProductMongo.Username,
		Pass:  cnf.ProductMongo.Password,
		Db:    cnf.ProductMongo.Database,
		Query: cnf.ProductMongo.Query,
	})
	d, err := mongodb.New(uri, cnf.ProductMongo.Database)
	if err != nil {
		panic(err)
	}
	return d
}
