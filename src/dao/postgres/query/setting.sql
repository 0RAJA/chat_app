-- name: CreateSetting :exec
insert into setting (account_id, relation_id, nick_name, is_leader, is_self)
values ($1, $2, $3, $4, $5);

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
select s.*,
       a.id     as account_id,
       a.name   as account_name,
       a.avatar as account_avatar
from (select setting.relation_id, setting.nick_name, setting.pin_time
      from setting,
           relation
      where setting.account_id = $1
        and setting.is_pin = true
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select id from setting where relation_id = s.relation_id and (account_id != $1 or is_self = true))
order by s.pin_time;

-- name: GetFriendShowSettingsOrderByShowTime :many
select s.*,
       a.id     as account_id,
       a.name   as account_name,
       a.avatar as account_avatar
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.is_show = true
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select id from setting where relation_id = s.relation_id and (account_id != $1 or is_self = true))
order by s.pin_time;

-- name: GetFriendSettingsOrderByName :many
select s.*,
       a.id     as account_id,
       a.name   as account_name,
       a.avatar as account_avatar
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select id from setting where relation_id = s.relation_id and (account_id != $1 or is_self = true))
order by s.pin_time;

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

-- name: GetFriendSettingsByName :many
select s.*,
       a.id             as account_id,
       a.name           as account_name,
       a.avatar         as account_avatar,
       count(*) over () as total
from (select relation_id,
             nick_name,
             is_not_disturb,
             is_pin,
             pin_time,
             is_show,
             last_show,
             is_self
      from setting,
           relation
      where setting.account_id = $1
        and setting.relation_id = relation.id
        and relation.relation_type = 'friend') as s,
     account a
where a.id = (select id
              from setting
              where relation_id = s.relation_id
                and (account_id != $1 or is_self = true))
  and ((a.name like (@name::varchar || '%')) or (nick_name like (@name::varchar || '%')))
order by s.pin_time
limit $2 offset $3;
