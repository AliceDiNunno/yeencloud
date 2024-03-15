package apps

import (
	"back/src/adapters/audit"
	"back/src/adapters/cluster/k8s"
	"back/src/adapters/http/gin"
	"back/src/adapters/persistence/database/postgres"
	"back/src/adapters/validator/govalidator"
	"back/src/core/usecases"
	"github.com/rs/zerolog/log"
)

func MainBackend(bundle *ApplicationBundle) error {
	httpConfig := bundle.Config.GetHTTPConfig()
	databaseConfig := bundle.Config.GetDatabaseConfig()
	version := bundle.Config.GetVersionConfig()
	_ = bundle.Config.GetKubernetesConfig()

	validator := govalidator.NewValidator()

	// #YC-12 TODO: make database dependent on config in order to have a local database for tests
	log.Info().Msg("Connecting to database")
	database, err := postgres.StartGormDatabase(databaseConfig)
	if err != nil {
		log.Info().Err(err).Msg("Error connecting to database")
		return err
	}
	database.Migrate()

	// #YC-13 TODO: pass the kubernetes config to the k8s adapter
	cluster := k8s.NewCluster()

	auditer := audit.NewAuditer(nil)

	ucs := usecases.NewInteractor(cluster, bundle.Translator, validator, auditer,
		database, database, database, database, database, database,
		database)

	http := gin.NewServiceHttpServer(ucs, httpConfig, version, bundle.Translator, validator, auditer)

	return http.Listen()
}
