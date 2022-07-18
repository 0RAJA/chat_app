-- name: CreateApplication :exec
insert into application (account1_id, account2_id, apply_msg)
values ($1, $2, $3);

-- name: DeleteApplication :exec
delete
from application
where account1_id = $1
  and account2_id = $2;

-- name: UpdateApplication :exec
update application
set status = $1 and refuse_msg = $2
where account1_id = $3
  and account2_id = $4;

-- name: GetApplications :many
select *
from application
where account1_id = $1
   or account2_id = $1
limit $2 offset $3;
