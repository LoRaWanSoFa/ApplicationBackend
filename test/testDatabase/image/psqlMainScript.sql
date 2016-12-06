--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

-- Started on 2016-12-01 12:44:40


CREATE DATABASE lorawan;
\c lorawan

CREATE ROLE admin LOGIN
  SUPERUSER INHERIT CREATEDB CREATEROLE NOREPLICATION;
  
CREATE USER docker WITH SUPERUSER PASSWORD 'docker';

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 1 (class 3079 OID 12393)
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- TOC entry 2217 (class 0 OID 0)
-- Dependencies: 1
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 189 (class 1259 OID 16441)
-- Name: gatewaynodes; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE gatewaynodes (
    gatewayaddress character varying(20) NOT NULL,
    deveui character varying(20) NOT NULL
);


ALTER TABLE gatewaynodes OWNER TO docker;

--
-- TOC entry 188 (class 1259 OID 16436)
-- Name: gateways; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE gateways (
    gatewayaddress character varying(20) NOT NULL,
    region character varying(10)
);


ALTER TABLE gateways OWNER TO docker;

--
-- TOC entry 198 (class 1259 OID 16830)
-- Name: message_payloads; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE message_payloads (
    message_id bigint NOT NULL,
    sensor_id bigint,
    payload_order smallint DEFAULT '1'::smallint NOT NULL,
    payload character(255) NOT NULL
);


ALTER TABLE message_payloads OWNER TO docker;

--
-- TOC entry 187 (class 1259 OID 16394)
-- Name: messages; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE messages (
    id bigint NOT NULL,
    deveui character varying(20) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    down boolean NOT NULL
);


ALTER TABLE messages OWNER TO docker;

--
-- TOC entry 186 (class 1259 OID 16392)
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE messages_id_seq OWNER TO docker;

--
-- TOC entry 2218 (class 0 OID 0)
-- Dependencies: 186
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE messages_id_seq OWNED BY messages.id;


--
-- TOC entry 190 (class 1259 OID 16538)
-- Name: migrations; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE migrations (
    migration character varying(255) NOT NULL,
    batch integer NOT NULL
);


ALTER TABLE migrations OWNER TO docker;

--
-- TOC entry 185 (class 1259 OID 16387)
-- Name: nodes; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE nodes (
    deveui character varying(20) NOT NULL,
    devaddr character varying(20) NOT NULL,
    appskey character varying(20) NOT NULL,
    nwkskey character varying(20) NOT NULL,
    operationtype character(1) NOT NULL,
    "interval" integer,
    latitude numeric(10,6) NOT NULL,
    longitude numeric(10,6) NOT NULL
);


ALTER TABLE nodes OWNER TO docker;

--
-- TOC entry 193 (class 1259 OID 16554)
-- Name: password_resets; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE password_resets (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);


ALTER TABLE password_resets OWNER TO docker;

--
-- TOC entry 197 (class 1259 OID 16814)
-- Name: sensors; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE sensors (
    id bigint NOT NULL,
    sensortype_id bigint NOT NULL,
    deveui character(20) NOT NULL,
    io_type integer NOT NULL,
    io_address integer NOT NULL,
    number_of_values integer NOT NULL,
    lenght_of_values integer NOT NULL,
    header_order integer NOT NULL
);


ALTER TABLE sensors OWNER TO docker;

--
-- TOC entry 196 (class 1259 OID 16812)
-- Name: sensors_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE sensors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sensors_id_seq OWNER TO docker;

--
-- TOC entry 2219 (class 0 OID 0)
-- Dependencies: 196
-- Name: sensors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE sensors_id_seq OWNED BY sensors.id;


--
-- TOC entry 195 (class 1259 OID 16803)
-- Name: sensortypes; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE sensortypes (
    id bigint NOT NULL,
    description character varying(255) NOT NULL,
    conversion_expression character varying(255) NOT NULL,
    data_type integer DEFAULT 0 NOT NULL
);


ALTER TABLE sensortypes OWNER TO docker;

--
-- TOC entry 194 (class 1259 OID 16801)
-- Name: sensortypes_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE sensortypes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sensortypes_id_seq OWNER TO docker;

--
-- TOC entry 2220 (class 0 OID 0)
-- Dependencies: 194
-- Name: sensortypes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE sensortypes_id_seq OWNED BY sensortypes.id;


--
-- TOC entry 192 (class 1259 OID 16543)
-- Name: users; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE users (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    remember_token character varying(100),
    created_at timestamp(0) without time zone,
    updated_at timestamp(0) without time zone
);


ALTER TABLE users OWNER TO docker;

--
-- TOC entry 191 (class 1259 OID 16541)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO docker;

--
-- TOC entry 2221 (class 0 OID 0)
-- Dependencies: 191
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- TOC entry 2050 (class 2604 OID 16397)
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY messages ALTER COLUMN id SET DEFAULT nextval('messages_id_seq'::regclass);


