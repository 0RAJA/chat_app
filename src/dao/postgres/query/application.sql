-- name: CreateApplication :exec
insert into application (account1_id, account2_id, apply_msg, refuse_msg)
values ($1, $2, $3, '');

-- name: DeleteApplication :exec
delete
from application
where account1_id = $1
  and account2_id = $2;

-- name: UpdateApplication :exec
update application
set status     = $1,
    refuse_msg = $2
where account1_id = $3
  and account2_id = $4;

-- name: GetApplicationByID :one
select *
from application
where account1_id = $1
  and account2_id = $2
limit 1;

-- name: ExistsApplicationByIDWithLock :one
select exists(
               select 1
               from application
               where (account1_id = $1 and account2_id = $2)
                  or (account1_id = $2 and account2_id = $1)
                   for update);

-- name: ExistsApplicationByID :one
select exists(
               select 1
               from application
               where (account1_id = $1 and account2_id = $2)
                  or (account1_id = $2 and account2_id = $1));

-- name: GetApplications :many
select app.*,
       a1.avatar as account1_avatar,
       a1.name   as account1_name,
       a2.avatar as account2_avatar,
       a2.name   as account2_name
from account a1,
     account a2,
     (select *, count(*) over () as total
      from application
      where account1_id = @account_id
         or account2_id = @account_id
      order by create_at desc
      limit $1 offset $2) as app
where a1.id = app.account1_id
  and a2.id = app.account2_id;
