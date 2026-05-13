package repository

import "go-clean-architecture/internal/post"

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

func (r *PostMemory) GetByID(id int64) (*post.Post, error) {
	if post, ok := r.data[id]; ok {
		return post, nil
	}
	return nil, nil
}

func (r *PostMemory) GetAll() ([]*post.Post, error) {
	posts := make([]*post.Post, 0, len(r.data))
	for _, post := range r.data {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostMemory) Create(post *post.Post) error {
	r.id++
	post.ID = r.id
	r.data[post.ID] = post
	return nil
}

func (r *PostMemory) Update(post *post.Post) error {
	if _, ok := r.data[post.ID]; !ok {
		return nil
	}
	r.data[post.ID] = post
	return nil
}

func (r *PostMemory) Delete(id int64) error {
	delete(r.data, id)
	return nil
}
