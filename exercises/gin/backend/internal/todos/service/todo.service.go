package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos []string

func Lists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func ListItem(c *gin.Context) {
	errormessage := "index out of range"
	indexstring := c.Param("index")
	if index, err := strconv.Atoi(indexstring); err == nil && index < len(todos) {
		c.JSON(http.StatusOK, gin.H{"item": todos[index]})
	} else {
		if err != nil {
			errormessage = fmt.Sprintf("number expected: %s\n", indexstring)
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func AddListItem(c *gin.Context) {
	item := c.PostForm("item")
	todos = append(todos, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
}
