package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"HOMEWORK/MyType"

	"github.com/gin-gonic/gin"
)

var roles = []MyType.Role{}

func main() {
	//讀取JSON資料
	initData()

	router := gin.Default()

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

	router.Run(":8080")
}

func initData() {
	jsonData := `[
		{
			"id": 1,
			"name": "阿修羅",
			"summary": "死國魖族最強者，在歷經多場戰役統一死國，直接挑戰天者權威，不服其領導，要求讓三族和平共處，因此被天者以增加資源為名，前往打通連接火宅佛獄、死國及苦境的莫汗走廊，未料在工程半途被楓岫主人所乘駕的隕石撞毀，阿修羅在天者刻意算計下意外身亡，身葬苦境與死國間的異次元。",
			"skills": [
				{
					"id": 1,
					"type": "武學",
					"name": "天之爆"
				},
				{
					"id": 2,
					"type": "武學",
					"name": "魔之狂"
				},
				{
					"id": 3,
					"type": "武學",
					"name": "天之渦"
				},
				{
					"id": 4,
					"type": "武學",
					"name": "闇之爆"
				},
				{
					"id": 5,
					"type": "法術",
					"name": "山河凝元·天地共引"
				},
				{
					"id": 6,
					"type": "法術",
					"name": "地之火·九天滅絕"
				}
			]
		},
		{
			"id": 2,
			"name": "白塵子",
			"summary": "火宅佛獄凱旋侯的副體之一，本名黑枒君，臥底中原武林，向佛獄通風報信，最後被素還真所殺並冒充其身份一探佛獄之秘。",
			"skills": [
				{
					"id": 7,
					"type": "武學",
					"name": "凝宇化空"
				},
				{
					"id": 8,
					"type": "武學",
					"name": "反神源"
				}
			]
		}
	]`
	json.Unmarshal([]byte(jsonData), &roles)
}
