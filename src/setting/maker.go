package setting

import (
	"github.com/0RAJA/Rutils/pkg/token"
	"github.com/0RAJA/chat_app/src/global"
)

type maker struct {
}

// Init tokenMaker初始化
func (maker) Init() {
	var err error
	global.Maker, err = token.NewPasetoMaker([]byte(global.PvSettings.Token.Key))
	if err != nil {
		panic(err)
	}
}
