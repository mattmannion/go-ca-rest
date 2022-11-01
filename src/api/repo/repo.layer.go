package repo

import (
	"_/src/api/repo/pg_repo"
	"_/src/clients/pg/pg_client"
	"_/src/types/repo_types"
)

type RepoLayer struct{ PostRepo repo_types.IPostRepo }

func NewRepoLayer() *RepoLayer {
	return &RepoLayer{PostRepo: pg_repo.NewPostRepo(pg_client.Db)}
}
