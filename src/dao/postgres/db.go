package postgres

import (
	"context"

	db "github.com/0RAJA/chat_app/src/dao/postgres/sqlc"
	"github.com/0RAJA/chat_app/src/dao/postgres/tx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	tx.TXer
	db.Querier
}

func Init(dataSourceName string) DB {
	pool, err := pgxpool.Connect(context.Background(), dataSourceName)
	if err != nil {
		panic(err)
	}
	return &tx.SqlStore{Queries: db.New(pool), DB: pool}
}
