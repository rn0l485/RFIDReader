package products


import (
	"go-RC522/Config"
	"net/http"
	"github.com/gin-gonic/gin"
)


func InitRouter(R *gin.Engine) *gin.Engine {
	p := R.Group("/products")
	{
		// persion page "/persion"
		p.GET("", func(c *gin.Context){
			c.HTML( http.StatusOK, "products.html", config.Config)
		})

	}
	return R
}
