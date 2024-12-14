package api

import "github.com/gin-gonic/gin"

func AuthRoutes(route *gin.Engine) {
	user := route.Group("/auth")
	user.POST("/signup")
	user.POST("/login")
	user.GET("/info")
}

func DocRoutes(route *gin.Engine) {
	doc := route.Group("/doc")
	doc.POST("/create")
	doc.GET("/get")
	doc.GET("/getAll")
	doc.GET("/edit")
	doc.GET("/delete")
}
