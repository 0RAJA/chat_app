package setting

type group struct {
	Auto      auto
	Dao       mDao
	Log       log
	Maker     maker
	Page      page
	Worker    worker
	Config    config
	EmailMark emailMark
	Load      load
	GenID     genID
	Chat      chat
}

var Group = new(group)

func AllInit() {
	Group.Config.Init()
	Group.Dao.Init()
	Group.Maker.Init()
	Group.Log.Init()
	Group.Page.Init()
	Group.Worker.Init()
	Group.EmailMark.Init()
	Group.GenID.Init()
	Group.Chat.Init()
	Group.Load.Init()
	Group.Auto.Init()
}
