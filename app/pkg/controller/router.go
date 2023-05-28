package controller

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		u := UserController{}
		users.GET("", u.Index)
		users.GET("/:id", u.GetUser)
		users.POST("", u.CreateUser)
		users.PATCH("/:id", u.UpdateUser)
		users.DELETE("/:id", u.DeleteUser)
	}
	items := r.Group("/items")
	{
		i := ItemController{}
		items.POST("", i.CreateItem)
		items.GET("", i.Index)
		items.GET("/:id", i.GetItem)
		items.PATCH("/:id", i.UpdateItem)
	}
	genres := r.Group("/genres")
	{
		g := GenreController{}
		genres.POST("", g.CreateGenre)
		genres.PATCH("/:id", g.UpdateGenre)
		genres.GET("/:id", g.GetGenre)
		genres.GET("", g.Index)
	}
}