--
-- TOC entry 2054 (class 2604 OID 16817)
-- Name: sensors id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors ALTER COLUMN id SET DEFAULT nextval('sensors_id_seq'::regclass);


--
-- TOC entry 2052 (class 2604 OID 16806)
-- Name: sensortypes id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensortypes ALTER COLUMN id SET DEFAULT nextval('sensortypes_id_seq'::regclass);


--
-- TOC entry 2051 (class 2604 OID 16546)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- TOC entry 2201 (class 0 OID 16441)
-- Dependencies: 189
-- Data for Name: gatewaynodes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY gatewaynodes (gatewayaddress, deveui) FROM stdin;
GATE123	A4C12BF
\.


--
-- TOC entry 2200 (class 0 OID 16436)
-- Dependencies: 188
-- Data for Name: gateways; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY gateways (gatewayaddress, region) FROM stdin;
GATE123	EUW123
\.


--
-- TOC entry 2210 (class 0 OID 16830)
-- Dependencies: 198
-- Data for Name: message_payloads; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY message_payloads (message_id, sensor_id, payload_order, payload) FROM stdin;
\.


--
-- TOC entry 2199 (class 0 OID 16394)
-- Dependencies: 187
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY messages (id, deveui, created_at, down) FROM stdin;
1	A4C12BF	2016-11-17 00:00:00	t
2	A4C12BF	2016-11-24 12:15:52.090433	f
3	A4C12BF	2016-11-24 12:16:17.618165	f
4	A4C12BF	2016-11-24 12:16:38.193585	f
7	A4C12BF	2016-11-24 20:47:00.166488	f
8	A4C12BF	2016-11-24 20:48:40.862602	f
10	A4C12BF	2016-11-24 20:49:40.171542	f
12	A4C12BF	2016-11-24 20:51:22.052245	f
14	A4C12BF	2016-11-24 20:51:51.346991	f
16	A4C12BF	2016-11-24 20:56:48.12764	f
18	A4C12BF	2016-11-24 21:00:47.566878	f
20	A4C12BF	2016-11-24 21:03:22.77461	f
22	A4C12BF	2016-11-24 21:03:45.717692	f
24	A4C12BF	2016-11-24 21:04:02.74158	f
26	A4C12BF	2016-11-29 13:05:48.602461	f
28	A4C12BF	2016-12-01 08:56:02.613455	f
30	A4C12BF	2016-12-01 08:56:50.405866	f
32	A4C12BF	2016-12-01 09:03:54.637014	f
34	A4C12BF	2016-12-01 09:05:35.041351	f
36	A4C12BF	2016-12-01 09:06:46.582832	f
38	A4C12BF	2016-12-01 09:07:04.813297	f
40	A4C12BF	2016-12-01 09:07:54.440649	f
42	A4C12BF	2016-12-01 09:09:54.1336	f
44	A4C12BF	2016-12-01 09:37:54.663616	f
46	A4C12BF	2016-12-01 09:38:31.627896	f
48	A4C12BF	2016-12-01 09:38:31.713778	f
50	A4C12BF	2016-12-01 09:50:00.276057	f
52	A4C12BF	2016-12-01 09:50:47.240952	f
\.


--
-- TOC entry 2222 (class 0 OID 0)
-- Dependencies: 186
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('messages_id_seq', 53, true);


--
-- TOC entry 2202 (class 0 OID 16538)
-- Dependencies: 190
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY migrations (migration, batch) FROM stdin;
2014_10_12_000000_create_users_table	1
2014_10_12_100000_create_password_resets_table	1
2016_10_11_141724_added_admin_user	1
2016_10_20_113141_alter_table_messages_date	1
2016_10_20_114540_alter_table_messages_date_type	1
2016_11_17_132212_drop_table_sensornodes	1
2016_11_17_133136_drop_table_sensor	2
2016_11_17_142048_recreate_sensorstypes	3
2016_11_17_144908_create_table_sensors	3
2016_11_18_091837_alter_table_messages	3
2016_11_18_101804_create_table_messagePayload	3
2016_11_24_085957_alter_table_sensor_types_add_datatype	4
\.


--
-- TOC entry 2197 (class 0 OID 16387)
-- Dependencies: 185
-- Data for Name: nodes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY nodes (deveui, devaddr, appskey, nwkskey, operationtype, "interval", latitude, longitude) FROM stdin;
A4C12BF	1345ABCD	70B426335	1325746823	A	5	23.456700	54.124200
\.


--
-- TOC entry 2205 (class 0 OID 16554)
-- Dependencies: 193
-- Data for Name: password_resets; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY password_resets (email, token, created_at) FROM stdin;
\.


--
-- TOC entry 2209 (class 0 OID 16814)
-- Dependencies: 197
-- Data for Name: sensors; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY sensors (id, sensortype_id, deveui, io_type, io_address, number_of_values, lenght_of_values, header_order) FROM stdin;
1	1	A4C12BF             	1	1	1	1	1
2	2	A4C12BF             	1	2	2	2	2
\.


