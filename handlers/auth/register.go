package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Email    string `json:"username"`
	Password string `json:"password"`
}

func (h authHandler) register(ctx *gin.Context) {
	request := LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	if h.oauthService.UserExists(request.Email) {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	if _, err := h.oauthService.CreateUser(
		request.Email,
		request.Password,
	); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.JSON(http.StatusOK, "")
}
