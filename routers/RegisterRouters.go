package routers

import (
	. "HA-back-end/routers/admin"
	"github.com/gin-gonic/gin"
)

var Eng *gin.Engine

func Init(){
	RegisterAdminRouters()
}
func RegisterAdminRouters()  {
	action:=Eng.Group("/")
	{
		action:=action.Group("/admin")
		{
			action.POST("login",Login)
			action.POST("register", CreateUser)
		}
	}
}