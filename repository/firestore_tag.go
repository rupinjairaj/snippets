package repository

import (
	"context"
	"errors"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/rupinjairaj/snippet/entity"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type firestoreTagRepo struct {
	uuidGen UUID
}

func NewFirestoreTagRepo(ug UUID) TagRepository {
	return &firestoreTagRepo{
		uuidGen: ug,
	}
}

func (r *firestoreTagRepo) Save(tagName string) (*entity.Tag, error) {

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

	uuid, err := r.uuidGen.GenerateUUID()
	if err != nil {
		log.Printf("Failed to generate UUID.")
		return nil, errors.New("Internal error occurred. Retry again.")
	}

	tag = &entity.Tag{
		Id:    uuid,
		Name:  tagName,
		Count: 0,
	}
	_, err = client.Collection(tagsCollectionName).Doc(tag.Name).Set(ctx, map[string]interface{}{
		"id":    tag.Id,
		"name":  tag.Name,
		"count": tag.Count,
	})
	if err != nil {
		log.Printf("Failed to add a new tag: %v\n", err)
		return nil, err
	}

	return tag, nil
}

func (r *firestoreTagRepo) FindByName(tagName string) (*entity.Tag, error) {

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create a firestore client: %v\n", err)
		return nil, err
	}
	defer client.Close()

	doc, err := client.Collection(tagsCollectionName).Doc(tagName).Get(ctx)
	if err != nil && status.Code(err) == codes.NotFound {
		log.Printf("No document found with '%s' tagName\n", tagName)
		return nil, nil
	}

	tag := entity.Tag{
		Id:    doc.Data()["id"].(string),
		Name:  doc.Data()["name"].(string),
		Count: doc.Data()["count"].(int64),
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

func (r *firestoreTagRepo) UpdateTag(tag *entity.Tag) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Printf("Failed to create a firestore client: %v\n", err)
		return err
	}
	defer client.Close()

	_, err = client.Collection(tagsCollectionName).Doc(tag.Name).Update(ctx, []firestore.Update{
		{
			Path:  "count",
			Value: tag.Count + 1,
		},
	})
	if err != nil {
		log.Printf("Failed to update tag count: %v\n", err)
		return err
	}

	return nil
}
