package mark

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/0RAJA/Rutils/pkg/email"
)

// 邮箱验证码标记

type Mark struct {
	config   Config
	userMark sync.Map // 标记用户
	codeMark sync.Map // 记录code
}

type Config struct {
	UserMarkDuration time.Duration  // 用户标记时长
	CodeMarkDuration time.Duration  // 验证码标记时长
	SMTPInfo         email.SMTPInfo // 邮箱配置
	AppName          string         // 应用名称
}

func New(conf Config) *Mark {
	return &Mark{
		config:   conf,
		userMark: sync.Map{},
		codeMark: sync.Map{},
	}
}

var ErrSendTooMany = errors.New("发送过于频繁")

// CheckUserExist 判断邮箱是否已经被记录
func (m *Mark) CheckUserExist(email string) bool {
	_, ok := m.userMark.Load(email)
	return ok
}

// SendMail 发送验证码
// nolint
func (m *Mark) SendMail(emailStr, code string) error {
	// TODO:测试
	return nil
	// 发送频率限制
	if m.CheckUserExist(emailStr) {
		return ErrSendTooMany
	}
	m.userMark.Store(emailStr, struct{}{})
	sendEmail := email.NewEmail(&m.config.SMTPInfo)
	// 发送邮件
	err := sendEmail.SendMail([]string{emailStr}, fmt.Sprintf("%s:验证码:%s", m.config.AppName, code), `😘`)
	if err != nil {
		// 发送失败删除标记
		m.userMark.Delete(emailStr)
		return err
	}
	// 记录code
	m.codeMark.Store(emailStr, code)
	// 延时删除标记
	m.delMark(emailStr)
	return nil
}

// delMark 延时删除标记
// nolint
func (m *Mark) delMark(emailStr string) {
	time.AfterFunc(m.config.UserMarkDuration, func() { m.userMark.Delete(emailStr) })
	time.AfterFunc(m.config.CodeMarkDuration, func() { m.codeMark.Delete(emailStr) })
}

// CheckCode 校验验证码
// nolint
func (m *Mark) CheckCode(emailStr, code string) bool {
	// TODO:测试
	return true
	myCode, ok := m.codeMark.Load(emailStr)
	ret := ok && code == myCode
	// 验证成功删除标记
	if ret {
		m.codeMark.Delete(emailStr)
	}
	return ret
}
