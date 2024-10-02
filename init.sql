--
-- PostgreSQL database dump
--

-- Dumped from database version 15.6 (Homebrew)
-- Dumped by pg_dump version 15.6 (Homebrew)

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
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    token text NOT NULL,
    data bytea NOT NULL,
    expiry timestamp with time zone NOT NULL
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- Name: snippets; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.snippets (
    id integer NOT NULL,
    title character varying(100) NOT NULL,
    content text NOT NULL,
    created timestamp with time zone DEFAULT CURRENT_DATE NOT NULL,
    expires date NOT NULL
);


ALTER TABLE public.snippets OWNER TO postgres;

--
-- Name: snippets_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.snippets_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.snippets_id_seq OWNER TO postgres;

--
-- Name: snippets_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.snippets_id_seq OWNED BY public.snippets.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    hashed_password character(60) NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: snippets id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.snippets ALTER COLUMN id SET DEFAULT nextval('public.snippets_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sessions (token, data, expiry) FROM stdin;
taWdEI_Rw1rw8zSz-DglYIcDS69uMt7Sq7Y2H_fB53E	\\x257f030102ff800001020108446561646c696e6501ff8200010656616c75657301ff8400000010ff810501010454696d6501ff8200000027ff83040101176d61705b737472696e675d696e74657266616365207b7d01ff8400010c0110000032ff80010f010000000eddfc7fc20a44c8c0ffff01011361757468656e7469636174656455736572494403696e740402000a00	2024-06-12 22:50:58.17228-07
\.


--
-- Data for Name: snippets; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.snippets (id, title, content, created, expires) FROM stdin;
26	The Way	~ Accept everything just the way it is.\r\n\r\n~ Do not, under any circumstances, depend on a partial feeling.\r\n\r\n~ Think lightly of yourself and deeply of the world.\r\n\r\n~ Do not regret what you have done.\r\n\r\n~ Never stray from the Way.\r\n\r\nMiyamoto Musashi\r\n	2024-06-12 10:54:07.45178-07	2025-06-12
27	The Flow	~ The usefulness of the pot lies in its emptiness.\r\n\r\n~ To the mind that is still, the whole universe surrenders.\r\n\r\n~ When I let go of what I am, I become what I might be.\r\n\r\n\r\nLao Tzu	2024-06-12 11:02:33.19108-07	2025-06-12
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, hashed_password, created) FROM stdin;
1	Test_User	test.test@gmail.cc	$2a$12$ITmZpyim93BBTdFuYhHA8ODRQWjrlGGi/DpwyMEmd/P4XGchD6PeK	2024-02-27 13:06:09.251212-08
2	user2	user2@example.com	$2a$12$zhFYInVPEMkzay07g4KCI.UfvW9J0oLSr9L9y8RxKSzrXvHqHI1nW	2024-02-27 14:34:11.421628-08
4	user	user@user.com	$2a$12$dq2qHlHdsd3pyMq7qo0dIOBgKE6/5ZB0gWc0KoJyMHhwio68Q9Y7u	2024-03-02 15:46:45.762476-08
5	Musashi	hello@yesme.dev	$2a$12$69ZKHgY0edCvn1r1CGQzEOYIy3hIseDQWPBXeaY6kQxEW/PrYNeDe	2024-06-12 10:50:36.231668-07
\.


--
-- Name: snippets_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.snippets_id_seq', 27, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 5, true);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);


--
-- Name: snippets snippets_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.snippets
    ADD CONSTRAINT snippets_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_uc_email; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_uc_email UNIQUE (email);


--
-- Name: idx_snippets_created; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_snippets_created ON public.snippets USING btree (created);


--
-- Name: sessions_expiry_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX sessions_expiry_idx ON public.sessions USING btree (expiry);


--
-- Name: TABLE sessions; Type: ACL; Schema: public; Owner: postgres
--

GRANT ALL ON TABLE public.sessions TO web;


--
-- Name: TABLE snippets; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT,DELETE,UPDATE ON TABLE public.snippets TO web;


--
-- Name: SEQUENCE snippets_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.snippets_id_seq TO web;


--
-- Name: TABLE users; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,INSERT ON TABLE public.users TO web;


--
-- Name: SEQUENCE users_id_seq; Type: ACL; Schema: public; Owner: postgres
--

GRANT SELECT,USAGE ON SEQUENCE public.users_id_seq TO web;


--
-- PostgreSQL database dump complete
--

