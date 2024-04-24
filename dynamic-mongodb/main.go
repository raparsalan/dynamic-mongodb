package main

import (
	"dynamic-mongodb/controllers"
	"dynamic-mongodb/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/create/:database/:collection", func(c *gin.Context) {
		database := c.Param("database")
		collection := c.Param("collection")

		service, err := services.NewMongoDBDataService("mongodb://localhost:27017", database, collection)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		controller := controllers.NewDataController(service)
		controller.CreateData(c)
	})

	r.GET("/api/get/:database/:collection/:identify/:name", func(c *gin.Context) {
		database := c.Param("database")
		collection := c.Param("collection")

		service, err := services.NewMongoDBDataService("mongodb://localhost:27017", database, collection)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		controller := controllers.NewDataController(service)
		controller.GetData(c)
	})

	r.GET("/api/getall/:database/:collection", func(c *gin.Context) {
		database := c.Param("database")
		collection := c.Param("collection")

		service, err := services.NewMongoDBDataService("mongodb://localhost:27017", database, collection)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		controller := controllers.NewDataController(service)
		controller.GetAllData(c)
	})

	r.PATCH("/api/update/:database/:collection/:indetify/:name", func(c *gin.Context) {
		database := c.Param("database")
		collection := c.Param("collection")

		service, err := services.NewMongoDBDataService("mongodb://localhost:27017", database, collection)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		controller := controllers.NewDataController(service)
		controller.UpdateData(c)
	})

	r.DELETE("/api/delete/:database/:collection/:identify/:name", func(c *gin.Context) {
		database := c.Param("database")
		collection := c.Param("collection")

		service, err := services.NewMongoDBDataService("mongodb://localhost:27017", database, collection)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		controller := controllers.NewDataController(service)
		controller.DeleteData(c)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
