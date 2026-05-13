package http

import "github.com/gin-gonic/gin"

type HandleGroup struct {
	Author *AuthorHandler
	Post   *PostHandler
}

func NewRouter(h *HandleGroup) *gin.Engine {
	r := gin.Default()
	r.GET("/authors", h.Author.GetAll)
	r.GET("/authors/:id", h.Author.GetByID)
	r.POST("/authors", h.Author.Create)
	r.PUT("/authors/:id", h.Author.Update)
	r.DELETE("/authors/:id", h.Author.Delete)

	r.GET("/posts", h.Post.GetAll)
	r.GET("/posts/:id", h.Post.GetByID)
	r.POST("/posts", h.Post.Create)
	r.PUT("/posts/:id", h.Post.Update)
	r.DELETE("/posts/:id", h.Post.Delete)
	return r
}
