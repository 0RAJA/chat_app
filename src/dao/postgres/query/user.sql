-- name: CreateUser :one
insert into "user" (email, password)
values ($1, $2)
returning *;

-- name: DeleteUser :exec
delete
from "user"
where "id" = $1;

-- name: GetUserByEmail :one
select *
from "user"
where email = $1
limit 1;

-- name: GetUserByID :one
select *
from "user"
where id = $1
limit 1;

-- name: UpdateUser :exec
update "user"
set "email"    = $1,
    "password" = $2
where "id" = $3;

-- name: ExistEmail :one
select exists(select 1 from "user" where email = $1);

-- name: GetAllEmails :many
select email
from "user";

-- name: ExistsUserByID :one
select exists(select 1 from "user" where id = $1);

-- name: GetAccountIDsByUserID :many
select id
from account
where user_id = $1;
