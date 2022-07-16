-- name: CreateGroupRelation :exec
insert into relation (relation_type, group_type)
values ('group', ROW (@name::varchar(50), @description::varchar(255), @avatar::varchar(255)));

-- name: UpdateGroupRelation :exec
update relation
set (group_type) = (ROW (@name::varchar(50), @description::varchar(255), @avatar::varchar(255)))
where relation_type = 'group'
  and id = @id;

-- name: CreateFriendRelation :exec
insert into relation (relation_type, friend_type)
values ('friend', ROW (@account1_id::bigint, @account2_id::bigint));

-- name: DeleteRelation :exec
delete
from relation
where id = @id;

-- name: GetGroupRelationByID :one
select id,
       relation_type,
       (group_type).name        as name,
       (group_type).description as description,
       (group_type).avatar      as avatar,
       create_at
from relation
where relation_type = 'group'
  and id = @id;
