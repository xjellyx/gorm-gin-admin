-- DROP TABLE IF EXISTS user_bases;
create table if not exists user_bases
(
id bigserial  not  null primary key ,
uid varchar(36) default '' unique  ,
username varchar(32) default ''  ,
login_passwd varchar(16) default '' ,
pay_passwd varchar(16) default '',
phone varchar(11) default '' unique ,
nick_name varchar(16) default '' ,
email varchar(32) default '' ,
status int ,
created_at timestamp  ,
updated_at timestamp  ,
deleted_at timestamp
);
create index if not exists user_bases_status on user_bases(status);
create index if not exists user_bases_username on user_bases(username);
COMMENT ON  TABLE  user_bases is '用户基本信息';
comment on column user_bases.uid is '唯一uid';
comment on column user_bases.username is '唯一用户名';
comment on column user_bases.login_passwd is '登录密码';
comment on column user_bases.pay_passwd is '支付密码';
comment on column user_bases.phone is '手机号码';
comment on column user_bases.email is '邮箱';
comment on column user_bases.status is '状态' ;
comment on column user_bases.created_at is '创建时间';
comment on column user_bases.updated_at is '更新时间';
comment on column user_bases.deleted_at is '删除时间';
