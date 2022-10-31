package pg_client

import (
	"_/src/envs"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// this error declaration is needed to stop a null pointer exception
	err error
	Db  *pgxpool.Pool
	Orm *gorm.DB
)

func init() {
	// Pg Init
	Db, err = pgxpool.New(context.Background(), envs.PgConn)
	if err != nil {
		fmt.Printf("Db Driver failed to connect. Err: %v\n", err)
	}

	err := Db.Ping(context.Background())
	if err != nil {
		fmt.Printf("Db Driver failed to connect. Err: %v\n", err)
		return
	}
	fmt.Println("Database connected")

	// Gorm Init
	Orm, err = gorm.Open(postgres.Open(envs.PgConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Printf("Orm Driver failed to connect. Err: %v\n", err)
		return
	}

	fmt.Println("Orm connected")
	// err = DB.AutoMigrate(&models.Users{})
	// if err != nil {
	// 	fmt.Printf("Auto Migration Failed. Err: %v\n", err)
	// }

	// DB Dev Reset
}
