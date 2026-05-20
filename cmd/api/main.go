package main

import (
	"go-clean-architecture/internal/author"
	"go-clean-architecture/internal/config"
	httpDelivery "go-clean-architecture/internal/delivery/http"

	// memoryRepo "go-clean-architecture/internal/infras/repository/memory"
	database "go-clean-architecture/internal/infras/database/postgres"
	postgresRepo "go-clean-architecture/internal/infras/repository/postgres"
	"go-clean-architecture/internal/post"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		panic(err)
	}
	authorRepo := postgresRepo.NewAuthor(db)
	authorUC := author.NewUseCase(authorRepo)
	authorHandler := httpDelivery.NewAuthorHandler(authorUC)

	postRepo := postgresRepo.NewPost(db)
	postUC := post.NewUseCase(postRepo)
	postHandler := httpDelivery.NewPostHandler(postUC)

	handleGroup := &httpDelivery.HandleGroup{
		Author: authorHandler,
		Post:   postHandler,
	}

	r := httpDelivery.NewRouter(handleGroup)
	r.Run(":8080")
}
