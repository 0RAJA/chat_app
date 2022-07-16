-- name: CreateApplication :exec
insert into application (account1_id, account2_id, apply_msg)
values (?, ?, ?);

-- name: DeleteApplication :exec
delete
from application
where account1_id = ?
  and account2_id = ?;

-- name: UpdateApplication :exec
update application
set status = ? and refuse_msg = ?
where account1_id = ?
  and account2_id = ?;

-- name: GetApplications :many
select *
from application
where account1_id = $1
   or account2_id = $1
limit $2 offset $3;
