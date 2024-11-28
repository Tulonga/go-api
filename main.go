package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Item struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var items = []Item{
    {ID: "1", Name: "Item One"},
    {ID: "2", Name: "Item Two"},
}

func main() {
    r := gin.Default()

    r.GET("/items", getItems)
    r.POST("/items", createItem)

    r.Run(":8080") 
}

func getItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

func createItem(c *gin.Context) {
    var newItem Item
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    items = append(items, newItem)
    c.JSON(http.StatusCreated, newItem)
}
