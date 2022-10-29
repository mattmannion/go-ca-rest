package firestore_repo

import (
	"_/cmd/envs"
	"_/cmd/models"
	"_/cmd/types/repo_types"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Repo struct{}

func NewPostRepo() repo_types.IPostRepo {
	return &Repo{}
}

func (*Repo) Save(post *models.Post) (*models.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, envs.FirestoreProjectName, option.WithCredentialsFile("./firebase.json"))
	if err != nil {
		log.Fatalln("Failed to create firestore...")
		return &models.Post{}, err
	}

	defer client.Close()

	_, _, err = client.Collection(envs.FirestoreCollectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalln("Failed to create firestore...")
		return nil, err
	}

	return post, nil
}

func (*Repo) FindAll() ([]models.Post, error) {

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, envs.FirestoreProjectName, option.WithCredentialsFile("./firebase.json"))
	if err != nil {
		log.Fatalln("Failed to create firestore...")
		return []models.Post{}, err
	}

	if err != nil {
		log.Fatalln("Failed to create firestore...")
		return nil, err
	}

	defer client.Close()

	var posts []models.Post

	iterator := client.Collection(envs.FirestoreCollectionName).Documents(ctx)

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
