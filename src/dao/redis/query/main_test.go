package query_test

import (
	"os"
	"testing"

	"github.com/0RAJA/chat_app/src/setting"
)

func TestMain(m *testing.M) {
	setting.Group.Config.Init()
	setting.Group.Dao.Init()
	os.Exit(m.Run())
}
