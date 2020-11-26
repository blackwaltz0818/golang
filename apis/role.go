package apis

import (
	"net/http"

	. "HOMEWORK/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

func GetRolesApi(c *gin.Context) {
	c.JSON(http.StatusOK, GetAllRoles())
}

/*
	router.GET("/role", func(c *gin.Context) {
		c.JSON(http.StatusOK, roles)
	})

	router.GET("/role/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		for i := 0; i < len(roles); i++ {
			if roles[i].ID == id {
				c.JSON(http.StatusOK, roles[i])
				break
			}
		}
	})

	router.POST("/role", func(c *gin.Context) {
		var r MyType.Role
		if err := c.ShouldBind(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		roles = append(roles, r)
		//c.Status(http.StatusOK)
		c.JSON(http.StatusOK, roles)
	})

	router.PUT("/role/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		var r MyType.Role
		if err := c.ShouldBind(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		for i := 0; i < len(roles); i++ {
			if roles[i].ID == id {
				roles[i].Name = r.Name
				roles[i].Summary = r.Summary
				break
			}
		}
		//c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, r)
	})

	router.DELETE("/role/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		for i, role := range roles {
			if role.ID == id {
				roles = append(roles[0:i], roles[i+1:]...)
				break
			}
		}

		c.Status(http.StatusNoContent)
	})
*/
