--
-- Name: hr; Type: SCHEMA; Schema: hr; Owner: postgres
--

CREATE SCHEMA hr;


ALTER SCHEMA hr OWNER TO postgres;


--
-- Name: department; Type: TABLE; Schema: hr; Owner: postgres
--



CREATE TABLE hr.department (
    dept_id integer NOT NULL,
    dept_name character varying(50),
    dept_modified_date timestamp without time zone
);


ALTER TABLE hr.department OWNER TO postgres;

--
-- Name: department_dept_id_seq; Type: SEQUENCE; Schema: hr; Owner: postgres
--

CREATE SEQUENCE hr.department_dept_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE hr.department_dept_id_seq OWNER TO postgres;

--
-- Name: department_dept_id_seq; Type: SEQUENCE OWNED BY; Schema: hr; Owner: postgres
--

ALTER SEQUENCE hr.department_dept_id_seq OWNED BY hr.department.dept_id;


--
-- Name: employee; Type: TABLE; Schema: hr; Owner: postgres
--

CREATE TABLE hr.employee (
    emp_entity_id integer NOT NULL,
    emp_emp_number character varying(25),
    emp_national_id character varying(25),
    emp_birth_date date,
    emp_marital_status character(1),
    emp_gender character(1),
    emp_hire_date timestamp without time zone,
    emp_salaried_flag character(1),
    emp_vacation_hours smallint,
    emp_sickleave_hours smallint,
    emp_current_flag smallint,
    emp_modified_date timestamp without time zone,
    emp_type character varying(15),
    emp_joro_id integer,
    emp_emp_entity_id integer,
    CONSTRAINT employee_emp_current_flag_check CHECK ((emp_current_flag = ANY (ARRAY[0, 1]))),
    CONSTRAINT employee_emp_gender_check CHECK ((emp_gender = ANY (ARRAY['M'::bpchar, 'F'::bpchar]))),
    CONSTRAINT employee_emp_marital_status_check CHECK ((emp_marital_status = ANY (ARRAY['M'::bpchar, 'S'::bpchar]))),
    CONSTRAINT employee_emp_salaried_flag_check CHECK ((emp_salaried_flag = ANY (ARRAY['0'::bpchar, '1'::bpchar]))),
    CONSTRAINT employee_emp_type_check CHECK (((emp_type)::text = ANY ((ARRAY['Internal'::character varying, 'Outsource'::character varying])::text[])))
);


ALTER TABLE hr.employee OWNER TO postgres;

--
-- Name: employee_client_contract; Type: TABLE; Schema: hr; Owner: postgres
--

CREATE TABLE hr.employee_client_contract (
    ecco_id integer NOT NULL,
    ecco_entity_id integer NOT NULL,
    ecco_contract_no character varying(55),
    ecco_contract_date timestamp without time zone,
    ecco_start_date timestamp without time zone,
    ecco_end_date timestamp without time zone,
    ecco_notes character varying(512),
    ecco_modified_date timestamp without time zone,
    ecco_media_link character varying(255),
    ecco_joty_id integer,
    ecco_account_manager integer,
    ecco_clit_id integer,
    ecco_status character varying(15)
);


ALTER TABLE hr.employee_client_contract OWNER TO postgres;

--
-- Name: employee_client_contract_ecco_id_seq; Type: SEQUENCE; Schema: hr; Owner: postgres
--

CREATE SEQUENCE hr.employee_client_contract_ecco_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE hr.employee_client_contract_ecco_id_seq OWNER TO postgres;

--
-- Name: employee_client_contract_ecco_id_seq; Type: SEQUENCE OWNED BY; Schema: hr; Owner: postgres
--

ALTER SEQUENCE hr.employee_client_contract_ecco_id_seq OWNED BY hr.employee_client_contract.ecco_id;


--
-- Name: employee_department_history; Type: TABLE; Schema: hr; Owner: postgres
--

CREATE TABLE hr.employee_department_history (
    edhi_id integer NOT NULL,
    edhi_entity_id integer NOT NULL,
    edhi_start_date timestamp without time zone,
    edhi_end_date timestamp without time zone,
    edhi_modified_date timestamp without time zone,
    edhi_dept_id integer
);


ALTER TABLE hr.employee_department_history OWNER TO postgres;

--
-- Name: employee_department_history_edhi_id_seq; Type: SEQUENCE; Schema: hr; Owner: postgres
--

CREATE SEQUENCE hr.employee_department_history_edhi_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE hr.employee_department_history_edhi_id_seq OWNER TO postgres;

--
-- Name: employee_department_history_edhi_id_seq; Type: SEQUENCE OWNED BY; Schema: hr; Owner: postgres
--

ALTER SEQUENCE hr.employee_department_history_edhi_id_seq OWNED BY hr.employee_department_history.edhi_id;


--
-- Name: employee_pay_history; Type: TABLE; Schema: hr; Owner: postgres
--

CREATE TABLE hr.employee_pay_history (
    ephi_entity_id integer NOT NULL,
    ephi_rate_change_date timestamp without time zone NOT NULL,
    ephi_rate_salary integer,
    ephi_pay_frequence smallint,
    ephi_modified_date timestamp without time zone,
    CONSTRAINT employee_pay_history_ephi_pay_frequence_check CHECK ((ephi_pay_frequence = ANY (ARRAY[1, 2])))
);


ALTER TABLE hr.employee_pay_history OWNER TO postgres;