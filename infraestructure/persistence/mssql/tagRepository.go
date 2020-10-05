package persistence

import (
	"example.com/mittaus/blog/domain"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) domain.ITagRepository {
	return &TagRepository{db}
}

func (repo TagRepository) Search(name string) ([]*domain.Tag, error) {
	var tags []*domain.Tag
	err := repo.db.Where(&domain.Tag{Name: name}).Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func (repo TagRepository) Get(id int) (*domain.Tag, error) {
	var tag domain.Tag
	err := repo.db.Find(&tag, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &tag, nil
}
