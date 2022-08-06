package setting

import (
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/manager"
)

type chat struct {
}

func (chat) Init() {
	global.ChatMap = manager.NewChatMap()
}
