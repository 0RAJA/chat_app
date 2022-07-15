package dao

import (
	"github.com/0RAJA/chat_app/src/dao/postgres"
	"github.com/0RAJA/chat_app/src/dao/redis/query"
)

type group struct {
	DB    postgres.DB
	Redis *query.Queries
}

var Group = new(group)
