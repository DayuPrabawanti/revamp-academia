// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package gen

import (
	"database/sql"
)

type UsersBusinessEntity struct {
	EntityID int32 `db:"entity_id" json:"entityId"`
}

type UsersPhoneNumberType struct {
	PontyCode         string       `db:"ponty_code" json:"pontyCode"`
	PontyModifiedDate sql.NullTime `db:"ponty_modified_date" json:"pontyModifiedDate"`
}

type UsersRole struct {
	RoleID           int32          `db:"role_id" json:"roleId"`
	RoleName         sql.NullString `db:"role_name" json:"roleName"`
	RoleType         sql.NullString `db:"role_type" json:"roleType"`
	RoleModifiedDate sql.NullTime   `db:"role_modified_date" json:"roleModifiedDate"`
}

type UsersUser struct {
	UserEntityID       int32                 `db:"user_entity_id" json:"userEntityId"`
	UserName           sql.NullString        `db:"user_name" json:"userName"`
	UserPassword       sql.NullString        `db:"user_password" json:"userPassword"`
	UserFirstName      sql.NullString        `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString        `db:"user_last_name" json:"userLastName"`
	UserBirthDate      sql.NullTime          `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion sql.NullInt32         `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   sql.NullTime          `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          sql.NullString        `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    sql.NullInt32         `db:"user_current_role" json:"userCurrentRole"`
}

type UsersUsersAddress struct {
	EtadAddrID       int32         `db:"etad_addr_id" json:"etadAddrId"`
	EtadModifiedDate sql.NullTime  `db:"etad_modified_date" json:"etadModifiedDate"`
	EtadEntityID     sql.NullInt32 `db:"etad_entity_id" json:"etadEntityId"`
	EtadAdtyID       sql.NullInt32 `db:"etad_adty_id" json:"etadAdtyId"`
}

type UsersUsersEducation struct {
	UsduID           int32          `db:"usdu_id" json:"usduId"`
	UsduEntityID     int32          `db:"usdu_entity_id" json:"usduEntityId"`
	UsduSchool       sql.NullString `db:"usdu_school" json:"usduSchool"`
	UsduDegree       sql.NullString `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   sql.NullString `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGraduateYear sql.NullString `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduStartDate    sql.NullTime   `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      sql.NullTime   `db:"usdu_end_date" json:"usduEndDate"`
	UsduGrade        sql.NullString `db:"usdu_grade" json:"usduGrade"`
	UsduActivities   sql.NullString `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  sql.NullString `db:"usdu_description" json:"usduDescription"`
	UsduModifiedDate sql.NullTime   `db:"usdu_modified_date" json:"usduModifiedDate"`
}

type UsersUsersEmail struct {
	PmailEntityID     int32          `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32          `db:"pmail_id" json:"pmailId"`
	PmailAddress      sql.NullString `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate sql.NullTime   `db:"pmail_modified_date" json:"pmailModifiedDate"`
}

type UsersUsersExperience struct {
	UsexID              int32          `db:"usex_id" json:"usexId"`
	UsexEntityID        int32          `db:"usex_entity_id" json:"usexEntityId"`
	UsexTitle           sql.NullString `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline sql.NullString `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexEmploymentType  sql.NullString `db:"usex_employment_type" json:"usexEmploymentType"`
	UsexCompanyName     sql.NullString `db:"usex_company_name" json:"usexCompanyName"`
	UsexIsCurrent       sql.NullString `db:"usex_is_current" json:"usexIsCurrent"`
	UsexStartDate       sql.NullTime   `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         sql.NullTime   `db:"usex_end_date" json:"usexEndDate"`
	UsexIndustry        sql.NullString `db:"usex_industry" json:"usexIndustry"`
	UsexDescription     sql.NullString `db:"usex_description" json:"usexDescription"`
	UsexExperienceType  sql.NullString `db:"usex_experience_type" json:"usexExperienceType"`
	UsexCityID          sql.NullInt32  `db:"usex_city_id" json:"usexCityId"`
}

type UsersUsersExperiencesSkill struct {
	UeskUsexID int32 `db:"uesk_usex_id" json:"ueskUsexId"`
	UeskUskiID int32 `db:"uesk_uski_id" json:"ueskUskiId"`
}

type UsersUsersLicense struct {
	UsliID           int32          `db:"usli_id" json:"usliId"`
	UsliLicenseCode  sql.NullString `db:"usli_license_code" json:"usliLicenseCode"`
	UsliModifiedDate sql.NullTime   `db:"usli_modified_date" json:"usliModifiedDate"`
	UsliStatus       sql.NullString `db:"usli_status" json:"usliStatus"`
	UsliEntityID     int32          `db:"usli_entity_id" json:"usliEntityId"`
}

type UsersUsersMedium struct {
	UsmeID           int32          `db:"usme_id" json:"usmeId"`
	UsmeEntityID     int32          `db:"usme_entity_id" json:"usmeEntityId"`
	UsmeFileLink     sql.NullString `db:"usme_file_link" json:"usmeFileLink"`
	UsmeFilename     sql.NullString `db:"usme_filename" json:"usmeFilename"`
	UsmeFilesize     sql.NullInt32  `db:"usme_filesize" json:"usmeFilesize"`
	UsmeFiletype     sql.NullString `db:"usme_filetype" json:"usmeFiletype"`
	UsmeNote         sql.NullString `db:"usme_note" json:"usmeNote"`
	UsmeModifiedDate sql.NullTime   `db:"usme_modified_date" json:"usmeModifiedDate"`
}

type UsersUsersPhone struct {
	UspoEntityID     int32          `db:"uspo_entity_id" json:"uspoEntityId"`
	UspoNumber       string         `db:"uspo_number" json:"uspoNumber"`
	UspoModifiedDate sql.NullTime   `db:"uspo_modified_date" json:"uspoModifiedDate"`
	UspoPontyCode    sql.NullString `db:"uspo_ponty_code" json:"uspoPontyCode"`
}

type UsersUsersRole struct {
	UsroEntityID     int32        `db:"usro_entity_id" json:"usroEntityId"`
	UsroRoleID       int32        `db:"usro_role_id" json:"usroRoleId"`
	UsroModifiedDate sql.NullTime `db:"usro_modified_date" json:"usroModifiedDate"`
}

type UsersUsersSkill struct {
	UskiID           int32          `db:"uski_id" json:"uskiId"`
	UskiEntityID     int32          `db:"uski_entity_id" json:"uskiEntityId"`
	UskiModifiedDate sql.NullTime   `db:"uski_modified_date" json:"uskiModifiedDate"`
	UskiSktyName     sql.NullString `db:"uski_skty_name" json:"uskiSktyName"`
}
