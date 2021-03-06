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

-- name: GetAccountsByName :many
select a.*, r.id as relation_id, count(*) over () as total
from (select id, name, avatar from account where name like (@name::varchar || '%')) as a
         left join relation r on (r.relation_type = 'friend' and
                                  (((r.friend_type).account1_id = a.id and
                                    (r.friend_type).account2_id = @account_id::bigint) or
                                   (r.friend_type).account2_id = a.id and
                                   (r.friend_type).account1_id = @account_id::bigint))
limit $1 offset $2;
