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

-- name: GetFriendPinSettingsOrderByPinTime :many
select relation_id,
       account1_id,
       account1_avatar,
       account2_id,
       account2_avatar,
       nick_name,
       pin_time
from setting_friend_info s
where account_id = $1
  and is_pin = true
order by pin_time;

-- name: GetFriendShowSettingsOrderByShowTime :many
select s.*
from setting_friend_info s
where account_id = $1
  and is_show = true
order by last_show desc;

-- name: GetFriendSettingsOrderByName :many
select s.*
from setting_friend_info s
where account_id = $1
order by nick_name;

-- name: ExistsFriendSetting :one
select exists(
               select 1
               from setting s,
                    relation r
               where r.relation_type = 'friend'
                 and ((r.friend_type).account1_id = @account1_id::bigint and
                      (r.friend_type).account2_id = @account2_id::bigint)
                 and s.account_id = @account1_id
           );
