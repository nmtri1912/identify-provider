package main

import (
	"identify-provider/db"
	"identify-provider/handlers/auth"
	"identify-provider/services/oauth"
	"identify-provider/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {
	utils.LoadConfiguration()

	gormDB := db.NewDB()

	oauthService := oauth.NewOauthService(gormDB)

	router := gin.Default()
	router.Use(cors.New(CORSConfig()))

	webGroup := router.Group("/api")
	auth.RegisterWebHandler(webGroup, oauthService)

	router.Run("localhost:8080")
}
