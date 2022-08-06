package setting

import "github.com/0RAJA/chat_app/src/logic"

type auto struct {
}

func (auto) Init() {
	logic.Group.Auto.Work()
}
