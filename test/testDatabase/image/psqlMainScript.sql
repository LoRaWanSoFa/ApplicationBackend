--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

-- Started on 2016-12-15 12:20:58

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
-- TOC entry 2218 (class 0 OID 0)
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
-- TOC entry 198 (class 1259 OID 16723)
-- Name: message_payloads; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE message_payloads (
    message_id bigint NOT NULL,
    sensor_id bigint,
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
-- TOC entry 2219 (class 0 OID 0)
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
-- TOC entry 193 (class 1259 OID 16682)
-- Name: password_resets; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE password_resets (
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp(0) without time zone
);


ALTER TABLE password_resets OWNER TO docker;

--
-- TOC entry 197 (class 1259 OID 16707)
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
    header_order integer NOT NULL,
    soft_deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE sensors OWNER TO docker;

--
-- TOC entry 196 (class 1259 OID 16705)
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
-- TOC entry 2220 (class 0 OID 0)
-- Dependencies: 196
-- Name: sensors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE sensors_id_seq OWNED BY sensors.id;


--
-- TOC entry 195 (class 1259 OID 16696)
-- Name: sensortypes; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE sensortypes (
    id bigint NOT NULL,
    description character varying(255) NOT NULL,
    conversion_expression character varying(255) NOT NULL,
    data_type integer DEFAULT '-1'::integer NOT NULL,
    sensor_type integer DEFAULT 0 NOT NULL
);


ALTER TABLE sensortypes OWNER TO docker;

--
-- TOC entry 194 (class 1259 OID 16694)
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
-- TOC entry 2221 (class 0 OID 0)
-- Dependencies: 194
-- Name: sensortypes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE sensortypes_id_seq OWNED BY sensortypes.id;


--
-- TOC entry 192 (class 1259 OID 16671)
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
-- TOC entry 191 (class 1259 OID 16669)
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
-- TOC entry 2222 (class 0 OID 0)
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
-- TOC entry 2055 (class 2604 OID 16710)
-- Name: sensors id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors ALTER COLUMN id SET DEFAULT nextval('sensors_id_seq'::regclass);


--
-- TOC entry 2052 (class 2604 OID 16699)
-- Name: sensortypes id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensortypes ALTER COLUMN id SET DEFAULT nextval('sensortypes_id_seq'::regclass);


--
-- TOC entry 2051 (class 2604 OID 16674)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- TOC entry 2202 (class 0 OID 16441)
-- Dependencies: 189
-- Data for Name: gatewaynodes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY gatewaynodes (gatewayaddress, deveui) FROM stdin;
GATE123	A4C12BF
\.


--
-- TOC entry 2201 (class 0 OID 16436)
-- Dependencies: 188
-- Data for Name: gateways; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY gateways (gatewayaddress, region) FROM stdin;
GATE123	EUW123
\.


--
-- TOC entry 2211 (class 0 OID 16723)
-- Dependencies: 198
-- Data for Name: message_payloads; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY message_payloads (message_id, sensor_id, payload) FROM stdin;
1	1	Test Payload 1                                                                                                                                                                                                                                                 
2	1	Test Payload 2                                                                                                                                                                                                                                                 
3	1	Test Payload 3                                                                                                                                                                                                                                                 
2	1	rain is comming                                                                                                                                                                                                                                                
251	1	Howdee1                                                                                                                                                                                                                                                        
251	2	Howdee2                                                                                                                                                                                                                                                        
252	\N	DOWNLINKMESSAGE                                                                                                                                                                                                                                                
\.


--
-- TOC entry 2200 (class 0 OID 16394)
-- Dependencies: 187
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY messages (id, deveui, created_at, down) FROM stdin;
1	Y4C75XD	2016-12-10 00:00:00	f
2	A4C12BF	2016-11-15 00:00:00	f
3	A4C12BF	2016-11-17 00:00:00	f
250	A4C12BF	2016-12-15 11:10:59.162217	f
251	A4C12BF	2016-12-15 11:10:59.171424	f
252	A4C12BF	2016-12-15 11:10:59	t
\.


--
-- TOC entry 2223 (class 0 OID 0)
-- Dependencies: 186
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('messages_id_seq', 252, true);


--
-- TOC entry 2203 (class 0 OID 16538)
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
2016_11_17_133136_drop_table_sensor	1
2016_11_17_142048_recreate_sensorstypes	1
2016_11_17_144908_create_table_sensors	1
2016_11_18_091837_alter_table_messages	1
2016_11_18_101804_create_table_messagePayload	1
2016_11_24_085957_alter_table_sensor_types_add_datatype	1
2016_12_09_105238_messagePayloadDropOrder	1
2016_12_09_105302_sensorsAddSoftDelete	1
2016_12_13_085143_addSensortypecolumn	2
2016_12_15_104219_setDefaulDatatypeSensorType	3
\.


--
-- TOC entry 2198 (class 0 OID 16387)
-- Dependencies: 185
-- Data for Name: nodes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY nodes (deveui, devaddr, appskey, nwkskey, operationtype, "interval", latitude, longitude) FROM stdin;
A4C12BF	1345ABCD	70B426335	1325746823	A	5	23.456700	54.124200
Y4C75XD	1345ABCD	70B426335	1325746823	A	5	23.456700	54.124200
00000000ABCDEF12	ABCDEF12	70B426334	1325746822	C	6	40.730610	-73.935242
00000000AF1294E5	AF1294E5	70B426333	1325746823	C	5	23.456700	54.124200
\.


--
-- TOC entry 2206 (class 0 OID 16682)
-- Dependencies: 193
-- Data for Name: password_resets; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY password_resets (email, token, created_at) FROM stdin;
\.


--
-- TOC entry 2210 (class 0 OID 16707)
-- Dependencies: 197
-- Data for Name: sensors; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY sensors (id, sensortype_id, deveui, io_type, io_address, number_of_values, lenght_of_values, header_order, soft_deleted) FROM stdin;
1	1	A4C12BF             	1	1	1	1	1	f
2	2	A4C12BF             	1	2	1	1	2	f
3	3	00000000ABCDEF12    	1	1	2	4	1	f
4	4	00000000ABCDEF12    	1	2	1	1	2	f
\.


--
-- TOC entry 2224 (class 0 OID 0)
-- Dependencies: 196
-- Name: sensors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('sensors_id_seq', 2, true);


--
-- TOC entry 2208 (class 0 OID 16696)
-- Dependencies: 195
-- Data for Name: sensortypes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY sensortypes (id, description, conversion_expression, data_type, sensor_type) FROM stdin;
1	EXAMPLE rain sensor.	/9 +0.9	-1	0
2	EXAMPLE wind sensor.	/-1	-1	0
3	GPS	0	2	0
4	BOOL	0	5	0
\.


--
-- TOC entry 2225 (class 0 OID 0)
-- Dependencies: 194
-- Name: sensortypes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('sensortypes_id_seq', 1, false);


--
-- TOC entry 2205 (class 0 OID 16671)
-- Dependencies: 192
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY users (id, name, email, password, remember_token, created_at, updated_at) FROM stdin;
1	Admin	admin@admin.com	$2y$10$m6zn.nsUmq9Tr0cNlMcZtOnSzqRqCftGIjlE2ZMy1GzTotbPlY2YG	\N	2016-12-09 11:31:35	2016-12-09 11:31:35
\.


--
-- TOC entry 2226 (class 0 OID 0)
-- Dependencies: 191
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('users_id_seq', 1, true);


--
-- TOC entry 2062 (class 2606 OID 16440)
-- Name: gateways gateways_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY gateways
    ADD CONSTRAINT gateways_pkey PRIMARY KEY (gatewayaddress);


--
-- TOC entry 2060 (class 2606 OID 16399)
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- TOC entry 2058 (class 2606 OID 16391)
-- Name: nodes nodes_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY nodes
    ADD CONSTRAINT nodes_pkey PRIMARY KEY (deveui);


--
-- TOC entry 2072 (class 2606 OID 16712)
-- Name: sensors sensors_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors
    ADD CONSTRAINT sensors_pkey PRIMARY KEY (id);


--
-- TOC entry 2070 (class 2606 OID 16704)
-- Name: sensortypes sensortypes_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensortypes
    ADD CONSTRAINT sensortypes_pkey PRIMARY KEY (id);


--
-- TOC entry 2064 (class 2606 OID 16681)
-- Name: users users_email_unique; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_email_unique UNIQUE (email);


--
-- TOC entry 2066 (class 2606 OID 16679)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 2073 (class 1259 OID 16737)
-- Name: message_payloads_message_id_sensor_id_index; Type: INDEX; Schema: public; Owner: docker
--

CREATE INDEX message_payloads_message_id_sensor_id_index ON message_payloads USING btree (message_id, sensor_id);


--
-- TOC entry 2067 (class 1259 OID 16688)
-- Name: password_resets_email_index; Type: INDEX; Schema: public; Owner: docker
--

CREATE INDEX password_resets_email_index ON password_resets USING btree (email);


--
-- TOC entry 2068 (class 1259 OID 16689)
-- Name: password_resets_token_index; Type: INDEX; Schema: public; Owner: docker
--

CREATE INDEX password_resets_token_index ON password_resets USING btree (token);


--
-- TOC entry 2075 (class 2606 OID 16444)
-- Name: gatewaynodes gatewaynode_deveui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY gatewaynodes
    ADD CONSTRAINT gatewaynode_deveui_fkey FOREIGN KEY (deveui) REFERENCES nodes(deveui) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2076 (class 2606 OID 16449)
-- Name: gatewaynodes gatewaynode_gatewayaddress_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY gatewaynodes
    ADD CONSTRAINT gatewaynode_gatewayaddress_fkey FOREIGN KEY (gatewayaddress) REFERENCES gateways(gatewayaddress) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2079 (class 2606 OID 16727)
-- Name: message_payloads message_payloads_message_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY message_payloads
    ADD CONSTRAINT message_payloads_message_id_foreign FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE;


--
-- TOC entry 2080 (class 2606 OID 16732)
-- Name: message_payloads message_payloads_sensor_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY message_payloads
    ADD CONSTRAINT message_payloads_sensor_id_foreign FOREIGN KEY (sensor_id) REFERENCES sensors(id) ON DELETE CASCADE;


--
-- TOC entry 2074 (class 2606 OID 16400)
-- Name: messages messages_deveui_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY messages
    ADD CONSTRAINT messages_deveui_fkey FOREIGN KEY (deveui) REFERENCES nodes(deveui) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 2077 (class 2606 OID 16713)
-- Name: sensors sensors_deveui_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors
    ADD CONSTRAINT sensors_deveui_foreign FOREIGN KEY (deveui) REFERENCES nodes(deveui) ON DELETE CASCADE;


--
-- TOC entry 2078 (class 2606 OID 16718)
-- Name: sensors sensors_sensortype_id_foreign; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY sensors
    ADD CONSTRAINT sensors_sensortype_id_foreign FOREIGN KEY (sensortype_id) REFERENCES sensortypes(id) ON DELETE CASCADE;


-- Completed on 2016-12-15 12:20:58

--
-- PostgreSQL database dump complete
--

