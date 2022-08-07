package setting

import (
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/dao/postgres"
	"github.com/0RAJA/chat_app/src/dao/redis"
	"github.com/0RAJA/chat_app/src/global"
)

type mDao struct {
}

// Init 持久化层初始化
func (m mDao) Init() {
	dao.Group.DB = postgres.Init(global.PvSettings.Postgresql.SourceName)
	dao.Group.Redis = redis.Init(
		global.PvSettings.Redis.Address,
		global.PvSettings.Redis.Password,
		global.PvSettings.Redis.PoolSize,
		global.PvSettings.Redis.DB,
	)

}
