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

func (load) Init() {
	var err error
	// 加载所有邮件到redis
	err = setting.DoThat(err, LoadAllEmailsToRedis)
	if err != nil {
		panic(err)
	}
}
