package main

import (
	. "HOMEWORK/apis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", IndexApi)

	router.GET("/role", GetRolesApi)

	router.GET("/role/:id", GetRoleApi)

	router.POST("/role", PostRoleApi)

	router.PUT("/role/:id", PutRoleApi)

	router.DELETE("/role/:id", DeleteRoleApi)

	return router
}
