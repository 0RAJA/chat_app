// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: setting.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

type CreateManySettingParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
}

const createSetting = `-- name: CreateSetting :exec
insert into setting (account_id, relation_id, nick_name, is_leader, is_self)
values ($1, $2, '', $3, $4)
`

type CreateSettingParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
	IsLeader   bool  `json:"is_leader"`
	IsSelf     bool  `json:"is_self"`
}

func (q *Queries) CreateSetting(ctx context.Context, arg *CreateSettingParams) error {
	_, err := q.db.Exec(ctx, createSetting,
		arg.AccountID,
		arg.RelationID,
		arg.IsLeader,
		arg.IsSelf,
	)
	return err
}

const deleteGroup = `-- name: DeleteGroup :exec
delete
from setting
where relation_id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, relationID int64) error {
	_, err := q.db.Exec(ctx, deleteGroup, relationID)
	return err
}

const deleteSetting = `-- name: DeleteSetting :exec
delete
from setting
where account_id = $1
  and relation_id = $2
`

type DeleteSettingParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
}

func (q *Queries) DeleteSetting(ctx context.Context, arg *DeleteSettingParams) error {
	_, err := q.db.Exec(ctx, deleteSetting, arg.AccountID, arg.RelationID)
	return err
}

const deleteSettingsByAccountID = `-- name: DeleteSettingsByAccountID :many
delete
from setting
where account_id = $1
returning relation_id
`

func (q *Queries) DeleteSettingsByAccountID(ctx context.Context, accountID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, deleteSettingsByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var relation_id int64
		if err := rows.Scan(&relation_id); err != nil {
			return nil, err
		}
		items = append(items, relation_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const existsFriendSetting = `-- name: ExistsFriendSetting :one
select exists(
               select 1
               from setting s,
                    relation r
               where r.relation_type = 'friend'
                 and ((r.friend_type).account1_id = $1::bigint and
                      (r.friend_type).account2_id = $2::bigint)
                 and s.account_id = $1
           )
`

type ExistsFriendSettingParams struct {
	Account1ID int64 `json:"account1_id"`
	Account2ID int64 `json:"account2_id"`
}

func (q *Queries) ExistsFriendSetting(ctx context.Context, arg *ExistsFriendSettingParams) (bool, error) {
	row := q.db.QueryRow(ctx, existsFriendSetting, arg.Account1ID, arg.Account2ID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsGroupLeaderByAccountIDWithLock = `-- name: ExistsGroupLeaderByAccountIDWithLock :one
select exists(select 1 from setting where account_id = $1 and is_leader = true) for update
`

func (q *Queries) ExistsGroupLeaderByAccountIDWithLock(ctx context.Context, accountID int64) (bool, error) {
	row := q.db.QueryRow(ctx, existsGroupLeaderByAccountIDWithLock, accountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsIsLeader = `-- name: ExistsIsLeader :one
select exists(
               select 1
               from setting
               where relation_id = $1
                 and account_id = $2
                 and is_leader is true
           )
`

type ExistsIsLeaderParams struct {
	RelationID int64 `json:"relation_id"`
	AccountID  int64 `json:"account_id"`
}

func (q *Queries) ExistsIsLeader(ctx context.Context, arg *ExistsIsLeaderParams) (bool, error) {
	row := q.db.QueryRow(ctx, existsIsLeader, arg.RelationID, arg.AccountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsSetting = `-- name: ExistsSetting :one
select exists(
               select 1
               from setting
               where account_id = $1
                 and relation_id = $2
           )
`

type ExistsSettingParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
}

func (q *Queries) ExistsSetting(ctx context.Context, arg *ExistsSettingParams) (bool, error) {
	row := q.db.QueryRow(ctx, existsSetting, arg.AccountID, arg.RelationID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAccountIDsByRelationID = `-- name: GetAccountIDsByRelationID :many
