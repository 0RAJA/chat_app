-- name: CreateGroupNotify :one
insert into group_notify
(id, relation_id, msg_content, msg_expand, account_id, create_at, read_ids)
values ($1, $2, $3, $4, $5, $6, $7)
returning *;
-- name: DeleteGroupNotify :exec
delete
from group_notify
where id = $1;

-- name: UpdateGroupNotify :one
update group_notify
set id         = $1,
    relation_id= $2,
    msg_content= $3,
    msg_expand = $4,
    account_id = $5,
    create_at  = $6,
    read_ids   = $7
where id=$8;

-- name: UpdateGroupNotify :one
update group_notify
set id         = $1,
    relation_id= $2,
    msg_content= $3,
    msg_expand = $4,
    account_id = $5,
    create_at  = $6,
    read_ids   = $7
where id=$8
    returning *;

-- name: GetGroupNotifyByID :one
select *
from group_notify
where id = $1
limit 1;





