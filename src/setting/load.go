package setting

import (
	"context"

	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/pkg/setting"
)

// 所有需要在启动时初始化的配置

type load struct {
}

// LoadAllEmailsToRedis 加载所有邮件到redis
func LoadAllEmailsToRedis() error {
	emails, err := dao.Group.DB.GetAllEmails(context.Background())
	if err != nil {
		return err
	}
	if err := dao.Group.Redis.ReloadEmails(context.Background(), emails...); err != nil {
		return err
	}
	return nil
}

// LoadAllGroupRelationToRedis 加载所有群组关系名单到redis
// :nolint
func LoadAllGroupRelationToRedis() error {
	// 群ID和成员IDs
	var relations map[int64][]int64
	// TODO: 加载所有群ID和成员IDs到redis
	// relations, err := dao.Group.DB.GetAllGroupRelation(context.Background())
	// if err != nil {
	// 	return err
	// }
	if err := dao.Group.Redis.ReloadGroupRelationIDs(context.Background(), relations); err != nil {
		return err
	}
	return nil
}

func (load) Init() {
	var err error
	// 加载所有邮件到redis
	err = setting.DoThat(err, LoadAllEmailsToRedis)
	if err != nil {
		panic(err)
	}
}
