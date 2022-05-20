package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"

	"cloud.google.com/go/firestore"
	"github.com/rupinjairaj/snippet/entity"
	"google.golang.org/api/iterator"
)

type firestoreTagRepo struct{}

func NewFirestoreTagRepo() TagRepository {
	return &firestoreTagRepo{}
}

func (r *firestoreTagRepo) Save(tagName string) (*entity.Tag, error) {

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

	tag, err := r.FindByName(tagName)
	if err != nil {
		log.Printf("Failed to check if tag already exists %v\n", err)
		return nil, err
	}

	if tag != nil && tag.Name == tagName {
		log.Printf("Not inserting tag as tag with name '%v' already exists\n", tagName)
		return tag, nil
	}

	tag = &entity.Tag{
		Id:   fmt.Sprint(rand.Int63()),
		Name: tagName,
	}
	_, _, err = client.Collection(tagsCollectionName).Add(ctx, map[string]interface{}{
		"id":   tag.Id,
		"name": tag.Name,
	})
	if err != nil {
		log.Printf("Failed to add a new tag: %v\n", err)
		return nil, err
	}

	return tag, nil
}

func (r *firestoreTagRepo) FindByName(tagName string) (*entity.Tag, error) {

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

	q := client.Collection(tagsCollectionName).Select("id").Where("name", "==", tagName)
	doc, err := q.Documents(ctx).Next()

	if err != nil && err == iterator.Done {
		log.Printf("tag not found: %v\n", err)
		return nil, nil
	}

	tag := entity.Tag{
		Id:   doc.Data()["id"].(string),
		Name: tagName,
	}

	return &tag, nil
}

func (r *firestoreTagRepo) FindAll() ([]entity.Tag, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create a firestore client: %v\n", err)
		return nil, err
	}
	defer client.Close()

	var tags []entity.Tag

	iter := client.Collection(tagsCollectionName).Documents(ctx)
	for {
		doc, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Printf("Failed to iterate the list of tags: %v\n", err)
			return nil, err
		}

		tag := entity.Tag{
			Id:   doc.Data()["id"].(string),
			Name: doc.Data()["name"].(string),
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
