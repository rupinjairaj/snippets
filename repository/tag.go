package repository

import "github.com/rupinjairaj/snippet/entity"

type TagRepository interface {
	Save(tagName string) (*entity.Tag, error)
	FindAll() ([]entity.Tag, error)
	FindByName(tagName string) (*entity.Tag, error)
	UpdateTag(tag *entity.Tag) error
}
