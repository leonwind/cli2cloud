--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
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
-- Name: cli_storage; Type: TABLE; Schema: public; Owner: cli2cloud
--

CREATE TABLE IF NOT EXISTS public.cli_storage (
    clientid text NOT NULL,
    content text NOT NULL,
    line integer NOT NULL
);


ALTER TABLE public.cli_storage OWNER TO cli2cloud;

--
-- Data for Name: cli_storage; Type: TABLE DATA; Schema: public; Owner: cli2cloud
--

COPY public.cli_storage (clientid, content, line) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
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
-- Name: cli_storage; Type: TABLE; Schema: public; Owner: cli2cloud
--

CREATE TABLE IF NOT EXISTS public.cli_storage (
    clientid text NOT NULL,
    content text NOT NULL,
    line integer NOT NULL
);


ALTER TABLE public.cli_storage OWNER TO cli2cloud;

--
-- Name: clients; Type: TABLE; Schema: public; Owner: cli2cloud
--

CREATE TABLE public.clients (
    id text NOT NULL,
    encrypted boolean NOT NULL,
    salt text,
    iv text,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.clients OWNER TO cli2cloud;

--
-- Data for Name: cli_storage; Type: TABLE DATA; Schema: public; Owner: cli2cloud
--

COPY public.cli_storage (clientid, content, line) FROM stdin;
\.


--
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: cli2cloud
--

COPY public.clients (id, encrypted, salt, iv, created) FROM stdin;
\.


--
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: cli2cloud
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (id);


--
-- Name: cli_storage cli_storage_clientid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: cli2cloud
--

ALTER TABLE ONLY public.cli_storage
    ADD CONSTRAINT cli_storage_clientid_fkey FOREIGN KEY (clientid) REFERENCES public.clients(id);


--
-- PostgreSQL database dump complete
--

