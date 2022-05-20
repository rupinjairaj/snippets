package repository

import "github.com/rupinjairaj/snippet/entity"

type SnippetRepository interface {
	Save(snippet *entity.SnippetClient) (*entity.SnippetFirestore, error)
	FindByTag(tagName string) ([]entity.SnippetFirestore, error)
}
