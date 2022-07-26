-- 使用中文分词
--连接到目标数据库，创建zhparser解析器
-- CREATE EXTENSION zhparser;
-- 将zhparser解析器作为全文检索配置项
CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);
--普遍情况下，我们只需要按照名词(n)，动词(v)，形容词(a)，成语(i),叹词(e)和习用语(l)6种方式对句子进行划分就可以了，词典使用的是内置的simple词典，即仅做小写转换
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l WITH simple;
-- 性别类型
create type Gender As ENUM ('男','女','未知');
-- 申请状态
create type ApplicationStatus As ENUM ('已申请','已同意','已拒绝');
-- 消息通知类型
create type MsgNotifyType As ENUM ('system','common');
-- 消息类型
create type MsgType As ENUM ('text','file');
-- 群或好友关系的类型
create type RelationType As ENUM ('group','friend');
-- 群类型
create type GroupType As
(
    name        varchar(50),  -- 群名称
    description varchar(255), -- 群描述
    avatar      varchar(255)  -- 群头像
);
-- 好友类型
create type FriendType As
(
    account1_id bigint, -- 好友1的账号id
    account2_id bigint  -- 好友2的账号id
);
create type FileType As ENUM ('img','file');

-- 用户
create table "user"
(
    id        bigserial primary key,              -- 用户id
    email     varchar(255) not null unique,       -- 邮箱
    password  varchar(255) not null,              -- 密码
    create_at timestamptz  not null default now() -- 创建时间
);

-- 账号
create table account
(
    id        bigint primary key,                                                               -- 账号id
    user_id   bigint       not null references "user" (id) on delete cascade on update cascade, -- 用户id
    name      varchar(255) not null,                                                            -- 账号名
    avatar    varchar(255) not null,                                                            -- 账号头像
    gender    Gender       not null default '未知',                                               -- 账号性别
    signature text         not null default '这个用户很懒,什么也没留下',                                    -- 签名
    create_at timestamptz  not null default now(),                                              -- 创建时间
    constraint account_unique_name unique (user_id, name)                                       -- 一个用户的不同账号名不能重复
);

-- 账户名和头像索引
create index account_index_name_avatar on account (name, avatar);

-- 群组或好友
create table relation
(
    id            bigserial primary key,                             -- id
    relation_type RelationType not null,                             -- 关系类型 group:群组,friend:好友
    group_type    GroupType,                                         -- 群组信息 只有群组才有这个字段 否则为null
    friend_type   FriendType,                                        -- 好友信息 只有好友才有这个字段 否则为null
    create_at     timestamptz default now(),                         -- 创建时间
    check ( ((group_type is null) or (friend_type is null)) and
            ((group_type is not null) or (friend_type is not null))) -- 只能存在一种信息
);

-- 账号对群组或好友关系的设置
create table setting
(
    account_id     bigint       not null references account (id) on delete cascade on update cascade,  -- 账号id
    relation_id    bigint       not null references relation (id) on delete cascade on update cascade, -- 关系id
    nick_name      varchar(255) not null,                                                              -- 昵称，默认是账号名或者群组名
    is_not_disturb boolean      not null default false,                                                -- 是否免打扰
    is_pin         boolean      not null default false,                                                -- 是否置顶
    pin_time       timestamptz  not null default now(),                                                -- 置顶时间
    is_show        boolean      not null default true,                                                 -- 是否显示
    last_show      timestamptz  not null default now(),                                                -- 最后一次显示时间
    is_leader      boolean      not null default false,                                                -- 是否群主 仅仅对群组有效
    is_self        boolean      not null default false                                                 -- 是否是自己对自己的关系 仅仅对好友有效
);
-- 昵称索引
create index relation_setting_nickname on setting (nick_name);
create index setting_idx_account_id_relation_id on setting (account_id, relation_id);

-- 好友申请
create table application
(
    account1_id bigint            not null references account (id) on delete cascade on update cascade, -- 申请者账号id
    account2_id bigint            not null references account (id) on delete cascade on update cascade, -- 被申请者账号id
    apply_msg   text              not null,                                                             -- 申请信息
    refuse_msg  text              not null,                                                             -- 拒绝信息
    status      ApplicationStatus not null default '已申请',                                               -- 申请状态
    create_at   timestamptz       not null default now(),                                               -- 创建时间
    update_at   timestamptz       not null default now(),                                               -- 更新时间
    constraint f_a_pk primary key (account1_id, account2_id)
);

