package main

import "github.com/gin-gonic/gin"

func routes() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/api/student")
	{
		v1.POST("", createStudent)
		v1.GET("", fetchAllStudent)
		v1.GET("/:id", fetchSingleStudent)
		v1.PUT("/:id", updateStudent)
		v1.DELETE("/:id", deleteStudent)
	}

	return router
}