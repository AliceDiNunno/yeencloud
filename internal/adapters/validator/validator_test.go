package validator

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ValidationTestSuite struct {
	suite.Suite

	validator Validator
}

func (suite *ValidationTestSuite) SetupSuite() {
	suite.validator = NewValidator()
}

func TestValidatorSuite(t *testing.T) {
	//Testing toDomain and fromDomain
	suite.Run(t, new(ValidationTestSuite))
}
