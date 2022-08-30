package manager

import (
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

func NewChatMap() *ChatMap {
	return &ChatMap{
		m: sync.Map{},
	}
}

type ChatMap struct {
	m   sync.Map // k: accountID v: ConnMap
	sID sync.Map // k: sID v: accountID
}

type ConnMap struct {
	m sync.Map // k: sID v: socketio.Conn
}

// Link 添加设备
func (c *ChatMap) Link(s socketio.Conn, accountID int64) {
	c.m.Store(s.ID(), accountID) // 存入SID和accountID对应关系
	cm, ok := c.m.Load(accountID)
	if !ok {
		cm := &ConnMap{}
		cm.m.Store(s.ID(), s)
		c.m.Store(accountID, cm)
		return
	}
	cm.(*ConnMap).m.Store(s.ID(), s)
}

// Leave 去除设备
func (c *ChatMap) Leave(s socketio.Conn) {
	accountID, ok := c.sID.LoadAndDelete(s.ID())
	if !ok {
		return
	}
	cm, ok := c.m.Load(accountID)
	if !ok {
		return
	}
	cm.(*ConnMap).m.Delete(s.ID())
	length := 0
	cm.(*ConnMap).m.Range(func(k, v interface{}) bool {
		length++
		return true
	})
	if length == 0 {
		c.m.Delete(accountID)
	}
}

// Send 给指定账号的全部设备推送消息
func (c *ChatMap) Send(accountID int64, event string, args ...interface{}) {
	cm, ok := c.m.Load(accountID)
	if !ok {
		return
	}
	cm.(*ConnMap).m.Range(func(k, v interface{}) bool {
		v.(socketio.Conn).Emit(event, args...)
		return true
	})
}

// SendMany 给指定多个账号的全部设备推送消息
func (c *ChatMap) SendMany(accountIDs []int64, event string, args ...interface{}) {
	for accountID := range accountIDs {
		cm, ok := c.m.Load(accountID)
		if !ok {
			return
		}
		cm.(*ConnMap).m.Range(func(k, v interface{}) bool {
			v.(socketio.Conn).Emit(event, args...)
			return true
		})
	}
}

// SendAll 给全部设备推送消息
func (c *ChatMap) SendAll(event string, args ...interface{}) {
	c.m.Range(func(k, v interface{}) bool {
		v.(*ConnMap).m.Range(func(k, v interface{}) bool {
			v.(socketio.Conn).Emit(event, args...)
			return true
		})
		return true
	})
}

type EachFunc socketio.EachFunc

// ForEach 遍历指定账号的全部设备
func (c *ChatMap) ForEach(accountID int64, f EachFunc) {
	cm, ok := c.m.Load(accountID)
	if !ok {
		return
	}
	cm.(*ConnMap).m.Range(func(k, v interface{}) bool {
		f(v.(socketio.Conn))
		return true
	})
}
