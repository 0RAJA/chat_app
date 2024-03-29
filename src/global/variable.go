package global

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/Rutils/pkg/createid/snowflake"
	"github.com/0RAJA/Rutils/pkg/goroutine/work"
	"github.com/0RAJA/Rutils/pkg/logger"
	"github.com/0RAJA/Rutils/pkg/token"
	upload "github.com/0RAJA/Rutils/pkg/upload/oss"
	"github.com/0RAJA/chat_app/src/manager"
	"github.com/0RAJA/chat_app/src/model/config"
	"github.com/0RAJA/chat_app/src/pkg/mark"
)

var (
	Logger     *logger.Log          // 日志
	PbSettings config.Public        // Public配置
	PvSettings config.Private       // Private配置
	Page       *app.Page            // 分页
	Worker     *work.Worker         // 工作池
	Maker      token.Maker          // token
	EmailMark  *mark.Mark           // 邮箱标记
	GenID      *snowflake.Snowflake // snowflake生成id
	ChatMap    *manager.ChatMap     // 聊天链接管理器
	OSS        upload.OSS
)
