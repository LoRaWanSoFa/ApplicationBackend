--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

-- Started on 2016-12-08 12:53:05

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
76	1	1	p1                                                                                                                                                                                                                                                             
76	1	2	p2                                                                                                                                                                                                                                                             
76	2	3	p3                                                                                                                                                                                                                                                             
\.


--
-- TOC entry 2199 (class 0 OID 16394)
-- Dependencies: 187
-- Data for Name: messages; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY messages (id, deveui, created_at, down) FROM stdin;
115	A4C12BF	2016-12-02 12:36:38.298558	f
116	A4C12BF	2016-12-02 12:56:38.154857	f
117	A4C12BF	2016-12-02 12:57:08.117854	f
118	A4C12BF	2016-12-02 12:57:38.557752	f
119	A4C12BF	2016-12-06 10:55:40.535019	f
120	A4C12BF	2016-12-06 10:56:16.43603	f
121	A4C12BF	2016-12-06 10:56:28.327102	f
122	A4C12BF	2016-12-06 10:56:43.598245	f
123	A4C12BF	2016-12-06 11:50:57.424101	f
124	A4C12BF	2016-12-06 11:50:59.724856	f
125	A4C12BF	2016-12-06 11:51:41.583033	f
126	A4C12BF	2016-12-06 11:52:30.653562	f
127	A4C12BF	2016-12-06 12:36:47.519213	f
128	A4C12BF	2016-12-06 12:39:26.683418	f
129	A4C12BF	2016-12-06 12:39:29.986743	f
130	A4C12BF	2016-12-06 12:40:19.509091	f
131	A4C12BF	2016-12-06 12:40:26.24116	f
132	A4C12BF	2016-12-06 12:40:33.945643	f
133	A4C12BF	2016-12-06 12:40:37.418786	f
134	A4C12BF	2016-12-06 12:40:46.148126	f
135	A4C12BF	2016-12-06 12:41:47.380918	f
136	A4C12BF	2016-12-06 12:41:54.630515	f
137	A4C12BF	2016-12-06 12:42:02.035045	f
138	A4C12BF	2016-12-06 12:47:21.401807	f
139	A4C12BF	2016-12-06 12:47:29.0956	f
140	A4C12BF	2016-12-06 13:07:07.507141	f
141	A4C12BF	2016-12-06 13:07:28.322412	f
142	A4C12BF	2016-12-06 13:10:37.782576	f
143	A4C12BF	2016-12-06 13:10:42.939163	f
144	A4C12BF	2016-12-06 13:15:35.606047	f
145	A4C12BF	2016-12-06 13:24:08.374788	f
146	A4C12BF	2016-12-06 13:26:36.316284	f
147	A4C12BF	2016-12-06 13:26:50.275587	f
148	A4C12BF	2016-12-06 13:34:20.076432	f
150	A4C12BF	2016-12-08 08:25:16.533052	f
151	A4C12BF	2016-12-08 08:30:24.298231	f
152	A4C12BF	2016-12-08 09:08:32.62949	f
76	A4C12BF	2016-12-01 13:56:21.462866	f
153	A4C12BF	2016-12-08 09:08:59.504495	f
154	A4C12BF	2016-12-08 09:37:34.819011	f
155	A4C12BF	2016-12-08 09:38:28.033172	f
156	A4C12BF	2016-12-08 09:38:40.444844	f
157	A4C12BF	2016-12-08 09:39:35.066355	f
158	A4C12BF	2016-12-08 09:39:45.357558	f
159	A4C12BF	2016-12-08 09:39:47.15094	f
160	A4C12BF	2016-12-08 09:41:29.928009	f
161	A4C12BF	2016-12-08 09:41:31.48498	f
162	A4C12BF	2016-12-08 09:41:32.76606	f
163	A4C12BF	2016-12-08 09:41:41.001162	f
164	A4C12BF	2016-12-08 09:41:58.467691	f
165	A4C12BF	2016-12-08 10:02:56.30854	f
166	A4C12BF	2016-12-08 10:19:01.658509	f
167	A4C12BF	2016-12-08 10:21:14.917733	f
168	A4C12BF	2016-12-08 10:21:27.779988	f
169	A4C12BF	2016-12-08 10:21:46.210119	f
170	A4C12BF	2016-12-08 10:22:18.354839	f
171	A4C12BF	2016-12-08 10:22:34.229935	f
172	A4C12BF	2016-12-08 10:32:05.320805	f
173	A4C12BF	2016-12-08 10:37:19.037818	f
175	A4C12BF	2016-12-08 10:37:39.791364	f
177	A4C12BF	2016-12-08 10:38:26.500802	f
178	00000000ABCDEF12	2016-12-08 10:45:52.937038	f
179	A4C12BF	2016-12-08 10:45:53.033343	f
180	A4C12BF	2016-12-08 10:52:05.948179	f
181	A4C12BF	2016-12-08 10:52:12.140259	f
182	00000000ABCDEF12	2016-12-08 11:37:19.14501	f
183	00000000ABCDEF12	2016-12-08 11:37:30.249535	f
184	A4C12BF	2016-12-08 11:45:46.34836	f
185	A4C12BF	2016-12-08 11:45:59.07122	f
186	A4C12BF	2016-12-08 11:46:24.619158	f
187	A4C12BF	2016-12-08 11:46:42.604792	f
188	A4C12BF	2016-12-08 11:48:36.122436	f
189	A4C12BF	2016-12-08 11:49:07.911551	f
\.


--
-- TOC entry 2222 (class 0 OID 0)
-- Dependencies: 186
-- Name: messages_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('messages_id_seq', 189, true);


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
00000000ABCDEF12	ABCDEF12	70B426334	1325746822	C	6	40.730610	-73.935242
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
3	3	00000000ABCDEF12    	1	1	2	4	1
4	4	00000000ABCDEF12    	1	2	1	1	2
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
3	GPS	0	2
4	BOOL	0	5
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


-- Completed on 2016-12-08 12:53:05

--
-- PostgreSQL database dump complete
--

