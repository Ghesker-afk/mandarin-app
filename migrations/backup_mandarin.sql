--
-- PostgreSQL database dump
--

\restrict pgiLrbrNwA9QRb0iJf3oxkLHpBbb75d4i7NjYWOONVtek29CuwUMeXh4EwVg83g

-- Dumped from database version 15.19
-- Dumped by pg_dump version 15.19

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
-- Name: ideograms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ideograms (
    id integer NOT NULL,
    "character" character varying(10) NOT NULL,
    pinyin character varying(50) NOT NULL,
    meaning text NOT NULL,
    example_sentence text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    meanings text[] DEFAULT '{}'::text[] NOT NULL,
    radical character varying(10),
    stroke_count integer
);


ALTER TABLE public.ideograms OWNER TO postgres;

--
-- Name: ideograms_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ideograms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ideograms_id_seq OWNER TO postgres;

--
-- Name: ideograms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ideograms_id_seq OWNED BY public.ideograms.id;


--
-- Name: ideograms id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ideograms ALTER COLUMN id SET DEFAULT nextval('public.ideograms_id_seq'::regclass);


--
-- Data for Name: ideograms; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ideograms (id, "character", pinyin, meaning, example_sentence, created_at, meanings, radical, stroke_count) FROM stdin;
1	??	Ni hao	Ol�	??,???????	2026-10-11 02:58:03.628756	{}	\N	\N
\.


--
-- Name: ideograms_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ideograms_id_seq', 1, true);


--
-- Name: ideograms ideograms_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ideograms
    ADD CONSTRAINT ideograms_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

\unrestrict pgiLrbrNwA9QRb0iJf3oxkLHpBbb75d4i7NjYWOONVtek29CuwUMeXh4EwVg83g

