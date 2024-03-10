package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/idgenms/controller"
	"github.com/prakash-p-3121/idgenms/repository/impl"
	"github.com/prakash-p-3121/mysqllib"
)

func main() {

	databaseInst, err := mysqllib.CreateDatabaseConnection("conf/database.toml")
	if err != nil {
		panic(err)
	}
	impl.SetDatabaseInstance(databaseInst)

	router := gin.Default()
	routerGroup := router.Group("/idgen")
	routerGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	routerGroup.POST("/next-id", controller.NextID)
	router.Run("127.0.0.1:3001")
}
