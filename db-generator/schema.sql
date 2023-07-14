--
-- Name: bootcamp; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA bootcamp;


ALTER SCHEMA bootcamp OWNER TO postgres;

--
-- Name: curriculum; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA curriculum;


ALTER SCHEMA curriculum OWNER TO postgres;

--
-- Name: hr; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA hr;


ALTER SCHEMA hr OWNER TO postgres;

--
-- Name: jobhire; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA jobhire;


ALTER SCHEMA jobhire OWNER TO postgres;

--
-- Name: master; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA master;


ALTER SCHEMA master OWNER TO postgres;

--
-- Name: payment; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA payment;


ALTER SCHEMA payment OWNER TO postgres;

--
-- Name: sales; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA sales;


ALTER SCHEMA sales OWNER TO postgres;

--
-- Name: users; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA users;


ALTER SCHEMA users OWNER TO postgres;


--
-- Name: batch; Type: TABLE; Schema: bootcamp; Owner: postgres
--

CREATE TABLE bootcamp.batch (
    batch_id integer NOT NULL,
    batch_entity_id integer NOT NULL,
    batch_name character varying(15),
    batch_description character varying(125),
    batch_start_date timestamp without time zone,
    batch_end_date timestamp without time zone,
    batch_reason character varying(256),
    batch_type character varying(15),
    batch_modified_date timestamp without time zone,
    batch_status character varying(15),
    batch_pic_id integer,
    CONSTRAINT batch_batch_type_check CHECK (((batch_type)::text = ANY ((ARRAY['offline'::character varying, 'online'::character varying, 'corporate'::character varying])::text[])))
);


ALTER TABLE bootcamp.batch OWNER TO postgres;

--
-- Name: batch_batch_id_seq; Type: SEQUENCE; Schema: bootcamp; Owner: postgres
--

CREATE SEQUENCE bootcamp.batch_batch_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE bootcamp.batch_batch_id_seq OWNER TO postgres;

--
-- Name: batch_batch_id_seq; Type: SEQUENCE OWNED BY; Schema: bootcamp; Owner: postgres
--

ALTER SEQUENCE bootcamp.batch_batch_id_seq OWNED BY bootcamp.batch.batch_id;


--
-- Name: batch_trainee; Type: TABLE; Schema: bootcamp; Owner: postgres
--

CREATE TABLE bootcamp.batch_trainee (
    batr_id integer NOT NULL,
    batr_status character varying(15),
    batr_certificated character(1),
    batre_certificate_link character varying(255),
    batr_access_token character varying(255),
    batr_access_grant character(1),
    batr_review character varying(1024),
    batr_total_score integer,
    batr_modified_date timestamp without time zone,
    batr_trainee_entity_id integer,
    batr_batch_id integer NOT NULL,
    CONSTRAINT batch_trainee_batr_access_grant_check CHECK ((batr_access_grant = ANY (ARRAY['0'::bpchar, '1'::bpchar]))),
    CONSTRAINT batch_trainee_batr_certificated_check CHECK ((batr_certificated = ANY (ARRAY['0'::bpchar, '1'::bpchar]))),
    CONSTRAINT batch_trainee_batr_status_check CHECK (((batr_status)::text = ANY ((ARRAY['passed'::character varying, 'failed'::character varying, 'resign'::character varying, 'running'::character varying])::text[])))
);


ALTER TABLE bootcamp.batch_trainee OWNER TO postgres;

--
-- Name: batch_trainee_batr_id_seq; Type: SEQUENCE; Schema: bootcamp; Owner: postgres
--

CREATE SEQUENCE bootcamp.batch_trainee_batr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE bootcamp.batch_trainee_batr_id_seq OWNER TO postgres;

--
-- Name: batch_trainee_batr_id_seq; Type: SEQUENCE OWNED BY; Schema: bootcamp; Owner: postgres
--

ALTER SEQUENCE bootcamp.batch_trainee_batr_id_seq OWNED BY bootcamp.batch_trainee.batr_id;


--
-- Name: batch_trainee_evaluation; Type: TABLE; Schema: bootcamp; Owner: postgres
--

CREATE TABLE bootcamp.batch_trainee_evaluation (
    btev_id integer NOT NULL,
    btev_type character varying(15),
    btev_header character varying(256),
    btev_section character varying(256),
    btev_skill character varying(256),
    btev_week integer,
    btev_skor integer,
    btev_note character varying(256),
    btev_modified_date timestamp without time zone,
    btev_batch_id integer,
    btev_trainee_entity_id integer,
    CONSTRAINT batch_trainee_evaluation_btev_type_check CHECK (((btev_type)::text = ANY ((ARRAY['hardskill'::character varying, 'softskill'::character varying])::text[])))
);


ALTER TABLE bootcamp.batch_trainee_evaluation OWNER TO postgres;

--
-- Name: batch_trainee_evaluation_btev_id_seq; Type: SEQUENCE; Schema: bootcamp; Owner: postgres
--

CREATE SEQUENCE bootcamp.batch_trainee_evaluation_btev_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE bootcamp.batch_trainee_evaluation_btev_id_seq OWNER TO postgres;

--
-- Name: batch_trainee_evaluation_btev_id_seq; Type: SEQUENCE OWNED BY; Schema: bootcamp; Owner: postgres
--

