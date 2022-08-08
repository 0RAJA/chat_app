package global

import (
	"context"
)

// DefaultContextWithTimeOut 获取默认限时连接上下文
// 成功: 返回上下文和终止函数
func DefaultContextWithTimeOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), PbSettings.Server.DefaultContextTimeout)
}
