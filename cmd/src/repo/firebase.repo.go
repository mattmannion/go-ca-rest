package repo

import (
	"_/cmd/src/model"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// repo constructor
type repo struct{}

func NewFirestoreRepo() PostRepo {
	return &repo{}
}

const pn = "go-ca-e59c4"
const cn = "posts"

func (*repo) Save(post *model.Post) (*model.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./firebase.json")
	client, err := firestore.NewClient(ctx, pn, opt)

	if err != nil {
		log.Fatalln("Failed to create firestore...")
		return &model.Post{}, err
	}

	defer client.Close()

	_, _, err = client.Collection(cn).Add(ctx, map[string]interface{}{
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

func (*repo) FindAll() ([]model.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./firebase.json")
	client, err := firestore.NewClient(ctx, pn, opt)

	if err != nil {
		log.Fatalln(err)
		return []model.Post{}, err
	}

	defer client.Close()

	var posts []model.Post

	iterator := client.Collection(cn).Documents(ctx)

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

	return posts, err
}
