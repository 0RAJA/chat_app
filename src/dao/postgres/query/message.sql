-- name: CreateMsg :one
insert into message (notify_type, msg_type, msg_content, msg_extend, file_id, account_id, rly_msg_id, relation_id)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning id, msg_content, msg_extend, file_id,create_at;

-- name: UpdateMsgPin :exec
update message
set is_pin = $2
where id = $1;

-- name: UpdateMsgRevoke :exec
update message
set is_revoke = $2
where id = $1;

-- name: UpdateMsgReads :many
update message m
set read_ids = array_append(read_ids, @accountID::bigint)
where id = any (@msgIDs::bigint[])
  and @accountID::bigint != ANY (read_ids)
  and relation_id = $1
returning id,account_id::bigint;

-- name: GetMsgsByRelationIDAndTime :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
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
       count(*) over ()                                                                      as total,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count
from message m1
where m1.relation_id = $1
  and m1.create_at < $2
order by m1.create_at
limit $3 offset $4;

-- name: GetRlyMsgsInfoByMsgID :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count,
       count(*) over ()                                                                      as total
from message m1
where m1.relation_id = $1
  and m1.rly_msg_id = @rly_msg_id::bigint
order by m1.create_at
limit $2 offset $3;

-- name: GetPinMsgsByRelationID :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count,
       count(*) over ()                                                                      as total
from message m1
where m1.relation_id = $1
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
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       m1.is_revoke,
       m1.is_top,
       m1.is_pin,
       m1.pin_time,
       m1.read_ids,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = $1) as reply_count,
       count(*) over ()                                                                      as total
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

-- name: UpdateMsgTopFalseByMsgID :exec
update message
set is_top = false
where id = $1;

-- name: GetMsgByID :one
select id,
       notify_type,
       msg_type,
       msg_content,
       msg_extend,
       file_id,
       account_id,
       rly_msg_id,
       relation_id,
       create_at,
       is_revoke,
       is_top,
       is_pin,
       pin_time,
       read_ids
from message
where id = $1
limit 1;

-- name: GetMsgsByContent :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       count(*) over () as total
from message m1
         join setting s on m1.relation_id = s.relation_id and s.account_id = $1
where (not m1.is_revoke)
  and m1.msg_content_tsv @@ plainto_tsquery(@content::varchar)
order by m1.create_at desc
limit $2 offset $3;

-- name: GetMsgsByContentAndRelation :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
       m1.file_id,
       m1.account_id,
       m1.relation_id,
       m1.create_at,
       count(*) over () as total
from message m1
         join setting s on m1.relation_id = $1 and m1.relation_id = s.relation_id and s.account_id = $2
where (not m1.is_revoke)
  and m1.msg_content_tsv @@ plainto_tsquery(@content::varchar)
order by m1.create_at desc
limit $3 offset $4;

-- name: FeedMsgsByAccountIDAndTime :many
select m1.id,
       m1.notify_type,
       m1.msg_type,
       m1.msg_content,
       m1.msg_extend,
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
       count(*) over ()                                                                                  as total,
       (select count(id) from message where rly_msg_id = m1.id and message.relation_id = m1.relation_id) as reply_count,
       (@accountID::bigint = m1.account_id or @accountID::bigint = any (m1.read_ids))::boolean           as has_read
from message m1
         join setting s on m1.relation_id = s.relation_id and s.account_id = @accountID::bigint
where m1.create_at > $1
limit $2 offset $3;
