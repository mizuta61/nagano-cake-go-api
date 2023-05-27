package controller

import (
	"github.com/fuhiz/docker-go-sample/app/pkg/connecter"
	"github.com/fuhiz/docker-go-sample/app/pkg/model"
	"github.com/gin-gonic/gin"
)

type GenreController struct{}

type GenreParam struct {
	Name string `json:"name"`
}

func (*GenreController) CreateGenre(c *gin.Context) {
	var param GenreParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newGenre := model.NewGenre(param.Name)
	genre, err := model.CreateGenre(connecter.DB(), newGenre)
	if err != nil {
		c.JSON(400, gin.H{"error": "genre create failed"})
	}
	c.JSON(200, gin.H{"genre": genre})
}
