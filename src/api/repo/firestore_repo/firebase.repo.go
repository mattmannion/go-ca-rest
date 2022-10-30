package firestore_repo

import (
	"_/src/envs"
	"_/src/models"
	"_/src/types/repo_types"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type NewFSC func(ctx context.Context, projectID string, opts ...option.ClientOption) (*firestore.Client, error)

type PostRepo struct {
	NewClient NewFSC
}

func NewPostRepo(client NewFSC) repo_types.IPostRepo {
	return &PostRepo{NewClient: client}
}

func (pr *PostRepo) Save(post *models.Post) (*models.Post, error) {
	ctx := context.Background()

	client, err := pr.NewClient(ctx, envs.Cfg.FirestoreProjectName, option.WithCredentialsFile(envs.Cfg.FirestoreJson))
	if err != nil {
		log.Fatalln("Failed to create firestore client")
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(envs.Cfg.FirestoreCollectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalln("Failed to connect to firestore collection")
		return nil, err
	}

	return post, nil
}

func (pr *PostRepo) FindAll() ([]models.Post, error) {

	ctx := context.Background()

	client, err := pr.NewClient(ctx, envs.Cfg.FirestoreProjectName, option.WithCredentialsFile(envs.Cfg.FirestoreJson))
	if err != nil {
		log.Fatalln("Failed to create firestore client")
		return []models.Post{}, err
	}

	defer client.Close()

	var posts []models.Post

	iterator := client.Collection(envs.Cfg.FirestoreCollectionName).Documents(ctx)

	for {
		doc, _ := iterator.Next()
		if doc == nil {
			break
		}

		posts = append(posts, models.Post{
			Id:    int(doc.Data()["Id"].(int64)),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		})
	}

	return posts, nil
}
