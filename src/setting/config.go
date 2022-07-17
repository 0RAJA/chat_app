package setting

import (
	"flag"
	"strings"

	"github.com/0RAJA/Rutils/pkg/setting"
	"github.com/0RAJA/chat_app/src/global"
	setting2 "github.com/0RAJA/chat_app/src/pkg/setting"
)

// 配置文件绑定到全局结构体上(默认加载)

var (
	configPaths       string // 配置文件路径
	privateConfigName string // private配置文件名
	publicConfigName  string // public配置文件名
	configType        string // 配置文件类型
)

func setupFlag() {
	// 命令行参数绑定
	flag.StringVar(&privateConfigName, "private_name", "private", "private配置文件名")
	flag.StringVar(&publicConfigName, "public_name", "public", "public配置文件名")
	flag.StringVar(&configType, "type", "yml", "配置文件类型")
	flag.StringVar(&configPaths, "path", global.RootDir+"/config/app", "指定要使用的配置文件路径,多个路径用逗号隔开")
	flag.Parse()
}

type config struct {
}

// Init 读取配置文件
func (config) Init() {
	setupFlag()
	var (
		err            error
		publicSetting  *setting.Setting
		privateSetting *setting.Setting
	)
	// 在调用其他组件的Init时，这个init会首先执行并且把配置文件绑定到全局的结构体上
	err = setting2.DoThat(err, func() error {
		publicSetting, err = setting.NewSetting(publicConfigName, configType, strings.Split(configPaths, ",")...) // 引入配置文件路径
		return setting2.DoThat(err, func() error { return publicSetting.BindAll(&global.PbSettings) })
	})
	err = setting2.DoThat(err, func() error {
		privateSetting, err = setting.NewSetting(privateConfigName, configType, strings.Split(configPaths, ",")...) // 引入配置文件路径
		return setting2.DoThat(err, func() error { return privateSetting.BindAll(&global.PvSettings) })
	})
	if err != nil {
		panic("初始化配置文件有误:" + err.Error())
	}
}
