package http

import (
	"github.com/gin-gonic/gin"

	"go-clean-architecture/internal/author"
)

type AuthorHandler struct {
	uc *author.UseCase
}

func NewAuthorHandler(uc *author.UseCase) *AuthorHandler {
	return &AuthorHandler{uc: uc}
}

func (h *AuthorHandler) Create(c *gin.Context) (*author.Author, error) {
	var req struct {
	}
}
