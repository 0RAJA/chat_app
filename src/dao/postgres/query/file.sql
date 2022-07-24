-- name: CreateFile :one
insert into file
    (file_name, file_type, file_size, key, url, relation_id, account_id)
values ($1, $2, $3, $4, $5, $6, $7) returning *;

-- name: DeleteFileByID :exec
delete
from file
where id = $1;

-- name: GetFileByRelationID :many
select *
from file
where relation_id = $1;

-- name: GetFileKeyByID :one
select key
from file
where id = $1;
