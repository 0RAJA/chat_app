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
select exists(select 1 from file where account_id= $1 and file_name= 'AccountAvatar');
-- name: GetGroupAvatar :one
select *
from file
where relation_id = $1
  and account_id is null;

-- name: UpdateGroupAvatar :exec
update file
set url= $1
where relation_id = $2 and file_name = 'groupAvatar' ;
-- name: UpdateAccountFile :exec
update  file
set  url = $1
where account_id = $2 and file_name = 'AccountAvatar';
-- name: GetAllRelationsOnFile :many
select relation_id
from file
group by relation_id;

-- name: GetFileByRelationIDIsNUll :many
select id, key
from file
where relation_id is null and file_name != 'AccountAvatar' ;

-- name: GetFileDetailsByID :one
select *
from file
where id = $1;