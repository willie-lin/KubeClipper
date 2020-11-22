package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "123",
		})
	})


	return router

}