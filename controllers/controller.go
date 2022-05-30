package controllers

import (
	"fmt"
	"log"
	"net/http"
	"newassignmen2/db"
	"newassignmen2/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItems(c *gin.Context) {

	var (
		items  []models.Items
		orders models.CreateOrders
	)

	db := db.GetDB()

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	insertOrder := models.Orders{
		CustomerName: orders.CustomerName,
		OrderedAt:    orders.OrderedAt,
	}

	db.Create(&insertOrder)
	orderID := insertOrder.ID

	for _, v := range orders.Item {
		item := models.Items{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderId:     insertOrder.ID,
		}

		items = append(items, item)
	}

	result := db.Create(&items)
	log.Println(orderID, result.RowsAffected)

	responseData := models.Orders{
		ID:           items[0].OrderId,
		CustomerName: orders.CustomerName,
		OrderedAt:    time.Now(),
		Item:         items,
	}

	c.JSON(http.StatusOK, responseData)
}

func GetItems(c *gin.Context) {
	db := db.GetDB()

	orders := []models.Orders{}
	err := db.Preload("Item").Find(&orders).Error

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func DeleteItems(c *gin.Context) {
	db := db.GetDB()

	strId := c.Param("orderID")
	id, _ := strconv.Atoi(strId)

	orders := models.Orders{}
	items := models.Items{}

	err := db.Where("order_id = ?", id).Delete(&items).Error
	rowsAff := db.Where("ID = ?", id).Delete(&orders).RowsAffected

	if rowsAff == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ID not found",
		})
		return
	}

	if err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Deleted",
	})
}

func UpdateItems(c *gin.Context) {
	db := db.GetDB()

	strId := c.Param("orderID")
	id, _ := strconv.Atoi(strId)

	var orders = models.Orders{}

	if err := c.ShouldBindJSON(&orders); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updateOrder := models.Orders{
		ID:           uint(id),
		CustomerName: orders.CustomerName,
		OrderedAt:    orders.OrderedAt,
		Item:         orders.Item,
	}

	db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateOrder)

	c.JSON(http.StatusOK, updateOrder)
}
