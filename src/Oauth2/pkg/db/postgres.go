package db

import (
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	postgresClient   *bun.DB
	initOnce         sync.Once
	initPostgresDone chan int
)

func initPostgres() {
	dsn := "postgres://postgres:password@localhost:5432/myoayth2_db?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	postgresClient = bun.NewDB(sqldb, pgdialect.New())
	close(initPostgresDone)
}

func GetDB() *bun.DB {
	initOnce.Do(func() {
		initPostgresDone = make(chan int)
		initPostgres()
	})
	if postgresClient == nil {
		initPostgresDone = make(chan int)
		initPostgres()
	}

	<-initPostgresDone
	return postgresClient
}
