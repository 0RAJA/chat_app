// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: message.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgtype"
)

const createMsg = `-- name: CreateMsg :one
insert into message (notify_type, msg_type, msg_content, msg_extend, file_id, account_id, rly_msg_id, relation_id)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning id, msg_content, msg_extend, file_id,create_at
`

type CreateMsgParams struct {
	NotifyType Msgnotifytype `json:"notify_type"`
	MsgType    string        `json:"msg_type"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	RlyMsgID   sql.NullInt64 `json:"rly_msg_id"`
	RelationID int64         `json:"relation_id"`
}

type CreateMsgRow struct {
	ID         int64         `json:"id"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	CreateAt   time.Time     `json:"create_at"`
}

func (q *Queries) CreateMsg(ctx context.Context, arg *CreateMsgParams) (*CreateMsgRow, error) {
	row := q.db.QueryRow(ctx, createMsg,
		arg.NotifyType,
		arg.MsgType,
		arg.MsgContent,
		arg.MsgExtend,
		arg.FileID,
		arg.AccountID,
		arg.RlyMsgID,
		arg.RelationID,
	)
	var i CreateMsgRow
	err := row.Scan(
		&i.ID,
		&i.MsgContent,
		&i.MsgExtend,
		&i.FileID,
		&i.CreateAt,
	)
	return &i, err
}

const getMsgByID = `-- name: GetMsgByID :one
select id,
       notify_type,
       msg_type,
       msg_content,
       msg_extend,
       file_id,
       account_id,
       rly_msg_id,
       relation_id,
       create_at,
       is_revoke,
       is_top,
       is_pin,
       pin_time,
       read_ids
from message
where id = $1
limit 1
`

type GetMsgByIDRow struct {
	ID         int64         `json:"id"`
	NotifyType Msgnotifytype `json:"notify_type"`
	MsgType    string        `json:"msg_type"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	RlyMsgID   sql.NullInt64 `json:"rly_msg_id"`
	RelationID int64         `json:"relation_id"`
	CreateAt   time.Time     `json:"create_at"`
	IsRevoke   bool          `json:"is_revoke"`
	IsTop      bool          `json:"is_top"`
	IsPin      bool          `json:"is_pin"`
	PinTime    time.Time     `json:"pin_time"`
	ReadIds    []int64       `json:"read_ids"`
}

func (q *Queries) GetMsgByID(ctx context.Context, id int64) (*GetMsgByIDRow, error) {
	row := q.db.QueryRow(ctx, getMsgByID, id)
	var i GetMsgByIDRow
	err := row.Scan(
		&i.ID,
		&i.NotifyType,
		&i.MsgType,
		&i.MsgContent,
		&i.MsgExtend,
		&i.FileID,
		&i.AccountID,
		&i.RlyMsgID,
		&i.RelationID,
		&i.CreateAt,
		&i.IsRevoke,
		&i.IsTop,
		&i.IsPin,
		&i.PinTime,
		&i.ReadIds,
	)
	return &i, err
}

const getMsgsByRelationIDAndTime = `-- name: GetMsgsByRelationIDAndTime :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.rly_msg_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       count(*) over ()                                                                      as total,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count
from message m1
where m1.relation_id = $1
  and m1.create_at < $2
order by m1.create_at
limit $3 offset $4
`

type GetMsgsByRelationIDAndTimeParams struct {
	RelationID int64     `json:"relation_id"`
	CreateAt   time.Time `json:"create_at"`
	Limit      int32     `json:"limit"`
	Offset     int32     `json:"offset"`
}

type GetMsgsByRelationIDAndTimeRow struct {
	ID         int64         `json:"id"`
	NotifyType Msgnotifytype `json:"notify_type"`
	MsgType    string        `json:"msg_type"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	RlyMsgID   sql.NullInt64 `json:"rly_msg_id"`
	RelationID int64         `json:"relation_id"`
	CreateAt   time.Time     `json:"create_at"`
	IsRevoke   bool          `json:"is_revoke"`
	IsTop      bool          `json:"is_top"`
	IsPin      bool          `json:"is_pin"`
	PinTime    time.Time     `json:"pin_time"`
	ReadIds    []int64       `json:"read_ids"`
	Total      int64         `json:"total"`
	ReplyCount int64         `json:"reply_count"`
}

