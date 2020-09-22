--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4 (Debian 12.4-1.pgdg100+1)
-- Dumped by pg_dump version 12.4 (Debian 12.4-1.pgdg100+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: api_groups; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.api_groups (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    path character varying(100),
    description character varying(64),
    api_group character varying(36),
    method character varying(10)
);


ALTER TABLE public.api_groups OWNER TO business;

--
-- Name: api_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.api_groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.api_groups_id_seq OWNER TO business;

--
-- Name: api_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.api_groups_id_seq OWNED BY public.api_groups.id;


--
-- Name: behavior_records; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.behavior_records (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    uid character varying(36),
    username character varying(16),
    behavior character varying(500),
    method character varying(12),
    path character varying(120),
    ip character varying(20)
);


ALTER TABLE public.behavior_records OWNER TO business;

--
-- Name: behavior_records_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.behavior_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.behavior_records_id_seq OWNER TO business;

--
-- Name: behavior_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.behavior_records_id_seq OWNED BY public.behavior_records.id;


--
-- Name: casbin_rule; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.casbin_rule (
    id bigint NOT NULL,
    p_type character varying(100),
    v0 character varying(100),
    v1 character varying(100),
    v2 character varying(100),
    v3 character varying(100),
    v4 character varying(100),
    v5 character varying(100)
);


ALTER TABLE public.casbin_rule OWNER TO business;

--
-- Name: casbin_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.casbin_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.casbin_rule_id_seq OWNER TO business;

--
-- Name: casbin_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.casbin_rule_id_seq OWNED BY public.casbin_rule.id;


--
-- Name: menus; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.menus (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(36),
    parent_id bigint,
    path character varying(36),
    component character varying(36),
    sort bigint,
    meta json,
    hidden boolean
);


ALTER TABLE public.menus OWNER TO business;

--
-- Name: menus_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menus_id_seq OWNER TO business;

--
-- Name: menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.menus_id_seq OWNED BY public.menus.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    role character varying(36),
    level character varying(1) DEFAULT '0'::character varying
);


ALTER TABLE public.roles OWNER TO business;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO business;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: user_bases; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.user_bases (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    uid character varying(36),
    username character varying(16),
    phone character varying(11),
    login_pwd character varying(128),
    pay_pwd character varying(128),
    email character varying(32),
    nickname character varying(12),
    head_icon text,
    sign character varying(256),
    status bigint,
    role_refer bigint
);


ALTER TABLE public.user_bases OWNER TO business;

--
-- Name: user_bases_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.user_bases_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_bases_id_seq OWNER TO business;

--
-- Name: user_bases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.user_bases_id_seq OWNED BY public.user_bases.id;


--
-- Name: user_cards; Type: TABLE; Schema: public; Owner: business
--

CREATE TABLE public.user_cards (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    uid character varying(36),
    name text,
    card_id character varying(18),
    issue_org text,
    birthday character varying(12),
    valid_period character varying(24),
    card_id_addr character varying(64),
    sex bigint,
    nation text
);


ALTER TABLE public.user_cards OWNER TO business;

--
-- Name: user_cards_id_seq; Type: SEQUENCE; Schema: public; Owner: business
--

CREATE SEQUENCE public.user_cards_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_cards_id_seq OWNER TO business;

--
-- Name: user_cards_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: business
--

ALTER SEQUENCE public.user_cards_id_seq OWNED BY public.user_cards.id;


--
-- Name: api_groups id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.api_groups ALTER COLUMN id SET DEFAULT nextval('public.api_groups_id_seq'::regclass);


--
-- Name: behavior_records id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.behavior_records ALTER COLUMN id SET DEFAULT nextval('public.behavior_records_id_seq'::regclass);


--
-- Name: casbin_rule id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.casbin_rule ALTER COLUMN id SET DEFAULT nextval('public.casbin_rule_id_seq'::regclass);


