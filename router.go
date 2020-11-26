package main

import (
	. "HOMEWORK/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/role", GetRolesApi)
	/*
		router.GET("/role", GetPersonsApi)

		router.POST("/role", AddPersonApi)

		router.PUT("/role/:id", ModPersonApi)

		router.DELETE("/role/:id", DelPersonApi)
	*/
	return router
}
