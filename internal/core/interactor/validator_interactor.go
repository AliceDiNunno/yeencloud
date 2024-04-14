package interactor

import (
	"github.com/AliceDiNunno/yeencloud/internal/adapters/validator"
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
)

type Validator interface {
	Validate(s interface{}) (bool, domain.ValidationErrors)
	RegisterValidation(tag string, fn validator.ValidationFunc)
}