--
-- Name: menus id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.menus ALTER COLUMN id SET DEFAULT nextval('public.menus_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: user_bases id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.user_bases ALTER COLUMN id SET DEFAULT nextval('public.user_bases_id_seq'::regclass);


--
-- Name: user_cards id; Type: DEFAULT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.user_cards ALTER COLUMN id SET DEFAULT nextval('public.user_cards_id_seq'::regclass);


--
-- Data for Name: api_groups; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.api_groups (id, created_at, updated_at, deleted_at, path, description, api_group, method) FROM stdin;
2	2020-09-01 17:29:56.170657+08	2020-09-03 10:06:38.450223+08	\N	/api/v1/admin/editApiGroup	编辑接口	api	put
4	2020-09-03 10:10:00.046753+08	2020-09-03 10:10:00.046753+08	\N	/api/v1/admin/removeApiGroup	删除接口	api	delete
3	2020-09-03 10:08:42.259244+08	2020-09-03 10:51:08.130929+08	\N	/api/v1/admin/getApiGroupList	获取接口列表	api	get
1	2020-09-01 17:13:25.409182+08	2020-09-03 10:52:59.324116+08	\N	/api/v1/admin/addApiGroup	添加接口	api	post
7	2020-09-03 11:07:09.557597+08	2020-09-03 11:21:40.630727+08	\N	/api/v1/captcha	获取验证码	captcha	post
8	2020-09-03 11:34:19.851617+08	2020-09-03 11:34:19.851617+08	\N	/api/v1/admin/getMenuList	获取菜单列表	menu	get
9	2020-09-03 11:34:46.706629+08	2020-09-03 11:34:46.706629+08	\N	/api/v1/admin/addMenu	添加菜单	menu	post
10	2020-09-03 11:35:15.706411+08	2020-09-03 11:35:15.706411+08	\N	/api/v1/admin/delMenu	删除菜单	menu	delete
12	2020-09-03 11:36:44.019502+08	2020-09-03 11:36:44.019502+08	\N	/api/v1/admin/login	管理员登录	user	post
11	2020-09-03 11:35:39.337759+08	2020-09-03 11:36:53.659572+08	\N	/api/v1/admin/menu	编辑菜单	menu	put
13	2020-09-03 11:37:25.446396+08	2020-09-03 11:38:09.553387+08	\N	/api/v1/user/profile	用户信息	user	get
14	2020-09-03 11:38:51.864979+08	2020-09-03 11:38:51.864979+08	\N	/api/v1/user/modifyPwd	修改登录密码	user	put
17	2020-09-03 14:21:04.330658+08	2020-09-03 14:21:04.330658+08	\N	/api/v1/admin/deleteUser	删除用户	user	delete
19	2020-09-03 14:22:09.017293+08	2020-09-03 14:22:28.815854+08	\N	/api/v1/admin/getUserKV	获取用户一些状态信息	user	get
18	2020-09-03 14:21:22.647945+08	2020-09-03 14:22:44.145525+08	\N	/api/v1/admin/editUser	编辑用户	user	put
20	2020-09-08 18:00:04.217012+08	2020-09-08 18:00:04.217012+08	\N	/api/v1/admin/getRoleList	获取角色列表	role	get
21	2020-09-08 18:00:47.838141+08	2020-09-08 18:00:47.838141+08	\N	/api/v1/admin/getRoleLevel	获取角色等级列表	role	get
22	2020-09-08 18:01:18.409594+08	2020-09-08 18:01:18.409594+08	\N	/api/v1/admin/editRole	编辑角色	role	put
23	2020-09-08 18:01:47.348512+08	2020-09-08 18:01:47.348512+08	\N	/api/v1/admin/removeRole	删除角色	role	delete
24	2020-09-08 18:02:18.796049+08	2020-09-08 18:02:18.796049+08	\N	/api/v1/admin/addRole	增加角色	role	post
16	2020-09-03 14:19:24.614455+08	2020-09-08 20:19:21.486459+08	\N	/api/v1/admin/userTotal	获取用户总数	user	get
15	2020-09-03 11:39:29.062578+08	2020-09-15 14:52:21.951604+08	\N	/api/v1/admin/listUser	获取用户列表	user	get
25	2020-09-17 11:02:29.460261+08	2020-09-17 11:02:29.460261+08	\N	/api/v1/admin/getRoleApiList	獲取角色ａｐｉ權限	role	get
26	2020-09-17 11:05:21.370992+08	2020-09-17 11:05:21.370992+08	\N	/api/v1/admin/addRoleApiPerm	角色添加權限	role	post
27	2020-09-17 11:10:06.41631+08	2020-09-17 11:10:06.41631+08	\N	/api/v1/admin/behaviorCount	獲取操作記錄總數	behavior	get
28	2020-09-17 11:11:10.950778+08	2020-09-17 11:11:10.950778+08	\N	/api/v1/admin/getBehaviorList	獲取數據列表	behavior	get
29	2020-09-17 11:11:42.895352+08	2020-09-17 11:11:42.895352+08	\N	/api/v1/admin/removeBehaviors	刪除操作記錄	behavior	delete
30	2020-09-17 11:19:12.174338+08	2020-09-17 11:19:12.174338+08	\N	 /api/v1/admin/removeRoleApiPerm	刪除角色權限	role	delete
31	2020-09-21 16:08:56.651647+08	2020-09-21 16:08:56.651647+08	\N	/api/v1/admin/getSettings	获取配置文件信息	setting	get
33	2020-09-21 16:29:24.37912+08	2020-09-21 16:29:24.37912+08	\N	/api/v1/admin/getApiGroupListAll	获取全部ＡＰＩ数据	api	get
\.


--
-- Data for Name: behavior_records; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.behavior_records (id, created_at, updated_at, deleted_at, uid, username, behavior, method, path, ip) FROM stdin;
10	2020-09-17 09:55:02.927536+08	2020-09-17 09:55:02.927536+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  role test	put	/api/v1/admin/editRole	127.0.0.1:47346
11	2020-09-17 11:02:29.462146+08	2020-09-17 11:02:29.462146+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	172.25.0.1:40336
12	2020-09-17 11:05:21.372321+08	2020-09-17 11:05:21.372321+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	172.25.0.1:40424
13	2020-09-17 11:06:08.413238+08	2020-09-17 11:06:08.413238+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm test 	post	/api/v1/admin/addRoleApiPerm	172.25.0.1:40474
14	2020-09-17 11:06:16.084987+08	2020-09-17 11:06:16.084987+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  role test	put	/api/v1/admin/editRole	172.25.0.1:40480
15	2020-09-17 11:08:30.066445+08	2020-09-17 11:08:30.066445+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission test 	delete	/api/v1/admin/removeRoleApiPerm	172.25.0.1:40528
16	2020-09-17 11:10:06.417742+08	2020-09-17 11:10:06.417742+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	172.25.0.1:40580
17	2020-09-17 11:11:10.952296+08	2020-09-17 11:11:10.952296+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	172.25.0.1:40600
18	2020-09-17 11:11:42.896726+08	2020-09-17 11:11:42.896726+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	172.25.0.1:40614
19	2020-09-17 11:11:54.090965+08	2020-09-17 11:11:54.090965+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm root 	post	/api/v1/admin/addRoleApiPerm	172.25.0.1:40638
20	2020-09-17 11:12:43.462673+08	2020-09-17 11:12:43.462673+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  behavior record  [] 	delete	/api/v1/admin/removeBehaviors	172.25.0.1:40688
21	2020-09-17 11:19:12.417145+08	2020-09-17 11:19:12.417145+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	127.0.0.1:44190
22	2020-09-17 11:19:32.079578+08	2020-09-17 11:19:32.079578+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm root 	post	/api/v1/admin/addRoleApiPerm	127.0.0.1:44416
23	2020-09-17 11:20:19.020806+08	2020-09-17 11:20:19.020806+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission root 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:44834
24	2020-09-17 11:20:23.906732+08	2020-09-17 11:20:23.906732+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission root 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:44884
25	2020-09-17 11:20:28.494047+08	2020-09-17 11:20:28.494047+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission root 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:44932
26	2020-09-17 11:20:36.526852+08	2020-09-17 11:20:36.526852+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm root 	post	/api/v1/admin/addRoleApiPerm	127.0.0.1:45010
27	2020-09-17 11:20:36.623927+08	2020-09-17 11:20:36.623927+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission root 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:45012
28	2020-09-17 15:47:11.593813+08	2020-09-17 15:47:11.593813+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:51158
29	2020-09-17 15:47:21.719821+08	2020-09-17 15:47:21.719821+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:51298
30	2020-09-18 09:07:22.053888+08	2020-09-18 09:07:22.053888+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  role test	put	/api/v1/admin/editRole	172.25.0.1:53558
31	2020-09-21 14:29:23.332721+08	2020-09-21 14:29:23.332721+08	\N	c60bbe71-9c87-45e0-96ee-aa1f24695002	test111	add  user 	post	/api/v1/admin/addUser	127.0.0.1:37130
32	2020-09-21 14:37:23.827999+08	2020-09-21 14:37:23.827999+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:39970
33	2020-09-21 14:37:27.066159+08	2020-09-21 14:37:27.066159+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:39982
34	2020-09-21 14:38:34.15788+08	2020-09-21 14:38:34.15788+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:40050
35	2020-09-21 14:41:01.274369+08	2020-09-21 14:41:01.274369+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:41474
36	2020-09-21 14:46:38.334555+08	2020-09-21 14:46:38.334555+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  user a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	put	/api/v1/admin/editUser	127.0.0.1:44316
37	2020-09-21 16:08:57.031176+08	2020-09-21 16:08:57.031176+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	127.0.0.1:57732
38	2020-09-21 16:24:57.339538+08	2020-09-21 16:24:57.339538+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission root 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:37852
39	2020-09-21 16:24:57.714836+08	2020-09-21 16:24:57.714836+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm root 	post	/api/v1/admin/addRoleApiPerm	127.0.0.1:37850
40	2020-09-21 16:29:24.618775+08	2020-09-21 16:29:24.618775+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  api 	post	/api/v1/admin/addApiGroup	127.0.0.1:38824
41	2020-09-21 16:29:34.249988+08	2020-09-21 16:29:34.249988+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission root 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:38946
42	2020-09-21 16:29:34.519773+08	2020-09-21 16:29:34.519773+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm root 	post	/api/v1/admin/addRoleApiPerm	127.0.0.1:38944
43	2020-09-21 16:54:36.47886+08	2020-09-21 16:54:36.47886+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	remove  role permission test 	delete	/api/v1/admin/removeRoleApiPerm	127.0.0.1:39556
44	2020-09-21 16:54:39.261845+08	2020-09-21 16:54:39.261845+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	add  role api perm test 	post	/api/v1/admin/addRoleApiPerm	127.0.0.1:39554
45	2020-09-21 17:02:36.672382+08	2020-09-21 17:02:36.672382+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  role test	put	/api/v1/admin/editRole	127.0.0.1:39802
46	2020-09-21 17:21:19.509018+08	2020-09-21 17:21:19.509018+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	edit  role test	put	/api/v1/admin/editRole	127.0.0.1:40750
\.


--
-- Data for Name: casbin_rule; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.casbin_rule (id, p_type, v0, v1, v2, v3, v4, v5) FROM stdin;
18	p	root	/api/v1/admin/addApiGroup	post			
19	p	root	/api/v1/admin/editApiGroup	put			
20	p	root	/api/v1/admin/getApiGroupList	get			
21	p	root	/api/v1/admin/removeApiGroup	delete			
23	p	root	/api/v1/admin/getMenuList	get			
24	p	root	/api/v1/admin/addMenu	post			
25	p	root	/api/v1/admin/delMenu	delete			
26	p	root	/api/v1/admin/menu	put			
27	p	root	/api/v1/admin/login	post			
28	p	root	/api/v1/user/profile	get			
29	p	root	/api/v1/user/modifyPwd	put			
31	p	root	/api/v1/admin/userTotal	get			
32	p	root	/api/v1/admin/deleteUser	delete			
33	p	root	/api/v1/admin/editUser	put			
34	p	root	/api/v1/admin/getUserKV	get			
35	p	root	/api/v1/captcha	post			
36	p	root	/api/v1/admin/getRoleList	get			
37	p	root	/api/v1/admin/getRoleLevel	get			
38	p	root	/api/v1/admin/editRole	put			
39	p	root	/api/v1/admin/removeRole	delete			
40	p	root	/api/v1/admin/addRole	post			
45	p	test	/api/v1/captcha	post			
48	p	test	/api/v1/admin/login	post			
30	p	root	/api/v1/admin/listUser	get			
49	p	root	/api/v1/admin/getRoleApiList	get	\N	\N	\N
50	p	root	/api/v1/admin/addRoleApiPerm	post	\N	\N	\N
51	p	test	/api/v1/admin/getRoleApiList	get			
52	p	root	/api/v1/admin/removeRoleApiPerm	delete	\N	\N	\N
53	p	root	/api/v1/admin/behaviorCount	get			
54	p	root	/api/v1/admin/getBehaviorList	get			
55	p	root	/api/v1/admin/removeBehaviors	delete			
57	p	root	 /api/v1/admin/removeRoleApiPerm	delete			
58	p	root	/api/v1/admin/getSettings	get			
59	p	root	/api/v1/admin/getApiGroupListAll	get			
60	p	test	/api/v1/admin/getApiGroupList	get			
61	p	test	/api/v1/admin/getApiGroupListAll	get			
62	p	test	/api/v1/admin/getMenuList	get			
63	p	test	/api/v1/user/profile	get			
64	p	test	/api/v1/admin/listUser	get			
65	p	test	/api/v1/admin/userTotal	get			
66	p	test	/api/v1/admin/getUserKV	get			
67	p	test	/api/v1/admin/getRoleList	get			
68	p	test	/api/v1/admin/getRoleLevel	get			
69	p	test	/api/v1/admin/behaviorCount	get			
70	p	test	/api/v1/admin/getBehaviorList	get			
71	p	test	/api/v1/admin/getSettings	get			
\.


--
-- Data for Name: menus; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.menus (id, created_at, updated_at, deleted_at, name, parent_id, path, component, sort, meta, hidden) FROM stdin;
2	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	profile	0	/layout/profile	profile/index	\N	{ "title":"Profile" ,"icon":"people" }	t
4	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	menu	3	/layout/permission/menu	permission/menu/index.vue	1	{"title":"Menu","icon":"menu"}	f
5	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	user	3	/layout/permission/user	permission/user/index	2	{"title":"User","icon":"user"}	f
6	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	tool	0	/layout/tool	tool/index	3	{"title":"Tool","icon":"tool"}	f
34	2020-08-31 09:25:32.284174+08	2020-08-31 09:25:32.284174+08	2020-08-31 09:36:07.592648+08	sss	1			0	{"icon":"","title":"","affix":""}	f
35	2020-08-31 09:43:01.701389+08	2020-08-31 09:43:01.701389+08	2020-08-31 09:43:11.653073+08		1			0	{"icon":"","title":"","affix":""}	f
37	2020-09-01 14:12:48.923728+08	2020-09-01 14:12:48.923728+08	2020-09-01 14:14:26.850974+08	permission	0	/layout/permission	permission/index.vue	2	{"icon":"permssion","title":"Permission","affix":""}	f
7	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	about	0	/layout/about	about/index	4	{"title":"About","icon":"about"}	f
3	2020-08-23 11:28:34.803501+08	2020-09-01 14:32:33.907162+08	\N	permission	0	/layout/permission	permission/index.vue	2	{"affix":"","icon":"permission","title":"Permission"}	f
41	2020-09-14 15:46:59.618985+08	2020-09-14 15:57:32.331893+08	\N	behavior	3	/layout/permission/behavior	permission/behavior/index	5	{"affix":"","icon":"behavior","title":"Behavor Record"}	f
36	2020-09-01 11:54:31.864752+08	2020-09-01 14:36:31.38394+08	\N	api	3	/layout/permission/api	permission/api/index	3	{"affix":"","icon":"api","title":"API"}	f
1	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	dashboard	0	/layout/dashboard	dashboard/index.vue	1	{ "title":"Dashboard" ,"icon":"dashboard",\n"affix":"true"}	f
38	2020-09-04 16:13:34.59235+08	2020-09-04 17:00:00.040999+08	\N	role	3	/layout/permission/role	permission/role/index	4	{"affix":"","icon":"role","title":"Role"}	f
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.roles (id, created_at, updated_at, deleted_at, role, level) FROM stdin;
1	2020-09-04 01:40:37.67+08	2020-09-04 01:40:39.398+08	\N	root	9
2	2020-09-08 17:38:54.705945+08	2020-09-21 17:21:19.310503+08	\N	test	5
\.


--
-- Data for Name: user_bases; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.user_bases (id, created_at, updated_at, deleted_at, uid, username, phone, login_pwd, pay_pwd, email, nickname, head_icon, sign, status, role_refer) FROM stdin;
3	2020-08-31 00:15:16.122+08	2020-09-01 11:23:50.851306+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	test333	13866686669	$2a$10$bHEKRlN4eLKVuuNJJYIGneJcES9vLgs/94ck9FFkfvJaPzG83he8K	\N	\N		./public/static/a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9.png		2	2
1	2020-08-23 11:28:34.803501+08	2020-09-04 14:16:41.434622+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	13899995555	$2a$10$bHEKRlN4eLKVuuNJJYIGneJcES9vLgs/94ck9FFkfvJaPzG83he8K				./public/static/9f0a39b0-8868-4f1f-a3d0-bc882e693ac0.png		1	1
6	2020-09-10 17:03:49.912896+08	2020-09-10 17:03:49.912896+08	\N	c60bbe71-9c87-45e0-96ee-aa1f24695002	test111	15655621514	$2a$10$FhACduB1R4O9JiKUDXUAMOpB8363LJqx5DEzH22.ZdFPWdbWYO0ey				./public/static/f56c1a47-0669-4238-ad9f-7851aff8dc42.png		1	2
7	2020-09-21 14:29:23.112699+08	2020-09-21 14:29:23.112699+08	\N	2a5e0067-de2b-4336-b3c4-f50ae2daecce	test222	18747246115	$2a$10$Ugj.Sf.Y6rST6jk6bGnbj.OPihhun7WTJXV4Tr8SeAlDICl.NpXPq						0	2
\.


--
-- Data for Name: user_cards; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.user_cards (id, created_at, updated_at, deleted_at, uid, name, card_id, issue_org, birthday, valid_period, card_id_addr, sex, nation) FROM stdin;
\.


--
-- Name: api_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.api_groups_id_seq', 33, true);


--
-- Name: behavior_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.behavior_records_id_seq', 46, true);


--
-- Name: casbin_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.casbin_rule_id_seq', 71, true);


--
-- Name: menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.menus_id_seq', 41, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.roles_id_seq', 2, true);


