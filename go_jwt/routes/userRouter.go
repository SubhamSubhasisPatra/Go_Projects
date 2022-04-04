package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("user/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}