ALTER SEQUENCE bootcamp.batch_trainee_evaluation_btev_id_seq OWNED BY bootcamp.batch_trainee_evaluation.btev_id;


--
-- Name: instructor_programs; Type: TABLE; Schema: bootcamp; Owner: postgres
--

CREATE TABLE bootcamp.instructor_programs (
    batch_id integer NOT NULL,
    inpro_entity_id integer NOT NULL,
    inpro_emp_entity_id integer NOT NULL,
    inpro_modified_date timestamp without time zone
);


ALTER TABLE bootcamp.instructor_programs OWNER TO postgres;

--
-- Name: program_apply; Type: TABLE; Schema: bootcamp; Owner: postgres
--

CREATE TABLE bootcamp.program_apply (
    prap_user_entity_id integer NOT NULL,
    prap_prog_entity_id integer NOT NULL,
    prap_test_score integer,
    prap_gpa integer,
    prap_iq_test integer,
    prap_review character varying(256),
    prap_modified_date timestamp without time zone,
    prap_status character varying(15)
);


ALTER TABLE bootcamp.program_apply OWNER TO postgres;

--
-- Name: program_apply_progress; Type: TABLE; Schema: bootcamp; Owner: postgres
--

CREATE TABLE bootcamp.program_apply_progress (
    parog_id integer NOT NULL,
    parog_user_entity_id integer NOT NULL,
    parog_prog_entity_id integer NOT NULL,
    parog_action_date timestamp without time zone,
    parog_modified_date timestamp without time zone,
    parog_comment character varying(512),
    parog_progress_name character varying(15),
    parog_emp_entity_id integer,
    parog_status character varying(15)
);


ALTER TABLE bootcamp.program_apply_progress OWNER TO postgres;

--
-- Name: program_apply_progress_parog_id_seq; Type: SEQUENCE; Schema: bootcamp; Owner: postgres
--

CREATE SEQUENCE bootcamp.program_apply_progress_parog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE bootcamp.program_apply_progress_parog_id_seq OWNER TO postgres;

--
-- Name: program_apply_progress_parog_id_seq; Type: SEQUENCE OWNED BY; Schema: bootcamp; Owner: postgres
--

ALTER SEQUENCE bootcamp.program_apply_progress_parog_id_seq OWNED BY bootcamp.program_apply_progress.parog_id;


--
-- Name: program_entity; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.program_entity (
    prog_entity_id integer NOT NULL,
    prog_title character varying(256),
    prog_headline character varying(512),
    prog_type character varying(15),
    prog_learning_type character varying(15),
    prog_rating integer,
    prog_total_trainee integer,
    prog_modified_date timestamp without time zone,
    prog_image character varying(256),
    prog_best_seller character(1),
    prog_price integer,
    prog_language character varying(35),
    prog_duration integer,
    prog_duration_type character varying(15),
    prog_tag_skill character varying(512),
    prog_city_id integer,
    prog_cate_id integer,
    prog_created_by integer,
    prog_status character varying(15),
    CONSTRAINT program_entity_prog_best_seller_check CHECK ((prog_best_seller = ANY (ARRAY['0'::bpchar, '1'::bpchar]))),
    CONSTRAINT program_entity_prog_duration_type_check CHECK (((prog_duration_type)::text = ANY ((ARRAY['month'::character varying, 'week'::character varying, 'days'::character varying])::text[]))),
    CONSTRAINT program_entity_prog_language_check CHECK (((prog_language)::text = ANY ((ARRAY['english'::character varying, 'bahasa'::character varying])::text[]))),
    CONSTRAINT program_entity_prog_learning_type_check CHECK (((prog_learning_type)::text = ANY ((ARRAY['online'::character varying, 'offline'::character varying, 'both'::character varying])::text[]))),
    CONSTRAINT program_entity_prog_type_check CHECK (((prog_type)::text = ANY ((ARRAY['bootcamp'::character varying, 'course'::character varying])::text[])))
);


ALTER TABLE curriculum.program_entity OWNER TO postgres;

--
-- Name: program_entity_description; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.program_entity_description (
    pred_prog_entity_id integer NOT NULL,
    pred_item_learning json,
    pred_item_include json,
    pred_requirement json,
    pred_description json,
    pred_target_level json
);


ALTER TABLE curriculum.program_entity_description OWNER TO postgres;

--
-- Name: program_entity_prog_entity_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.program_entity_prog_entity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE curriculum.program_entity_prog_entity_id_seq OWNER TO postgres;

--
-- Name: program_entity_prog_entity_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.program_entity_prog_entity_id_seq OWNED BY curriculum.program_entity.prog_entity_id;


--
-- Name: program_reviews; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.program_reviews (
    prow_user_entity_id integer NOT NULL,
    prow_prog_entity_id integer NOT NULL,
    prow_review character varying(512),
    prow_rating integer,
    prow_modified_date timestamp without time zone
);


ALTER TABLE curriculum.program_reviews OWNER TO postgres;

