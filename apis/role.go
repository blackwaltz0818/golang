package apis

import (
	"net/http"
	"strconv"

	. "HOMEWORK/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func GetRolesApi(c *gin.Context) {
	InitData()
	c.JSON(http.StatusOK, SelectRoles())
}

func GetRoleApi(c *gin.Context) {
	InitData()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, SelectRole(id))
}

func PostRoleApi(c *gin.Context) {
	InitData()
	var r Role
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, InsertRole(r))
}

func PutRoleApi(c *gin.Context) {
	InitData()
	var r Role
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, UpdateRole(r))
}

func DeleteRoleApi(c *gin.Context) {
	InitData()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, DeleteRole(id))
}
