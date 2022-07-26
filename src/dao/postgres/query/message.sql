;-- name: CreateMsg :one
insert into message (notify_type, msg_type, msg_content, msg_expand, file_id, account_id, rly_msg_id, relation_id)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: UpdateMsgPin :exec
update message
set is_pin = $2
where id = $1;

-- name: UpdateMsgTop :exec
update message
set is_top = $2
where id = $1;

-- name: UpdateMsgRevoke :exec
update message
set msg_type    = null,
    is_revoke   = $2,
    msg_content = ''
where id = $1;

-- name: UpdateMsgReads :exec
update message
set read_ids = array_append(read_ids, $2)
where id = $1
  and $2 != ANY (read_ids);

-- name: GetMsgsByRelationIDAndTime :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_expand,
       m1.file_id,
       m1.account_id,
       m1.rly_msg_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id) as reply_count,
       m2.msg_type                                              as rly_msg_type,
       m2.msg_content                                           as rly_msg_content,
       m2.msg_expand                                            as rly_msg_expand
from message m1
         left join message m2 on m1.relation_id = $1 and m1.create_at < $2 and m1.rly_msg_id = m2.id
order by m1.create_at desc
limit $3 offset $4;

-- name: GetRlyMsgsInfoByMsgID :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_expand,
       m1.file_id,
       m1.account_id,
       m1.rly_msg_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids
from message m1
where m1.rly_msg_id = $1
  and m1.relation_id = $2
order by m1.create_at
limit $3 offset $4;

-- name: GetPinMsgsByRelationID :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_expand,
       m1.file_id,
       m1.account_id,
       m1.rly_msg_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids
from message m1
where m1.rly_msg_id = $1
  and m1.is_pin = true
order by m1.pin_time desc
limit $2 offset $3;

-- name: HasReadMsg :one
select exists(select 1 from message where id = @msg_id and @account_id = ANY (read_ids));

-- name: GetTopMsgByRelationID :one
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_expand,
       m1.file_id,
       m1.account_id,
       m1.rly_msg_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids
from message m1
where m1.relation_id = $1
  and m1.is_top = true
limit 1;

-- name: UpdateMsgTopFalseByRelationID :exec
update message
set is_top = false
where relation_id = $1
  and is_top = true;

-- name: UpdateMsgTopTrueByMsgID :exec
update message
set is_top = true
where id = $1;
