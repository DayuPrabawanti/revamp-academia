-- name: GetBusinessEntity :one
SELECT * FROM users.business_entity
WHERE entity_id = $1;

-- name: ListBusinessEntity :many
SELECT * FROM users.business_entity
ORDER BY entity_id;

-- name: CreateBusinessEntity :one
INSERT INTO users.business_entity 
(entity_id)
VALUES($1)
RETURNING entity_id;

-- name: DeleteBusinessEntity :exec
DELETE FROM users.business_entity
WHERE entity_id = $1;


-- users

-- name: GetUsers :one
SELECT * FROM users.users
WHERE user_entity_id = $1;

-- name: ListUsers :many
SELECT * FROM users.users
ORDER BY user_name;

-- name: CreateUsers :one

INSERT INTO users.users 
(user_entity_id, user_name, user_password, user_first_name, 
user_last_name, user_birth_date, user_email_promotion, user_demographic, 
user_modified_date, user_photo, user_current_role)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
RETURNING user_entity_id;

-- name: DeleteUsers :exec
DELETE FROM users.users
WHERE user_entity_id = $1;

-- name: UpdateUsers :exec
UPDATE users.users
  set user_name = $2,
  user_password=$3,
  user_first_name= $4,
  user_last_name =$5,
  user_birth_date=$6,
  user_email_promotion=$7,
  user_demographic=$8,
  user_modified_date=$9,
  user_photo=$10,
  user_current_role=$11
WHERE user_entity_id = $1;

-- roles

-- name: GetRoles :one
SELECT * FROM users.roles
WHERE role_id = $1;

-- name: ListRoles :many
SELECT * FROM users.roles
ORDER BY role_name;

-- name: CreateRoles :one

INSERT INTO users.roles 
(role_id, role_name, role_type, role_modified_date)
VALUES($1,$2,$3,$4)
RETURNING role_id;

-- name: DeleteRoles :exec
DELETE FROM users.roles
WHERE role_id = $1;

-- name: UpdateRoles :exec
UPDATE users.roles
  set role_name = $2,
  role_type=$3,
  role_modified_date=$4
WHERE role_id = $1;

-- users_roles

-- name: GetUsersRoles :one
SELECT * FROM users.users_roles
WHERE usro_entity_id = $1;

-- name: ListUsersRoles :many
SELECT * FROM users.users_roles
ORDER BY usro_modified_date;

-- name: CreateUsersRoles :one

INSERT INTO users.users_roles
(usro_entity_id, usro_role_id, usro_modified_date)
VALUES($1,$2,$3)
RETURNING usro_entity_id;

-- name: DeleteUsersRoles :exec
DELETE FROM users.users_roles
WHERE usro_entity_id = $1;

-- name: UpdateUsersRoles :exec
UPDATE users.users_roles
  set usro_role_id = $2,
  usro_modified_date=$3
WHERE usro_entity_id = $1;

-- Users Phone Number Type

-- name: GetUsersPhoneNumberType :one
SELECT * FROM users.phone_number_type
WHERE ponty_code = $1;

-- name: ListUsersPhoneNumberType :many
SELECT * FROM users.phone_number_type
ORDER BY ponty_modified_date;

-- name: CreateUsersPhoneNumberType :one

INSERT INTO users.phone_number_type
(ponty_code, ponty_modified_date)
VALUES($1,$2)
RETURNING ponty_code;

-- name: DeleteUsersPhoneNumberType :exec
DELETE FROM users.phone_number_type
WHERE ponty_code = $1;

-- name: UpdateUsersPhoneNumberType :exec
UPDATE users.phone_number_type
  set ponty_modified_date = $2
WHERE ponty_code = $1;

-- Users Addrress

-- name: GetAddress :one
SELECT * FROM users.users_address
WHERE etad_addr_id = $1;

-- name: ListAddress :many
SELECT * FROM users.users_address
ORDER BY etad_modified_date;

-- name: CreateAddrees :one

INSERT INTO users.users_address
(etad_addr_id, etad_modified_date, etad_entity_id, etad_adty_id)
VALUES($1,$2,$3,$4)
RETURNING etad_addr_id;

-- name: DeleteAddress :exec
DELETE FROM users.users_address
WHERE etad_addr_id = $1;

-- name: UpdateAddress :exec
UPDATE users.users_address
  set etad_modified_date = $2,
  etad_entity_id = $3,
  etad_adty_id = $4
WHERE etad_addr_id = $1;


-- Users Education

-- name: GetEducation :one
SELECT * FROM users.users_education
WHERE usdu_id = $1;

-- name: ListEducation :many
SELECT * FROM users.users_education
ORDER BY usdu_entity_id;

-- name: CreateEducation :one

