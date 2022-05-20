package repository

import (
	"context"
	"errors"
	"log"
	"math/rand"

	"cloud.google.com/go/firestore"
	"github.com/rupinjairaj/snippet/entity"
	"google.golang.org/api/iterator"
)

type firestoreSnippetRepo struct {
	tagRepo TagRepository
}

func NewFirestoreSnippetRepo(firestoreTagRepo TagRepository) SnippetRepository {

	return &firestoreSnippetRepo{
		tagRepo: firestoreTagRepo,
	}
}

func (r *firestoreSnippetRepo) Validate(snippet *entity.SnippetClient) error {

	if snippet.Content == "" {
		return errors.New("Snippet does not contain any Content.")
	}
	if len(snippet.Tags) <= 0 {
		return errors.New("Snippet does not contain any Tags.")
	}
	return nil
}

func (r *firestoreSnippetRepo) Save(snippet *entity.SnippetClient) (*entity.SnippetFirestore, error) {

	// validate incoming snippet
	err := r.Validate(snippet)
	if err != nil {
		log.Printf("Snippet validation failed: %v", err)
		return nil, err
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create a firestore client: %v\n", err)
		return nil, err
	}
	defer client.Close()

	var tagIds []string

	for _, tagName := range snippet.Tags {
		tag, err := r.tagRepo.Save(tagName)
		if err != nil {
			log.Printf("Error occurred trying to save tag")
			return nil, err
		}

		tagIds = append(tagIds, tag.Id)
	}

	newSnippet := &entity.SnippetFirestore{
		Id:      rand.Int63(),
		TagIds:  tagIds,
		Content: snippet.Content,
	}

	_, _, err = client.Collection(snippetCollectionName).Add(ctx, map[string]interface{}{
		"id":      newSnippet.Id,
		"tagIds":  newSnippet.TagIds,
		"content": newSnippet.Content,
	})
	if err != nil {
		log.Printf("Failed to add a new tag: %v\n", err)
		return nil, err
	}

	return newSnippet, nil
}

func (r *firestoreSnippetRepo) FindByTag(tagName string) ([]entity.SnippetFirestore, error) {

	if tagName == "" {
		return nil, errors.New("Tag name not provided.")
	}

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create a firestore client: %v\n", err)
		return nil, err
	}
	defer client.Close()

	var snippets []entity.SnippetFirestore

	tag, err := r.tagRepo.FindByName(tagName)
	if err != nil {
		log.Printf("Failed to find tag: %v\n", err)
		return nil, err
	}

	q := client.Collection(snippetCollectionName).Where("tagIds", "array-contains", tag.Id)
	iter := q.Documents(ctx)

	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			log.Printf("Iterator is done!: %v", err)
			break
		}

		if err != nil {
			log.Printf("Failed to iterate the list of tags: %v\n", err)
			return nil, err
		}

		snippet := entity.SnippetFirestore{
			Id:      doc.Data()["id"].(int64),
			Content: doc.Data()["content"].(string),
			TagIds:  nil,
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}
