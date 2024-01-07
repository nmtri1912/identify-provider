package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

func (h authHandler) login(ctx *gin.Context) {
	request := LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	user, err := h.oauthService.GetAuthUser(
		request.Email,
		request.Password,
	)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	// ADD SESSION AUTHENTICATION

	// // Redirect to the authorize page by default but allow redirection to other
	// // pages by specifying a path with login_redirect_uri query string param
	// loginRedirectURI := ctx.Request.URL.Query().Get("login_redirect_uri")
	// if loginRedirectURI == "" {
	// 	loginRedirectURI = "/web/admin"
	// }

	ctx.JSON(http.StatusOK, user)
}
