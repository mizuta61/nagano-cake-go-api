package controller

import (
	"strconv"

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

func (*GenreController) UpdateGenre(c *gin.Context) {
	ID := c.Params.ByName("id")
	genreID, _ := strconv.Atoi(ID)
	genre, err := model.GetGenreById(connecter.DB(), genreID)
	if err != nil {
		c.JSON(400, gin.H{"error": "genre not found"})
		return
	}
	var param GenreParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updateParam := map[string]interface{}{
		"name": param.Name,
	}
	_, err = genre.Update(connecter.DB(), updateParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "genre update failed"})
	}
	c.JSON(200, gin.H{"genre": genre})
}

func (*GenreController) GetGenre(c *gin.Context) {
	ID := c.Params.ByName("id")
	genreID, _ := strconv.Atoi(ID)
	genre, err := model.GetGenreById(connecter.DB(), genreID)
	if err != nil {
		c.JSON(400, gin.H{"error": "genre not found"})
		return
	}
	c.JSON(200, gin.H{"genre": genre})
}

func (*GenreController) Index(c *gin.Context) {
	genres, err := model.GetGenres(connecter.DB())
	if err != nil {
		c.JSON(400, gin.H{"error": "genre search failed"})
		return
	}
	c.JSON(200, gin.H{"genres": genres})
}
