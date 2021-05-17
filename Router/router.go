package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-RC522/Config"
)

var R *gin.Engine

func InitRouter() {
	R = gin.Default()

	R.LoadHTMLGlob("./Router/templates/*.html")
	R.Static("/static", "./Router/static")

	R.GET("/", templateHTML)

	R.NoRoute(pageNotFound)
	R.NoMethod(pageNotFound)
}




func templateHTML( c *gin.Context) {
	c.HTML( http.StatusOK, "index.html", config.Config)
}

func pageNotFound(c *gin.Context){
	c.HTML( http.StatusNotFound, "error.html", config.Config)
}