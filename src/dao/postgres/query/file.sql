-- name: CreateFile :one
insert into file
    (file_name, file_type, file_size, key, url, relation_id, account_id)
values (?, ?, ?, ?, ?, ?, ?)
returning *;

-- name: DeleteFileByID :exec
delete
from file
where id=?;

-- name: GetFileByRelationID :many
select *
from file
where relation_id=?;

-- name: GetFileByID :one
select *
from file
where id=?;