--
-- Name: user_bases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.user_bases_id_seq', 7, true);


--
-- Name: user_cards_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.user_cards_id_seq', 1, false);


--
-- Name: api_groups api_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.api_groups
    ADD CONSTRAINT api_groups_pkey PRIMARY KEY (id);


--
-- Name: behavior_records behavior_records_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.behavior_records
    ADD CONSTRAINT behavior_records_pkey PRIMARY KEY (id);


--
-- Name: casbin_rule casbin_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.casbin_rule
    ADD CONSTRAINT casbin_rule_pkey PRIMARY KEY (id);


--
-- Name: menus menus_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.menus
    ADD CONSTRAINT menus_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: user_bases user_bases_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.user_bases
    ADD CONSTRAINT user_bases_pkey PRIMARY KEY (id);


--
-- Name: user_cards user_cards_pkey; Type: CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.user_cards
    ADD CONSTRAINT user_cards_pkey PRIMARY KEY (id);


--
-- Name: idx_api_groups_deleted_at; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_api_groups_deleted_at ON public.api_groups USING btree (deleted_at);


--
-- Name: idx_api_groups_path_method; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_api_groups_path_method ON public.api_groups USING btree (method, path);


--
-- Name: idx_behavior_records_deleted_at; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_behavior_records_deleted_at ON public.behavior_records USING btree (deleted_at);


