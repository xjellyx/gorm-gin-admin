--DROP TABLE  user_bases;
--drop table  user_cards ;


create table if not exists user_cards(
    id bigserial not null primary key ,
    name varchar not null ,
    card_id varchar(18) unique ,
    issue_org varchar,
    birthday varchar(12),
    valid_period varchar(12),
    card_id_addr varchar(64),
    sex int ,
    nation varchar,
    created_at timestamp  ,
    updated_at timestamp  ,
    deleted_at timestamp
);
create index if not exists user_cards_card_id on user_cards(card_id);

comment on column user_cards.name is '姓名';
comment on column user_cards.card_id is '身份证号码';
comment on column user_cards.issue_org is '身份证发证机关';
comment on column user_cards.birthday is '出生日期';
comment on column user_cards.valid_period is '有效时期';
comment on column user_cards.card_id_addr is '身份证地址';
comment on column user_cards.sex is '性别';
comment on column user_cards.nation is '贯藉';
comment on column user_cards.created_at is '创建时间';
comment on column user_cards.updated_at is '更新时间';
comment on column user_cards.deleted_at is '删除时间';


create table if not exists user_bases
(
id bigserial  not  null primary key ,
uid varchar(36) default '' unique  ,
username varchar(16) default '' unique ,
login_pwd varchar(64) default '' ,
pay_pwd varchar(64) default '',
phone varchar(11) default '' unique ,
nickname varchar(12) default '' ,
email varchar(32) default '' ,
head_icon varchar default '',
sign  varchar(256) default '',
status int,
-- card_id int references user_cards(id),
created_at timestamp  ,
updated_at timestamp  ,
deleted_at timestamp
);
create index if not exists user_bases_status on user_bases(status);
create index if not exists user_bases_username on user_bases(username);
--alter table user_bases drop constraint if exists user_bases_card_id_fkey;
--alter table user_bases add  foreign key (card_id)  references user_cards(id) ;
COMMENT ON  TABLE  user_bases is '用户基本信息';
comment on column user_bases.uid is '唯一uid';
comment on column user_bases.username is '唯一用户名';
comment on column user_bases.login_pwd is '登录密码';
comment on column user_bases.pay_pwd is '支付密码';
comment on column user_bases.phone is '手机号码';
comment on column user_bases.email is '邮箱';
comment on column user_bases.status is '状态' ;
comment on column user_bases.created_at is '创建时间';
comment on column user_bases.updated_at is '更新时间';
comment on column user_bases.deleted_at is '删除时间';


