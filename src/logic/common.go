package logic

import (
	"github.com/0RAJA/Rutils/pkg/token"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model"
	"github.com/0RAJA/chat_app/src/pkg/retry"
)

// 尝试重试
func reTry(name string, f func() error) {
	go func() {
		d := global.PbSettings.Auto.Retry.Duration
		times := global.PbSettings.Auto.Retry.MaxTimes
		report := <-retry.NewTry(name, f, d, times).Run()
		global.Logger.Error(report.Error())
	}()
}

// 新建token并返回token，payload, err
func newToken(t model.TokenType, id int64) (string, *token.Payload, error) {
	duration := global.PvSettings.Token.UserTokenDuration
	if t == model.AccountToken {
		duration = global.PvSettings.Token.AccountTokenDuration
	}
	data, err := model.NewTokenContent(t, id).Marshal()
	if err != nil {
		return "", nil, err
	}
	result, payload, err := global.Maker.CreateToken(data, duration)
	if err != nil {
		return "", nil, err
	}
	return result, payload, nil
}