--
-- Name: idx_behavior_records_ip; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_behavior_records_ip ON public.behavior_records USING btree (ip);


--
-- Name: idx_behavior_records_method; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_behavior_records_method ON public.behavior_records USING btree (method);


--
-- Name: idx_behavior_records_path; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_behavior_records_path ON public.behavior_records USING btree (path);


--
-- Name: idx_behavior_records_uid; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_behavior_records_uid ON public.behavior_records USING btree (uid);


--
-- Name: idx_behavior_records_username; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_behavior_records_username ON public.behavior_records USING btree (username);


--
-- Name: idx_menus_deleted_at; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_menus_deleted_at ON public.menus USING btree (deleted_at);


--
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


--
-- Name: idx_roles_role; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_roles_role ON public.roles USING btree (role);


--
-- Name: idx_user_bases_deleted_at; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_user_bases_deleted_at ON public.user_bases USING btree (deleted_at);


--
-- Name: idx_user_bases_phone; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_user_bases_phone ON public.user_bases USING btree (phone);


--
-- Name: idx_user_bases_uid; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_user_bases_uid ON public.user_bases USING btree (uid);


--
-- Name: idx_user_bases_username; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_user_bases_username ON public.user_bases USING btree (username);


--
-- Name: idx_user_cards_card_id; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_user_cards_card_id ON public.user_cards USING btree (card_id);


--
-- Name: idx_user_cards_deleted_at; Type: INDEX; Schema: public; Owner: business
--

CREATE INDEX idx_user_cards_deleted_at ON public.user_cards USING btree (deleted_at);


--
-- Name: idx_user_cards_uid; Type: INDEX; Schema: public; Owner: business
--

CREATE UNIQUE INDEX idx_user_cards_uid ON public.user_cards USING btree (uid);


--
-- Name: user_bases fk_user_bases_role; Type: FK CONSTRAINT; Schema: public; Owner: business
--

ALTER TABLE ONLY public.user_bases
    ADD CONSTRAINT fk_user_bases_role FOREIGN KEY (role_refer) REFERENCES public.roles(id);


--
-- PostgreSQL database dump complete
--

