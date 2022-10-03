package routes

import (
	_ "github.com/SpawNKZ/football-team-search-platform/docs"
	"github.com/SpawNKZ/football-team-search-platform/internal/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/user")
		{
			users.GET("", controller.GetUser)
		}
	}
}
