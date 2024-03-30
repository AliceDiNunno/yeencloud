package gin

import (
	"net/http"

	"github.com/AliceDiNunno/yeencloud/src/core/domain"
)

var (
	TranslatablePageNotFound     = domain.Translatable{Key: "PageNotFound"}
	TranslatableMethodNotAllowed = domain.Translatable{Key: "MethodNotAllowed"}
	TranslatableBadRequest       = domain.Translatable{Key: "BadRequest"}
	TranslatableInternalError    = domain.Translatable{Key: "InternalError"}

	TranslatableAuthenticationTokenMissing = domain.Translatable{Key: "AuthenticationTokenMissing"}
)

var (
	ErrorPageNotFound     = domain.ErrorDescription{HttpCode: http.StatusNotFound, Code: TranslatablePageNotFound}
	ErrorMethodNotAllowed = domain.ErrorDescription{HttpCode: http.StatusMethodNotAllowed, Code: TranslatableMethodNotAllowed}
	ErrorBadRequest       = domain.ErrorDescription{HttpCode: http.StatusBadRequest, Code: TranslatableBadRequest}
	ErrorInternal         = domain.ErrorDescription{HttpCode: http.StatusInternalServerError, Code: TranslatableInternalError}

	ErrorAuthenticationTokenMissing = domain.ErrorDescription{HttpCode: http.StatusUnauthorized, Code: TranslatableAuthenticationTokenMissing}
)
