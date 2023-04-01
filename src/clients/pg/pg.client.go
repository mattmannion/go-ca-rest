package pg

import (
	"_/src/clients/pg/pg_sql"
	"_/src/envs"
	"_/src/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	err error
	Db  *gorm.DB
)

func init() {
	// Orm Init
	Db, err = gorm.Open(postgres.Open(envs.PgConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf("Orm Driver failed to connect. Err: %v\n", err)
		return
	}

	fmt.Println("Database connected")

	// Orm Sync
	err = Db.AutoMigrate(&models.Post{})
	if err != nil {
		fmt.Printf("Auto Migration Failed. Err: %v\n", err)
		return
	}

	fmt.Println("Database Schema Syncohronized")

	// DB Seeding for Development and Testing
	if envs.Cfg.Env == envs.Dev {
		ResetAndSeedPgDb()

		fmt.Println("Database Reset and Seeded")
	}
}

func ResetAndSeedPgDb() {
	// Seed Posts
	Db.Exec(pg_sql.SeedPosts.TruncatePosts)
	Db.Exec(pg_sql.SeedPosts.ResetPostsId)
	Db.Exec(pg_sql.SeedPosts.InsertPosts)
}
