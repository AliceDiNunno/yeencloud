package apps

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/audit"
	"github.com/AliceDiNunno/yeencloud/src/adapters/cluster/k8s"
	"github.com/AliceDiNunno/yeencloud/src/adapters/http/gin"
	"github.com/AliceDiNunno/yeencloud/src/adapters/log"
	"github.com/AliceDiNunno/yeencloud/src/adapters/log/reporting/rollbar"
	"github.com/AliceDiNunno/yeencloud/src/adapters/log/terminal/zerolog"
	"github.com/AliceDiNunno/yeencloud/src/adapters/persistence/database/postgres"
	"github.com/AliceDiNunno/yeencloud/src/adapters/validator"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
	"github.com/AliceDiNunno/yeencloud/src/core/usecases"
)

func MainBackend(bundle *ApplicationBundle) error {
	httpConfig := bundle.Config.GetHTTPConfig()
	databaseConfig := bundle.Config.GetDatabaseConfig()
	version := bundle.Config.GetVersionConfig()
	rollbarConfig := bundle.Config.GetRollbarConfig()
	runContext := bundle.Config.GetRunContextConfig()

	_ = bundle.Config.GetKubernetesConfig()

	validator := validator.NewValidator()
	logger := log.Logger()

	zlogmiddle := zerolog.NewZeroLogMiddleware()
	rollbarmiddle := rollbar.NewRollbarMiddleware(rollbarConfig, runContext, version)

	logger.AddMiddleware(zlogmiddle)
	logger.AddMiddleware(rollbarmiddle)

	logger.Log(domain.LogLevelDebug).
		WithField(domain.LogFieldConfigVersion, version).
		WithField(domain.LogFieldConfigDatabase, databaseConfig).
		WithField(domain.LogFieldConfigHTTP, httpConfig).
		WithField(domain.LogFieldConfigRunContext, runContext).
		Msg("Starting backend")

	// #YC-12 TODO: make database dependent on config in order to have a local database for tests
	logger.Log(domain.LogLevelInfo).Msg("Connecting to database")

	database, err := postgres.StartGormDatabase(logger, databaseConfig)
	if err != nil {
		logger.Log(domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Error connecting to database")
		return err
	}
	err = database.Migrate()
	if err != nil {
		logger.Log(domain.LogLevelError).WithField(domain.LogFieldError, err).Msg("Error migrating database")
		return err
	}

	// #YC-13 TODO: pass the kubernetes config to the k8s adapter
	cluster := k8s.NewCluster()

	auditer := audit.NewAuditer(logger, nil)

	persistence := usecases.NewPersistence(database, database, database, database, database, database)

	ucs := usecases.NewUsecases(cluster, bundle.Translator, validator, auditer, persistence)

	http := gin.NewServiceHttpServer(ucs, httpConfig, logger, version, bundle.Translator, validator, auditer)

	return http.Listen()
}
