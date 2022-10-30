package repo

import (
	"_/src/api/repo/firestore_repo"
	"_/src/types/repo_types"

	"cloud.google.com/go/firestore"
)

// type Deps struct{ DataBase repo_types.IPostRepo }

type RepoLayer struct{ PostRepo repo_types.IPostRepo }

func NewRepoLayer() *RepoLayer {
	return &RepoLayer{PostRepo: firestore_repo.NewPostRepo(firestore.NewClient)}
}
