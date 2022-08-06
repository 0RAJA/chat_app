-- name: CreateFile :one
insert into file
    (file_name, file_type, file_size, key, url, relation_id, account_id)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;

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

-- name: GetAvatar :one
select f.url, f.file_name
from file f
where file_name =
      (select max(file_name) from file f1 where f1.account_id = $1 and f1.relation_id is null);

-- name: GetGroupAvatar :one
select *
from file
where relation_id = $1
  and account_id is null;

-- name: UpdateGroupAvatar :exec
update file
set url= $1
where relation_id = $2;

-- name: GetAllRelationsOnFile :many
select relation_id from file group by relation_id;

-- name: GetFileByRelationIDIsNUll :many
select id,key from file where relation_id is null ;