select DISTINCT account_id
from setting
where relation_id = $1
`

func (q *Queries) GetAccountIDsByRelationID(ctx context.Context, relationID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, getAccountIDsByRelationID, relationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var account_id int64
		if err := rows.Scan(&account_id); err != nil {
			return nil, err
		}
		items = append(items, account_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFriendPinSettingsOrderByPinTime = `-- name: GetFriendPinSettingsOrderByPinTime :many
select s.relation_id, s.nick_name, s.pin_time,
       a.id     as account_id,
       a.name   as account_name,
       a.avatar as account_avatar
from (select setting.relation_id, setting.nick_name, setting.pin_time
      from setting,
           relation
      where setting.account_id = $1
        and setting.is_pin = true
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select account_id from setting where relation_id = s.relation_id and (account_id != $1 or is_self = true))
order by s.pin_time
`

type GetFriendPinSettingsOrderByPinTimeRow struct {
	RelationID    int64     `json:"relation_id"`
	NickName      string    `json:"nick_name"`
	PinTime       time.Time `json:"pin_time"`
	AccountID     int64     `json:"account_id"`
	AccountName   string    `json:"account_name"`
	AccountAvatar string    `json:"account_avatar"`
}

func (q *Queries) GetFriendPinSettingsOrderByPinTime(ctx context.Context, accountID int64) ([]*GetFriendPinSettingsOrderByPinTimeRow, error) {
	rows, err := q.db.Query(ctx, getFriendPinSettingsOrderByPinTime, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFriendPinSettingsOrderByPinTimeRow{}
	for rows.Next() {
		var i GetFriendPinSettingsOrderByPinTimeRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.PinTime,
			&i.AccountID,
			&i.AccountName,
			&i.AccountAvatar,
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

const getFriendSettingsByName = `-- name: GetFriendSettingsByName :many
select s.relation_id, s.nick_name, s.is_not_disturb, s.is_pin, s.pin_time, s.is_show, s.last_show, s.is_self,
       a.id             as account_id,
       a.name           as account_name,
       a.avatar         as account_avatar,
       count(*) over () as total
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select account_id
              from setting
              where relation_id = s.relation_id
                and (account_id != $1 or is_self = true))
  and ((a.name like (%@name::varchar % || '%')) or (nick_name like ($4::varchar || '%')))
order by s.pin_time
limit $2 offset $3
`

type GetFriendSettingsByNameParams struct {
	AccountID int64  `json:"account_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	Name      string `json:"name"`
}

type GetFriendSettingsByNameRow struct {
	RelationID    int64     `json:"relation_id"`
	NickName      string    `json:"nick_name"`
	IsNotDisturb  bool      `json:"is_not_disturb"`
	IsPin         bool      `json:"is_pin"`
	PinTime       time.Time `json:"pin_time"`
	IsShow        bool      `json:"is_show"`
	LastShow      time.Time `json:"last_show"`
	IsSelf        bool      `json:"is_self"`
	AccountID     int64     `json:"account_id"`
	AccountName   string    `json:"account_name"`
	AccountAvatar string    `json:"account_avatar"`
	Total         int64     `json:"total"`
}

