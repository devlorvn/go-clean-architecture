package repository

import (
	"context"
	"go-clean-architecture/internal/post"
)

type PostMemory struct {
	data map[int64]*post.Post
	id   int64
}

func NewPostMemory() *PostMemory {
	return &PostMemory{
		data: make(map[int64]*post.Post),
		id:   0,
	}
}

func (r *PostMemory) GetByID(_ context.Context, id int64) (*post.Post, error) {
	if post, ok := r.data[id]; ok {
		return post, nil
	}
	return nil, nil
}

func (r *PostMemory) GetAll(_ context.Context) ([]*post.Post, error) {
	posts := make([]*post.Post, 0, len(r.data))
	for _, post := range r.data {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostMemory) Create(_ context.Context, post *post.Post) error {
	r.id++
	post.ID = r.id
	r.data[post.ID] = post
	return nil
}

func (r *PostMemory) Update(_ context.Context, post *post.Post) error {
	if _, ok := r.data[post.ID]; !ok {
		return nil
	}
	r.data[post.ID] = post
	return nil
}

func (r *PostMemory) Delete(_ context.Context, id int64) error {
	delete(r.data, id)
	return nil
}

func (r *PostMemory) DeactivateByAuthorID(_ context.Context, authorID int64) error {
	_ = authorID
	return nil
}

func (r *PostMemory) ActivateByAuthorID(_ context.Context, authorID int64) error {
	_ = authorID
	return nil
}
