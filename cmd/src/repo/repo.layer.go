package repo

import (
	"_/cmd/src/repo/repo_types"
)

type Deps struct{ DataBase repo_types.IPostRepo }

type RepoLayer struct{ PostRepo repo_types.IPostRepo }

func NewRepoLayer(deps Deps) *RepoLayer {
	return &RepoLayer{PostRepo: deps.DataBase}
}
