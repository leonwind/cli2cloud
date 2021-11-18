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
-- Name: core_storage; Type: TABLE; Schema: public; Owner: admin.
--

CREATE TABLE public.core_storage (
    clientid text NOT NULL,
    "row" integer NOT NULL,
    content text NOT NULL,
    encrypted boolean NOT NULL,
    created timestamp with time zone NOT NULL
);


ALTER TABLE public.core_storage OWNER TO "admin";

--
-- Data for Name: core_storage; Type: TABLE DATA; Schema: public; Owner: leon.windheuser
--

COPY public.core_storage (clientid, "row", content, encrypted, created) FROM stdin;
\.


--
-- PostgreSQL database dump complete
--

