package setting

import (
	"github.com/0RAJA/Rutils/pkg/goroutine/work"
	"github.com/0RAJA/chat_app/src/global"
)

type worker struct {
}

func (worker) Init() {
	global.Worker = work.Init(work.Config{
		TaskChanCapacity:   global.PbSettings.Worker.TaskChanCapacity,
		WorkerChanCapacity: global.PbSettings.Worker.WorkerChanCapacity,
		WorkerNum:          global.PbSettings.Worker.WorkerNum,
	})
}
