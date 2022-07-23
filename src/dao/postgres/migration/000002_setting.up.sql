-- 好友关系视图
create view setting_friend_info as
select account_id,
       relation_id,
       nick_name,
       is_not_disturb,
       is_pin,
       pin_time,
       is_show,
       last_show,
       (r.friend_type).account1_id::bigint as account1_id,
       a.avatar                            as account1_avatar,
       (r.friend_type).account2_id::bigint as account2_id,
       b.avatar                            as account2_avatar
from setting s,
     relation r,
     account a,
     account b
where relation_id = r.id
  and r.relation_type = 'friend'
  and a.id = (r.friend_type).account1_id
  and b.id = (r.friend_type).account2_id;
