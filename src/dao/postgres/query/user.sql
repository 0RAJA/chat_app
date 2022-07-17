-- name: CreateUser :one
insert into "user" ("email", "password")
values (?, ?)
returning *;

-- name: DeleteUser :exec
delete
from "user"
where "id" = ?;

-- name: GetUserByEmail :one
select *
from "user"
where email = ?
limit 1;

-- name: GetUserByID :one
select *
from "user"
where id = ?
limit 1;

-- name: UpdateUser :exec
update "user"
set "email"    = ?,
    "password" = ?
where "id" = ?;

-- name: ExistEmail :one
select exists(select 1 from "user" where email = ?);

-- name: GetAllEmails :many
select email
from "user";
