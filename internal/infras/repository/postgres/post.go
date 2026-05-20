package postgres

import (
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

func (p *Post) GetByID(id int64) (*post.Post, error) {
	var postModel postgresModel.Post

	err := p.db.First(&postModel, id).Error
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
	}, nil
}

func (p *Post) GetAll() ([]*post.Post, error) {
	var postModels []postgresModel.Post
	err := p.db.Find(&postModels).Error
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
		}
	}
	return posts, nil
}

func (p *Post) Create(post *post.Post) error {
	newPost := postgresModel.Post{
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}
	err := p.db.Create(&newPost).Error
	if err != nil {
		return err
	}
	post.ID = newPost.ID
	return nil
}

func (p *Post) Update(post *post.Post) error {
	return p.db.Model(&postgresModel.Post{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"title":     post.Title,
		"content":   post.Content,
		"author_id": post.AuthorID,
	}).Error
}

func (p *Post) Delete(id int64) error {
	return p.db.Delete(&postgresModel.Post{}, id).Error
}
