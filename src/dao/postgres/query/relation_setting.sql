-- name: CreateRelationSetting :one
insert into relation_setting
(account_id, relation_id, nick_name, is_not_disturb, is_pin, pin_time, is_show, last_show, is_leader)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
returning *;

-- name: DeleteRelationSetting :exec
delete
from relation_setting
where account_id=$1
  and relation_id=$2;

-- name: UpdateRelationSetting :one
update relation_setting
set nick_name= $1,
    is_not_disturb=$2,
    is_pin=$3,
    pin_time=$4,
    is_show=$5,
    last_show=$6
where account_id=$7
  and relation_id=$8
returning *;

-- name: GetRelationSetting :one

select *
from relation_setting
where account_id=$1
  and relation_id=$2
limit 1;


