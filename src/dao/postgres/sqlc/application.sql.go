// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: application.sql

package db

import (
	"context"
	"database/sql"
)

const createApplication = `-- name: CreateApplication :exec
insert into application (account1_id, account2_id, apply_msg)
values (?, ?, ?)
`

type CreateApplicationParams struct {
	Account1ID int64          `json:"account1_id"`
	Account2ID int64          `json:"account2_id"`
	ApplyMsg   sql.NullString `json:"apply_msg"`
}

func (q *Queries) CreateApplication(ctx context.Context, arg *CreateApplicationParams) error {
	_, err := q.db.Exec(ctx, createApplication, arg.Account1ID, arg.Account2ID, arg.ApplyMsg)
	return err
}

const deleteApplication = `-- name: DeleteApplication :exec
delete
from application
where account1_id = ?
  and account2_id = ?
`

type DeleteApplicationParams struct {
	Account1ID int64 `json:"account1_id"`
	Account2ID int64 `json:"account2_id"`
}

func (q *Queries) DeleteApplication(ctx context.Context, arg *DeleteApplicationParams) error {
	_, err := q.db.Exec(ctx, deleteApplication, arg.Account1ID, arg.Account2ID)
	return err
}

const getApplications = `-- name: GetApplications :many
select account1_id, account2_id, apply_msg, refuse_msg, status, create_at, update_at
from application
where account1_id = $1
   or account2_id = $1
limit $2 offset $3
`

type GetApplicationsParams struct {
	Account1ID int64 `json:"account1_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) GetApplications(ctx context.Context, arg *GetApplicationsParams) ([]*Application, error) {
	rows, err := q.db.Query(ctx, getApplications, arg.Account1ID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Application{}
	for rows.Next() {
		var i Application
		if err := rows.Scan(
			&i.Account1ID,
			&i.Account2ID,
			&i.ApplyMsg,
			&i.RefuseMsg,
			&i.Status,
			&i.CreateAt,
			&i.UpdateAt,
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

const updateApplication = `-- name: UpdateApplication :exec
update application
set status = ? and refuse_msg = ?
where account1_id = ?
  and account2_id = ?
`

type UpdateApplicationParams struct {
	Status     Applicationstatus `json:"status"`
	RefuseMsg  sql.NullString    `json:"refuse_msg"`
	Account1ID int64             `json:"account1_id"`
	Account2ID int64             `json:"account2_id"`
}

func (q *Queries) UpdateApplication(ctx context.Context, arg *UpdateApplicationParams) error {
	_, err := q.db.Exec(ctx, updateApplication,
		arg.Status,
		arg.RefuseMsg,
		arg.Account1ID,
		arg.Account2ID,
	)
	return err
}
