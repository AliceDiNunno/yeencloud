package apps

import (
	"back/src/adapters/cluster/k8s"
	"back/src/adapters/http/gin"
	"back/src/adapters/persistence/database/postgres"
	"back/src/adapters/validator/govalidator"
	"back/src/core/domain"
	"back/src/core/usecases"
	"github.com/rs/zerolog/log"
)

func MainBackend(bundle *domain.ApplicationBundle) error {
	httpConfig := bundle.Config.GetHTTPConfig()
	databaseConfig := bundle.Config.GetDatabaseConfig()
	_ = bundle.Config.GetKubernetesConfig()
	_ = bundle.Config.GetVersionConfig()

	validator := govalidator.NewValidator()

	//TODO: make database dependent on config in order to have a local database for tests
	log.Info().Msg("Connecting to database")
	database := postgres.StartGormDatabase(databaseConfig, "default")
	database.Migrate()

	//TODO: pass the kubernetes config to the k8s adapter
	cluster := k8s.NewCluster()

	ucs := usecases.NewInteractor(cluster, database, database, database, database, database, bundle.Translator, validator)

	http := gin.NewServiceHttpServer(ucs, httpConfig, bundle.Translator, validator)

	return http.Listen()
}
