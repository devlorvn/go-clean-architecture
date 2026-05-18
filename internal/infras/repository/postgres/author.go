package postgres

import (
	"go-clean-architecture/internal/author"
	postgresModel "go-clean-architecture/internal/infras/database/postgres/model"

	"gorm.io/gorm"
)

type Author struct {
	db *gorm.DB
}

func NewAuthor(db *gorm.DB) *Author {
	return &Author{db: db}
}

func (a *Author) Create(author *author.Author) error {
	newAuthor := postgresModel.Author{
		Name: author.Name,
	}
	err := a.db.Create(&newAuthor).Error
	if err != nil {
		return err
	}
	author.ID = newAuthor.ID
	return nil
}

func (a *Author) GetByID(id int64) (*author.Author, error) {
	var authorModel postgresModel.Author
	err := a.db.First(&authorModel, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &author.Author{
		ID:   authorModel.ID,
		Name: authorModel.Name,
	}, nil
}

func (a *Author) GetAll() ([]*author.Author, error) {
	var authorModels []postgresModel.Author
	err := a.db.Find(&authorModels).Error
	if err != nil {
		return nil, err
	}
	authors := make([]*author.Author, len(authorModels))
	for i, m := range authorModels {
		authors[i] = &author.Author{
			ID:   m.ID,
			Name: m.Name,
		}
	}
	return authors, nil
}

func (a *Author) Update(author *author.Author) error {
	return a.db.Model(&postgresModel.Author{}).Where("id = ?", author.ID).Updates(map[string]interface{}{
		"name": author.Name,
	}).Error
}

func (a *Author) Delete(id int64) error {
	return a.db.Delete(&postgresModel.Author{}, id).Error
}
