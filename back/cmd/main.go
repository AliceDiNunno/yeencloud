package main

import (
	"back/src/apps"
	"back/src/core/config/env"
	"fmt"
	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func initConfig(c *cli.Context) *env.Config {
	configFile := c.String("config")

	if configFile == "" {
		configFile = ".env"
	}

	log.Info().Msg("Loading configuration from %s + configFile")
	err := godotenv.Load(configFile)
	if err != nil {
		log.Err(err).Msg("Error loading configuration")
	}

	cfg := env.NewConfig(c)

	return cfg
}

func loadTranslator() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	translationDir := "./src/locale"

	files, err := os.ReadDir(translationDir)
	if err != nil {
		log.Err(err).Msg("Error reading translation directory")
	}

	for _, file := range files {
		path := fmt.Sprintf("%s/%s", translationDir, file.Name())

		if strings.HasSuffix(path, ".toml") {
			bundle.MustLoadMessageFile(path)
		}
	}

	langs := []string{}
	for _, availableLang := range bundle.LanguageTags() {
		langs = append(langs, availableLang.String())
	}

	log.Info().Strs("languages", langs).Msg("Loaded languages")

	return bundle
}

func loadLogger() {
	isDev := os.Getenv("ENV") != "prod" && os.Getenv("ENV") != "production"

	if isDev {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// Short caller (file:line)
	zerolog.CallerMarshalFunc = func(_ uintptr, file string, line int) string {
		file = strings.TrimPrefix(file, "/app/")

		return fmt.Sprintf("%s:%d", file, line)
	}

	logFormat := os.Getenv("LOG_FORMAT")

	if logFormat == "" {
		if isDev {
			logFormat = "console"
		} else {
			logFormat = "json"
		}
	}

	if logFormat != "json" && logFormat != "JSON" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	log.Logger = log.With().Caller().Logger()

	if isDev {
		log.Warn().Msg("Running in development mode logging can be verbose or contain sensitive information")
		log.Warn().Msg("To run in production mode set the ENV environment variable to `prod` or `production`")
	}
}

func initApp(c *cli.Context) *apps.ApplicationBundle {
	cfg := initConfig(c)

	loadLogger()
	i18nTranslator := loadTranslator()

	return &apps.ApplicationBundle{
		Config:     cfg,
		Translator: i18nTranslator,
	}
}

func main() {
	zerolog.DurationFieldInteger = false
	zerolog.DurationFieldUnit = time.Millisecond

	app := &cli.App{
		Name:  "backend",
		Usage: "start backend server",
		Action: func(c *cli.Context) error {
			cfg := initApp(c)
			log.Info().Msg("Starting backend server")
			return apps.MainBackend(cfg)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("An error occurred while running the application")
	}
}
