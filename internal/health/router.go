package health

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {

	route := r.Group("/health")
	route.GET("", getHealth)

}
