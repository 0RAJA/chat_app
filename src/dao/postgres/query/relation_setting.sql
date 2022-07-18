-- name: CreateRelationSetting :one
insert into relation_setting
(account_id, relation_id, nick_name, is_not_disturb, is_pin, pin_time, is_show, last_show, is_leader)
values (?, ?, ?, ?, ?, ?, ?, ?, ?)
returning *;

-- name: DeleteRelationSetting :exec
delete
from relation_setting
where account_id=?
  and relation_id=?;

-- name: UpdateRelationSetting :one
update relation_setting
set nick_name= ?,
    is_not_disturb=?,
    is_pin=?,
    pin_time=?,
    is_show=?,
    last_show=?
where account_id=?
  and relation_id=?
returning *;

-- name: GetRelationSetting :one

select *
from relation_setting
where account_id=?
  and relation_id=?
limit 1;


