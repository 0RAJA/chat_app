// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const countAccountByUserID = `-- name: CountAccountByUserID :one
select count(id)::int
from account
where user_id = $1
`

func (q *Queries) CountAccountByUserID(ctx context.Context, userID int64) (int32, error) {
	row := q.db.QueryRow(ctx, countAccountByUserID, userID)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const createAccount = `-- name: CreateAccount :exec
insert into account (id, user_id, name, avatar)
values ($1, $2, $3, $4)
`

type CreateAccountParams struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg *CreateAccountParams) error {
	_, err := q.db.Exec(ctx, createAccount,
		arg.ID,
		arg.UserID,
		arg.Name,
		arg.Avatar,
	)
	return err
}

const deleteAccount = `-- name: DeleteAccount :exec
delete
from account
where id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const deleteAccountsByUserID = `-- name: DeleteAccountsByUserID :many
delete
from account
where user_id = $1
returning id
`

func (q *Queries) DeleteAccountsByUserID(ctx context.Context, userID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, deleteAccountsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const existsAccountByID = `-- name: ExistsAccountByID :one
select exists(
               select 1
               from account
               where id = $1
           )
`

func (q *Queries) ExistsAccountByID(ctx context.Context, id int64) (bool, error) {
	row := q.db.QueryRow(ctx, existsAccountByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsAccountByNameAndUserID = `-- name: ExistsAccountByNameAndUserID :one
select exists(
               select 1
               from account
               where user_id = $1
                 and name = $2
           )
`

type ExistsAccountByNameAndUserIDParams struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}

func (q *Queries) ExistsAccountByNameAndUserID(ctx context.Context, arg *ExistsAccountByNameAndUserIDParams) (bool, error) {
	row := q.db.QueryRow(ctx, existsAccountByNameAndUserID, arg.UserID, arg.Name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAccountByID = `-- name: GetAccountByID :one
select id, user_id, name, avatar, gender, signature, create_at
from account
where id = $1
limit 1
`

func (q *Queries) GetAccountByID(ctx context.Context, id int64) (*Account, error) {
	row := q.db.QueryRow(ctx, getAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Avatar,
		&i.Gender,
		&i.Signature,
		&i.CreateAt,
	)
	return &i, err
}

const getAccountsByName = `-- name: GetAccountsByName :many
select a.id, a.name, a.avatar, r.id as relation_id, count(*) over () as total
from (select id, name, avatar from account where name like ($3::varchar || '%')) as a
         left join relation r on (r.relation_type = 'friend' and
                                  (((r.friend_type).account1_id = a.id and
                                    (r.friend_type).account2_id =@ account_id::bigint) or
                                   (r.friend_type).account2_id = a.id and
                                   (r.friend_type).account1_id = $4::bigint))
limit $1 offset $2
`

type GetAccountsByNameParams struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	Name      string `json:"name"`
	AccountID int64  `json:"account_id"`
}

type GetAccountsByNameRow struct {
	ID         int64         `json:"id"`
	Name       string        `json:"name"`
	Avatar     string        `json:"avatar"`
	RelationID sql.NullInt64 `json:"relation_id"`
	Total      int64         `json:"total"`
}

func (q *Queries) GetAccountsByName(ctx context.Context, arg *GetAccountsByNameParams) ([]*GetAccountsByNameRow, error) {
	rows, err := q.db.Query(ctx, getAccountsByName,
		arg.Limit,
		arg.Offset,
		arg.Name,
		arg.AccountID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetAccountsByNameRow{}
	for rows.Next() {
		var i GetAccountsByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Avatar,
			&i.RelationID,
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

const getAccountsByUserID = `-- name: GetAccountsByUserID :many
select id, name, avatar
from account
where user_id = $1
`

type GetAccountsByUserIDRow struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func (q *Queries) GetAccountsByUserID(ctx context.Context, userID int64) ([]*GetAccountsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, getAccountsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*GetAccountsByUserIDRow{}
	for rows.Next() {
		var i GetAccountsByUserIDRow
		if err := rows.Scan(&i.ID, &i.Name, &i.Avatar); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :exec
update account
set name      = $1,
    avatar    = $2,
    gender    = $3,
    signature = $4
where id = $5
`

type UpdateAccountParams struct {
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Gender    Gender `json:"gender"`
	Signature string `json:"signature"`
	ID        int64  `json:"id"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg *UpdateAccountParams) error {
	_, err := q.db.Exec(ctx, updateAccount,
		arg.Name,
		arg.Avatar,
		arg.Gender,
		arg.Signature,
		arg.ID,
	)
	return err
}
