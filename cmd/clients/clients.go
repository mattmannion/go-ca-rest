package clients

import (
	"_/cmd/envs"
	"_/cmd/src/repo/firestore_repo"
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	FirestoreClient, FirestoreErr = firestore.NewClient(
		context.Background(),
		envs.FirestoreProjectName,
		option.WithCredentialsFile("./firebase.json"),
	)

	Firestore = firestore_repo.NewPostRepo(
		context.Background(),
		*FirestoreClient,
		FirestoreErr,
		envs.FirestoreCollectionName,
	)
)
