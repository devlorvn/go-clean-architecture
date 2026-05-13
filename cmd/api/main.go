package main

import (
	"go-clean-architecture/internal/author"
	httpDelivery "go-clean-architecture/internal/delivery/http"
	memoryRepo "go-clean-architecture/internal/infras/repository/memory"
	"go-clean-architecture/internal/post"
)

func main() {
	authorRepo := memoryRepo.NewAuthorMemory()
	authorUC := author.NewUseCase(authorRepo)
	authorHandler := httpDelivery.NewAuthorHandler(authorUC)

	postRepo := memoryRepo.NewPostMemory()
	postUC := post.NewUseCase(postRepo)
	postHandler := httpDelivery.NewPostHandler(postUC)

	handleGroup := &httpDelivery.HandleGroup{
		Author: authorHandler,
		Post:   postHandler,
	}

	r := httpDelivery.NewRouter(handleGroup)
	r.Run(":8080")
}
