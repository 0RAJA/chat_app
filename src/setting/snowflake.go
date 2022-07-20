package setting

import (
	"time"

	"github.com/0RAJA/Rutils/pkg/createid/snowflake"
	"github.com/0RAJA/chat_app/src/global"
)

type genID struct {
}

func (genID) Init() {
	var err error
	global.GenID, err = snowflake.Init(time.Now(), global.PbSettings.App.MachineID)
	if err != nil {
		panic(err)
	}
}
