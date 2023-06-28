package webServer

import (
	"log"

	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/mercadolibre/fury_go-toolkit-kvs/pkg/kvs"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/application/fruit"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/domain/general"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/infra/config"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/infra/repositories"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/infra/webServer/controllers"
	"github.com/mercadolibre/fury_rampup-vitormoschetta/cmd/internal/infra/webServer/middlewares"
)

func Run() error {
	app, err := fury.NewWebApplication()
	if err != nil {
		log.Fatal(err)
	}

	config.LoadInitialConfig()

	app.Use(middlewares.LoggingMiddleware())

	// log.Print(os.Getenv(config.KvsContainerNameKey))
	// log.Print(os.Getenv(config.KvsContainerKey))
	// log.Print(os.Getenv(config.KvsHostReadKey))
	// log.Print(os.Getenv(config.KvsHostWriteKey))
	// log.Print(os.Getenv(config.KvsEnabledKey))

	log.Print("KvsEnabledKey: " + config.ConfigMap[config.KvsEnabledKey])
	log.Print("KvsContainerNameKey: " + config.ConfigMap[config.KvsContainerNameKey])

	var repository general.IRepository
	if config.ConfigMap[config.KvsEnabledKey] == "false" {
		repository = repositories.NewMemoryRepository()
	} else {
		kvsClient, err := kvs.NewClient(config.ConfigMap[config.KvsContainerNameKey])
		if err != nil {
			log.Fatal("Error creating kvs client", err)
		}
		repository = repositories.NewKvsRepository(kvsClient)
	}

	fruitUseCase := fruit.NewFruitUseCase(repository)
	fruitController := controllers.NewFruitController(repository, fruitUseCase)
	helloController := controllers.NewHelloController()

	app.Get("/hello", helloController.HelloController)
	app.Post("/hello", helloController.HelloController)
	app.Post("/fruits", fruitController.PostFruitsController)
	app.Get("/fruits/{id}", fruitController.GetFruitController)
	app.Get("/fruits", fruitController.GetFruitsController)

	return app.Run()
}
