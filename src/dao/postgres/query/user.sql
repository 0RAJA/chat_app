-- name: CreateUser :one
insert into "user" ("email", "password")
values (?, ?)
returning *;

-- name: DeleteUser :exec
delete
from "user"
where "id" = ?;

-- name: GetUserByEmail :one
select id, email, password
from "user"
where email = ?
limit 1;

-- name: GetUserByID :one
select id, email, password
from "user"
where id = ?
limit 1;

-- name: UpdateUser :one
update "user"
set "email"    = ?,
    "password" = ?
where "id" = ?
returning *;

-- name: ExistEmail :one
select exists(select 1 from "user" where email = ?);

-- name: GetAllEmails :many
select email
from "user";