--
-- Name: section_detail; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.section_detail (
    secd_id integer NOT NULL,
    secd_title character varying(256),
    secd_preview character(1),
    secd_score integer,
    secd_note character varying(256),
    secd_minute integer,
    secd_modified_date timestamp without time zone,
    secd_sect_id integer,
    CONSTRAINT section_detail_secd_preview_check CHECK ((secd_preview = ANY (ARRAY['0'::bpchar, '1'::bpchar])))
);


ALTER TABLE curriculum.section_detail OWNER TO postgres;

--
-- Name: section_detail_material; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.section_detail_material (
    sedm_id integer NOT NULL,
    sedm_filename character varying(55),
    sedm_filesize integer,
    sedm_filetype character varying(15),
    sedm_filelink character varying(255),
    sedm_modified_date timestamp without time zone,
    sedm_secd_id integer,
    CONSTRAINT section_detail_material_sedm_filetype_check CHECK (((sedm_filetype)::text = ANY ((ARRAY['video'::character varying, 'image'::character varying, 'text'::character varying, 'link'::character varying])::text[])))
);


ALTER TABLE curriculum.section_detail_material OWNER TO postgres;

--
-- Name: section_detail_material_sedm_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.section_detail_material_sedm_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE curriculum.section_detail_material_sedm_id_seq OWNER TO postgres;

--
-- Name: section_detail_material_sedm_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.section_detail_material_sedm_id_seq OWNED BY curriculum.section_detail_material.sedm_id;


--
-- Name: section_detail_secd_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.section_detail_secd_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE curriculum.section_detail_secd_id_seq OWNER TO postgres;

--
-- Name: section_detail_secd_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.section_detail_secd_id_seq OWNED BY curriculum.section_detail.secd_id;


--
-- Name: sections; Type: TABLE; Schema: curriculum; Owner: postgres
--

CREATE TABLE curriculum.sections (
    sect_id integer NOT NULL,
    sect_prog_entity_id integer NOT NULL,
    sect_title character varying(100),
    sect_description character varying(256),
    sect_total_section integer,
    sect_total_lecture integer,
    sect_total_minute integer,
    sect_modified_date timestamp without time zone
);


ALTER TABLE curriculum.sections OWNER TO postgres;

--
-- Name: sections_sect_id_seq; Type: SEQUENCE; Schema: curriculum; Owner: postgres
--

CREATE SEQUENCE curriculum.sections_sect_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE curriculum.sections_sect_id_seq OWNER TO postgres;

--
-- Name: sections_sect_id_seq; Type: SEQUENCE OWNED BY; Schema: curriculum; Owner: postgres
--

ALTER SEQUENCE curriculum.sections_sect_id_seq OWNED BY curriculum.sections.sect_id;


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

--
-- Name: client; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.client (
    clit_id integer NOT NULL,
    clit_name character varying(256),
    clit_about character varying(512),
    clit_modified_date timestamp without time zone,
    clit_addr_id integer,
    clit_emra_id integer
);


ALTER TABLE jobhire.client OWNER TO postgres;

--
-- Name: employee_range; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.employee_range (
    emra_id integer NOT NULL,
    emra_range_min integer,
    emra_range_max integer,
    emra_modified_date timestamp without time zone
);


ALTER TABLE jobhire.employee_range OWNER TO postgres;

--
-- Name: employee_range_emra_id_seq; Type: SEQUENCE; Schema: jobhire; Owner: postgres
--

CREATE SEQUENCE jobhire.employee_range_emra_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE jobhire.employee_range_emra_id_seq OWNER TO postgres;

--
-- Name: employee_range_emra_id_seq; Type: SEQUENCE OWNED BY; Schema: jobhire; Owner: postgres
--

ALTER SEQUENCE jobhire.employee_range_emra_id_seq OWNED BY jobhire.employee_range.emra_id;


--
-- Name: job_category; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.job_category (
    joca_id integer NOT NULL,
    joca_name character varying(255),
    joca_modified_date timestamp without time zone
);


ALTER TABLE jobhire.job_category OWNER TO postgres;

--
-- Name: job_category_joca_id_seq; Type: SEQUENCE; Schema: jobhire; Owner: postgres
--

CREATE SEQUENCE jobhire.job_category_joca_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE jobhire.job_category_joca_id_seq OWNER TO postgres;

--
-- Name: job_category_joca_id_seq; Type: SEQUENCE OWNED BY; Schema: jobhire; Owner: postgres
--

ALTER SEQUENCE jobhire.job_category_joca_id_seq OWNED BY jobhire.job_category.joca_id;


--
-- Name: job_photo; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.job_photo (
    jopho_id integer NOT NULL,
    jopho_filename character varying(55),
    jopho_filesize integer,
    jopho_filetype character varying(15),
    jopho_modified_date timestamp without time zone,
    jopho_entity_id integer,
    CONSTRAINT job_photo_jopho_filetype_check CHECK (((jopho_filetype)::text = ANY ((ARRAY['png'::character varying, 'jpeg'::character varying])::text[])))
);


ALTER TABLE jobhire.job_photo OWNER TO postgres;

--
-- Name: job_photo_jopho_id_seq; Type: SEQUENCE; Schema: jobhire; Owner: postgres
--

CREATE SEQUENCE jobhire.job_photo_jopho_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE jobhire.job_photo_jopho_id_seq OWNER TO postgres;

