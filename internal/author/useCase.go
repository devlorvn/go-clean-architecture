package author

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) Create(name string) (*Author, error) {
	author := &Author{Name: name, Active: true}

	err := u.repo.Create(author)

	return author, err
}

func (u *UseCase) GetByID(id int64) (*Author, error) {
	return u.repo.GetByID(id)
}

func (u *UseCase) GetAll() ([]*Author, error) {
	return u.repo.GetAll()
}

func (u *UseCase) Update(id int64, name string) (*Author, error) {
	author, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	author.Name = name

	err = u.repo.Update(author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (u *UseCase) Delete(id int64) error {
	return u.repo.Delete(id)
}

func (u *UseCase) Deactivate(id int64) (*Author, error) {
	author, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	author.Active = false

	err = u.repo.Update(author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (u *UseCase) Activate(id int64) (*Author, error) {
	author, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	author.Active = true

	err = u.repo.Update(author)
	if err != nil {
		return nil, err
	}
	return author, nil
}
