-- name: CreateGroupNotify :one
insert into group_notify
(id, relation_id, msg_content, msg_expand, account_id, create_at, read_ids)
values (?, ?, ?, ?, ?, ?, ?)
returning *;
-- name: DeleteGroupNotify :exec
delete
from group_notify
where id = ?;

-- name: UpdateGroupNotify :one
update group_notify
set id         = ?,
    relation_id= ?,
    msg_content= ?,
    msg_expand = ?,
    account_id = ?,
    create_at  = ?,
    read_ids   = ?
where id=?
returning *;

-- name: GetGroupNotifyByID :one
select *
from group_notify
where id = ?
limit 1;