--
-- Name: job_photo_jopho_id_seq; Type: SEQUENCE OWNED BY; Schema: jobhire; Owner: postgres
--

ALTER SEQUENCE jobhire.job_photo_jopho_id_seq OWNED BY jobhire.job_photo.jopho_id;


--
-- Name: job_post; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.job_post (
    jopo_entity_id integer NOT NULL,
    jopo_number character varying(25),
    jopo_title character varying(256),
    jopo_start_date timestamp without time zone,
    jopo_end_date timestamp without time zone,
    jopo_min_salary integer,
    jopo_max_salary integer,
    jopo_min_experience integer,
    jopo_max_experience integer,
    jopo_primary_skill character varying(256),
    jopo_secondary_skill character varying(256),
    jopo_publish_date timestamp without time zone,
    jopo_modified_date timestamp without time zone,
    jopo_emp_entity_id integer,
    jopo_clit_id integer,
    jopo_joro_id integer,
    jopo_joty_id integer,
    jopo_joca_id integer,
    jopo_addr_id integer,
    jopo_work_code character varying(15),
    jopo_edu_code character varying(5),
    jopo_indu_code character varying(15),
    jopo_status character varying(15)
);


ALTER TABLE jobhire.job_post OWNER TO postgres;

--
-- Name: job_post_desc; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.job_post_desc (
    jopo_entity_id integer NOT NULL,
    jopo_description json,
    jopo_responsibility json,
    jopo_target json,
    jopo_benefit json
);


ALTER TABLE jobhire.job_post_desc OWNER TO postgres;

--
-- Name: job_post_jopo_entity_id_seq; Type: SEQUENCE; Schema: jobhire; Owner: postgres
--

CREATE SEQUENCE jobhire.job_post_jopo_entity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE jobhire.job_post_jopo_entity_id_seq OWNER TO postgres;

--
-- Name: job_post_jopo_entity_id_seq; Type: SEQUENCE OWNED BY; Schema: jobhire; Owner: postgres
--

ALTER SEQUENCE jobhire.job_post_jopo_entity_id_seq OWNED BY jobhire.job_post.jopo_entity_id;


--
-- Name: talent_apply; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.talent_apply (
    taap_user_entity_id integer NOT NULL,
    taap_entity_id integer NOT NULL,
    taap_intro character varying,
    taap_scoring integer,
    taap_modified_date timestamp without time zone,
    taap_status character varying(15)
);


ALTER TABLE jobhire.talent_apply OWNER TO postgres;

--
-- Name: talent_apply_progress; Type: TABLE; Schema: jobhire; Owner: postgres
--

CREATE TABLE jobhire.talent_apply_progress (
    tapr_id integer NOT NULL,
    taap_user_entity_id integer NOT NULL,
    taap_entity_id integer NOT NULL,
    tapr_modified_date timestamp without time zone,
    tapr_status character varying(15),
    tapr_comment character varying(256),
    tapr_progress_name character varying(55)
);


ALTER TABLE jobhire.talent_apply_progress OWNER TO postgres;

--
-- Name: talent_apply_progress_tapr_id_seq; Type: SEQUENCE; Schema: jobhire; Owner: postgres
--

CREATE SEQUENCE jobhire.talent_apply_progress_tapr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE jobhire.talent_apply_progress_tapr_id_seq OWNER TO postgres;

--
-- Name: talent_apply_progress_tapr_id_seq; Type: SEQUENCE OWNED BY; Schema: jobhire; Owner: postgres
--

ALTER SEQUENCE jobhire.talent_apply_progress_tapr_id_seq OWNED BY jobhire.talent_apply_progress.tapr_id;


--
-- Name: address; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.address (
    addr_id integer NOT NULL,
    addr_line1 character varying(255),
    addr_line2 character varying(255),
    addr_postal_code character varying(10),
    addr_spatial_location character varying,
    addr_modified_date timestamp without time zone,
    addr_city_id integer
);


ALTER TABLE master.address OWNER TO postgres;

--
-- Name: address_addr_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.address_addr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.address_addr_id_seq OWNER TO postgres;

--
-- Name: address_addr_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.address_addr_id_seq OWNED BY master.address.addr_id;


--
-- Name: address_type; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.address_type (
    adty_id integer NOT NULL,
    adty_name character varying(15),
    adty_modified_date timestamp without time zone
);


ALTER TABLE master.address_type OWNER TO postgres;

--
-- Name: address_type_adty_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.address_type_adty_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.address_type_adty_id_seq OWNER TO postgres;

--
-- Name: address_type_adty_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.address_type_adty_id_seq OWNED BY master.address_type.adty_id;


--
-- Name: category; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.category (
    cate_id integer NOT NULL,
    cate_name character varying(255),
    cate_cate_id integer,
    cate_modified_date timestamp without time zone
);


ALTER TABLE master.category OWNER TO postgres;

--
-- Name: category_cate_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.category_cate_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.category_cate_id_seq OWNER TO postgres;

--
-- Name: category_cate_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.category_cate_id_seq OWNED BY master.category.cate_id;


--
-- Name: city; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.city (
    city_id integer NOT NULL,
    city_name character varying(155),
    city_modified_date timestamp without time zone,
    city_prov_id integer
);


