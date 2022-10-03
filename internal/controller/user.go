package controller

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	c.String(200, "You got user")
}
