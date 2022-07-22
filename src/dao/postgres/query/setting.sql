-- name: CreateSetting :exec
insert into setting (account_id, relation_id, nick_name, is_leader)
values ($1, $2, $3, $4);

-- name: DeleteSetting :exec
delete
from setting
where account_id = $1
  and relation_id = $2;

-- name: UpdateSettingNickName :exec
update setting
set nick_name = $3
where account_id = $1
  and relation_id = $2;

-- name: UpdateSettingDisturb :exec
update setting
set is_not_disturb = $3
where account_id = $1
  and relation_id = $2;

-- name: UpdateSettingPin :exec
update setting
set is_pin = $3
where account_id = $1
  and relation_id = $2;

-- name: UpdateSettingLeader :exec
update setting
set is_leader = $3
where account_id = $1
  and relation_id = $2;

-- name: GetSettingByID :one
select *
from setting
where account_id = $1
  and relation_id = $2;
