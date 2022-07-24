package repository

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/rupinjairaj/snippet/entity"
	"google.golang.org/api/iterator"
)

type firestoreSnippetRepo struct {
	tagRepo TagRepository
	uuidGen UUID
}

func NewFirestoreSnippetRepo(firestoreTagRepo TagRepository, ug UUID) SnippetRepository {

	return &firestoreSnippetRepo{
		tagRepo: firestoreTagRepo,
		uuidGen: ug,
	}
}

func (r *firestoreSnippetRepo) Save(snippet *entity.SnippetClient) (*entity.SnippetFirestore, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create a firestore client: %v\n", err)
		return nil, err
	}
	defer client.Close()

	var tagIds []string
	var tags []entity.Tag

	for _, tagName := range snippet.Tags {
		tag, err := r.tagRepo.Save(tagName)
		if err != nil {
			log.Printf("Error occurred trying to save tag")
			return nil, err
		}

		tags = append(tags, *tag)
		tagIds = append(tagIds, tag.Id)
	}

	id, err := r.uuidGen.GenerateUUID()
	if err != nil {
		log.Printf("Failed to generate UUID.")
		return nil, errors.New("Internal error occurred. Retry again.")
	}

	newSnippet := &entity.SnippetFirestore{
		Id:      id,
		Name:    snippet.Name,
		TagIds:  tagIds,
		Content: snippet.Content,
	}

	_, err = client.Collection(snippetCollectionName).Doc(newSnippet.Name).Set(ctx, map[string]interface{}{
		"id":      newSnippet.Id,
		"name":    newSnippet.Name,
		"tagIds":  newSnippet.TagIds,
		"content": newSnippet.Content,
	})

	if err != nil {
		log.Printf("Failed to insert the snippet: %v\n", err)
		return nil, err
	}

	for _, tag := range tags {
		err = r.tagRepo.UpdateTag(&tag)
		if err != nil {
			log.Printf("Failed to update count for tag %s: %v\n", tag.Name, err)
		}
	}

	return newSnippet, nil
}

func (r *firestoreSnippetRepo) FindByTag(tagName string) ([]entity.SnippetFirestore, error) {
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

	q := client.Collection(snippetCollectionName).Where("tagIds", "array-contains", tag.Name)
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
			Id:      doc.Data()["id"].(string),
			Name:    doc.Data()["name"].(string),
			Content: doc.Data()["content"].(string),
			TagIds:  nil,
		}

		snippets = append(snippets, snippet)
	}

	return snippets, nil
}
