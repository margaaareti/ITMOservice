package delivery

import (
	"github.com/gin-gonic/gin"
	"itmo/auth"
)

func RegisterHTTPEndPoints(router *gin.Engine, service auth.AuthService) {

	h := CreateNewHandler(service)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/register", h.SignUp)
	}

}
