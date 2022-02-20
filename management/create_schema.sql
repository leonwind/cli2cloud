--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
-- Name: cli_output; Type: TABLE; Schema: public; Owner: leon.windheuser
--

CREATE TABLE public.cli_output (
    userid text NOT NULL,
    "row" integer NOT NULL,
    content text NOT NULL
);


ALTER TABLE public.cli_output OWNER TO "leon.windheuser";

--
-- Name: users; Type: TABLE; Schema: public; Owner: leon.windheuser
--

CREATE TABLE public.users (
    id text NOT NULL,
    encrypted boolean NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.users OWNER TO "leon.windheuser";

--
-- Data for Name: cli_output; Type: TABLE DATA; Schema: public; Owner: leon.windheuser
--

COPY public.cli_output (userid, "row", content) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: leon.windheuser
--

COPY public.users (id, encrypted, created) FROM stdin;
\.


--
-- Name: cli_output cli_output_pkey; Type: CONSTRAINT; Schema: public; Owner: leon.windheuser
--

ALTER TABLE ONLY public.cli_output
    ADD CONSTRAINT cli_output_pkey PRIMARY KEY (userid);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: leon.windheuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