INSERT INTO users.users_education
(usdu_id, usdu_entity_id, usdu_school, usdu_degree, usdu_field_study,
usdu_graduate_year, usdu_start_date, usdu_end_date, usdu_grade,
usdu_activities, usdu_description, usdu_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING usdu_id;

-- name: DeleteEducation :exec
DELETE FROM users.users_education
WHERE usdu_id = $1;

-- name: UpdateEducation :exec
UPDATE users.users_education
  set usdu_entity_id = $2,
  usdu_school = $3,
  usdu_degree = $4,
  usdu_field_study = $5,
  usdu_graduate_year = $6,
  usdu_start_date = $7,
  usdu_end_date = $8,
  usdu_grade = $9,
  usdu_activities = $10,
  usdu_description = $11,
  usdu_modified_date = $12
WHERE usdu_id = $1;


-- Users Email

-- name: GetEmail :one
SELECT * FROM users.users_email
WHERE pmail_id = $1;

-- name: ListEmail :many
SELECT * FROM users.users_email
ORDER BY pmail_id;

-- name: CreateEmail :one

INSERT INTO users.users_email
(pmail_entity_id, pmail_id, pmail_address, pmail_modified_date)
VALUES($1,$2,$3,$4)
RETURNING pmail_id;

-- name: DeleteEmail :exec
DELETE FROM users.users_email
WHERE pmail_id = $1;

-- name: UpdateEmail :exec
UPDATE users.users_email
  set pmail_entity_id = $2,
  pmail_address   = $3,
  pmail_modified_date = $4
WHERE pmail_id = $1;

-- Users Experience

-- name: GetExperience :one
SELECT * FROM users.users_experiences
WHERE usex_id = $1;

-- name: ListExperience :many
SELECT * FROM users.users_experiences
ORDER BY usex_title;

-- name: CreateExperience :one

INSERT INTO users.users_experiences
(usex_id, usex_entity_id, usex_title, usex_profile_headline, usex_employment_type,
usex_company_name, usex_is_current, usex_start_date, usex_end_date, usex_industry,
usex_description, usex_experience_type, usex_city_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
RETURNING usex_id;

-- name: DeleteExperience :exec
DELETE FROM users.users_experiences
WHERE usex_id = $1;

-- name: UpdateExperience :exec
UPDATE users.users_experiences
  set 
  usex_entity_id = $2,
  usex_title      = $3,
  usex_profile_headline    = $4,
  usex_employment_type     = $5,
  usex_company_name        = $6,
  usex_is_current          = $7,
  usex_start_date         = $8,
  usex_end_date           = $9,
  usex_industry            = $10,
  usex_description       = $11,
  usex_experience_type     = $12,
  usex_city_id             = $13
WHERE usex_id = $1;


-- Users Experience Skill

-- name: GetExperienceSkill :one
SELECT * FROM users.users_experiences_skill
WHERE uesk_usex_id = $1;

-- name: ListExperienceSkill :many
SELECT * FROM users.users_experiences_skill
ORDER BY uesk_uski_id;

-- name: CreateExperienceSkill :one

INSERT INTO users.users_experiences_skill
(uesk_usex_id, uesk_uski_id)
VALUES($1,$2)
RETURNING uesk_usex_id;

-- name: DeleteExperienceSkill :exec
DELETE FROM users.users_experiences_skill
WHERE uesk_usex_id = $1;

-- name: UpdateExperienceSkill :exec
UPDATE users.users_experiences_skill
  set uesk_uski_id = $2
WHERE uesk_usex_id = $1;


-- Users License

-- name: GetLicense :one
SELECT * FROM users.users_license
WHERE usli_id = $1;

-- name: ListLicense :many
SELECT * FROM users.users_license
ORDER BY usli_id;

-- name: CreateLicense :one

INSERT INTO users.users_license
(usli_id, usli_license_code, usli_modified_date, usli_status, usli_entity_id)
VALUES($1,$2,$3,$4,$5)
RETURNING usli_id;

-- name: DeleteLicense :exec
DELETE FROM users.users_license
WHERE usli_id = $1;

-- name: UpdateLicense :exec
UPDATE users.users_license
  set usli_license_code = $2,
  usli_modified_date = $3,
  usli_status = $4,
  usli_entity_id = $5
WHERE usli_id = $1;


-- Users Media

-- name: GetMedia :one
SELECT * FROM users.users_media
WHERE usme_id = $1;

-- name: ListMedia :many
SELECT * FROM users.users_media
ORDER BY usme_id;

-- name: CreateMedia :one

INSERT INTO users.users_media
(usme_id, usme_entity_id, usme_file_link, usme_filename,
usme_filesize, usme_filetype, usme_note, usme_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING usme_id;

-- name: DeleteMedia :exec
DELETE FROM users.users_media
WHERE usme_id = $1;

-- name: UpdateMedia :exec
UPDATE users.users_media
  set usme_entity_id = $2,
  usme_file_link = $3,
  usme_filename = $4,
  usme_filesize = $5,
  usme_filetype = $6,
  usme_note = $7,
  usme_modified_date = $8
WHERE usme_id = $1;


-- Users Phones

-- name: GetPhones :one
SELECT * FROM users.users_phones
WHERE uspo_entity_id = $1;

-- name: ListPhones :many
SELECT * FROM users.users_phones
ORDER BY uspo_entity_id;

-- name: CreatePhones :one

INSERT INTO users.users_phones
(uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code)
VALUES($1,$2,$3,$4)
RETURNING uspo_entity_id;

-- name: DeletePhones :exec
DELETE FROM users.users_phones
WHERE uspo_entity_id = $1;

-- name: UpdatePhones :exec
UPDATE users.users_phones
  set uspo_number = $2,
  uspo_modified_date = $3,
  uspo_ponty_code=$4
WHERE uspo_entity_id = $1;


-- Users Skill

-- name: GetSkill :one
SELECT * FROM users.users_skill
WHERE uski_id = $1;

-- name: ListSkill :many
SELECT * FROM users.users_skill
ORDER BY uski_id;

-- name: CreateSkill :one

INSERT INTO users.users_skill
(uski_id, uski_entity_id, uski_modified_date, uski_skty_name)
VALUES($1,$2,$3,$4)
RETURNING uski_id;

-- name: DeleteSkill :exec
DELETE FROM users.users_skill
WHERE uski_id = $1;

-- name: UpdateSkill :exec
UPDATE users.users_skill
  set uski_entity_id = $2,
  uski_modified_date = $3,
  uski_skty_name= $4
WHERE uski_id = $1;


-- hr.employee

-- name: GetEmployee :one
SELECT * FROM hr.employee
WHERE emp_entity_id = $1;

-- name: ListEmployees :many
SELECT * FROM hr.employee
ORDER BY emp_emp_number;

-- name: CreateEmployee :one

INSERT INTO hr.employee (emp_entity_id, emp_emp_number, emp_national_id, emp_birth_date, emp_marital_status, emp_gender, emp_hire_date, emp_salaried_flag, emp_vacation_hours, emp_sickleave_hours, emp_current_flag, emp_modified_date, emp_type, emp_joro_id, emp_emp_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
RETURNING emp_entity_id;

-- name: DeleteEmployee :exec
DELETE FROM hr.employee
WHERE emp_entity_id = $1;

UPDATE hr.employee
  set emp_marital_status = $2,
  emp_gender = $3
WHERE emp_entity_id = $1;

-- hr.employee_department_history

-- name: GetEmployeeDepartmentHistory :one
SELECT * FROM hr.employee_department_history
WHERE edhi_id = $1;

-- name: ListEmployeeDepartmentHistory :many
SELECT * FROM hr.employee_department_history
ORDER BY edhi_id;

-- name: CreateEmployeeDepartmentHistory :one

INSERT INTO hr.employee_department_history 
(edhi_id, edhi_entity_id, edhi_start_date, edhi_end_date, edhi_modified_date, edhi_dept_id)
VALUES($1,$2,$3,$4,$5,$6)
RETURNING edhi_id;

-- name: DeleteEmployeeDepartmentHistory :exec
DELETE FROM hr.employee_department_history
WHERE edhi_id = $1;

-- name: UpdateEmployeeDepartmentHistory :exec
UPDATE hr.employee_department_history
  set edhi_start_date = $2,
  edhi_end_date = $3
WHERE edhi_id = $1;

-- hr.department

-- name: GetDepartment :one
SELECT * FROM hr.department
WHERE dept_id = $1;

-- name: ListDepartment :many
SELECT * FROM hr.department
ORDER BY dept_id;

-- name: CreateDepartment :one

INSERT INTO hr.department 
(dept_id, dept_name, dept_modified_date)
VALUES($1,$2,$3)
RETURNING dept_id;

-- name: DeleteDepartment :exec
DELETE FROM hr.department
WHERE dept_id = $1;

-- name: UpdateDepartment :exec
UPDATE hr.department
  set dept_name = $2
WHERE dept_id = $1;

-- hr.employee_client_contract

-- name: GetClientContract :one
SELECT * FROM hr.employee_client_contract
WHERE ecco_id = $1;

-- name: ListClientContract :many
SELECT * FROM hr.employee_client_contract
ORDER BY ecco_id;

-- name: CreateClientContract :one

INSERT INTO hr.employee_client_contract (ecco_id, ecco_entity_id, ecco_contract_no, ecco_contract_date, ecco_start_date, ecco_end_date, ecco_notes, ecco_modified_date, ecco_media_link, ecco_joty_id, ecco_account_manager, ecco_clit_id, ecco_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING ecco_id;

-- name: DeleteClientContract :exec
DELETE FROM hr.employee_client_contract
WHERE ecco_id = $1;

-- name: UpdateClientContract :exec
UPDATE hr.employee_client_contract
  set ecco_start_date = $2,
  ecco_end_date = $3
WHERE ecco_id = $1;

-- hr.employee_pay_history

-- name: GetPayHistory :one
SELECT * FROM hr.employee_pay_history
WHERE ephi_entity_id = $1;

-- name: ListPayHistory :many
SELECT * FROM hr.employee_pay_history
ORDER BY ephi_entity_id;

-- name: CreatePayHistory :one

INSERT INTO hr.employee_pay_history (ephi_entity_id, ephi_rate_change_date, ephi_rate_salary, ephi_pay_frequence, ephi_modified_date) VALUES ($1, $2, $3, $4, $5)
RETURNING ephi_entity_id;

-- name: DeletePayHistory :exec
DELETE FROM hr.employee_pay_history
WHERE ephi_entity_id = $1;

-- name: UpdatePayHistory :exec
UPDATE hr.employee_pay_history
  set ephi_rate_salary = $2,
  ephi_pay_frequence = $3
WHERE ephi_entity_id = $1;


-- name: Getprogram_entity :one
SELECT * FROM curriculum.program_entity
WHERE prog_entity_id = $1;

-- name: Listprogram_entity :many
SELECT * FROM curriculum.program_entity
ORDER BY prog_title;

-- name: Createprogram_entity :one

INSERT INTO curriculum.program_entity (prog_entity_id, 
prog_title, 
prog_headline, 
prog_type, 
prog_learning_type, 
prog_rating, 
prog_total_trainee, 
prog_modified_date, 
prog_image, 
prog_best_seller, 
prog_price, 
prog_language, 
prog_duration, 
prog_duration_type, 
prog_tag_skill, 
prog_city_id, 
prog_cate_id, 
prog_created_by, 
prog_status) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
RETURNING prog_entity_id;

-- name: Deleteprogram_entity :exec
DELETE FROM curriculum.program_entity
WHERE prog_entity_id = $1;

-- name: Updateprogram_entity :exec
UPDATE curriculum.program_entity
  set prog_title = $2,
  prog_headline = $3
WHERE prog_entity_id = $1;

-- curriculum.program_reviews

-- name: Getprogram_reviews :one
SELECT * FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1;

-- name: Listprogram_reviews :many
SELECT * FROM curriculum.program_reviews
ORDER BY prow_rating;

-- name: Createprogram_reviews :one

INSERT INTO curriculum.program_reviews (prow_user_entity_id, 
prow_prog_entity_id, 
prow_review, 
prow_rating, 
prow_modified_date) 

VALUES($1,$2,$3,$4,$5)
RETURNING prow_user_entity_id;

-- name: Deleteprogram_reviews :exec
DELETE FROM curriculum.program_reviews
WHERE prow_user_entity_id = $1;

-- name: Updateprogram_reviews :exec
UPDATE curriculum.program_reviews
  set prow_review = $2,
  prow_rating = $3
WHERE prow_user_entity_id = $1;

-- curriculum.sections

-- name: Getsections :one
SELECT * FROM curriculum.sections
WHERE sect_id = $1;

-- name: Listsections :many
SELECT * FROM curriculum.sections
ORDER BY sect_title;

-- name: Createsections :one

INSERT INTO curriculum.sections (sect_id, 
sect_prog_entity_id, 
sect_title, 
sect_description, 
sect_total_section, 
sect_total_lecture, 
sect_total_minute, 
sect_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING sect_id;

-- name: Deletesections :exec
DELETE FROM curriculum.sections
WHERE sect_id = $1;

-- name: Updatesections :exec
UPDATE curriculum.sections
  set sect_title = $2,
  sect_description = $3
WHERE sect_id = $1;


-- curriculum.section_detail

-- name: Getsection_detail :one
SELECT * FROM curriculum.section_detail
WHERE secd_id = $1;

-- name: Listsection_detail :many
SELECT * FROM curriculum.section_detail
ORDER BY secd_title;

-- name: Createsection_detail :one

INSERT INTO curriculum.section_detail (secd_id, 
secd_title, 
secd_preview, 
secd_score, 
secd_note, 
secd_minute, 
secd_modified_date, 
secd_sect_id)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING secd_id;

-- name: Deletesection_detail :exec
DELETE FROM curriculum.section_detail
WHERE secd_id = $1;

-- name: Updatesection_detail :exec
UPDATE curriculum.section_detail
  set secd_title = $2,
  secd_preview = $3
WHERE secd_id = $1;


-- curriculum.section_detail_material

-- name: Getsection_detail_material :one
SELECT * FROM curriculum.section_detail_material
WHERE sedm_id = $1;

-- name: Listsection_detail_material :many
SELECT * FROM curriculum.section_detail_material
ORDER BY sedm_filename;

-- name: Createsection_detail_material :one

INSERT INTO curriculum.section_detail_material (sedm_id, 
sedm_filename, 
sedm_filesize, 
sedm_filetype, 
sedm_filelink, 
sedm_modified_date, 
sedm_secd_id)

VALUES($1,$2,$3,$4,$5,$6,$7)
RETURNING sedm_id;

-- name: Deletesection_detail_material :exec
DELETE FROM curriculum.section_detail_material
WHERE sedm_id = $1;

-- name: Updatesection_detail_material :exec
UPDATE curriculum.section_detail_material
  set sedm_filename = $2,
  sedm_filesize = $3
WHERE sedm_id = $1;




-- curriculum.program_entity_description

-- name: Getprogram_entity_description :one
SELECT * FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1;

-- name: Listprogram_entity_description :many
SELECT * FROM curriculum.program_entity_description
ORDER BY pred_prog_entity_id;

-- name: Createprogram_entity_description :one

INSERT INTO curriculum.program_entity_description (pred_prog_entity_id, 
pred_item_learning, 
pred_item_include, 
pred_requirement, 
pred_description, 
pred_target_level) 
VALUES($1,$2,$3,$4,$5,$6)
RETURNING pred_prog_entity_id;

-- name: Deleteprogram_entity_description :exec
DELETE FROM curriculum.program_entity_description
WHERE pred_prog_entity_id = $1;

-- name: Updateprogram_entity_description :exec
UPDATE curriculum.program_entity_description
  set pred_item_learning= $2,
  pred_item_include = $3
WHERE pred_prog_entity_id= $1;

-- name: GetBatch :one 
SELECT * FROM bootcamp.batch
WHERE batch_id = $1;

-- name: ListBatchs :many
SELECT * FROM bootcamp.batch
ORDER BY batch_name;

-- name: CreateBatch :one
INSERT INTO bootcamp.batch
(batch_id, batch_entity_id, batch_name, batch_description, batch_start_date, batch_end_date, batch_reason, batch_type, batch_modified_date, batch_status, batch_pic_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING batch_id;

-- name: DeleteBatch :exec
DELETE  FROM bootcamp.batch
WHERE batch_id = $1;

-- name: UpdateBatch :exec
UPDATE bootcamp.batch
SET batch_name = $2,
    batch_description = $3
WHERE batch_id = $1;

-- name: GetBatchTrainee :one
SELECT * FROM bootcamp.batch_trainee
WHERE batr_id = $1;

-- name: ListBatchTrinee :many
SELECT * FROM bootcamp.batch_trainee
ORDER BY batr_id;

-- name: CreateBatchTrainee :one
INSERT INTO bootcamp.batch_trainee
(batr_id, batr_status, batr_certificated, batre_certificate_link, batr_access_token, batr_access_grant, batr_review, batr_total_score, batr_modified_date, batr_trainee_entity_id, batr_batch_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING batr_id;

-- name: DeleteBatchTrainee :exec
DELETE FROM bootcamp.batch_trainee
WHERE batr_id = $1;

-- name: UpdateBatchTrainee :exec
UPDATE bootcamp.batch_trainee
SET batr_status = $2,
    batr_review = $3
WHERE batr_id = $1;

-- name: GetBatchTraineeEvaluation :one
SELECT * FROM bootcamp.batch_trainee_evaluation
WHERE btev_id = $1;

-- name: ListBatchTraineeEvaluations :many
SELECT * FROM bootcamp.batch_trainee_evaluation
ORDER BY btev_id;

-- name: CreateBatchTraineeEvaluation :one
INSERT INTO bootcamp.batch_trainee_evaluation
(btev_id, btev_type, btev_header, btev_section, btev_skill, btev_week, btev_skor, btev_note, btev_modified_date, btev_batch_id, btev_trainee_entity_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING btev_id;

-- name: DeleteBatchTraineeEvaluation :exec
DELETE FROM bootcamp.batch_trainee_evaluation
WHERE btev_id = $1;

-- name: UpdateBatchTraineeEvaluation :exec
UPDATE bootcamp.batch_trainee_evaluation
SET btev_type = $2,
    btev_header = $3
WHERE btev_id = $1;

-- name: GetInstructorProgram :one
SELECT * FROM bootcamp.instructor_programs
WHERE batch_id = $1;

-- name: ListInstructorPrograms :many
SELECT * FROM bootcamp.instructor_programs
ORDER BY batch_id;

-- name: CreateInstructorProgram :one
INSERT INTO bootcamp.instructor_programs
(batch_id, inpro_entity_id, inpro_emp_entity_id, inpro_modified_date)
VALUES ($1, $2, $3, $4)
RETURNING batch_id;

-- name: DeleteInstructorProgram :exec
DELETE FROM bootcamp.instructor_programs
WHERE batch_id = $1;

-- name: UpdateInstructorProgram :exec
UPDATE bootcamp.instructor_programs
SET inpro_entity_id = $2,
    inpro_modified_date = $3
WHERE batch_id = $1;

-- name: GetProgramApply :one
SELECT * FROM bootcamp.program_apply
WHERE prap_user_entity_id = $1;

-- name: ListProgramApplies :many
SELECT * FROM bootcamp.program_apply
ORDER BY prap_user_entity_id;

-- name: CreateProgramApply :one
INSERT INTO bootcamp.program_apply
(prap_user_entity_id, prap_prog_entity_id, prap_test_score, prap_gpa, prap_iq_test, prap_review, prap_modified_date, prap_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING prap_user_entity_id;

-- name: DeleteProgramApply :exec
DELETE FROM bootcamp.program_apply
WHERE prap_user_entity_id = $1;

-- name: UpdateProgramApply :exec
UPDATE bootcamp.program_apply
SET prap_prog_entity_id = $2,
    prap_modified_date = $3
WHERE prap_user_entity_id = $1;

-- name: GetProgramApplyProgress :one
SELECT * FROM bootcamp.program_apply_progress
WHERE parog_id = $1;

-- name: ListProgramApplyProgresses :many
SELECT * FROM bootcamp.program_apply_progress
ORDER BY parog_id;

-- name: CreateProgramApplyProgress :one
INSERT INTO bootcamp.program_apply_progress
(parog_id, parog_user_entity_id, parog_prog_entity_id, parog_action_date, parog_modified_date, parog_comment, parog_progress_name, parog_emp_entity_id, parog_status)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING parog_id;

-- name: DeleteProgramApplyProgress :exec
DELETE FROM bootcamp.program_apply_progress
WHERE parog_id = $1;

-- name: UpdateProgramApplyProgress :exec
UPDATE bootcamp.program_apply_progress
SET parog_user_entity_id = $2,
    parog_modified_date = $3
WHERE parog_id = $1;


-- name: GetCategory :one
SELECT * FROM jobHire.job_category
WHERE joca_id = $1;

-- name: ListCategories :many
SELECT * FROM jobHire.job_category
ORDER BY joca_name;

-- name: CreateCategory :one

INSERT INTO jobHire.job_category(joca_id, joca_name, joca_modified_date) VALUES ($1, $2, $3)
RETURNING joca_id;

-- name: DeleteCategory :exec
DELETE FROM jobHire.job_category
WHERE joca_id = $1;

-- name: UpdateCategory :exec
UPDATE jobHire.job_category
  set joca_name = $2,
  joca_modified_date = $3
WHERE joca_id = $1;

-- name: GetEmployeeRange :one
SELECT * FROM jobHire.employee_range
WHERE emra_id = $1;

-- name: ListEmployeesRange :many
SELECT * FROM jobHire.employee_range
ORDER BY emra_range_max;

-- name: CreateEmployeeRange :one
INSERT INTO jobHire.employee_range(emra_id, emra_range_min, emra_range_max, emra_modified_date) VALUES ($1, $2, $3, $4)
RETURNING emra_id;

-- name: DeleteEmployeeRange :exec
DELETE FROM jobHire.employee_range
WHERE emra_id = $1;

-- name: UpdateEmployeeRange :exec
UPDATE jobHire.employee_range
  set emra_range_min = $2,
  emra_range_max = $3
WHERE emra_id = $1;


-- name: GetClient :one
SELECT * FROM jobHire.client
WHERE clit_id = $1;

-- name: ListClient :many
SELECT * FROM jobHire.client
ORDER BY clit_name;

-- name: CreateClient :one

INSERT INTO jobHire.client(clit_id, clit_name, clit_about, clit_modified_date, clit_addr_id, clit_emra_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING clit_id;

-- name: DeleteClient :exec
DELETE FROM jobHire.client
WHERE clit_id = $1;

-- name: UpdateClient :exec
UPDATE jobHire.client
  set clit_name = $2,
  clit_about = $3
WHERE clit_id = $1;


-- name: GetJobPost :one
SELECT * FROM jobHire.job_post
WHERE jopo_entity_id = $1;

-- name: ListJobPost :many
SELECT * FROM jobHire.job_post
ORDER BY jopo_title;

-- name: CreateJobPost :one
INSERT INTO jobHire.job_post(jopo_entity_id, jopo_number, jopo_title, jopo_start_date, jopo_end_date, jopo_min_salary, jopo_max_salary, jopo_min_experience, jopo_max_experience, jopo_primary_skill, jopo_secondary_skill, jopo_publish_date, jopo_modified_date, jopo_emp_entity_id, jopo_clit_id, jopo_joro_id, jopo_joty_id, jopo_joca_id, jopo_addr_id, jopo_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING jopo_entity_id;

-- name: DeleteJobPost :exec
DELETE FROM jobHire.job_post
WHERE jopo_entity_id = $1;

-- name: UpdateJobPost :exec
UPDATE jobHire.job_post
  set jopo_title = $2
WHERE jopo_entity_id = $1;


-- name: GetJobPostDesc :one
SELECT * FROM jobHire.job_post_desc
WHERE jopo_entity_id = $1;

-- name: ListJobPostDesc :many
SELECT * FROM jobHire.job_post_desc
ORDER BY jopo_entity_id;

-- name: CreateJobPostDesc :one
INSERT INTO jobHire.job_post_desc(jopo_entity_id, jopo_description, jopo_responsibility, jopo_target, jopo_benefit) VALUES ($1, $2, $3, $4, $5)
RETURNING jopo_entity_id;

-- name: DeleteJobPostDesc :exec
DELETE FROM jobHire.job_post_desc
WHERE jopo_entity_id = $1;

-- name: UpdateJobPostDesc :exec
UPDATE jobHire.job_post_desc
  set jopo_description = $2,
  jopo_responsibility = $3
WHERE jopo_entity_id = $1;


-- name: GetTalentApply :one
SELECT * FROM jobHire.talent_apply
WHERE taap_user_entity_id = $1;

-- name: ListTalentApply :many
SELECT * FROM jobHire.talent_apply
ORDER BY taap_entity_id;

-- name: CreateTalentApply :one
INSERT INTO jobHire.talent_apply(taap_user_entity_id, taap_entity_id, taap_intro, taap_scoring, taap_modified_date, taap_status) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING taap_user_entity_id;

-- name: DeleteTalentApply :exec
DELETE FROM jobHire.talent_apply
WHERE taap_user_entity_id = $1;

-- name: UpdateTalentApply :exec
UPDATE jobHire.talent_apply
  set taap_intro = $2,
  taap_scoring = $3
WHERE taap_user_entity_id = $1;


-- name: GetTalentApplyProgress :one
SELECT * FROM jobHire.talent_apply_progress
WHERE tapr_id = $1;

-- name: ListTalentApplyProgress :many
SELECT * FROM jobHire.talent_apply_progress
ORDER BY tapr_status;

-- name: CreateTalentApplyProgress :one
INSERT INTO jobHire.talent_apply_progress(tapr_id, taap_user_entity_id, taap_entity_id, tapr_modified_date, tapr_status) VALUES ($1, $2, $3, $4, $5)
RETURNING tapr_id;

-- name: DeleteTalentApplyProgress :exec
DELETE FROM jobHire.talent_apply_progress
WHERE tapr_id = $1;

-- name: UpdateTalentApplyProgress :exec
UPDATE jobHire.talent_apply_progress
  set tapr_comment = $2,
  tapr_progress_name = $3
WHERE tapr_id = $1;

-- name: Getcart_items :one
SELECT * FROM sales.cart_items
WHERE cait_id = $1;

-- name: ListCart_item :many
SELECT * FROM sales.cart_items
ORDER BY cait_quantity;


-- name: CreateCartItems :one
INSERT INTO sales.cart_items(cait_id, cait_quantity, cait_unit_price, cait_modified_date,cait_user_entity_id,cait_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING cait_id;

-- name: DeleteCart_item :exec
DELETE FROM sales.cart_items
WHERE cait_id = $1;

-- name: UpdateCart_item :exec
UPDATE sales.cart_items
  set cait_quantity = $2,
  cait_unit_price = $3
WHERE cait_id = $1;

-- name: Getsales_order_detail :one
SELECT * FROM sales.sales_order_detail
WHERE sode_id = $1;
-- name: ListSales_order_detail :many
SELECT * FROM sales.sales_order_detail
ORDER BY sode_qty;

-- name: CreateSales_order_detail :one
INSERT INTO sales.sales_order_detail(sode_id, sode_qty, sode_unit_price, sode_unit_discount,sode_line_total,sode_modified_date,sode_sohe_id,sode_prog_entity_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING sode_id;

-- name: DeleteSales_order_detail :exec
DELETE FROM sales.sales_order_detail
WHERE sode_id = $1;

-- name: UpdateSales_order_detail :exec
UPDATE sales.sales_order_detail
  set sode_qty = $2,
  sode_unit_price = $3
WHERE sode_id = $1;


-- name: Getsales_order_header :one
SELECT * FROM sales.sales_order_header
WHERE sohe_id = $1;

-- name: ListSales_order_header :many
SELECT * FROM sales.sales_order_header
ORDER BY  sohe_order_date;

-- name: CreateSales_order_header :one
INSERT INTO sales.sales_order_header(sohe_id, sohe_order_date, sohe_due_date, sohe_ship_date,sohe_order_number,sohe_account_number,sohe_trpa_code_number,sohe_subtotal,sohe_tax,sohe_total_due,sohe_license_code,sohe_modified_date,sohe_user_entity_id,sohe_status) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING sohe_id;

-- name: DeleteSales_order_header :exec
DELETE FROM sales.sales_order_header
WHERE sohe_id = $1;

-- name: UpdateSales_order_header :exec
UPDATE sales.sales_order_header
  set sohe_order_date = $2,
  sohe_due_date = $3
WHERE sohe_id = $1;


-- name: GetcSpecial_offer :one
SELECT * FROM sales.special_offer
WHERE spof_id = $1;

-- name: ListSpecial_offer :many
SELECT * FROM sales.special_offer
ORDER BY  spof_description;

-- name: CreateSales_special_offer :one
INSERT INTO sales.special_offer(spof_id, spof_description, spof_discount, spof_type,spof_start_date,spof_end_date,spof_min_qty,spof_max_qty,spof_modified_date,spof_cate_id) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING spof_id;

-- name: DeleteSales_special_offer :exec
DELETE FROM sales.special_offer
WHERE spof_id = $1;

-- name: UpdateSpecial_offer :exec
UPDATE sales.special_offer
  set spof_description = $2,
  spof_discount = $3
WHERE spof_id = $1;


-- name: GetcSpecial_offer_programs :one
SELECT * FROM sales.special_offer_programs
WHERE soco_id = $1;

-- name: ListSpecial_offer_programs :many
SELECT * FROM sales.special_offer_programs
ORDER BY soco_spof_id;

-- name: CreateSpecial_offer_programs :one
INSERT INTO sales.special_offer_programs(soco_id, soco_spof_id, soco_prog_entity_id, soco_status,soco_modified_date) VALUES ($1, $2, $3, $4, $5)
RETURNING soco_id;

-- name: DeleteSpecial_offer_programs :exec
DELETE FROM sales.special_offer_programs
WHERE soco_id = $1;

-- name: UpdateSpecial_offer_programs :exec
UPDATE sales.special_offer_programs
  set soco_spof_id = $2,
  soco_prog_entity_id = $3
WHERE soco_id = $1;


-- payment.bank

-- name: GetPaymentBank :one

SELECT * FROM payment.bank WHERE bank_entity_id = $1;

-- name: ListPaymentBank :many

SELECT * FROM payment.bank ORDER BY bank_name;

-- name: CreatePaymentBank :one

INSERT INTO
    payment.bank(
        bank_entity_id,
        bank_code,
        bank_name,
        bank_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING bank_entity_id;

-- name: DeletePaymentBank :exec

DELETE FROM payment.bank WHERE bank_entity_id = $1;

UPDATE payment.bank
set
    bank_code = $2,
    bank_name = $3
WHERE bank_entity_id = $1;

-- payment.fintech

-- name: GetPaymentFintech :one

SELECT * FROM payment.fintech WHERE fint_entity_id = $1;

-- name: ListPaymentFintech :many

SELECT * FROM payment.fintech ORDER BY fint_name;

-- name: CreatePaymentFintech :one

INSERT INTO
    payment.fintech (
        fint_entity_id,
        fint_code,
        fint_name,
        fint_modified_date
    )
VALUES ($1, $2, $3, $4) RETURNING fint_entity_id;

-- name: DeletePaymentFintech :exec

DELETE FROM payment.fintech WHERE fint_entity_id = $1;

-- name: UpdatePaymentFintech :exec

UPDATE payment.fintech
set
    fint_code = $2,
    fint_name = $3
WHERE fint_entity_id = $1;

-- payment.transaction_payment

-- name: GetPaymentTransaction_payment :one

SELECT * FROM payment.transaction_payment WHERE trpa_id = $1;

-- name: ListPaymentTransaction_payment :many

SELECT * FROM payment.transaction_payment ORDER BY trpa_code_number;

-- name: CreatePaymentTransaction_payment :one

INSERT INTO
    payment.transaction_payment (
        trpa_id,
        trpa_code_number,
        trpa_order_number,
        trpa_debit,
        trpa_credit,
        trpa_type,
        trpa_note,
        trpa_modified_date,
        trpa_source_id,
        trpa_target_id,
        trpa_user_entity_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10,
        $11
    ) RETURNING trpa_id;

-- name: DeletePaymentTransaction_payment :exec

DELETE FROM payment.transaction_payment WHERE trpa_id = $1;

-- name: UpdatePaymentTransaction_payment :exec

UPDATE
    payment.transaction_payment
set
    trpa_code_number = $2,
    trpa_order_number = $3,
    trpa_debit = $4,
    trpa_credit = $5,
    trpa_type = $6,
    trpa_note = $7,
    trpa_modified_date = $8,
    trpa_source_id = $9,
    trpa_target_id = $10,
    trpa_user_entity_id = $11
WHERE trpa_id = $1;

-- payment.users_account

-- name: GetPaymentUsers_account :one

SELECT * FROM payment.users_account WHERE usac_bank_entity_id = $1;

-- name: ListPaymentUsers_account :many

SELECT * FROM payment.users_account ORDER BY usac_account_number;

-- name: CreatePaymentUsers_account :one

INSERT INTO
    payment.users_account (
        usac_bank_entity_id,
        usac_user_entity_id,
        usac_account_number,
        usac_saldo,
        usac_type,
        usac_start_date,
        usac_end_date,
        usac_modified_date,
        usac_status
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING usac_bank_entity_id;

-- name: DeletePaymentUsers_account :exec

DELETE FROM payment.users_account WHERE usac_bank_entity_id = $1;

-- name: UpdatePaymentUsers_account :exec

UPDATE payment.users_account
set
    usac_user_entity_id = $2,
    usac_account_number = $3,
    usac_saldo = $4,
    usac_type = $5,
    usac_start_date = $6,
    usac_end_date = $7,
    usac_modified_date = $8,
    usac_status = $9
WHERE usac_bank_entity_id = $1;