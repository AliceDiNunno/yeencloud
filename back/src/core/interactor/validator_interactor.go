package interactor

import (
	"github.com/AliceDiNunno/yeencloud/src/adapters/validator"
	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

type Validator interface {
	Validate(s interface{}) (bool, domain.ValidationErrors)
	RegisterValidation(tag string, fn validator.ValidationFunc)
}
