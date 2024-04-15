package gin

import (
	"github.com/AliceDiNunno/yeencloud/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (server *ServiceHTTPServer) setTokenRoutes(tokenGroup *gin.RouterGroup) {
	tokens := tokenGroup.Group("")

	tokens.POST("validate_mail", server.validateMailHandler)
	tokens.POST("forgotten_password", server.forgottenPasswordHandler)
	tokens.POST("recover_password", server.recoverPasswordHandler)
}

func (server *ServiceHTTPServer) validateMailHandler(ctx *gin.Context) {
	var validateMailRequest domain.ValidateMail

	if err := ctx.ShouldBindJSON(&validateMailRequest); err != nil {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	if !server.validate(ctx, validateMailRequest) {
		return
	}

	audit := server.getTrace(ctx)

	session, err := server.usecases(ctx).ValidateMail(audit, validateMailRequest)

	if err != nil {
		server.abortWithError(ctx, err)
		return
	}

	server.created(ctx, session)
}

func (server *ServiceHTTPServer) forgottenPasswordHandler(ctx *gin.Context) {
	var validateMailRequest domain.RequestNewPassword

	if err := ctx.ShouldBindJSON(&validateMailRequest); err != nil {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	if !server.validate(ctx, validateMailRequest) {
		return
	}

	audit := server.getTrace(ctx)

	_ = audit
}

func (server *ServiceHTTPServer) recoverPasswordHandler(ctx *gin.Context) {
	var recoverPasswordRequest domain.RecoverPassword

	if err := ctx.ShouldBindJSON(&recoverPasswordRequest); err != nil {
		server.abortWithError(ctx, &BadRequestError{})
		return
	}

	if !server.validate(ctx, recoverPasswordRequest) {
		return
	}

	audit := server.getTrace(ctx)

	_ = audit
}