func (q *Queries) GetMsgsByRelationIDAndTime(ctx context.Context, arg *GetMsgsByRelationIDAndTimeParams) ([]*GetMsgsByRelationIDAndTimeRow, error) {
	rows, err := q.db.Query(ctx, getMsgsByRelationIDAndTime,
		arg.RelationID,
		arg.CreateAt,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetMsgsByRelationIDAndTimeRow{}
	for rows.Next() {
		var i GetMsgsByRelationIDAndTimeRow
		if err := rows.Scan(
			&i.ID,
			&i.NotifyType,
			&i.MsgType,
			&i.MsgContent,
			&i.MsgExtend,
			&i.FileID,
			&i.AccountID,
			&i.RlyMsgID,
			&i.RelationID,
			&i.CreateAt,
			&i.IsRevoke,
			&i.IsTop,
			&i.IsPin,
			&i.PinTime,
			&i.ReadIds,
			&i.Total,
			&i.ReplyCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPinMsgsByRelationID = `-- name: GetPinMsgsByRelationID :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count,
       count(*) over ()                                                                      as total
from message m1
where m1.relation_id = $1
  and m1.is_pin = true
order by m1.pin_time desc
limit $2 offset $3
`

type GetPinMsgsByRelationIDParams struct {
	RelationID int64 `json:"relation_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

type GetPinMsgsByRelationIDRow struct {
	ID         int64         `json:"id"`
	NotifyType Msgnotifytype `json:"notify_type"`
	MsgType    string        `json:"msg_type"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	RelationID int64         `json:"relation_id"`
	CreateAt   time.Time     `json:"create_at"`
	IsRevoke   bool          `json:"is_revoke"`
	IsTop      bool          `json:"is_top"`
	IsPin      bool          `json:"is_pin"`
	PinTime    time.Time     `json:"pin_time"`
	ReadIds    []int64       `json:"read_ids"`
	ReplyCount int64         `json:"reply_count"`
	Total      int64         `json:"total"`
}

func (q *Queries) GetPinMsgsByRelationID(ctx context.Context, arg *GetPinMsgsByRelationIDParams) ([]*GetPinMsgsByRelationIDRow, error) {
	rows, err := q.db.Query(ctx, getPinMsgsByRelationID, arg.RelationID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetPinMsgsByRelationIDRow{}
	for rows.Next() {
		var i GetPinMsgsByRelationIDRow
		if err := rows.Scan(
			&i.ID,
			&i.NotifyType,
			&i.MsgType,
			&i.MsgContent,
			&i.MsgExtend,
			&i.FileID,
			&i.AccountID,
			&i.RelationID,
			&i.CreateAt,
			&i.IsRevoke,
			&i.IsTop,
			&i.IsPin,
			&i.PinTime,
			&i.ReadIds,
			&i.ReplyCount,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRlyMsgsInfoByMsgID = `-- name: GetRlyMsgsInfoByMsgID :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count,
       count(*) over ()                                                                      as total
from message m1
where m1.relation_id = $1
  and m1.rly_msg_id = $4::bigint
order by m1.create_at
limit $2 offset $3
`

type GetRlyMsgsInfoByMsgIDParams struct {
	RelationID int64 `json:"relation_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	RlyMsgID   int64 `json:"rly_msg_id"`
}

type GetRlyMsgsInfoByMsgIDRow struct {
	ID         int64         `json:"id"`
	NotifyType Msgnotifytype `json:"notify_type"`
	MsgType    string        `json:"msg_type"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	RelationID int64         `json:"relation_id"`
	CreateAt   time.Time     `json:"create_at"`
	IsRevoke   bool          `json:"is_revoke"`
	IsTop      bool          `json:"is_top"`
	IsPin      bool          `json:"is_pin"`
	PinTime    time.Time     `json:"pin_time"`
	ReadIds    []int64       `json:"read_ids"`
	ReplyCount int64         `json:"reply_count"`
	Total      int64         `json:"total"`
}

func (q *Queries) GetRlyMsgsInfoByMsgID(ctx context.Context, arg *GetRlyMsgsInfoByMsgIDParams) ([]*GetRlyMsgsInfoByMsgIDRow, error) {
	rows, err := q.db.Query(ctx, getRlyMsgsInfoByMsgID,
		arg.RelationID,
		arg.Limit,
		arg.Offset,
		arg.RlyMsgID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetRlyMsgsInfoByMsgIDRow{}
	for rows.Next() {
		var i GetRlyMsgsInfoByMsgIDRow
		if err := rows.Scan(
			&i.ID,
			&i.NotifyType,
			&i.MsgType,
			&i.MsgContent,
			&i.MsgExtend,
			&i.FileID,
			&i.AccountID,
			&i.RelationID,
			&i.CreateAt,
			&i.IsRevoke,
			&i.IsTop,
			&i.IsPin,
			&i.PinTime,
			&i.ReadIds,
			&i.ReplyCount,
			&i.Total,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTopMsgByRelationID = `-- name: GetTopMsgByRelationID :one
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count,
       count(*) over ()                                                                      as total
from message m1
where m1.relation_id = $1
  and m1.is_top = true
limit 1
`

type GetTopMsgByRelationIDRow struct {
	ID         int64         `json:"id"`
	NotifyType Msgnotifytype `json:"notify_type"`
	MsgType    string        `json:"msg_type"`
	MsgContent string        `json:"msg_content"`
	MsgExtend  pgtype.JSON   `json:"msg_extend"`
	FileID     sql.NullInt64 `json:"file_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
	RelationID int64         `json:"relation_id"`
	CreateAt   time.Time     `json:"create_at"`
	IsRevoke   bool          `json:"is_revoke"`
	IsTop      bool          `json:"is_top"`
	IsPin      bool          `json:"is_pin"`
	PinTime    time.Time     `json:"pin_time"`
	ReadIds    []int64       `json:"read_ids"`
	ReplyCount int64         `json:"reply_count"`
	Total      int64         `json:"total"`
}

func (q *Queries) GetTopMsgByRelationID(ctx context.Context, relationID int64) (*GetTopMsgByRelationIDRow, error) {
	row := q.db.QueryRow(ctx, getTopMsgByRelationID, relationID)
	var i GetTopMsgByRelationIDRow
	err := row.Scan(
		&i.ID,
		&i.NotifyType,
		&i.MsgType,
		&i.MsgContent,
		&i.MsgExtend,
		&i.FileID,
		&i.AccountID,
		&i.RelationID,
		&i.CreateAt,
		&i.IsRevoke,
		&i.IsTop,
		&i.IsPin,
		&i.PinTime,
		&i.ReadIds,
		&i.ReplyCount,
		&i.Total,
	)
	return &i, err
}

const hasReadMsg = `-- name: HasReadMsg :one
select exists(select 1 from message where id = $1 and $2 = ANY (read_ids))
`

type HasReadMsgParams struct {
	MsgID     int64       `json:"msg_id"`
	AccountID interface{} `json:"account_id"`
}

func (q *Queries) HasReadMsg(ctx context.Context, arg *HasReadMsgParams) (bool, error) {
	row := q.db.QueryRow(ctx, hasReadMsg, arg.MsgID, arg.AccountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const updateMsgPin = `-- name: UpdateMsgPin :exec
update message
set is_pin = $2
where id = $1
`

type UpdateMsgPinParams struct {
	ID    int64 `json:"id"`
	IsPin bool  `json:"is_pin"`
}

func (q *Queries) UpdateMsgPin(ctx context.Context, arg *UpdateMsgPinParams) error {
	_, err := q.db.Exec(ctx, updateMsgPin, arg.ID, arg.IsPin)
	return err
}

const updateMsgReads = `-- name: UpdateMsgReads :exec
update message
set read_ids = array_append(read_ids, $2::bigint)
where id = $1
  and $2::bigint != ANY (read_ids)
`

type UpdateMsgReadsParams struct {
	ID        int64 `json:"id"`
	Accountid int64 `json:"accountid"`
}

func (q *Queries) UpdateMsgReads(ctx context.Context, arg *UpdateMsgReadsParams) error {
	_, err := q.db.Exec(ctx, updateMsgReads, arg.ID, arg.Accountid)
	return err
}

const updateMsgRevoke = `-- name: UpdateMsgRevoke :exec
update message
set is_revoke = $2
where id = $1
`

type UpdateMsgRevokeParams struct {
	ID       int64 `json:"id"`
	IsRevoke bool  `json:"is_revoke"`
}

func (q *Queries) UpdateMsgRevoke(ctx context.Context, arg *UpdateMsgRevokeParams) error {
	_, err := q.db.Exec(ctx, updateMsgRevoke, arg.ID, arg.IsRevoke)
	return err
}

const updateMsgTopFalseByMsgID = `-- name: UpdateMsgTopFalseByMsgID :exec
update message
set is_top = false
where id = $1
`

func (q *Queries) UpdateMsgTopFalseByMsgID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, updateMsgTopFalseByMsgID, id)
	return err
}

const updateMsgTopFalseByRelationID = `-- name: UpdateMsgTopFalseByRelationID :exec
update message
set is_top = false
where relation_id = $1
  and is_top = true
`

func (q *Queries) UpdateMsgTopFalseByRelationID(ctx context.Context, relationID int64) error {
	_, err := q.db.Exec(ctx, updateMsgTopFalseByRelationID, relationID)
	return err
}

const updateMsgTopTrueByMsgID = `-- name: UpdateMsgTopTrueByMsgID :exec
update message
set is_top = true
where id = $1
`

func (q *Queries) UpdateMsgTopTrueByMsgID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, updateMsgTopTrueByMsgID, id)
	return err
}