-- 文件记录
create table file
(
    id          bigserial primary key,                                                -- 文件id
    file_name   varchar(255) not null,                                                -- 文件名称
    file_type   FileType     not null,                                                -- 文件类型
    file_size   bigint       not null,                                                -- 文件大小 byte
    key         varchar(255) not null,                                                -- 文件key 用于oss中删除文件
    url         varchar(255) not null,                                                -- 文件url
    relation_id bigint references relation (id) on delete set null on update cascade, -- 关系id
    account_id  bigint references account (id) on delete set null on update cascade,  -- 发送账号id
    create_at   timestamptz  not null default now()                                   -- 创建时间
);
create index file_relation_id on file (relation_id);

-- 消息
create table message
(
    id              bigserial primary key,                                                               -- 消息id
    notify_type     MsgNotifyType not null,                                                              -- 消息通知类型 system:系统消息,common:普通消息
    msg_type        MsgType       not null,                                                              -- 消息类型 text:文本消息,file:文件消息
    msg_content     text          not null,                                                              -- 消息内容
    msg_expand      json,                                                                                -- 消息扩展信息
    file_id         bigint references file (id) on delete cascade on update cascade,                     -- 文件id 如果不是文件类型则为null
    account_id      bigint references account (id) on delete set null on update cascade,                 -- 发送账号id
    rly_msg_id      bigint references message (id) on delete cascade on update cascade,                  -- 回复消息id，没有则为null
    relation_id     bigint        not null references relation (id) on delete cascade on update cascade, -- 关系id
    create_at       timestamptz   not null default now(),                                                -- 创建时间
    is_revoke       boolean       not null default false,                                                -- 是否撤回
    is_top          boolean       not null default false,                                                -- 是否置顶
    is_pin          boolean       not null default false,                                                -- 是否置顶
    pin_time        timestamptz   not null default now(),                                                -- 置顶时间
    read_ids        bigint[]      not null default '{}'::bigint[],                                       -- 已读用户id集合
    msg_content_tsv tsvector,                                                                            -- 消息分词
    check (notify_type = 'common' or (notify_type = 'system' and account_id is null)),                   -- 系统消息时发送账号id为null
    check ( msg_type = 'text' or (msg_type = 'file' and file_id is not null))                            -- 文件消息时文件id不为null
);
create index msg_create_at on message (create_at);
-- 分词索引
create index message_msg_content_tsv on message using gin (to_tsvector('chinese', msg_content));

-- 触发器更新 message_msg_content_tsv
CREATE TRIGGER message_msg_content_tsv
    BEFORE INSERT OR UPDATE
    ON message
    FOR EACH ROW
EXECUTE PROCEDURE
    tsvector_update_trigger(msg_content_tsv, 'public.chinese', msg_content);

-- 群通知
create table group_notify
(
    id              bigserial primary key,                                               -- 群通知id
    relation_id     bigint references relation (id) on delete cascade on update cascade, -- 关系id
    msg_content     text        not null,                                                -- 消息内容
    msg_expand      json,                                                                -- 消息扩展信息
    account_id      bigint references account (id) on delete set null on update cascade, -- 发送账号id
    create_at       timestamptz not null default now(),                                  -- 创建时间
    read_ids        bigint[]    not null default '{}'::bigint[],                         -- 已读用户id集合
    msg_content_tsv tsvector                                                             -- 消息分词
);
-- 分词索引
create index group_notify_msg_content_tsv on group_notify using gin (to_tsvector('chinese', msg_content));
-- 触发器更新 group_notify_msg_content_tsv
CREATE TRIGGER group_notify_msg_content_tsv
    BEFORE INSERT OR UPDATE
    ON group_notify
    FOR EACH ROW
EXECUTE PROCEDURE
    tsvector_update_trigger(msg_content_tsv, 'public.chinese', msg_content);

-- 更新pin时间戳函数
create or replace function pin_timestamp() returns trigger as
              $$
begin
    if new.is_pin then new.pin_time = now(); end if; return new;
end;
$$ language plpgsql;

-- 更新关系设置pin时间戳触发器
create trigger pin_timestamp_trigger
    before update of is_pin
    on setting
    for each row
execute procedure pin_timestamp();

-- 更新消息pin时间戳触发器
create trigger pin_timestamp_trigger
    before update of is_pin
    on message
    for each row
execute procedure pin_timestamp();

-- 更新时间戳函数
create or replace function cs_timestamp(
) returns trigger as
              $$
begin
    new.update_at = now();
return new;
end;
$$ language plpgsql;

-- 申请表更新时间戳触发器
create trigger application_update_at_trigger
    before update
    on application
    for each row
    execute procedure cs_timestamp();

-- 更新show时间戳函数
create or replace function show_timestamp() returns trigger as
              $$
begin
    if new.is_show then new.last_show = now(); end if; return new;
end;
$$ language plpgsql;

-- 更新关系设置show时间戳触发器
create trigger show_timestamp_trigger
    before update of is_show
    on setting
    for each row
execute procedure show_timestamp();