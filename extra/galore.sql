--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: todos; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    content text,
    user_id integer,
    created timestamp without time zone,
    modified timestamp without time zone
);


ALTER TABLE todos OWNER TO postgres;

--
-- Name: user_token; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE user_token (
    user_id integer,
    token character varying(1024),
    created timestamp without time zone,
    modified timestamp without time zone
);


ALTER TABLE user_token OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres; Tablespace: 
--

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    user_name character varying(255),
    password character varying(1024),
    created timestamp without time zone,
    modified timestamp without time zone
);


ALTER TABLE users OWNER TO postgres;

--
-- Data for Name: todos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY todos (id, content, user_id, created, modified) FROM stdin;
\.


--
-- Data for Name: user_token; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY user_token (user_id, token, created, modified) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY users (user_id, user_name, password, created, modified) FROM stdin;
\.


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--