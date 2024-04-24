package controllers

import (
	"dynamic-mongodb/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type DataController struct {
	DataService services.DataService
}

func NewDataController(dataService services.DataService) *DataController {
	return &DataController{
		DataService: dataService,
	}
}

func (dc *DataController) CreateData(c *gin.Context) {
	var data bson.M
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dc.DataService.CreateData(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data created successfully"})
}

func (dc *DataController) GetData(c *gin.Context) {
	identify := c.Param("identify")
	name := c.Param("name")
	data, err := dc.DataService.GetData(identify, name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (dc *DataController) GetAllData(c *gin.Context) {
	dataList, err := dc.DataService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dataList)
}

func (dc *DataController) UpdateData(c *gin.Context) {
	identify := c.Param("identify")
	name := c.Param("name")
	var data bson.M
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dc.DataService.UpdateData(identify, name, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func (dc *DataController) DeleteData(c *gin.Context) {
	identify := c.Param("identify")
	name := c.Param("name")
	fmt.Println(identify)
	fmt.Println(name)
	if err := dc.DataService.DeleteData(identify, name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
