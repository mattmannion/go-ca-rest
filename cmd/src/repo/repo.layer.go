package repo

import (
	"_/cmd/src/repo/firebase"
	"_/cmd/src/repo/types"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type RepoLayer struct {
	PostRepo types.IPostRepo
}

const pn = "go-ca-e59c4"
const cn = "posts"

var client, err = firestore.NewClient(context.Background(), pn, option.WithCredentialsFile("./firebase.json"))

// This is a collection of all the controllers
func NewRepoLayer() *RepoLayer {
	return &RepoLayer{
		PostRepo: firebase.NewFirestorePostRepo(
			context.Background(),
			*client,
			pn,
			cn,
			err,
		),
	}
}
