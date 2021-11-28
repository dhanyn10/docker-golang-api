--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1 (Ubuntu 14.1-1.pgdg18.04+1)
-- Dumped by pg_dump version 14.1 (Ubuntu 14.1-1.pgdg18.04+1)

-- Started on 2021-11-28 20:40:30 WIB

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
-- TOC entry 210 (class 1259 OID 16388)
-- Name: article; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.article (
    id bigint NOT NULL,
    author character varying(250) NOT NULL,
    title character varying(250) NOT NULL,
    body character varying(250) NOT NULL,
    created timestamp without time zone NOT NULL
);


ALTER TABLE public.article OWNER TO admin;

--
-- TOC entry 209 (class 1259 OID 16387)
-- Name: article_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.article_id_seq OWNER TO admin;

--
-- TOC entry 3291 (class 0 OID 0)
-- Dependencies: 209
-- Name: article_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.article_id_seq OWNED BY public.article.id;


--
-- TOC entry 3142 (class 2604 OID 16391)
-- Name: article id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.article ALTER COLUMN id SET DEFAULT nextval('public.article_id_seq'::regclass);


--
-- TOC entry 3285 (class 0 OID 16388)
-- Dependencies: 210
-- Data for Name: article; Type: TABLE DATA; Schema: public; Owner: admin
--

INSERT INTO public.article (id, author, title, body, created) VALUES (20, 'suparman', 'perjuangan', 'perang dunia dua menuju akhir', '2021-11-28 15:59:09.658958');
INSERT INTO public.article (id, author, title, body, created) VALUES (21, 'jostein', 'dunia suphie', 'novel tentang dunia filsafat untuk anak-anak', '2021-11-28 16:13:20.703631');
INSERT INTO public.article (id, author, title, body, created) VALUES (22, 'jostein', 'dunia suphie', 'novel tentang dunia filsafat', '2021-11-28 20:33:34.986473');


--
-- TOC entry 3292 (class 0 OID 0)
-- Dependencies: 209
-- Name: article_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.article_id_seq', 22, true);


--
-- TOC entry 3144 (class 2606 OID 16395)
-- Name: article article_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_pkey PRIMARY KEY (id);


-- Completed on 2021-11-28 20:40:30 WIB

--
-- PostgreSQL database dump complete
--

