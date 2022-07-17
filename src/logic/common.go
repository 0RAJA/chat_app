package logic

import (
	"github.com/0RAJA/chat_app/src/global"
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
