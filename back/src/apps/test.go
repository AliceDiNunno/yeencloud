package apps

import (
	"back/src/adapters/validator/govalidator"
	testConf "back/src/core/config/testConfig"
)

func TestInteractor() {
	testConfig := testConf.NewConfig()

	httpConfig := testConfig.GetHTTPConfig()
	databaseConfig := testConfig.GetDatabaseConfig()

	validator := govalidator.NewValidator()
	_ = validator

	_, _ = httpConfig, databaseConfig
}
