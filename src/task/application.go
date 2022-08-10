package task

import (
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/model/chat"
)

func Application(accountID int64) func() {
	return func() {
		global.ChatMap.Send(accountID, chat.ServerApplication)
	}
}
