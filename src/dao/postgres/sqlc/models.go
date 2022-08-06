// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

type Applicationstatus string

const (
	ApplicationstatusValue0 Applicationstatus = "已申请"
	ApplicationstatusValue1 Applicationstatus = "已同意"
	ApplicationstatusValue2 Applicationstatus = "已拒绝"
)

func (e *Applicationstatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Applicationstatus(s)
	case string:
		*e = Applicationstatus(s)
	default:
		return fmt.Errorf("unsupported scan type for Applicationstatus: %T", src)
	}
	return nil
}

type Filetype string

const (
	FiletypeImg  Filetype = "img"
	FiletypeFile Filetype = "file"
)

func (e *Filetype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Filetype(s)
	case string:
		*e = Filetype(s)
	default:
		return fmt.Errorf("unsupported scan type for Filetype: %T", src)
	}
	return nil
}

type Gender string

const (
	GenderValue0 Gender = "男"
	GenderValue1 Gender = "女"
	GenderValue2 Gender = "未知"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type Msgnotifytype string

const (
	MsgnotifytypeSystem Msgnotifytype = "system"
	MsgnotifytypeCommon Msgnotifytype = "common"
)

func (e *Msgnotifytype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Msgnotifytype(s)
	case string:
		*e = Msgnotifytype(s)
	default:
		return fmt.Errorf("unsupported scan type for Msgnotifytype: %T", src)
	}
	return nil
}

type Relationtype string

const (
	RelationtypeGroup  Relationtype = "group"
	RelationtypeFriend Relationtype = "friend"
)

func (e *Relationtype) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Relationtype(s)
	case string:
		*e = Relationtype(s)
	default:
		return fmt.Errorf("unsupported scan type for Relationtype: %T", src)
	}
	return nil
}

type Account struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Gender    Gender    `json:"gender"`
	Signature string    `json:"signature"`
	CreateAt  time.Time `json:"create_at"`
}

type Application struct {
	Account1ID int64             `json:"account1_id"`
	Account2ID int64             `json:"account2_id"`
	ApplyMsg   string            `json:"apply_msg"`
	RefuseMsg  string            `json:"refuse_msg"`
	Status     Applicationstatus `json:"status"`
	CreateAt   time.Time         `json:"create_at"`
	UpdateAt   time.Time         `json:"update_at"`
}

type File struct {
	ID         int64         `json:"id"`
	FileName   string        `json:"file_name"`
	FileType   Filetype      `json:"file_type"`
	FileSize   int64         `json:"file_size"`
	Key        string        `json:"key"`
	Url        string        `json:"url"`
	RelationID sql.NullInt64 `json:"relation_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	CreateAt   time.Time     `json:"create_at"`
}

type GroupNotify struct {
	ID            int64         `json:"id"`
	RelationID    sql.NullInt64 `json:"relation_id"`
	MsgContent    string        `json:"msg_content"`
	MsgExpand     pgtype.JSON   `json:"msg_expand"`
	AccountID     sql.NullInt64 `json:"account_id"`
	CreateAt      time.Time     `json:"create_at"`
	ReadIds       []int64       `json:"read_ids"`
}

type Message struct {
	ID            int64         `json:"id"`
	NotifyType    Msgnotifytype `json:"notify_type"`
	MsgType       string        `json:"msg_type"`
	MsgContent    string        `json:"msg_content"`
	MsgExtend     pgtype.JSON   `json:"msg_extend"`
	FileID        sql.NullInt64 `json:"file_id"`
	AccountID     sql.NullInt64 `json:"account_id"`
	RlyMsgID      sql.NullInt64 `json:"rly_msg_id"`
	RelationID    int64         `json:"relation_id"`
	CreateAt      time.Time     `json:"create_at"`
	IsRevoke      bool          `json:"is_revoke"`
	IsTop         bool          `json:"is_top"`
	IsPin         bool          `json:"is_pin"`
	PinTime       time.Time     `json:"pin_time"`
	ReadIds       []int64       `json:"read_ids"`
	MsgContentTsv interface{}   `json:"msg_content_tsv"`
}

type Relation struct {
	ID           int64          `json:"id"`
	RelationType Relationtype   `json:"relation_type"`
	GroupType    sql.NullString `json:"group_type"`
	FriendType   sql.NullString `json:"friend_type"`
	CreateAt     sql.NullTime   `json:"create_at"`
}

type Setting struct {
	AccountID    int64     `json:"account_id"`
	RelationID   int64     `json:"relation_id"`
	NickName     string    `json:"nick_name"`
	IsNotDisturb bool      `json:"is_not_disturb"`
	IsPin        bool      `json:"is_pin"`
	PinTime      time.Time `json:"pin_time"`
	IsShow       bool      `json:"is_show"`
	LastShow     time.Time `json:"last_show"`
	IsLeader     bool      `json:"is_leader"`
	IsSelf       bool      `json:"is_self"`
}

type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at"`
}
