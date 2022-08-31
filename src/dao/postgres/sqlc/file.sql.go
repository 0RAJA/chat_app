// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: file.sql

package db

import (
	"context"
	"database/sql"
)

const createFile = `-- name: CreateFile :one
insert into file
    (file_name, file_type, file_size, key, url, relation_id, account_id)
values ($1, $2, $3, $4, $5, $6, $7)
returning id, file_name, file_type, file_size, key, url, relation_id, account_id, create_at
`

type CreateFileParams struct {
	FileName   string        `json:"file_name"`
	FileType   Filetype      `json:"file_type"`
	FileSize   int64         `json:"file_size"`
	Key        string        `json:"key"`
	Url        string        `json:"url"`
	RelationID sql.NullInt64 `json:"relation_id"`
	AccountID  sql.NullInt64 `json:"account_id"`
}

func (q *Queries) CreateFile(ctx context.Context, arg *CreateFileParams) (*File, error) {
	row := q.db.QueryRow(ctx, createFile,
		arg.FileName,
		arg.FileType,
		arg.FileSize,
		arg.Key,
		arg.Url,
		arg.RelationID,
		arg.AccountID,
	)
	var i File
	err := row.Scan(
		&i.ID,
		&i.FileName,
		&i.FileType,
		&i.FileSize,
		&i.Key,
		&i.Url,
		&i.RelationID,
		&i.AccountID,
		&i.CreateAt,
	)
	return &i, err
}

const deleteFileByID = `-- name: DeleteFileByID :exec
delete
from file
where id = $1
`

func (q *Queries) DeleteFileByID(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteFileByID, id)
	return err
}

const getAllRelationsOnFile = `-- name: GetAllRelationsOnFile :many
select relation_id
from file
group by relation_id
`

func (q *Queries) GetAllRelationsOnFile(ctx context.Context) ([]sql.NullInt64, error) {
	rows, err := q.db.Query(ctx, getAllRelationsOnFile)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []sql.NullInt64{}
	for rows.Next() {
		var relation_id sql.NullInt64
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

const getAvatar = `-- name: GetAvatar :one
select exists(select 1 from file where account_id= $1 and file_name= 'AccountAvatar')
`

func (q *Queries) GetAvatar(ctx context.Context, accountID sql.NullInt64) (bool, error) {
	row := q.db.QueryRow(ctx, getAvatar, accountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getFileByRelationID = `-- name: GetFileByRelationID :many
select id, file_name, file_type, file_size, key, url, relation_id, account_id, create_at
from file
where relation_id = $1
`

func (q *Queries) GetFileByRelationID(ctx context.Context, relationID sql.NullInt64) ([]*File, error) {
	rows, err := q.db.Query(ctx, getFileByRelationID, relationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*File{}
	for rows.Next() {
		var i File
		if err := rows.Scan(
			&i.ID,
			&i.FileName,
			&i.FileType,
			&i.FileSize,
			&i.Key,
			&i.Url,
			&i.RelationID,
			&i.AccountID,
			&i.CreateAt,
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

const getFileByRelationIDIsNUll = `-- name: GetFileByRelationIDIsNUll :many
select id, key
from file
where relation_id is null and file_name != 'AccountAvatar'
`

type GetFileByRelationIDIsNUllRow struct {
	ID  int64  `json:"id"`
	Key string `json:"key"`
}

func (q *Queries) GetFileByRelationIDIsNUll(ctx context.Context) ([]*GetFileByRelationIDIsNUllRow, error) {
	rows, err := q.db.Query(ctx, getFileByRelationIDIsNUll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetFileByRelationIDIsNUllRow{}
	for rows.Next() {
		var i GetFileByRelationIDIsNUllRow
		if err := rows.Scan(&i.ID, &i.Key); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFileDetailsByID = `-- name: GetFileDetailsByID :one
select id, file_name, file_type, file_size, key, url, relation_id, account_id, create_at
from file
where id = $1
`

func (q *Queries) GetFileDetailsByID(ctx context.Context, id int64) (*File, error) {
	row := q.db.QueryRow(ctx, getFileDetailsByID, id)
	var i File
	err := row.Scan(
		&i.ID,
		&i.FileName,
		&i.FileType,
		&i.FileSize,
		&i.Key,
		&i.Url,
		&i.RelationID,
		&i.AccountID,
		&i.CreateAt,
	)
	return &i, err
}

const getFileKeyByID = `-- name: GetFileKeyByID :one
select key
from file
where id = $1
`

func (q *Queries) GetFileKeyByID(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRow(ctx, getFileKeyByID, id)
	var key string
	err := row.Scan(&key)
	return key, err
}

const getGroupAvatar = `-- name: GetGroupAvatar :one
select id, file_name, file_type, file_size, key, url, relation_id, account_id, create_at
from file
where relation_id = $1
  and account_id is null
`

func (q *Queries) GetGroupAvatar(ctx context.Context, relationID sql.NullInt64) (*File, error) {
	row := q.db.QueryRow(ctx, getGroupAvatar, relationID)
	var i File
	err := row.Scan(
		&i.ID,
		&i.FileName,
		&i.FileType,
		&i.FileSize,
		&i.Key,
		&i.Url,
		&i.RelationID,
		&i.AccountID,
		&i.CreateAt,
	)
	return &i, err
}

const updateAccountFile = `-- name: UpdateAccountFile :exec
update  file
set  url = $1
where account_id = $2 and file_name = 'AccountAvatar'
`

type UpdateAccountFileParams struct {
	Url       string        `json:"url"`
	AccountID sql.NullInt64 `json:"account_id"`
}

func (q *Queries) UpdateAccountFile(ctx context.Context, arg *UpdateAccountFileParams) error {
	_, err := q.db.Exec(ctx, updateAccountFile, arg.Url, arg.AccountID)
	return err
}

const updateGroupAvatar = `-- name: UpdateGroupAvatar :exec
update file
set url= $1
where relation_id = $2 and file_name = 'groupAvatar'
`

type UpdateGroupAvatarParams struct {
	Url        string        `json:"url"`
	RelationID sql.NullInt64 `json:"relation_id"`
}

func (q *Queries) UpdateGroupAvatar(ctx context.Context, arg *UpdateGroupAvatarParams) error {
	_, err := q.db.Exec(ctx, updateGroupAvatar, arg.Url, arg.RelationID)
	return err
}
