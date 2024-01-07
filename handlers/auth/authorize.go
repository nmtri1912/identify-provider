package auth

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type AuthorizeRequest struct {
	Allow        bool   `json:allow`
	Email        string `json:"username"`
	Password     string `json:"password"`
	ResponseType string `json:"response_type"`
}

const REDIRECT_URL = "localhost:3000/authorization/"

func (h authHandler) authorize(ctx *gin.Context) {
	request := AuthorizeRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	// TODO: GET SEESSION ID => GET USER USING SESSION ID
	user, err := h.oauthService.GetAuthUser(
		request.Email,
		request.Password,
	)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	if !request.Allow {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	// TODO: CHECK USER SCOPE

	redictUrl, _ := url.ParseRequestURI(REDIRECT_URL)
	query := redictUrl.Query()

	if request.ResponseType == "code" {
		authoCode, err := 
	}

	// TODO: token
	if request.ResponseType == "token" {

	}

	ctx.JSON(http.StatusOK, user)
}
