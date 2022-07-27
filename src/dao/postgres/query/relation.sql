-- name: CreateGroupRelation :exec
insert into relation (relation_type, group_type)
values ('group', ROW (@name::varchar(50), @description::varchar(255), @avatar::varchar(255)))
returning id;

-- name: UpdateGroupRelation :exec
update relation
set (group_type) = (ROW (@name::varchar(50), @description::varchar(255), @avatar::varchar(255)))
where relation_type = 'group'
  and id = @id;

-- name: CreateFriendRelation :one
insert into relation (relation_type, friend_type)
values ('friend', ROW (@account1_id::bigint, @account2_id::bigint))
returning id;

-- name: DeleteRelation :exec
delete
from relation
where id = @id;

-- name: GetGroupRelationByID :one
select id,
       relation_type,
       (group_type).name::varchar        as name,
       (group_type).description::varchar as description,
       (group_type).avatar::varchar      as avatar,
       create_at
from relation
where relation_type = 'group'
  and id = @id;

-- name: ExistsFriendRelation :one
select exists(select 1
              from relation
              where relation_type = 'friend'
                and (friend_type).account1_id = @account1_id::bigint
                and (friend_type).account2_id = @account2_id::bigint);

-- name: DeleteFriendRelationsByAccountID :exec
delete
from relation
where relation_type = 'friend'
  and ((friend_type).account1_id = @account_id::bigint or (friend_type).account2_id = @account_id::bigint);

-- name: DeleteFriendRelationsByAccountIDs :exec
delete
from relation
where relation_type = 'friend'
  and ((friend_type).account1_id = ANY (@account_ids::bigint[])
    or (friend_type).account2_id = ANY (@account_ids::bigint[]));

-- name: GetFriendRelationByID :one
select (friend_type).account1_id as account1_id,
       (friend_type).account2_id as account2_id,
       create_at
from relation
where relation_type = 'friend'
  and id = $1;
