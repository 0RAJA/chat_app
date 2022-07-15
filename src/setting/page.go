package setting

import (
	"github.com/0RAJA/Rutils/pkg/app"
	"github.com/0RAJA/chat_app/src/global"
)

type page struct {
}

func (page) Init() {
	global.Page = app.InitPage(
		global.PbSettings.Page.DefaultPageSize,
		global.PbSettings.Page.MaxPageSize,
		global.PbSettings.Page.PageKey,
		global.PbSettings.Page.PageSizeKey,
	)
}