ALTER TABLE master.city OWNER TO postgres;

--
-- Name: city_city_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.city_city_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.city_city_id_seq OWNER TO postgres;

--
-- Name: city_city_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.city_city_id_seq OWNED BY master.city.city_id;


--
-- Name: country; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.country (
    country_code character varying(3) NOT NULL,
    country_name character varying(85),
    country_modified_date timestamp without time zone
);


ALTER TABLE master.country OWNER TO postgres;

--
-- Name: education; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.education (
    edu_code character varying(5) NOT NULL,
    edu_name character varying(55)
);


ALTER TABLE master.education OWNER TO postgres;

--
-- Name: industry; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.industry (
    indu_code character varying(15) NOT NULL,
    indu_name character varying(85)
);


ALTER TABLE master.industry OWNER TO postgres;

--
-- Name: job_role; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.job_role (
    joro_id integer NOT NULL,
    joro_name character varying(55),
    joro_modified_date timestamp without time zone
);


ALTER TABLE master.job_role OWNER TO postgres;

--
-- Name: job_role_joro_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.job_role_joro_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.job_role_joro_id_seq OWNER TO postgres;

--
-- Name: job_role_joro_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.job_role_joro_id_seq OWNED BY master.job_role.joro_id;


--
-- Name: job_type; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.job_type (
    joty_id integer NOT NULL,
    joty_name character varying(55)
);


ALTER TABLE master.job_type OWNER TO postgres;

--
-- Name: job_type_joty_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.job_type_joty_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.job_type_joty_id_seq OWNER TO postgres;

--
-- Name: job_type_joty_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.job_type_joty_id_seq OWNED BY master.job_type.joty_id;


--
-- Name: modules; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.modules (
    module_name character varying(125) NOT NULL
);


ALTER TABLE master.modules OWNER TO postgres;

--
-- Name: province; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.province (
    prov_id integer NOT NULL,
    prov_code character varying(5),
    prov_name character varying(85),
    prov_modified_date timestamp without time zone,
    prov_country_code character varying(3)
);


ALTER TABLE master.province OWNER TO postgres;

--
-- Name: province_prov_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.province_prov_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.province_prov_id_seq OWNER TO postgres;

--
-- Name: province_prov_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.province_prov_id_seq OWNED BY master.province.prov_id;


--
-- Name: route_actions; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.route_actions (
    roac_id integer NOT NULL,
    roac_name character varying(30),
    roac_orderby integer,
    roac_display integer,
    roac_module_name character varying(125)
);


ALTER TABLE master.route_actions OWNER TO postgres;

--
-- Name: route_actions_roac_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.route_actions_roac_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.route_actions_roac_id_seq OWNER TO postgres;

--
-- Name: route_actions_roac_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.route_actions_roac_id_seq OWNED BY master.route_actions.roac_id;


--
-- Name: skill_template; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.skill_template (
    skte_id integer NOT NULL,
    skte_skill character varying(256),
    skte_description character varying(256),
    skte_week integer,
    skte_orderby integer,
    skte_modified_date timestamp without time zone,
    skty_name character varying(55),
    skte_skte_id integer
);


ALTER TABLE master.skill_template OWNER TO postgres;

--
-- Name: skill_template_skte_id_seq; Type: SEQUENCE; Schema: master; Owner: postgres
--

CREATE SEQUENCE master.skill_template_skte_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE master.skill_template_skte_id_seq OWNER TO postgres;

--
-- Name: skill_template_skte_id_seq; Type: SEQUENCE OWNED BY; Schema: master; Owner: postgres
--

ALTER SEQUENCE master.skill_template_skte_id_seq OWNED BY master.skill_template.skte_id;


--
-- Name: skill_type; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.skill_type (
    skty_name character varying(55) NOT NULL
);


ALTER TABLE master.skill_type OWNER TO postgres;

--
-- Name: status; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.status (
    status character varying(15) NOT NULL,
    status_modified_date timestamp without time zone,
    status_module character varying(25)
);


ALTER TABLE master.status OWNER TO postgres;

--
-- Name: working_type; Type: TABLE; Schema: master; Owner: postgres
--

CREATE TABLE master.working_type (
    woty_code character varying(15) NOT NULL,
    woty_name character varying(55)
);


ALTER TABLE master.working_type OWNER TO postgres;

--
-- Name: bank; Type: TABLE; Schema: payment; Owner: postgres
--

CREATE TABLE payment.bank (
    bank_entity_id integer NOT NULL,
    bank_code character varying(10),
    bank_name character varying(55),
    bank_modified_date timestamp without time zone
);


ALTER TABLE payment.bank OWNER TO postgres;

--
-- Name: fintech; Type: TABLE; Schema: payment; Owner: postgres
--

CREATE TABLE payment.fintech (
    fint_entity_id integer NOT NULL,
    fint_code character varying(10),
    fint_name character varying(55),
    fint_modified_date timestamp without time zone
);


ALTER TABLE payment.fintech OWNER TO postgres;

--
-- Name: transaction_payment; Type: TABLE; Schema: payment; Owner: postgres
--

