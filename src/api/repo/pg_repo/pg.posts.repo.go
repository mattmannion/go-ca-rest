package pg_repo

import (
	"_/src/clients/pg/pg_sql"
	"_/src/models"
	"_/src/types/repo_types"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Deps struct {
	Db *pgxpool.Pool
}

type PostRepo struct {
	Db *pgxpool.Pool
}

func NewPostRepo(deps Deps) repo_types.IPostRepo {
	return &PostRepo{Db: deps.Db}
}

func (pr *PostRepo) Insert(post *models.Post) (*models.Post, error) {
	pr.Db.QueryRow(
		context.Background(),
		pg_sql.Posts.InsertPost,
		post.Title,
		post.Text,
	).Scan(
		&post.Id,
		&post.Title,
		&post.Text,
	)

	return post, nil
}

func (pr *PostRepo) GetAll() ([]models.Post, error) {
	posts := []models.Post{}

	rows, _ := pr.Db.Query(context.Background(), pg_sql.Posts.GetAll)

	defer rows.Close()

	for rows.Next() {
		post := &models.Post{}

		rows.Scan(&post.Id, &post.Title, &post.Text)

		posts = append(posts, *post)
	}

	return posts, nil
}
