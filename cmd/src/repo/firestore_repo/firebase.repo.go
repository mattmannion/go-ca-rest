package firestore_repo

import (
	"_/cmd/src/models"
	"_/cmd/src/repo/repo_types"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// Repo constructor
type Repo struct {
	Ctx            context.Context
	Client         firestore.Client
	ProjectName    string
	CollectionName string
	Error          error
}

func NewPostRepo(
	ctx context.Context,
	client firestore.Client,
	err error,
	cn string,
) repo_types.IPostRepo {

	return &Repo{
		Ctx:            ctx,
		Client:         client,
		CollectionName: cn,
		Error:          err,
	}
}

func (r *Repo) Save(post *models.Post) (*models.Post, error) {
	if r.Error != nil {
		log.Fatalln("Failed to create firestore...")
		return &models.Post{}, r.Error
	}

	defer r.Client.Close()

	_, _, err := r.Client.Collection(r.CollectionName).Add(r.Ctx, map[string]interface{}{
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

func (r *Repo) FindAll() ([]models.Post, error) {
	if r.Error != nil {
		log.Fatalln("Failed to create firestore...")
		return nil, r.Error
	}

	defer r.Client.Close()

	var posts []models.Post

	iterator := r.Client.Collection(r.CollectionName).Documents(r.Ctx)

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
