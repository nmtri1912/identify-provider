package auth

import (
	"identify-provider/services/oauth"

	"github.com/gin-gonic/gin"
)

func RegisterWebHandler(router *gin.RouterGroup,
	oauthService oauth.OauthService) {
	handler := &authHandler{
		oauthService: oauthService,
	}

	router.POST("register", handler.register)

	router.POST("login", handler.login)
}

type authHandler struct {
	oauthService oauth.OauthService
}