func (q *Queries) GetFriendSettingsByName(ctx context.Context, arg *GetFriendSettingsByNameParams) ([]*GetFriendSettingsByNameRow, error) {
	rows, err := q.db.Query(ctx, getFriendSettingsByName,
		arg.AccountID,
		arg.Limit,
		arg.Offset,
		arg.Name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFriendSettingsByNameRow{}
	for rows.Next() {
		var i GetFriendSettingsByNameRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.IsNotDisturb,
			&i.IsPin,
			&i.PinTime,
			&i.IsShow,
			&i.LastShow,
			&i.IsSelf,
			&i.AccountID,
			&i.AccountName,
			&i.AccountAvatar,
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

const getFriendSettingsOrderByName = `-- name: GetFriendSettingsOrderByName :many
select s.relation_id, s.nick_name, s.is_not_disturb, s.is_pin, s.pin_time, s.is_show, s.last_show, s.is_self,
       a.id     as account_id,
       a.name   as account_name,
       a.avatar as account_avatar
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select account_id from setting where relation_id = s.relation_id and (account_id != $1 or is_self = true))
order by s.pin_time
`

type GetFriendSettingsOrderByNameRow struct {
	RelationID    int64     `json:"relation_id"`
	NickName      string    `json:"nick_name"`
	IsNotDisturb  bool      `json:"is_not_disturb"`
	IsPin         bool      `json:"is_pin"`
	PinTime       time.Time `json:"pin_time"`
	IsShow        bool      `json:"is_show"`
	LastShow      time.Time `json:"last_show"`
	IsSelf        bool      `json:"is_self"`
	AccountID     int64     `json:"account_id"`
	AccountName   string    `json:"account_name"`
	AccountAvatar string    `json:"account_avatar"`
}

func (q *Queries) GetFriendSettingsOrderByName(ctx context.Context, accountID int64) ([]*GetFriendSettingsOrderByNameRow, error) {
	rows, err := q.db.Query(ctx, getFriendSettingsOrderByName, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFriendSettingsOrderByNameRow{}
	for rows.Next() {
		var i GetFriendSettingsOrderByNameRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.IsNotDisturb,
			&i.IsPin,
			&i.PinTime,
			&i.IsShow,
			&i.LastShow,
			&i.IsSelf,
			&i.AccountID,
			&i.AccountName,
			&i.AccountAvatar,
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

const getFriendShowSettingsOrderByShowTime = `-- name: GetFriendShowSettingsOrderByShowTime :many
select s.relation_id, s.nick_name, s.is_not_disturb, s.is_pin, s.pin_time, s.is_show, s.last_show, s.is_self,
       a.id     as account_id,
       a.name   as account_name,
       a.avatar as account_avatar
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.is_show = true
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select account_id from setting where relation_id = s.relation_id and (account_id != $1 or is_self = true))
order by s.pin_time
`

type GetFriendShowSettingsOrderByShowTimeRow struct {
	RelationID    int64     `json:"relation_id"`
	NickName      string    `json:"nick_name"`
	IsNotDisturb  bool      `json:"is_not_disturb"`
	IsPin         bool      `json:"is_pin"`
	PinTime       time.Time `json:"pin_time"`
	IsShow        bool      `json:"is_show"`
	LastShow      time.Time `json:"last_show"`
	IsSelf        bool      `json:"is_self"`
	AccountID     int64     `json:"account_id"`
	AccountName   string    `json:"account_name"`
	AccountAvatar string    `json:"account_avatar"`
}

func (q *Queries) GetFriendShowSettingsOrderByShowTime(ctx context.Context, accountID int64) ([]*GetFriendShowSettingsOrderByShowTimeRow, error) {
	rows, err := q.db.Query(ctx, getFriendShowSettingsOrderByShowTime, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFriendShowSettingsOrderByShowTimeRow{}
	for rows.Next() {
		var i GetFriendShowSettingsOrderByShowTimeRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.IsNotDisturb,
			&i.IsPin,
			&i.PinTime,
			&i.IsShow,
			&i.LastShow,
			&i.IsSelf,
			&i.AccountID,
			&i.AccountName,
			&i.AccountAvatar,
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

const getGroupList = `-- name: GetGroupList :many
select s.relation_id, s.nick_name, s.is_not_disturb, s.is_pin, s.pin_time, s.is_show, s.last_show, s.is_self,
       r.id                       as relation_id,
       (r.group_type).name        as group_name,
       (r.group_type).avatar      as group_avatar,
       (r.group_type).description as description,
       count(*) over ()           as total
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.relation_id = relation.id
        and relation.relation_type = 'group') as s,
     relation r
where r.id = (select s.relation_id
              from setting
              where relation_id = s.relation_id
                and (setting.account_id = $1))
order by s.last_show
`

type GetGroupListRow struct {
	RelationID   int64       `json:"relation_id"`
	NickName     string      `json:"nick_name"`
	IsNotDisturb bool        `json:"is_not_disturb"`
	IsPin        bool        `json:"is_pin"`
	PinTime      time.Time   `json:"pin_time"`
	IsShow       bool        `json:"is_show"`
	LastShow     time.Time   `json:"last_show"`
	IsSelf       bool        `json:"is_self"`
	RelationID_2 int64       `json:"relation_id_2"`
	GroupName    interface{} `json:"group_name"`
	GroupAvatar  interface{} `json:"group_avatar"`
	Description  interface{} `json:"description"`
	Total        int64       `json:"total"`
}

func (q *Queries) GetGroupList(ctx context.Context, accountID int64) ([]*GetGroupListRow, error) {
	rows, err := q.db.Query(ctx, getGroupList, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetGroupListRow{}
	for rows.Next() {
		var i GetGroupListRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.IsNotDisturb,
			&i.IsPin,
			&i.PinTime,
			&i.IsShow,
			&i.LastShow,
			&i.IsSelf,
			&i.RelationID_2,
			&i.GroupName,
			&i.GroupAvatar,
			&i.Description,
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

const getGroupMembers = `-- name: GetGroupMembers :many
select account_id
from setting
where relation_id = $1
`

func (q *Queries) GetGroupMembers(ctx context.Context, relationID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, getGroupMembers, relationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var account_id int64
		if err := rows.Scan(&account_id); err != nil {
			return nil, err
		}
		items = append(items, account_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupPinSettingsOrderByPinTime = `-- name: GetGroupPinSettingsOrderByPinTime :many
select s.relation_id,
       s.nick_name,
       s.pin_time,
       r.id,
       r.group_type
from (select setting.relation_id, setting.nick_name, setting.pin_time
      from setting,
           relation
      where setting.account_id = $1
        and setting.is_pin = true
        and setting.relation_id = relation.id
        and relation.relation_type = 'group') as s,
     relation r
where r.id = (select relation_id from setting where relation_id = s.relation_id and account_id = $1)
order by s.pin_time
`

type GetGroupPinSettingsOrderByPinTimeRow struct {
	RelationID int64          `json:"relation_id"`
	NickName   string         `json:"nick_name"`
	PinTime    time.Time      `json:"pin_time"`
	ID         int64          `json:"id"`
	GroupType  sql.NullString `json:"group_type"`
}

func (q *Queries) GetGroupPinSettingsOrderByPinTime(ctx context.Context, accountID int64) ([]*GetGroupPinSettingsOrderByPinTimeRow, error) {
	rows, err := q.db.Query(ctx, getGroupPinSettingsOrderByPinTime, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetGroupPinSettingsOrderByPinTimeRow{}
	for rows.Next() {
		var i GetGroupPinSettingsOrderByPinTimeRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.PinTime,
			&i.ID,
			&i.GroupType,
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

const getGroupSettingsByName = `-- name: GetGroupSettingsByName :many
select s.relation_id, s.nick_name, s.is_not_disturb, s.is_pin, s.pin_time, s.is_show, s.last_show, s.is_self,
       r.id                       as relation_id,
       (r.group_type).name        as group_name,
       (r.group_type).avatar      as group_avatar,
       (r.group_type).description as description,
       count(*) over ()           as total
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.relation_id = relation.id
        and relation.relation_type = 'group') as s,
     relation r
where r.id = (select s.relation_id
              from setting
              where relation_id = s.relation_id
                and (setting.account_id = $1))
  and (((r.group_type).name like ('%' || $4::varchar || '%')))
order by (r.group_type).name
limit $2 offset $3
`

type GetGroupSettingsByNameParams struct {
	AccountID int64  `json:"account_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	Name      string `json:"name"`
}

type GetGroupSettingsByNameRow struct {
	RelationID   int64       `json:"relation_id"`
	NickName     string      `json:"nick_name"`
	IsNotDisturb bool        `json:"is_not_disturb"`
	IsPin        bool        `json:"is_pin"`
	PinTime      time.Time   `json:"pin_time"`
	IsShow       bool        `json:"is_show"`
	LastShow     time.Time   `json:"last_show"`
	IsSelf       bool        `json:"is_self"`
	RelationID_2 int64       `json:"relation_id_2"`
	GroupName    interface{} `json:"group_name"`
	GroupAvatar  interface{} `json:"group_avatar"`
	Description  interface{} `json:"description"`
	Total        int64       `json:"total"`
}

func (q *Queries) GetGroupSettingsByName(ctx context.Context, arg *GetGroupSettingsByNameParams) ([]*GetGroupSettingsByNameRow, error) {
	rows, err := q.db.Query(ctx, getGroupSettingsByName,
		arg.AccountID,
		arg.Limit,
		arg.Offset,
		arg.Name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetGroupSettingsByNameRow{}
	for rows.Next() {
		var i GetGroupSettingsByNameRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.IsNotDisturb,
			&i.IsPin,
			&i.PinTime,
			&i.IsShow,
			&i.LastShow,
			&i.IsSelf,
			&i.RelationID_2,
			&i.GroupName,
			&i.GroupAvatar,
			&i.Description,
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

const getGroupShowSettingsOrderByShowTime = `-- name: GetGroupShowSettingsOrderByShowTime :many
select s.relation_id, s.nick_name, s.is_not_disturb, s.is_pin, s.pin_time, s.is_show, s.last_show, s.is_self,
       r.id,
       r.group_type
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.is_pin = true
        and setting.relation_id = relation.id
        and relation.relation_type = 'group') as s,
     relation r
where r.id = (select relation_id from setting where relation_id = s.relation_id and account_id = $1)
order by s.last_show desc
`

type GetGroupShowSettingsOrderByShowTimeRow struct {
	RelationID   int64          `json:"relation_id"`
	NickName     string         `json:"nick_name"`
	IsNotDisturb bool           `json:"is_not_disturb"`
	IsPin        bool           `json:"is_pin"`
	PinTime      time.Time      `json:"pin_time"`
	IsShow       bool           `json:"is_show"`
	LastShow     time.Time      `json:"last_show"`
	IsSelf       bool           `json:"is_self"`
	ID           int64          `json:"id"`
	GroupType    sql.NullString `json:"group_type"`
}

func (q *Queries) GetGroupShowSettingsOrderByShowTime(ctx context.Context, accountID int64) ([]*GetGroupShowSettingsOrderByShowTimeRow, error) {
	rows, err := q.db.Query(ctx, getGroupShowSettingsOrderByShowTime, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetGroupShowSettingsOrderByShowTimeRow{}
	for rows.Next() {
		var i GetGroupShowSettingsOrderByShowTimeRow
		if err := rows.Scan(
			&i.RelationID,
			&i.NickName,
			&i.IsNotDisturb,
			&i.IsPin,
			&i.PinTime,
			&i.IsShow,
			&i.LastShow,
			&i.IsSelf,
			&i.ID,
			&i.GroupType,
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

const getSettingByID = `-- name: GetSettingByID :one
select account_id, relation_id, nick_name, is_not_disturb, is_pin, pin_time, is_show, last_show, is_leader, is_self
from setting
where account_id = $1
  and relation_id = $2
`

type GetSettingByIDParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
}

func (q *Queries) GetSettingByID(ctx context.Context, arg *GetSettingByIDParams) (*Setting, error) {
	row := q.db.QueryRow(ctx, getSettingByID, arg.AccountID, arg.RelationID)
	var i Setting
	err := row.Scan(
		&i.AccountID,
		&i.RelationID,
		&i.NickName,
		&i.IsNotDisturb,
		&i.IsPin,
		&i.PinTime,
		&i.IsShow,
		&i.LastShow,
		&i.IsLeader,
		&i.IsSelf,
	)
	return &i, err
}

const transferIsSelfFalse = `-- name: TransferIsSelfFalse :exec
update setting
set is_leader = false
where relation_id = $1
  and account_id = $2
`

type TransferIsSelfFalseParams struct {
	RelationID int64 `json:"relation_id"`
	AccountID  int64 `json:"account_id"`
}

func (q *Queries) TransferIsSelfFalse(ctx context.Context, arg *TransferIsSelfFalseParams) error {
	_, err := q.db.Exec(ctx, transferIsSelfFalse, arg.RelationID, arg.AccountID)
	return err
}

const transferIsSelfTrue = `-- name: TransferIsSelfTrue :exec
update setting
set is_leader = true
where relation_id = $1
  and account_id = $2
`

type TransferIsSelfTrueParams struct {
	RelationID int64 `json:"relation_id"`
	AccountID  int64 `json:"account_id"`
}

func (q *Queries) TransferIsSelfTrue(ctx context.Context, arg *TransferIsSelfTrueParams) error {
	_, err := q.db.Exec(ctx, transferIsSelfTrue, arg.RelationID, arg.AccountID)
	return err
}

const updateSettingDisturb = `-- name: UpdateSettingDisturb :exec
update setting
set is_not_disturb = $3
where account_id = $1
  and relation_id = $2
`

type UpdateSettingDisturbParams struct {
	AccountID    int64 `json:"account_id"`
	RelationID   int64 `json:"relation_id"`
	IsNotDisturb bool  `json:"is_not_disturb"`
}

func (q *Queries) UpdateSettingDisturb(ctx context.Context, arg *UpdateSettingDisturbParams) error {
	_, err := q.db.Exec(ctx, updateSettingDisturb, arg.AccountID, arg.RelationID, arg.IsNotDisturb)
	return err
}

const updateSettingLeader = `-- name: UpdateSettingLeader :exec
update setting
set is_leader = $3
where account_id = $1
  and relation_id = $2
`

type UpdateSettingLeaderParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
	IsLeader   bool  `json:"is_leader"`
}

func (q *Queries) UpdateSettingLeader(ctx context.Context, arg *UpdateSettingLeaderParams) error {
	_, err := q.db.Exec(ctx, updateSettingLeader, arg.AccountID, arg.RelationID, arg.IsLeader)
	return err
}

const updateSettingNickName = `-- name: UpdateSettingNickName :exec
update setting
set nick_name = $3
where account_id = $1
  and relation_id = $2
`

type UpdateSettingNickNameParams struct {
	AccountID  int64  `json:"account_id"`
	RelationID int64  `json:"relation_id"`
	NickName   string `json:"nick_name"`
}

func (q *Queries) UpdateSettingNickName(ctx context.Context, arg *UpdateSettingNickNameParams) error {
	_, err := q.db.Exec(ctx, updateSettingNickName, arg.AccountID, arg.RelationID, arg.NickName)
	return err
}

const updateSettingPin = `-- name: UpdateSettingPin :exec
update setting
set is_pin = $3
where account_id = $1
  and relation_id = $2
`

type UpdateSettingPinParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
	IsPin      bool  `json:"is_pin"`
}

func (q *Queries) UpdateSettingPin(ctx context.Context, arg *UpdateSettingPinParams) error {
	_, err := q.db.Exec(ctx, updateSettingPin, arg.AccountID, arg.RelationID, arg.IsPin)
	return err
}

const updateSettingShow = `-- name: UpdateSettingShow :exec
update setting
set is_show = $3
where account_id = $1
  and relation_id = $2
`

type UpdateSettingShowParams struct {
	AccountID  int64 `json:"account_id"`
	RelationID int64 `json:"relation_id"`
	IsShow     bool  `json:"is_show"`
}

func (q *Queries) UpdateSettingShow(ctx context.Context, arg *UpdateSettingShowParams) error {
	_, err := q.db.Exec(ctx, updateSettingShow, arg.AccountID, arg.RelationID, arg.IsShow)
	return err
}
