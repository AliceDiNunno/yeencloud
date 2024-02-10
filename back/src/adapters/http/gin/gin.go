package gin

import (
	"back/src/core/config"
	"back/src/core/usecases"
	"fmt"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
	"os"
	"time"
)

type ServiceHTTPServer struct {
	engine *gin.Engine
	config config.HTTPConfig

	ucs        usecases.Usecases
	translator *i18n.Bundle
	validator  usecases.Validator
}

func NewServiceHttpServer(ucs usecases.Usecases, config config.HTTPConfig, translator *i18n.Bundle, validator usecases.Validator) *ServiceHTTPServer {
	server := ServiceHTTPServer{
		config:     config,
		ucs:        ucs,
		translator: translator,
		validator:  validator,
	}

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug().Str("Method", httpMethod).Str("Handler", handlerName).Int("Handlers", nuHandlers).Msg(absolutePath)
	}

	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(cors.New(cors.Config{
		//TODO: Change this to the real origin in production while keeping * for development
		AllowOrigins:     []string{config.FrontendURL},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTION", "DELETE"},
		AllowHeaders:     []string{"Origin, Authorization, Content-Type, Accept-Language"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(ginzerolog.Logger("backend"))
	r.Use(gin.Recovery())

	server.engine = r

	server.SetRoutes()
	return &server
}

func (server *ServiceHTTPServer) Listen() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.config.ListeningAddress, server.config.ListeningPort))

	if err != nil {
		return err
	}

	log.Info().Str("Address", ln.Addr().String()).Msg("Now Listening !")

	err = http.Serve(ln, server.engine)
	if err != nil {
		return err
	}
	err = server.engine.Run()

	if err != nil {
		return err
	}

	return nil
}

// TODO: Add language list
// TODO: Add git commit hash
func (server *ServiceHTTPServer) getStatusHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "OK",
	})
}
