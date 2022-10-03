package main

import (
	"github.com/SpawNKZ/football-team-search-platform/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoutes(router)
	router.Run(":8080")
}
