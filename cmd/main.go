package main

import (
	"log"
	"os"

	"github.com/AliceDiNunno/yeencloud/internal/apps"
	"github.com/AliceDiNunno/yeencloud/internal/core/config/env"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func initConfig(c *cli.Context) *env.Config {
	configFile := c.String("config")

	if configFile == "" {
		configFile = ".env"
	}

	_ = godotenv.Load(configFile)

	cfg := env.NewConfig(c)

	return cfg
}

func initApp(c *cli.Context) *apps.ApplicationBundle {
	cfg := initConfig(c)

	return &apps.ApplicationBundle{
		Config: cfg,
	}
}

func main() {
	app := &cli.App{
		Name:  "backend",
		Usage: "start backend server",
		Action: func(c *cli.Context) error {
			cfg := initApp(c)
			return apps.MainBackend(cfg)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln("An error occurred while running the application : " + err.Error())
	}
}
