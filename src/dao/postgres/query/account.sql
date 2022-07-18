-- name: CreateAccount :one
insert into account (id, user_id, name, avatar)
values ($1, $2, $3, $4)
returning *;

-- name: DeleteAccount :exec
delete
from account
where id = $1;

-- name: UpdateAccount :one
update account
set name      = $1,
    avatar    = $2,
    gender    = $3,
    signature = $4
where id = $5
returning *;

-- name: GetAccountByID :one
select *
from account
where id = $1
limit 1;

-- name: GetAccountsByUserID :many
select id, name, avatar
from account
where user_id = $1;

-- name: ExistsAccountByID :one
select exists(
               select 1
               from account
               where id = $1
           );