--
-- TOC entry 2223 (class 0 OID 0)
-- Dependencies: 196
-- Name: sensors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('sensors_id_seq', 1, false);


--
-- TOC entry 2207 (class 0 OID 16803)
-- Dependencies: 195
-- Data for Name: sensortypes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY sensortypes (id, description, conversion_expression, data_type) FROM stdin;
1	This type of sensor is responsible for making Sander happy.	420 is life	0
2	TEMP	+33.8	0
\.


--
-- TOC entry 2224 (class 0 OID 0)
-- Dependencies: 194
-- Name: sensortypes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('sensortypes_id_seq', 2, true);


--
-- TOC entry 2204 (class 0 OID 16543)
-- Dependencies: 192
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY users (id, name, email, password, remember_token, created_at, updated_at) FROM stdin;
1	Admin	admin@admin.com	$2y$10$DeBNcbsdmUhS91/bp3QKrOiOfin6rUvZMoI789T5O5XBjfu5L1pCm	\N	2016-11-17 14:41:31	2016-11-17 14:41:31
\.


--
-- TOC entry 2225 (class 0 OID 0)
-- Dependencies: 191
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('users_id_seq', 1, true);


--
-- TOC entry 2061 (class 2606 OID 16440)
-- Name: gateways gateways_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY gateways
    ADD CONSTRAINT gateways_pkey PRIMARY KEY (gatewayaddress);


--
-- TOC entry 2059 (class 2606 OID 16399)
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- TOC entry 2057 (class 2606 OID 16391)
-- Name: nodes nodes_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY nodes
    ADD CONSTRAINT nodes_pkey PRIMARY KEY (deveui);


--
-- TOC entry 2071 (class 2606 OID 16819)
-- Name: sensors sensors_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors
    ADD CONSTRAINT sensors_pkey PRIMARY KEY (id);


--
-- TOC entry 2069 (class 2606 OID 16811)
-- Name: sensortypes sensortypes_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensortypes
    ADD CONSTRAINT sensortypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2063 (class 2606 OID 16553)
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- TOC entry 2065 (class 2606 OID 16551)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2072 (class 1259 OID 16844)
-- Name: message_payloads_message_id_sensor_id_index; Type: INDEX; Schema: public; Owner: docker
--

CREATE INDEX message_payloads_message_id_sensor_id_index ON message_payloads USING btree (message_id, sensor_id);


--
-- TOC entry 2066 (class 1259 OID 16560)
-- Name: password_resets_email_index; Type: INDEX; Schema: public; Owner: docker
--

CREATE INDEX password_resets_email_index ON password_resets USING btree (email);


--
-- TOC entry 2067 (class 1259 OID 16561)
-- Name: password_resets_token_index; Type: INDEX; Schema: public; Owner: docker
--

CREATE INDEX password_resets_token_index ON password_resets USING btree (token);


--
-- TOC entry 2074 (class 2606 OID 16444)
-- Name: gatewaynodes gatewaynode_deveui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY gatewaynodes
    ADD CONSTRAINT gatewaynode_deveui_fkey FOREIGN KEY (deveui) REFERENCES nodes(deveui) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2075 (class 2606 OID 16449)
-- Name: gatewaynodes gatewaynode_gatewayaddress_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY gatewaynodes
    ADD CONSTRAINT gatewaynode_gatewayaddress_fkey FOREIGN KEY (gatewayaddress) REFERENCES gateways(gatewayaddress) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2078 (class 2606 OID 16834)
-- Name: message_payloads message_payloads_message_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY message_payloads
    ADD CONSTRAINT message_payloads_message_id_foreign FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE;


--
-- TOC entry 2079 (class 2606 OID 16839)
-- Name: message_payloads message_payloads_sensor_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY message_payloads
    ADD CONSTRAINT message_payloads_sensor_id_foreign FOREIGN KEY (sensor_id) REFERENCES sensors(id) ON DELETE CASCADE;


--
-- TOC entry 2073 (class 2606 OID 16400)
-- Name: messages messages_deveui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_deveui_fkey FOREIGN KEY (deveui) REFERENCES nodes(deveui) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2076 (class 2606 OID 16820)
-- Name: sensors sensors_deveui_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors
    ADD CONSTRAINT sensors_deveui_foreign FOREIGN KEY (deveui) REFERENCES nodes(deveui) ON DELETE CASCADE;


--
-- TOC entry 2077 (class 2606 OID 16825)
-- Name: sensors sensors_sensortype_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors
    ADD CONSTRAINT sensors_sensortype_id_foreign FOREIGN KEY (sensortype_id) REFERENCES sensortypes(id) ON DELETE CASCADE;


-- Completed on 2016-12-01 12:44:40

--
-- PostgreSQL database dump complete
--
INSERT INTO nodes
VALUES ('00000000ABCDEF12','ABCDEF12','70B426334','1325746822','C',6,40.730610,-73.935242);

