package record

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {

	route := r.Group("/record")
	route.POST("/add", getHealth)
	route.POST("/remove", getHealth)

}
