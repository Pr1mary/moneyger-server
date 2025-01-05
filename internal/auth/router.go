package auth

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {

	route := r.Group("/user")
	route.GET("/login", getHealth)
	route.GET("/detail", getHealth)

}
