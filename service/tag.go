package service

import (
	"errors"

	"github.com/rupinjairaj/snippet/entity"
	"github.com/rupinjairaj/snippet/repository"
)

type TagService interface {
	Validate(tagName string) error
	Create(tagName string) (*entity.Tag, error)
	FindAll() ([]entity.Tag, error)
	FindByName(tagName string) (*entity.Tag, error)
}

type tagSer struct {
	tagRepo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagSer{
		tagRepo: repo,
	}
}

func (tagSer *tagSer) Validate(tagName string) error {
	if tagName == "" {
		return errors.New("Tag name is an empty string.")
	}

	return nil
}

func (tagSer *tagSer) Create(tagName string) (*entity.Tag, error) {
	err := tagSer.Validate(tagName)
	if err != nil {
		return nil, err
	}

	return tagSer.tagRepo.Save(tagName)
}

func (tagSer *tagSer) FindAll() ([]entity.Tag, error) {
	return tagSer.tagRepo.FindAll()
}

func (tagSer *tagSer) FindByName(tagName string) (*entity.Tag, error) {
	err := tagSer.Validate(tagName)
	if err != nil {
		return nil, err
	}

	return tagSer.FindByName(tagName)
}
