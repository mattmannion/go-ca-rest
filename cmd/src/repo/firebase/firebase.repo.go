package firebase

import (
	"_/cmd/src/model"
	"_/cmd/src/repo/types"
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// Repo constructor
type Repo struct {
	Ctx    context.Context
	Client firestore.Client
	PN     string
	CN     string
}

func NewFirestorePostRepo(
	ctx context.Context,
	client firestore.Client,
	pn string,
	cn string,
	err error,
) types.IPostRepo {

	if err != nil {
		log.Fatalln("Failed to create firestore...")
	}

	return &Repo{
		Ctx:    ctx,
		Client: client,
		PN:     pn,
		CN:     cn,
	}
}

func (r *Repo) Save(post *model.Post) (*model.Post, error) {
	defer r.Client.Close()

	_, _, err := r.Client.Collection(r.CN).Add(r.Ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalln("Failed to create firestore...")
		return &model.Post{}, err
	}

	return post, nil
}

func (r *Repo) FindAll() ([]model.Post, error) {
	defer r.Client.Close()

	var posts []model.Post

	iterator := r.Client.Collection(r.CN).Documents(r.Ctx)

	for {
		doc, _ := iterator.Next()
		if doc == nil {
			break
		}

		posts = append(posts, model.Post{
			Id:    int(doc.Data()["Id"].(int64)),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		})
	}

	return posts, nil
}
