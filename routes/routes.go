package routes

import (
	"halosus/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	userController := new(controllers.UserController)

	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world.")
	})
	v1 := router.Group("/v1")
	{
		user := v1.Group("/user")
		{
			it := user.Group("/it")
			{
				it.POST("/register", userController.CreateIT)
				it.POST("/login", userController.ITLogin)
			}
			nurse := user.Group("/nurse")
			{
				nurse.POST("/register", userController.CreateNurse)
				nurse.PUT("/:userId", userController.UpdateNurse)
				nurse.DELETE("/:userId", userController.DeleteNurse)
				nurse.POST("/:userId/access", userController.GrantNurseAccess)
			}
		}
	}

	return router
}