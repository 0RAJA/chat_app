package setting

type group struct {
	Dao    mDao
	Log    log
	Maker  maker
	Page   page
	Worker worker
}

var Group = new(group)

func AllInit() {
	Group.Dao.Init()
	Group.Maker.Init()
	Group.Log.Init()
	Group.Page.Init()
	Group.Worker.Init()
}
