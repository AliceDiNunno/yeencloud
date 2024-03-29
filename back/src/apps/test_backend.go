package apps

import (
	"back/src/adapters/validator/govalidator"
	testConf "back/src/core/config/testConfig"
)

func TestInteractor() /*usecases.Interactor*/ {
	testConfig := testConf.NewConfig()

	httpConfig := testConfig.GetHTTPConfig()
	databaseConfig := testConfig.GetDatabaseConfig()

	validator := govalidator.NewValidator()
	_ = validator

	_, _ = httpConfig, databaseConfig

	/*ucs := usecases.NewInteractor(a, b, c, d, e, f, g, h, i, j, k, l, m, n)

	return ucs*/
}
