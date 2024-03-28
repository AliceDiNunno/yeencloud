package main

import (
	"fmt"
	"github.com/AliceDiNunno/yeencloud/src/apps"
	"github.com/AliceDiNunno/yeencloud/src/core/config/env"
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/language"
	"log"
	"os"
	"strings"
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

func loadTranslator() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	translationDir := "./src/locale"

	files, err := os.ReadDir(translationDir)
	if err != nil {
		log.Fatalln("Error reading translation directory " + err.Error())
	}

	for _, file := range files {
		path := fmt.Sprintf("%s/%s", translationDir, file.Name())

		if strings.HasSuffix(path, ".toml") {
			bundle.MustLoadMessageFile(path)
		}
	}

	return bundle
}

func initApp(c *cli.Context) *apps.ApplicationBundle {
	cfg := initConfig(c)

	i18nTranslator := loadTranslator()

	return &apps.ApplicationBundle{
		Config:     cfg,
		Translator: i18nTranslator,
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
