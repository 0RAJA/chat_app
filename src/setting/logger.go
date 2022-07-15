package setting

import (
	"github.com/0RAJA/Rutils/pkg/logger"
	"github.com/0RAJA/chat_app/src/global"
)

type log struct {
}

// Init 日志初始化
func (log) Init() {
	global.Logger = logger.NewLogger(&logger.InitStruct{
		LogSavePath:   global.PbSettings.Log.LogSavePath,
		LogFileExt:    global.PbSettings.Log.LogFileExt,
		MaxSize:       global.PbSettings.Log.MaxSize,
		MaxBackups:    global.PbSettings.Log.MaxBackups,
		MaxAge:        global.PbSettings.Log.MaxAge,
		Compress:      global.PbSettings.Log.Compress,
		LowLevelFile:  global.PbSettings.Log.LowLevelFile,
		HighLevelFile: global.PbSettings.Log.HighLevelFile,
	}, global.PbSettings.Log.Level)
}
