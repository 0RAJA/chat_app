-- 表
drop table if exists relation_setting cascade;
drop table if exists friend_application cascade;
drop table if exists message cascade;
drop table if exists group_notify cascade;
drop table if exists file cascade;
drop table if exists account cascade;
drop table if exists "user" cascade;
drop table if exists relation cascade;
-- 类型
drop type if exists gender cascade;
drop type if exists applicationstatus cascade;
drop type if exists msgnotifytype cascade;
drop type if exists msgtype cascade;
drop type if exists relationtype cascade;
drop type if exists grouptype cascade;
drop type if exists friendtype cascade;
drop type if exists filetype cascade;
-- 方法
drop function if exists pin_timestamp() cascade;
drop function if exists cs_timestamp() cascade;
drop function if exists show_timestamp() cascade;
-- 语言
drop text search configuration chinese;
