package apps

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/validator"
	testConf "github.com/AliceDiNunno/yeencloud/internal/core/config/testConfig"
)

func TestInteractor() /*usecases.Interactor*/ {
	testConfig := testConf.NewConfig()

	httpConfig := testConfig.GetHTTPConfig()
	databaseConfig := testConfig.GetDatabaseConfig()

	validator := validator.NewValidator()
	_ = validator

	_, _ = httpConfig, databaseConfig

	/*ucs := usecases.NewInteractor(a, b, c, d, e, f, g, h, i, j, k, l, m, n)

	return ucs*/
}
