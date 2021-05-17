package persion


import (
	"go-RC522/Config"
	"net/http"
	"github.com/gin-gonic/gin"
)


func InitRouter(R *gin.Engine) *gin.Engine {
	p := R.Group("/persion")
	{
		// persion page "/persion"
		p.GET("", func(c *gin.Context){
			c.HTML( http.StatusOK, "persion.html", config.Config)
		})

	}
	return R
}

