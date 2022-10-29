package repo

import (
	"_/cmd/api/repo/firestore_repo"
	"_/cmd/types/repo_types"
)

// type Deps struct{ DataBase repo_types.IPostRepo }

type RepoLayer struct{ PostRepo repo_types.IPostRepo }

func NewRepoLayer() *RepoLayer {
	return &RepoLayer{PostRepo: firestore_repo.NewPostRepo()}
}
