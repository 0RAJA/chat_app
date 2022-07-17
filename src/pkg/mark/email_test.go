package mark_test

import (
	"log"
	"testing"
	"time"

	"github.com/0RAJA/Rutils/pkg/email"
	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/0RAJA/chat_app/src/global"
	"github.com/0RAJA/chat_app/src/pkg/mark"
	"github.com/stretchr/testify/require"
)

func TestSendEmailCode(t *testing.T) {
	emailConf := global.PvSettings.Email
	conf := mark.Config{
		UserMarkDuration: time.Second * 1,
		CodeMarkDuration: time.Second * 2,
		SMTPInfo: email.SMTPInfo{
			Host:     emailConf.Host,
			Port:     emailConf.Port,
			IsSSL:    emailConf.IsSSL,
			UserName: emailConf.UserName,
			Password: emailConf.Password,
			From:     emailConf.From,
		},
		AppName: global.PbSettings.App.Name,
	}
	var emailMark = mark.New(conf)
	emailNum := "1647193241@qq.com"
	code := utils.RandomString(6)
	// 测试发送
	log.Println("send1")
	err := emailMark.SendMail(emailNum, code)
	require.NoError(t, err)
	log.Println("check1")
	require.True(t, emailMark.CheckCode(emailNum, code))
	// 测试快速请求
	log.Println("send2")
	require.ErrorIs(t, emailMark.SendMail(emailNum, code), mark.ErrSendTooMany)
	// 测试慢请求
	time.Sleep(conf.UserMarkDuration)
	code = utils.RandomString(6)
	log.Println("send2")
	require.NoError(t, emailMark.SendMail(emailNum, code))
	log.Println("check2")
	require.True(t, emailMark.CheckCode(emailNum, code))
	// 测试清除验证码
	time.Sleep(conf.CodeMarkDuration)
	log.Println("check3")
	require.False(t, emailMark.CheckCode(emailNum, code))
}
