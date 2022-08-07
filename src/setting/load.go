package setting

import (
	"context"

	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/pkg/tool"
)

// 所有需要在启动时初始化的配置

type load struct {
}

// LoadAllEmailsToRedis 加载所有邮箱到redis
func LoadAllEmailsToRedis() error {
	emails, err := dao.Group.DB.GetAllEmails(context.Background())
	if err != nil {
		return err
	}
	if err := dao.Group.Redis.ReloadEmails(context.Background(), emails...); err != nil {
		return err
	}
	global.Logger.Info("邮箱加载完成")
	return nil
}

// LoadAllGroupRelationToRedis 加载所有关系名单到redis
func LoadAllGroupRelationToRedis() error {
	// 群ID和成员IDs
	relations := make(map[int64][]int64)
	relationIDs, err := dao.Group.DB.GetAllRelationIDs(context.Background())
	if err != nil {
		return err
	}
	for _, relationID := range relationIDs {
		accountIDs, err := dao.Group.DB.GetAccountIDsByRelationID(context.Background(), relationID)
		if err != nil {
			return err
		}
		relations[relationID] = accountIDs
	}
	if err := dao.Group.Redis.ReloadRelationIDs(context.Background(), relations); err != nil {
		return err
	}
	global.Logger.Info("关系名单加载完成")
	return nil
}

func (load) Init() {
	var err error
	// 加载所有邮件到redis
	err = tool.DoThat(err, LoadAllEmailsToRedis)
	err = tool.DoThat(err, LoadAllGroupRelationToRedis)
	if err != nil {
		panic(err)
	}
}
