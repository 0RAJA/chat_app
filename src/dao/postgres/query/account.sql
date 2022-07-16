-- name: CreateAccount :one
insert into account (id, user_id, name, avatar)
values (?, ?, ?, ?)
returning *;

-- name: DeleteAccount :exec
delete
from account
where id = ?;

-- name: UpdateAccount :one
update account
set name      = ?,
    avatar    = ?,
    gender    = ?,
    signature = ?
where id = ?
returning *;

-- name: GetAccountByID :one
select *
from account
where id = ?
limit 1;

-- name: GetAccountsByUserID :many
select id, name, avatar
from account
where user_id = ?;