CREATE TABLE payment.transaction_payment (
    trpa_id integer NOT NULL,
    trpa_code_number character varying(55),
    trpa_order_number character varying(25),
    trpa_debit numeric,
    trpa_credit numeric,
    trpa_type character varying(15),
    trpa_note character varying(255),
    trpa_modified_date timestamp without time zone,
    trpa_source_id character varying(25) NOT NULL,
    trpa_target_id character varying(25) NOT NULL,
    trpa_user_entity_id integer,
    CONSTRAINT transaction_payment_trpa_type_check CHECK (((trpa_type)::text = ANY ((ARRAY['top-up'::character varying, 'transfer'::character varying, 'order'::character varying, 'refund'::character varying])::text[])))
);


ALTER TABLE payment.transaction_payment OWNER TO postgres;

--
-- Name: transaction_payment_trpa_id_seq; Type: SEQUENCE; Schema: payment; Owner: postgres
--

CREATE SEQUENCE payment.transaction_payment_trpa_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE payment.transaction_payment_trpa_id_seq OWNER TO postgres;

--
-- Name: transaction_payment_trpa_id_seq; Type: SEQUENCE OWNED BY; Schema: payment; Owner: postgres
--

ALTER SEQUENCE payment.transaction_payment_trpa_id_seq OWNED BY payment.transaction_payment.trpa_id;


--
-- Name: users_account; Type: TABLE; Schema: payment; Owner: postgres
--

CREATE TABLE payment.users_account (
    usac_bank_entity_id integer NOT NULL,
    usac_user_entity_id integer NOT NULL,
    usac_account_number character varying(25),
    usac_saldo numeric,
    usac_type character varying(15),
    usac_start_date timestamp without time zone,
    usac_end_date timestamp without time zone,
    usac_modified_date timestamp without time zone,
    usac_status character varying(15),
    CONSTRAINT users_account_usac_status_check CHECK (((usac_status)::text = ANY ((ARRAY['active'::character varying, 'inactive'::character varying, 'blocked'::character varying])::text[]))),
    CONSTRAINT users_account_usac_type_check CHECK (((usac_type)::text = ANY ((ARRAY['debit'::character varying, 'credit-card'::character varying, 'payment'::character varying])::text[])))
);


ALTER TABLE payment.users_account OWNER TO postgres;

--
-- Name: cart_items; Type: TABLE; Schema: sales; Owner: postgres
--

CREATE TABLE sales.cart_items (
    cait_id integer NOT NULL,
    cait_quantity integer,
    cait_unit_price money,
    cait_modified_date timestamp without time zone,
    cait_user_entity_id integer,
    cait_prog_entity_id integer
);


ALTER TABLE sales.cart_items OWNER TO postgres;

--
-- Name: cart_items_cait_id_seq; Type: SEQUENCE; Schema: sales; Owner: postgres
--

CREATE SEQUENCE sales.cart_items_cait_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sales.cart_items_cait_id_seq OWNER TO postgres;

--
-- Name: cart_items_cait_id_seq; Type: SEQUENCE OWNED BY; Schema: sales; Owner: postgres
--

ALTER SEQUENCE sales.cart_items_cait_id_seq OWNED BY sales.cart_items.cait_id;


--
-- Name: sales_order_detail; Type: TABLE; Schema: sales; Owner: postgres
--

CREATE TABLE sales.sales_order_detail (
    sode_id integer NOT NULL,
    sode_qty integer,
    sode_unit_price money,
    sode_unit_discount money,
    sode_line_total integer,
    sode_modified_date timestamp without time zone,
    sode_sohe_id integer,
    sode_prog_entity_id integer
);


ALTER TABLE sales.sales_order_detail OWNER TO postgres;

--
-- Name: sales_order_header; Type: TABLE; Schema: sales; Owner: postgres
--

CREATE TABLE sales.sales_order_header (
    sohe_id integer NOT NULL,
    sohe_order_date timestamp without time zone,
    sohe_due_date timestamp without time zone,
    sohe_ship_date timestamp without time zone,
    sohe_order_number character varying(25),
    sohe_account_number character varying(25),
    sohe_trpa_code_number character varying(55),
    sohe_subtotal money,
    sohe_tax money,
    sohe_total_due integer,
    sohe_license_code character varying(512),
    sohe_modified_date timestamp without time zone,
    sohe_user_entity_id integer,
    sohe_status character varying(15)
);


ALTER TABLE sales.sales_order_header OWNER TO postgres;

--
-- Name: sales_order_header_sohe_id_seq; Type: SEQUENCE; Schema: sales; Owner: postgres
--

CREATE SEQUENCE sales.sales_order_header_sohe_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sales.sales_order_header_sohe_id_seq OWNER TO postgres;

--
-- Name: sales_order_header_sohe_id_seq; Type: SEQUENCE OWNED BY; Schema: sales; Owner: postgres
--

ALTER SEQUENCE sales.sales_order_header_sohe_id_seq OWNED BY sales.sales_order_header.sohe_id;


--
-- Name: special_offer; Type: TABLE; Schema: sales; Owner: postgres
--

CREATE TABLE sales.special_offer (
    spof_id integer NOT NULL,
    spof_description character varying(256),
    spof_discount integer,
    spof_type character varying(15),
    spof_start_date timestamp without time zone,
    spof_end_date timestamp without time zone,
    spof_min_qty integer,
    spof_max_qty integer,
    spof_modified_date timestamp without time zone,
    spof_cate_id integer
);


