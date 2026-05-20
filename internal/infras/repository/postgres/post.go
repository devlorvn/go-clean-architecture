package postgres

import (
	"context"
	database "go-clean-architecture/internal/infras/database/postgres"
	postgresModel "go-clean-architecture/internal/infras/database/postgres/model"
	"go-clean-architecture/internal/post"

	"gorm.io/gorm"
)

type Post struct {
	db *gorm.DB
}

func NewPost(db *gorm.DB) *Post {
	return &Post{db: db}
}

func (p *Post) dbFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := database.TxFromContext(ctx); ok {
		return tx.WithContext(ctx)
	}
	return p.db.WithContext(ctx)
}

func (p *Post) GetByID(ctx context.Context, id int64) (*post.Post, error) {
	var postModel postgresModel.Post

	err := p.dbFromContext(ctx).First(&postModel, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &post.Post{
		ID:       postModel.ID,
		Title:    postModel.Title,
		Content:  postModel.Content,
		AuthorID: postModel.AuthorID,
		Active:   postModel.Active,
	}, nil
}

func (p *Post) GetAll(ctx context.Context) ([]*post.Post, error) {
	var postModels []postgresModel.Post
	err := p.dbFromContext(ctx).Find(&postModels).Error
	if err != nil {
		return nil, err
	}
	posts := make([]*post.Post, len(postModels))
	for i, m := range postModels {
		posts[i] = &post.Post{
			ID:       m.ID,
			Title:    m.Title,
			Content:  m.Content,
			AuthorID: m.AuthorID,
			Active:   m.Active,
		}
	}
	return posts, nil
}

func (p *Post) Create(ctx context.Context, post *post.Post) error {
	newPost := postgresModel.Post{
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}
	err := p.dbFromContext(ctx).Create(&newPost).Error
	if err != nil {
		return err
	}
	post.ID = newPost.ID
	return nil
}

func (p *Post) Update(ctx context.Context, post *post.Post) error {
	return p.dbFromContext(ctx).Model(&postgresModel.Post{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"title":     post.Title,
		"content":   post.Content,
		"author_id": post.AuthorID,
		"active":    post.Active,
	}).Error
}

func (p *Post) Delete(ctx context.Context, id int64) error {
	return p.dbFromContext(ctx).Delete(&postgresModel.Post{}, id).Error
}

func (p *Post) DeactivateByAuthorID(ctx context.Context, authorID int64) error {
	return p.dbFromContext(ctx).Model(&postgresModel.Post{}).Where("author_id = ?", authorID).Update("active", false).Error
}

func (p *Post) ActivateByAuthorID(ctx context.Context, authorID int64) error {
	return p.dbFromContext(ctx).Model(&postgresModel.Post{}).Where("author_id = ?", authorID).Update("active", true).Error
}
