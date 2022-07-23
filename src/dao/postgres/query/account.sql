-- name: CreateAccount :exec
insert into account (id, user_id, name, avatar)
values ($1, $2, $3, $4);

-- name: DeleteAccount :exec
delete
from account
where id = $1;

-- name: UpdateAccount :exec
update account
set name      = $1,
    avatar    = $2,
    gender    = $3,
    signature = $4
where id = $5;

-- name: GetAccountByID :one
select *
from account
where id = $1
limit 1;

-- name: GetAccountsByUserID :many
select id, name, avatar
from account
where user_id = $1;

-- name: GetAccountsByName :many
select id, name, avatar, count(*) over () as total
from account
where name = $1
limit $2 offset $3;

-- name: ExistsAccountByID :one
select exists(
               select 1
               from account
               where id = $1
           );

-- name: ExistsAccountByNameAndUserID :one
select exists(
               select 1
               from account
               where user_id = $1
                 and name = $2
           );

-- name: CountAccountByUserID :one
select count(id)::int
from account
where user_id = $1;

-- name: DeleteAccountsByUserID :many
delete
from account
where user_id = $1
returning id;
