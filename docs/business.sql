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
    path character varying(24),
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
15	2020-09-03 11:39:29.062578+08	2020-09-03 11:39:29.062578+08	\N	/api/v1/admin/userList	获取用户列表	user	get
17	2020-09-03 14:21:04.330658+08	2020-09-03 14:21:04.330658+08	\N	/api/v1/admin/deleteUser	删除用户	user	delete
19	2020-09-03 14:22:09.017293+08	2020-09-03 14:22:28.815854+08	\N	/api/v1/admin/getUserKV	获取用户一些状态信息	user	get
18	2020-09-03 14:21:22.647945+08	2020-09-03 14:22:44.145525+08	\N	/api/v1/admin/editUser	编辑用户	user	put
20	2020-09-08 18:00:04.217012+08	2020-09-08 18:00:04.217012+08	\N	/api/v1/admin/getRoleList	获取角色列表	role	get
21	2020-09-08 18:00:47.838141+08	2020-09-08 18:00:47.838141+08	\N	/api/v1/admin/getRoleLevel	获取角色等级列表	role	get
22	2020-09-08 18:01:18.409594+08	2020-09-08 18:01:18.409594+08	\N	/api/v1/admin/editRole	编辑角色	role	put
23	2020-09-08 18:01:47.348512+08	2020-09-08 18:01:47.348512+08	\N	/api/v1/admin/removeRole	删除角色	role	delete
24	2020-09-08 18:02:18.796049+08	2020-09-08 18:02:18.796049+08	\N	/api/v1/admin/addRole	增加角色	role	post
16	2020-09-03 14:19:24.614455+08	2020-09-08 20:19:21.486459+08	\N	/api/v1/admin/userTotal	获取用户总数	user	get
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
30	p	root	/api/v1/admin/userList	get			
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
41	p	test	/api/v1/admin/getMenuList	get			
42	p	test	/api/v1/admin/addMenu	post			
43	p	test	/api/v1/admin/delMenu	delete			
44	p	test	/api/v1/admin/menu	put			
45	p	test	/api/v1/captcha	post			
46	p	test	/api/v1/admin/addApiGroup	post			
47	p	test	/api/v1/admin/editApiGroup	put			
48	p	test	/api/v1/admin/login	post			
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
36	2020-09-01 11:54:31.864752+08	2020-09-01 14:36:31.38394+08	\N	api	3	/layout/permission/api	permission/api/index	3	{"affix":"","icon":"api","title":"API"}	f
1	2020-08-23 11:28:34.803501+08	2020-08-23 11:28:34.803501+08	\N	dashboard	0	/layout/dashboard	dashboard/index.vue	1	{ "title":"Dashboard" ,"icon":"dashboard",\n"affix":"true"}	f
38	2020-09-04 16:13:34.59235+08	2020-09-04 17:00:00.040999+08	\N	role	3	/layout/permission/role	permission/role/index	4	{"affix":"","icon":"role","title":"Role"}	f
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.roles (id, created_at, updated_at, deleted_at, role, level) FROM stdin;
1	2020-09-04 01:40:37.67+08	2020-09-04 01:40:39.398+08	\N	root	9
2	2020-09-08 17:38:54.705945+08	2020-09-08 17:51:00.094054+08	\N	test	3
\.


--
-- Data for Name: user_bases; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.user_bases (id, created_at, updated_at, deleted_at, uid, username, phone, login_pwd, pay_pwd, email, nickname, head_icon, sign, status, role_refer) FROM stdin;
3	2020-08-31 00:15:16.122+08	2020-09-01 11:23:50.851306+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75wexdf	test	13866686669	$2a$10$bHEKRlN4eLKVuuNJJYIGneJcES9vLgs/94ck9FFkfvJaPzG83he8K	\N	\N	\N	./public/static/a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9.png	\N	2	2
6	2020-09-10 17:03:49.912896+08	2020-09-10 17:03:49.912896+08	\N	c60bbe71-9c87-45e0-96ee-aa1f24695002	test111	15655621514	$2a$10$FhACduB1R4O9JiKUDXUAMOpB8363LJqx5DEzH22.ZdFPWdbWYO0ey						0	2
1	2020-08-23 11:28:34.803501+08	2020-09-04 14:16:41.434622+08	\N	a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9	admin	13899995555	$2a$10$bHEKRlN4eLKVuuNJJYIGneJcES9vLgs/94ck9FFkfvJaPzG83he8K				./public/static/a47306ce-1d4e-495c-aa1a-c7e3f75cdcd9.png		1	1
\.


--
-- Data for Name: user_cards; Type: TABLE DATA; Schema: public; Owner: business
--

COPY public.user_cards (id, created_at, updated_at, deleted_at, uid, name, card_id, issue_org, birthday, valid_period, card_id_addr, sex, nation) FROM stdin;
\.


--
-- Name: api_groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.api_groups_id_seq', 24, true);


--
-- Name: casbin_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.casbin_rule_id_seq', 48, true);


--
-- Name: menus_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.menus_id_seq', 40, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.roles_id_seq', 2, true);


--
-- Name: user_bases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: business
--

SELECT pg_catalog.setval('public.user_bases_id_seq', 6, true);


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

