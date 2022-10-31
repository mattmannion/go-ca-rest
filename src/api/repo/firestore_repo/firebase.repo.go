package firestore_repo

import (
	"_/src/clients/pg/pg_client"
	"_/src/envs"
	"_/src/models"
	"_/src/types/repo_types"
	"context"
	"fmt"
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

type Thing struct {
	Id   int
	Text string
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

	// postgres

	things := []Thing{}

	rows, err := pg_client.Db.Query(context.Background(), "select * from stuff")
	if err != nil {
		fmt.Printf("err query: %v\n", err)
	}

	defer rows.Close()

	for rows.Next() {
		fmt.Println("here")
		thing := Thing{}
		errl := rows.Scan(&thing.Id, &thing.Text)
		if errl != nil {
			fmt.Printf("err for loop: %v\n", errl)
		}

		things = append(things, thing)
	}

	fmt.Println(things)

	// pg end

	return posts, nil
}
