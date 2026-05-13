package post

type UseCase struct {
	Repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{Repo: repo}
}

func (u *UseCase) Create(title, content string, authorID int) (*Post, error) {
	post := &Post{Title: title, Content: content, AuthorID: authorID}
	err := u.Repo.Create(post)

	return post, err
}

func (u *UseCase) GetByID(id int) (*Post, error) {
	return u.Repo.GetByID(id)
}

func (u *UseCase) GetAll() ([]*Post, error) {
	return u.Repo.GetAll()
}

func (u *UseCase) Update(id int, title, content string) (*Post, error) {
	post, err := u.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	post.Title = title
	post.Content = content

	err = u.Repo.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (u *UseCase) Delete(id int) error {
	return u.Repo.Delete(id)
}
