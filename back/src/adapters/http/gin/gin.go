package gin

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/domain/config"
	"github.com/AliceDiNunno/yeencloud/src/core/interactor"
	"github.com/AliceDiNunno/yeencloud/src/core/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ServiceHTTPServer struct {
	engine *gin.Engine

	config        config.HTTPConfig
	versionConfig config.VersionConfig
	log           interactor.Logger

	ucs        usecases.Usecases
	translator *i18n.Bundle
	validator  interactor.Validator
	auditer    interactor.Audit
}

func NewServiceHttpServer(ucs usecases.Usecases, config config.HTTPConfig, log interactor.Logger, version config.VersionConfig, translator *i18n.Bundle, validator interactor.Validator, auditer interactor.Audit) *ServiceHTTPServer {
	server := ServiceHTTPServer{
		config:        config,
		versionConfig: version,
		log:           log,

		ucs:        ucs,
		translator: translator,
		validator:  validator,
		auditer:    auditer,
	}

	gin.DebugPrintRouteFunc = server.printRoutes

	// TODO: use config
	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	allowHeaders := fmt.Sprintf("%s, %s, %s, %s", HeaderAuthorization, HeaderContentType, HeaderAcceptLanguage, HeaderUserAgent)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FrontendURL},
		AllowMethods:     []string{MethodPut, MethodPatch, MethodGet, MethodPost, MethodOption, MethodDelete},
		AllowHeaders:     []string{allowHeaders},
		ExposeHeaders:    []string{HeaderContentLength},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(server.ginLogger())
	r.Use(gin.Recovery())

	server.engine = r

	server.SetRoutes()
	return &server
}

func (server *ServiceHTTPServer) Listen() error {
	httpserver := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", server.config.ListeningAddress, server.config.ListeningPort),
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           server.engine,
	}

	err := httpserver.ListenAndServe()
	if err != nil {
		return err
	}

	server.log.Log(domain.LogLevelInfo).Msg("HTTP Server started")
	err = server.engine.Run()

	if err != nil {
		return err
	}

	return nil
}
