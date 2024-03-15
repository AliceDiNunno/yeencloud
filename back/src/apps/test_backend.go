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

	/*ucs := usecases.NewInteractor(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	return ucs*/
}