ALTER TABLE sales.special_offer OWNER TO postgres;

--
-- Name: special_offer_programs; Type: TABLE; Schema: sales; Owner: postgres
--

CREATE TABLE sales.special_offer_programs (
    soco_id integer NOT NULL,
    soco_spof_id integer NOT NULL,
    soco_prog_entity_id integer NOT NULL,
    soco_status character varying(15),
    soco_modified_date timestamp without time zone,
    CONSTRAINT special_offer_programs_soco_status_check CHECK (((soco_status)::text = ANY ((ARRAY['open'::character varying, 'cancelled'::character varying, 'closed'::character varying])::text[])))
);


ALTER TABLE sales.special_offer_programs OWNER TO postgres;

--
-- Name: special_offer_programs_soco_id_seq; Type: SEQUENCE; Schema: sales; Owner: postgres
--

CREATE SEQUENCE sales.special_offer_programs_soco_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sales.special_offer_programs_soco_id_seq OWNER TO postgres;

--
-- Name: special_offer_programs_soco_id_seq; Type: SEQUENCE OWNED BY; Schema: sales; Owner: postgres
--

ALTER SEQUENCE sales.special_offer_programs_soco_id_seq OWNED BY sales.special_offer_programs.soco_id;


--
-- Name: special_offer_programs_soco_spof_id_seq; Type: SEQUENCE; Schema: sales; Owner: postgres
--

CREATE SEQUENCE sales.special_offer_programs_soco_spof_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sales.special_offer_programs_soco_spof_id_seq OWNER TO postgres;

--
-- Name: special_offer_programs_soco_spof_id_seq; Type: SEQUENCE OWNED BY; Schema: sales; Owner: postgres
--

ALTER SEQUENCE sales.special_offer_programs_soco_spof_id_seq OWNED BY sales.special_offer_programs.soco_spof_id;


--
-- Name: special_offer_spof_id_seq; Type: SEQUENCE; Schema: sales; Owner: postgres
--

CREATE SEQUENCE sales.special_offer_spof_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sales.special_offer_spof_id_seq OWNER TO postgres;

--
-- Name: special_offer_spof_id_seq; Type: SEQUENCE OWNED BY; Schema: sales; Owner: postgres
--

ALTER SEQUENCE sales.special_offer_spof_id_seq OWNED BY sales.special_offer.spof_id;


--
-- Name: business_entity; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.business_entity (
    entity_id integer NOT NULL
);


ALTER TABLE users.business_entity OWNER TO postgres;

--
-- Name: business_entity_entity_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.business_entity_entity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.business_entity_entity_id_seq OWNER TO postgres;

--
-- Name: business_entity_entity_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.business_entity_entity_id_seq OWNED BY users.business_entity.entity_id;


--
-- Name: phone_number_type; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.phone_number_type (
    ponty_code character varying(15) NOT NULL,
    ponty_modified_date timestamp without time zone
);


ALTER TABLE users.phone_number_type OWNER TO postgres;

--
-- Name: roles; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.roles (
    role_id integer NOT NULL,
    role_name character varying(35),
    role_type character varying(15),
    role_modified_date timestamp without time zone
);


ALTER TABLE users.roles OWNER TO postgres;

--
-- Name: roles_role_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.roles_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.roles_role_id_seq OWNER TO postgres;

--
-- Name: roles_role_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.roles_role_id_seq OWNED BY users.roles.role_id;


--
-- Name: users; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users (
    user_entity_id integer NOT NULL,
    user_name character varying(15),
    user_password character varying(256),
    user_first_name character varying(50),
    user_last_name character varying(50),
    user_birth_date timestamp without time zone,
    user_email_promotion integer,
    user_demographic character varying,
    user_modified_date timestamp without time zone,
    user_photo character varying(255),
    user_current_role integer
);


ALTER TABLE users.users OWNER TO postgres;

--
-- Name: users_address; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_address (
    etad_addr_id integer NOT NULL,
    etad_modified_date timestamp without time zone,
    etad_entity_id integer,
    etad_adty_id integer
);


ALTER TABLE users.users_address OWNER TO postgres;

--
-- Name: users_education; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_education (
    usdu_id integer NOT NULL,
    usdu_entity_id integer NOT NULL,
    usdu_school character varying(255),
    usdu_degree character varying(15),
    usdu_field_study character varying(125),
    usdu_graduate_year character varying(4),
    usdu_start_date timestamp without time zone,
    usdu_end_date timestamp without time zone,
    usdu_grade character varying(5),
    usdu_activities character varying(512),
    usdu_description character varying(512),
    usdu_modified_date timestamp without time zone,
    CONSTRAINT users_eductaion_usdu_degree_check CHECK (((usdu_degree)::text = ANY ((ARRAY['Bachelor'::character varying, 'Diploma'::character varying])::text[])))
);


ALTER TABLE users.users_education OWNER TO postgres;

--
-- Name: users_eductaion_usdu_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.users_eductaion_usdu_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.users_eductaion_usdu_id_seq OWNER TO postgres;

--
-- Name: users_eductaion_usdu_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.users_eductaion_usdu_id_seq OWNED BY users.users_education.usdu_id;


