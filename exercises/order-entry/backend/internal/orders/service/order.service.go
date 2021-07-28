package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KernelGamut32/golang-codehouse/exercises/order-entry/backend/internal/orders"
	"github.com/gin-gonic/gin"
)

var orderSet []orders.Order

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"orders": orderSet})
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	if idValue, err := strconv.Atoi(id); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		for _, order := range orderSet {
			if order.ID == idValue {
				c.JSON(http.StatusOK, gin.H{"order": order})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"order": nil})
	}
}

func AddOrder(c *gin.Context) {
	var order orders.Order
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(orderSet) == 0 {
		order.ID = 1
	} else {
		order.ID = orderSet[len(orderSet)-1].ID + 1
	}
	orderSet = append(orderSet, order)
	c.Writer.WriteHeader(http.StatusCreated)
}

func UpdateOrder(c *gin.Context) {
	var order orders.Order
	err := json.NewDecoder(c.Request.Body).Decode(&order)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	id := c.Param("id")
	if idValue, err := strconv.Atoi(id); err != nil || idValue != order.ID {
		c.Writer.WriteHeader(http.StatusBadRequest)
	} else {
		for index, item := range orderSet {
			if item.ID == idValue {
				orderSet[index] = order
				c.Writer.WriteHeader(http.StatusOK)
			}
		}
	}
}
