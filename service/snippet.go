package service

import (
	"errors"

	"github.com/rupinjairaj/snippet/entity"
	"github.com/rupinjairaj/snippet/repository"
)

type SnippetService interface {
	Validate(snippet *entity.SnippetClient) error
	Save(snippet *entity.SnippetClient) (*entity.SnippetFirestore, error)
	FindByTag(tagName string) ([]entity.SnippetFirestore, error)
}

type snippetSer struct {
	snippetRepo repository.SnippetRepository
}

func NewSnippetService(repo repository.SnippetRepository) SnippetService {
	return &snippetSer{
		snippetRepo: repo,
	}
}

func (ss *snippetSer) Validate(snippet *entity.SnippetClient) error {
	if snippet.Content == "" {
		return errors.New("Snippet does not contain any Content.")
	}
	if len(snippet.Tags) <= 0 {
		return errors.New("Snippet does not contain any Tags.")
	}
	return nil
}

func (ss *snippetSer) Save(snippet *entity.SnippetClient) (*entity.SnippetFirestore, error) {
	err := ss.Validate(snippet)
	if err != nil {
		return nil, err
	}

	return ss.snippetRepo.Save(snippet)
}

func (ss *snippetSer) FindByTag(tagName string) ([]entity.SnippetFirestore, error) {
	if tagName == "" {
		return nil, errors.New("Tag name not provided.")
	}

	return ss.snippetRepo.FindByTag(tagName)
}