--
-- Name: users_email; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_email (
    pmail_entity_id integer NOT NULL,
    pmail_id integer NOT NULL,
    pmail_address character varying(50),
    pmail_modified_date timestamp without time zone
);


ALTER TABLE users.users_email OWNER TO postgres;

--
-- Name: users_email_pmail_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.users_email_pmail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.users_email_pmail_id_seq OWNER TO postgres;

--
-- Name: users_email_pmail_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.users_email_pmail_id_seq OWNED BY users.users_email.pmail_id;


--
-- Name: users_experiences; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_experiences (
    usex_id integer NOT NULL,
    usex_entity_id integer NOT NULL,
    usex_title character varying(255),
    usex_profile_headline character varying(512),
    usex_employment_type character varying(15),
    usex_company_name character varying(255),
    usex_is_current character(1),
    usex_start_date timestamp without time zone,
    usex_end_date timestamp without time zone,
    usex_industry character varying(15),
    usex_description character varying(512),
    usex_experience_type character varying(15),
    usex_city_id integer,
    CONSTRAINT users_experiences_usex_employment_type_check CHECK (((usex_employment_type)::text = ANY ((ARRAY['fulltime'::character varying, 'freelance'::character varying])::text[]))),
    CONSTRAINT users_experiences_usex_experience_type_check CHECK (((usex_experience_type)::text = ANY ((ARRAY['company'::character varying, 'certified'::character varying, 'voluntering'::character varying, 'organization'::character varying, 'reward'::character varying])::text[]))),
    CONSTRAINT users_experiences_usex_is_current_check CHECK ((usex_is_current = ANY (ARRAY['0'::bpchar, '1'::bpchar])))
);


ALTER TABLE users.users_experiences OWNER TO postgres;

--
-- Name: users_experiences_skill; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_experiences_skill (
    uesk_usex_id integer NOT NULL,
    uesk_uski_id integer NOT NULL
);


ALTER TABLE users.users_experiences_skill OWNER TO postgres;

--
-- Name: users_experiences_usex_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.users_experiences_usex_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.users_experiences_usex_id_seq OWNER TO postgres;

--
-- Name: users_experiences_usex_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.users_experiences_usex_id_seq OWNED BY users.users_experiences.usex_id;


--
-- Name: users_license; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_license (
    usli_id integer NOT NULL,
    usli_license_code character varying(512),
    usli_modified_date timestamp without time zone,
    usli_status character varying(15),
    usli_entity_id integer NOT NULL,
    CONSTRAINT users_license_usli_status_check CHECK (((usli_status)::text = ANY ((ARRAY['Active'::character varying, 'NonActive'::character varying])::text[])))
);


ALTER TABLE users.users_license OWNER TO postgres;

--
-- Name: users_license_usli_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.users_license_usli_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.users_license_usli_id_seq OWNER TO postgres;

--
-- Name: users_license_usli_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.users_license_usli_id_seq OWNED BY users.users_license.usli_id;


--
-- Name: users_media; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_media (
    usme_id integer NOT NULL,
    usme_entity_id integer NOT NULL,
    usme_file_link character varying(255),
    usme_filename character varying(55),
    usme_filesize integer,
    usme_filetype character varying(15),
    usme_note character varying(55),
    usme_modified_date timestamp without time zone,
    CONSTRAINT users_media_usme_filetype_check CHECK (((usme_filetype)::text = ANY ((ARRAY['jpg'::character varying, 'pdf'::character varying, 'word'::character varying])::text[])))
);


ALTER TABLE users.users_media OWNER TO postgres;

--
-- Name: users_media_usme_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.users_media_usme_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.users_media_usme_id_seq OWNER TO postgres;

--
-- Name: users_media_usme_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.users_media_usme_id_seq OWNED BY users.users_media.usme_id;


--
-- Name: users_phones; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_phones (
    uspo_entity_id integer NOT NULL,
    uspo_number character varying(15) NOT NULL,
    uspo_modified_date timestamp without time zone,
    uspo_ponty_code character varying(15)
);


ALTER TABLE users.users_phones OWNER TO postgres;

--
-- Name: users_roles; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_roles (
    usro_entity_id integer NOT NULL,
    usro_role_id integer NOT NULL,
    usro_modified_date timestamp without time zone
);


ALTER TABLE users.users_roles OWNER TO postgres;

--
-- Name: users_skill; Type: TABLE; Schema: users; Owner: postgres
--

CREATE TABLE users.users_skill (
    uski_id integer NOT NULL,
    uski_entity_id integer NOT NULL,
    uski_modified_date timestamp without time zone,
    uski_skty_name character varying(15)
);


ALTER TABLE users.users_skill OWNER TO postgres;

--
-- Name: users_skill_uski_id_seq; Type: SEQUENCE; Schema: users; Owner: postgres
--

CREATE SEQUENCE users.users_skill_uski_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users.users_skill_uski_id_seq OWNER TO postgres;

--
-- Name: users_skill_uski_id_seq; Type: SEQUENCE OWNED BY; Schema: users; Owner: postgres
--

ALTER SEQUENCE users.users_skill_uski_id_seq OWNED BY users.users_skill.uski_id;