package logic

import (
	"context"

	"github.com/0RAJA/Rutils/pkg/goroutine/task"
	"github.com/0RAJA/chat_app/src/dao"
	"github.com/0RAJA/chat_app/src/global"
)

type auto struct {
}

func (auto) Work() {
	ctx := context.Background()
	deleteExpiredFileTask := task.Task{
		Name:            "deleteExpiredFile",
		Ctx:             ctx,
		TaskDuration:    global.PbSettings.Auto.DeleteExpiredFileDuration,
		TimeoutDuration: global.PbSettings.Server.DefaultContextTimeout,
		F:               DeleteExpiredFile(),
	}
	startTask(deleteExpiredFileTask)
}

func startTask(tasks ...task.Task) {
	for i := range tasks {
		task.NewTickerTask(tasks[i])
	}
}

// DeleteExpiredFile 定时删除没有relation的文件
func DeleteExpiredFile() task.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("auto task run : deleteExpiredFile")
		ctx, cancel := context.WithTimeout(parentCtx, global.PbSettings.Server.DefaultContextTimeout)
		defer cancel()
		data, err := dao.Group.DB.GetAllRelationOnRelation(ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		relations := make(map[int64]int)
		for _, v := range data {
			relations[v] = 1
		}
		data2, err := dao.Group.DB.GetAllRelationsOnFile(ctx)
		if err != nil {
			global.Logger.Error(err.Error())
			return
		}
		for _, v := range data2 {
			_, ok := relations[v.Int64]
			if !ok {
				//err := dao.Group.DB.DeleteFileByRelationID(ctx)
				//if err != nil {
				//	global.Logger.Error(err.Error())
				//	return
				//}
				d, err := dao.Group.DB.GetFileByRelationIDIsNUll(ctx)
				if err != nil {
					global.Logger.Error(err.Error())
					return
				}
				for _, v := range d {
					_, err := Group.File.DeleteFile(ctx, v.ID)
					if err != nil {
						global.Logger.Error(err.Error())
						return
					}
				}

			}
		}
	}
}
