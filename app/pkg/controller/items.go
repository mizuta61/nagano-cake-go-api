package controller

import (
	"strconv"

	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/model"
	"github.com/gin-gonic/gin"
)

type ItemController struct{}

type ItemParam struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Introduction string `json:"introduction"`
	IsActive     bool   `json:"is_active"`
}

func (*ItemController) Index(c *gin.Context) {
	items, err := model.GetItems(connecter.DB())

	if err != nil {
		c.JSON(400, gin.H{"error": "item search failed"})
		return
	}

	c.JSON(200, gin.H{"items": items})
}

func (*ItemController) CreateItem(c *gin.Context) {
	var param ItemParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newItem := model.NewItem(param.Name, param.Introduction, param.Price)
	item, err := model.CreateItem(connecter.DB(), newItem)

	if err != nil {
		c.JSON(400, gin.H{"error": "item create failed"})
	}

	c.JSON(200, gin.H{"item": item})
}

func (*ItemController) GetItem(c *gin.Context) {
	ID := c.Params.ByName("id")
	itemID, _ := strconv.Atoi(ID)
	item, err := model.GetItemById(connecter.DB(), itemID)
	if err != nil {
		c.JSON(400, gin.H{"error": "item not found"})
		return
	}
	c.JSON(200, gin.H{"item": item})
}

func (*ItemController) UpdateItem(c *gin.Context) {
	ID := c.Params.ByName("id")
	itemID, _ := strconv.Atoi(ID)
	item, err := model.GetItemById(connecter.DB(), itemID)
	if err != nil {
		c.JSON(400, gin.H{"error": "item not found"})
		return
	}
	var param ItemParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updateParam := map[string]interface{}{
		"name":         param.Name,
		"price":        param.Price,
		"introduction": param.Introduction,
		"is_active":    param.IsActive,
	}
	_, err = item.Update(connecter.DB(), updateParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "item update failed"})
		return
	}
	c.JSON(200, gin.H{"item": item})
}
