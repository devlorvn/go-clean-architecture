package repository

import "go-clean-architecture/internal/author"

type AuthorMemory struct {
	data map[int64]*author.Author
	id   int64
}

func NewAuthorMemory() *AuthorMemory {
	return &AuthorMemory{
		data: make(map[int64]*author.Author),
		id:   0,
	}
}

func (r *AuthorMemory) GetByID(id int64) (*author.Author, error) {
	if author, ok := r.data[id]; ok {
		return author, nil
	}
	return nil, nil
}

func (r *AuthorMemory) GetAll() ([]*author.Author, error) {
	authors := make([]*author.Author, 0, len(r.data))
	for _, author := range r.data {
		authors = append(authors, author)
	}
	return authors, nil
}

func (r *AuthorMemory) Create(author *author.Author) error {
	r.id++
	author.ID = r.id
	r.data[author.ID] = author
	return nil
}

func (r *AuthorMemory) Update(author *author.Author) error {
	if _, ok := r.data[author.ID]; !ok {
		return nil
	}
	r.data[author.ID] = author
	return nil
}

func (r *AuthorMemory) Delete(id int64) error {
	delete(r.data, id)
	return nil
}
