package assets


import (
	"go-RC522/Config"
	"net/http"
	"github.com/gin-gonic/gin"
)


func InitRouter(R *gin.Engine) *gin.Engine {
	p := R.Group("/assets")
	{
		// persion page "/persion"
		p.GET("", func(c *gin.Context){
			c.HTML( http.StatusOK, "assets.html", config.Config)
		})

	}
	return R
}

