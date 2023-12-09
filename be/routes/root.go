package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.Engine) gin.HandlerFunc {
	fmt.Println("Home route")

	// load html js css
	router.LoadHTMLGlob("fe/*.html")
	router.Static("/static", "fe/static")

	return func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	}
}
