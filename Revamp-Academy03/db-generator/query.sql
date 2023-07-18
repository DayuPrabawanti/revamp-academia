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
  set usex_entity_id = $2,
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
ORDER BY uesk_usex_id;

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