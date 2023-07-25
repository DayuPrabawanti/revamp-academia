// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package models

import (
	"database/sql"
	"time"
	// "github.com/tabbed/pqtype"
)

type Gabung struct {
	CurriculumProgramEntity CurriculumProgramEntity
	CurriculumSection       CurriculumSection
	// CurriculumSectionDetail CurriculumSectionDetail
	// MasterCategory          MasterCategory
}
type GetGabung struct {
	CurriculumProgramEntity CurriculumProgramEntity
	CurriculumSectionGet    []CurriculumSectionGet
}
type CreateGabung struct {
	CurriculumProgramEntity CurriculumProgramEntity
	CurriculumSection       CurriculumSection
}
type CreateCurriculumProgramEntity struct {
	ProgEntityID     int32     `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string    `db:"prog_title" json:"progTitle"`
	ProgHeadline     string    `db:"prog_headline" json:"progHeadline"`
	ProgType         string    `db:"prog_type" json:"progType"`
	ProgLearningType string    `db:"prog_learning_type" json:"progLearningType"`
	ProgModifiedDate time.Time `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string    `db:"prog_image" json:"progImage"`
	ProgPrice        int32     `db:"prog_price" json:"progPrice"`
	ProgLanguage     string    `db:"prog_language" json:"progLanguage" check:"'bahasa' OR 'english'"`
	ProgDuration     int32     `db:"prog_duration" json:"progDuration"`
	ProgDurationType string    `db:"prog_duration_type" json:"progDurationType"`
	ProgCateID       int32     `db:"prog_cate_id" json:"progCateId"`

}


type MasterJobRoleList []MasterJobRole
type UsersUserList []UsersUser
type HrEmployeeList []HrEmployee
type BootcampBatchList []BootcampBatch
type JobhireTalentApplyList []JobhireTalentApply
type MasterStatusList []MasterStatus
type UsersUsersSkillList []UsersUsersSkill 

type TalentsDetailMockup struct{
	UsersUser 			UsersUser
	BootcampBatch 		BootcampBatch
	BootcampBatchTrainee BootcampBatchTrainee
	BootcampBatchTraineeEvaluation BootcampBatchTraineeEvaluation
	UsersUsersEmail		UsersUsersEmail
	UsersUsersPhone		UsersUsersPhone
	CurriculumProgramEntity CurriculumProgramEntity
	JobhireJobPost 		JobhireJobPost
	JobhireClient 		JobhireClient
	HrEmployeeClientContract HrEmployeeClientContract
	
}

type TalentsMockup struct{
	UsersUser 			UsersUser
	BootcampBatch 		BootcampBatch
	BootcampBatchTraineeEvaluation BootcampBatchTraineeEvaluation
	CurriculumProgramEntity		CurriculumProgramEntity
}

type EmployeeMockupModel struct{
	General 	HrEmployee
	Department 	HrEmployeePayHistory
	Assigment 	HrEmployeeDepartmentHistory
}


type MergeJobAndMaster struct{
	JobHirePost JobhireJobPost
	MasterAddress MasterAddress
	MasterCity MasterCity
	
}

type BootcampBatch struct {
	BatchID           int32          `db:"batch_id" json:"batchId"`
	BatchEntityID     int32          `db:"batch_entity_id" json:"batchEntityId"`
	BatchName         sql.NullString `db:"batch_name" json:"batchName"`
	BatchDescription  sql.NullString `db:"batch_description" json:"batchDescription"`
	BatchStartDate    sql.NullTime   `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate      sql.NullTime   `db:"batch_end_date" json:"batchEndDate"`
	BatchReason       sql.NullString `db:"batch_reason" json:"batchReason"`
	BatchType         sql.NullString `db:"batch_type" json:"batchType"`
	BatchModifiedDate sql.NullTime   `db:"batch_modified_date" json:"batchModifiedDate"`
	BatchStatus       sql.NullString `db:"batch_status" json:"batchStatus"`
	BatchPicID        sql.NullInt32  `db:"batch_pic_id" json:"batchPicId"`
}

type BootcampBatchTrainee struct {
	BatrID               int32          `db:"batr_id" json:"batrId"`
	BatrStatus           sql.NullString `db:"batr_status" json:"batrStatus"`
	BatrCertificated     sql.NullString `db:"batr_certificated" json:"batrCertificated"`
	BatreCertificateLink sql.NullString `db:"batre_certificate_link" json:"batreCertificateLink"`
	BatrAccessToken      sql.NullString `db:"batr_access_token" json:"batrAccessToken"`
	BatrAccessGrant      sql.NullString `db:"batr_access_grant" json:"batrAccessGrant"`
	BatrReview           sql.NullString `db:"batr_review" json:"batrReview"`
	BatrTotalScore       sql.NullInt32  `db:"batr_total_score" json:"batrTotalScore"`
	BatrModifiedDate     sql.NullTime   `db:"batr_modified_date" json:"batrModifiedDate"`
	BatrTraineeEntityID  sql.NullInt32  `db:"batr_trainee_entity_id" json:"batrTraineeEntityId"`
	BatrBatchID          int32          `db:"batr_batch_id" json:"batrBatchId"`
}

type BootcampBatchTraineeEvaluation struct {
	BtevID              int32          `db:"btev_id" json:"btevId"`
	BtevType            sql.NullString `db:"btev_type" json:"btevType"`
	BtevHeader          sql.NullString `db:"btev_header" json:"btevHeader"`
	BtevSection         sql.NullString `db:"btev_section" json:"btevSection"`
	BtevSkill           sql.NullString `db:"btev_skill" json:"btevSkill"`
	BtevWeek            sql.NullInt32  `db:"btev_week" json:"btevWeek"`
	BtevSkor            sql.NullInt32  `db:"btev_skor" json:"btevSkor"`
	BtevNote            sql.NullString `db:"btev_note" json:"btevNote"`
	BtevModifiedDate    sql.NullTime   `db:"btev_modified_date" json:"btevModifiedDate"`
	BtevBatchID         sql.NullInt32  `db:"btev_batch_id" json:"btevBatchId"`
	BtevTraineeEntityID sql.NullInt32  `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
}

type BootcampInstructorProgram struct {
	BatchID           int32        `db:"batch_id" json:"batchId"`
	InproEntityID     int32        `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32        `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate sql.NullTime `db:"inpro_modified_date" json:"inproModifiedDate"`
}

type BootcampProgramApply struct {
	PrapUserEntityID int32          `db:"prap_user_entity_id" json:"prapUserEntityId"`
	PrapProgEntityID int32          `db:"prap_prog_entity_id" json:"prapProgEntityId"`
	PrapTestScore    sql.NullInt32  `db:"prap_test_score" json:"prapTestScore"`
	PrapGpa          sql.NullInt32  `db:"prap_gpa" json:"prapGpa"`
	PrapIqTest       sql.NullInt32  `db:"prap_iq_test" json:"prapIqTest"`
	PrapReview       sql.NullString `db:"prap_review" json:"prapReview"`
	PrapModifiedDate sql.NullTime   `db:"prap_modified_date" json:"prapModifiedDate"`
	PrapStatus       sql.NullString `db:"prap_status" json:"prapStatus"`
}

type BootcampProgramApplyProgress struct {
	ParogID           int32          `db:"parog_id" json:"parogId"`
	ParogUserEntityID int32          `db:"parog_user_entity_id" json:"parogUserEntityId"`
	ParogProgEntityID int32          `db:"parog_prog_entity_id" json:"parogProgEntityId"`
	ParogActionDate   sql.NullTime   `db:"parog_action_date" json:"parogActionDate"`
	ParogModifiedDate sql.NullTime   `db:"parog_modified_date" json:"parogModifiedDate"`
	ParogComment      sql.NullString `db:"parog_comment" json:"parogComment"`
	ParogProgressName sql.NullString `db:"parog_progress_name" json:"parogProgressName"`
	ParogEmpEntityID  sql.NullInt32  `db:"parog_emp_entity_id" json:"parogEmpEntityId"`
	ParogStatus       sql.NullString `db:"parog_status" json:"parogStatus"`
}

type MergeMockUser struct {
	Users     UsersUser
	Education UsersUsersEducation
	Media     UsersUsersMedia
}



type CurriculumProgramEntity struct {
	ProgEntityID     int32          `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string         `db:"prog_title" json:"progTitle"`
	ProgHeadline     string         `db:"prog_headline" json:"progHeadline"`
	ProgType         string         `db:"prog_type" json:"progType"`
	ProgLearningType string         `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       sql.NullString `db:"prog_rating" json:"progRating"`
	ProgTotalTraniee sql.NullString `db:"prog_total_traniee" json:"progTotalTraniee"`
	ProgModifiedDate time.Time      `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string         `db:"prog_image" json:"progImage"`
	ProgBestSeller   sql.NullString `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        int32          `db:"prog_price" json:"progPrice"`
	ProgLanguage     string         `db:"prog_language" json:"progLanguage" check:"'bahasa' OR 'english'"`
	ProgDuration     int32          `db:"prog_duration" json:"progDuration"`
	ProgDurationType string         `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     sql.NullString `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       sql.NullString `db:"prog_city_id" json:"progCityId"`
	ProgCateID       int32          `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    sql.NullInt32  `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       string         `db:"prog_status" json:"progStatus"`
}


type CurriculumProgramEntityDescription struct {
	PredProgEntityID int32                 `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredItemInclude  sql.NullString `db:"pred_item_include" json:"predItemInclude"`
	PredRequirment  sql.NullString `db:"pred_requirement" json:"predRequirement"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
	PredTargetLevel  sql.NullString `db:"pred_target_level" json:"predTargetLevel"`
}

type CurriculumProgramReview struct {
	ProwUserEntityID int32     `db:"prow_user_entity_id" json:"prowUserEntityId"`
	ProwProgEntityID int32     `db:"prow_prog_entity_id" json:"prowProgEntityId"`
	ProwReview       string    `db:"prow_review" json:"prowReview"`
	ProwRating       int32     `db:"prow_rating" json:"prowRating"`
	ProwModifiedDate time.Time `db:"prow_modified_date" json:"prowModifiedDate"`
}

type CurriculumSection struct {
	SectID           int32          `db:"sect_id" json:"sectId"`
	SectProgEntityID int32          `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SectTitle        string         `db:"sect_title" json:"sectTitle"`
	SectDescription  sql.NullString `db:"sect_description" json:"sectDescription"`
	SectTotalSection int32          `db:"sect_total_section" json:"sectTotalSection"`
	SectTotalLecture int32          `db:"sect_total_lecture" json:"sectTotalLecture"`
	SectTotalMinute  int32          `db:"sect_total_minute" json:"sectTotalMinute"`
	SectModifiedDate time.Time      `db:"sect_modified_date" json:"sectModifiedDate"`
}

type CurriculumSectionGet struct {
	SectProgEntityID int32          `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SectTitle        string         `db:"sect_title" json:"sectTitle"`
	SectDescription  sql.NullString `db:"sect_description" json:"sectDescription"`
	SectTotalMinute  int32          `db:"sect_total_minute" json:"sectTotalMinute"`
	SectTotalSection sql.NullInt32  `db:"sect_total_section" json:"sectTotalSection"`
	SectTotalLecture sql.NullInt32  `db:"sect_total_lecture" json:"sectTotalLecture"`
	SectModifiedDate sql.NullTime   `db:"sect_modified_date" json:"sectModifiedDate"`
}

type CurriculumSectionDetail struct {
	SecdID           int32          `db:"secd_id" json:"secdId"`
	SecdTitle        string         `db:"secd_title" json:"secdTitle"`
	SecdPreview      string         `db:"secd_preview" json:"secdPreview"`
	SecdScore        int32          `db:"secd_score" json:"secdScore"`
	SecdNote         sql.NullString `db:"secd_note" json:"secdNote"`
	SecdMinute       int32          `db:"secd_minute" json:"secdMinute"`
	SecdModifiedDate time.Time      `db:"secd_modified_date" json:"secdModifiedDate"`
	SecdSectID       int32          `db:"secd_sect_id" json:"secdSectId"`
}

type CurriculumSectionDetailMaterial struct {
	SedmID           int32     `db:"sedm_id" json:"sedmId"`
	SedmFilename     string    `db:"sedm_filename" json:"sedmFilename"`
	SedmFilesize     int32     `db:"sedm_filesize" json:"sedmFilesize"`
	SedmFiletype     string    `db:"sedm_filetype" json:"sedmFiletype"`
	SedmFilelink     string    `db:"sedm_filelink" json:"sedmFilelink"`
	SedmModifiedDate time.Time `db:"sedm_modified_date" json:"sedmModifiedDate"`
	SedmSecdID       int32     `db:"sedm_secd_id" json:"sedmSecdId"`
}


type HrDepartment struct {
	DeptID           int32          `db:"dept_id" json:"deptId"`
	DeptName         sql.NullString `db:"dept_name" json:"deptName"`
	DeptModifiedDate sql.NullTime   `db:"dept_modified_date" json:"deptModifiedDate"`
}

type HrEmployee struct {
	EmpEntityID       int32          `db:"emp_entity_id" json:"empEntityId"`
	EmpEmpNumber      sql.NullString `db:"emp_emp_number" json:"empEmpNumber"`
	EmpNationalID     sql.NullString `db:"emp_national_id" json:"empNationalId"`
	EmpBirthDate      sql.NullTime   `db:"emp_birth_date" json:"empBirthDate"`
	EmpMaritalStatus  sql.NullString `db:"emp_marital_status" json:"empMaritalStatus"`
	EmpGender         sql.NullString `db:"emp_gender" json:"empGender"`
	EmpHireDate       sql.NullTime   `db:"emp_hire_date" json:"empHireDate"`
	EmpSalariedFlag   sql.NullString `db:"emp_salaried_flag" json:"empSalariedFlag"`
	EmpVacationHours  sql.NullInt16  `db:"emp_vacation_hours" json:"empVacationHours"`
	EmpSickleaveHours sql.NullInt16  `db:"emp_sickleave_hours" json:"empSickleaveHours"`
	EmpCurrentFlag    sql.NullInt16  `db:"emp_current_flag" json:"empCurrentFlag"`
	EmpModifiedDate   sql.NullTime   `db:"emp_modified_date" json:"empModifiedDate"`
	EmpType           sql.NullString `db:"emp_type" json:"empType"`
	EmpJoroID         sql.NullInt32  `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    sql.NullInt32  `db:"emp_emp_entity_id" json:"empEmpEntityId"`
}

type HrEmployeeClientContract struct {
	EccoID             int32          `db:"ecco_id" json:"eccoId"`
	EccoEntityID       int32          `db:"ecco_entity_id" json:"eccoEntityId"`
	EccoContractNo     sql.NullString `db:"ecco_contract_no" json:"eccoContractNo"`
	EccoContractDate   sql.NullTime   `db:"ecco_contract_date" json:"eccoContractDate"`
	EccoStartDate      sql.NullTime   `db:"ecco_start_date" json:"eccoStartDate"`
	EccoEndDate        sql.NullTime   `db:"ecco_end_date" json:"eccoEndDate"`
	EccoNotes          sql.NullString `db:"ecco_notes" json:"eccoNotes"`
	EccoModifiedDate   sql.NullTime   `db:"ecco_modified_date" json:"eccoModifiedDate"`
	EccoMediaLink      sql.NullString `db:"ecco_media_link" json:"eccoMediaLink"`
	EccoJotyID         sql.NullInt32  `db:"ecco_joty_id" json:"eccoJotyId"`
	EccoAccountManager sql.NullInt32  `db:"ecco_account_manager" json:"eccoAccountManager"`
	EccoClitID         sql.NullInt32  `db:"ecco_clit_id" json:"eccoClitId"`
	EccoStatus         sql.NullString `db:"ecco_status" json:"eccoStatus"`
}

type HrEmployeeDepartmentHistory struct {
	EdhiID           int32         `db:"edhi_id" json:"edhiId"`
	EdhiEntityID     int32         `db:"edhi_entity_id" json:"edhiEntityId"`
	EdhiStartDate    sql.NullTime  `db:"edhi_start_date" json:"edhiStartDate"`
	EdhiEndDate      sql.NullTime  `db:"edhi_end_date" json:"edhiEndDate"`
	EdhiModifiedDate sql.NullTime  `db:"edhi_modified_date" json:"edhiModifiedDate"`
	EdhiDeptID       sql.NullInt32 `db:"edhi_dept_id" json:"edhiDeptId"`
}

type HrEmployeePayHistory struct {
	EphiEntityID       int32         `db:"ephi_entity_id" json:"ephiEntityId"`
	EphiRateChangeDate time.Time     `db:"ephi_rate_change_date" json:"ephiRateChangeDate"`
	EphiRateSalary     sql.NullInt32 `db:"ephi_rate_salary" json:"ephiRateSalary"`
	EphiPayFrequence   sql.NullInt16 `db:"ephi_pay_frequence" json:"ephiPayFrequence"`
	EphiModifiedDate   sql.NullTime  `db:"ephi_modified_date" json:"ephiModifiedDate"`
}

type JobhireClient struct {
	ClitID           int32          `db:"clit_id" json:"clitId"`
	ClitName         sql.NullString `db:"clit_name" json:"clitName"`
	ClitAbout        sql.NullString `db:"clit_about" json:"clitAbout"`
	ClitModifiedDate sql.NullTime   `db:"clit_modified_date" json:"clitModifiedDate"`
	ClitAddrID       sql.NullInt32  `db:"clit_addr_id" json:"clitAddrId"`
	ClitEmraID       sql.NullInt32  `db:"clit_emra_id" json:"clitEmraId"`
}

type JobhireEmployeeRange struct {
	EmraID           int32         `db:"emra_id" json:"emraId"`
	EmraRangeMin     sql.NullInt32 `db:"emra_range_min" json:"emraRangeMin"`
	EmraRangeMax     sql.NullInt32 `db:"emra_range_max" json:"emraRangeMax"`
	EmraModifiedDate sql.NullTime  `db:"emra_modified_date" json:"emraModifiedDate"`
}

type JobhireJobCategory struct {
	JocaID           int32          `db:"joca_id" json:"jocaId"`
	JocaName         string `db:"joca_name" json:"jocaName"`
	JocaModifiedDate sql.NullTime   `db:"joca_modified_date" json:"jocaModifiedDate"`
}

type JobhireJobPhoto struct {
	JophoID           int32          `db:"jopho_id" json:"jophoId"`
	JophoFilename     sql.NullString `db:"jopho_filename" json:"jophoFilename"`
	JophoFilesize     sql.NullInt32  `db:"jopho_filesize" json:"jophoFilesize"`
	JophoFiletype     sql.NullString `db:"jopho_filetype" json:"jophoFiletype"`
	JophoModifiedDate sql.NullTime   `db:"jopho_modified_date" json:"jophoModifiedDate"`
	JophoEntityID     sql.NullInt32  `db:"jopho_entity_id" json:"jophoEntityId"`
}

type JobhireJobPost struct {
	JopoEntityID       int32          `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoNumber         sql.NullString `db:"jopo_number" json:"jopoNumber"`
	JopoTitle          sql.NullString `db:"jopo_title" json:"jopoTitle"`
	JopoStartDate      sql.NullTime   `db:"jopo_start_date" json:"jopoStartDate"`
	JopoEndDate        sql.NullTime   `db:"jopo_end_date" json:"jopoEndDate"`
	JopoMinSalary      sql.NullInt32  `db:"jopo_min_salary" json:"jopoMinSalary"`
	JopoMaxSalary      sql.NullInt32  `db:"jopo_max_salary" json:"jopoMaxSalary"`
	JopoMinExperience  sql.NullInt32  `db:"jopo_min_experience" json:"jopoMinExperience"`
	JopoMaxExperience  sql.NullInt32  `db:"jopo_max_experience" json:"jopoMaxExperience"`
	JopoPrimarySkill   sql.NullString `db:"jopo_primary_skill" json:"jopoPrimarySkill"`
	JopoSecondarySkill sql.NullString `db:"jopo_secondary_skill" json:"jopoSecondarySkill"`
	JopoPublishDate    sql.NullTime   `db:"jopo_publish_date" json:"jopoPublishDate"`
	JopoModifiedDate   sql.NullTime   `db:"jopo_modified_date" json:"jopoModifiedDate"`
	JopoEmpEntityID    sql.NullInt32  `db:"jopo_emp_entity_id" json:"jopoEmpEntityId"`
	JopoClitID         sql.NullInt32  `db:"jopo_clit_id" json:"jopoClitId"`
	JopoJoroID         sql.NullInt32  `db:"jopo_joro_id" json:"jopoJoroId"`
	JopoJotyID         sql.NullInt32  `db:"jopo_joty_id" json:"jopoJotyId"`
	JopoJocaID         sql.NullInt32  `db:"jopo_joca_id" json:"jopoJocaId"`
	JopoAddrID         int32  `db:"jopo_addr_id" json:"jopoAddrId"`
	JopoWorkCode       sql.NullString `db:"jopo_work_code" json:"jopoWorkCode"`
	JopoEduCode        sql.NullString `db:"jopo_edu_code" json:"jopoEduCode"`
	JopoInduCode       sql.NullString `db:"jopo_indu_code" json:"jopoInduCode"`
	JopoStatus         sql.NullString `db:"jopo_status" json:"jopoStatus"`
}

type JobhireJobPostDesc struct {
	JopoEntityID       int32                 `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoDescription    sql.NullString `db:"jopo_description" json:"jopoDescription"`
	JopoResponsibility sql.NullString `db:"jopo_responsibility" json:"jopoResponsibility"`
	JopoTarget         sql.NullString `db:"jopo_target" json:"jopoTarget"`
	JopoBenefit        sql.NullString `db:"jopo_benefit" json:"jopoBenefit"`
}

type JobhireTalentApply struct {
	TaapUserEntityID int32          `db:"taap_user_entity_id" json:"taapUserEntityId"`
	TaapEntityID     int32          `db:"taap_entity_id" json:"taapEntityId"`
	TaapIntro        sql.NullString `db:"taap_intro" json:"taapIntro"`
	TaapScoring      sql.NullInt32  `db:"taap_scoring" json:"taapScoring"`
	TaapModifiedDate sql.NullTime   `db:"taap_modified_date" json:"taapModifiedDate"`
	TaapStatus       sql.NullString `db:"taap_status" json:"taapStatus"`
}

type JobhireTalentApplyProgress struct {
	TaprID           int32          `db:"tapr_id" json:"taprId"`
	TaapUserEntityID int32          `db:"taap_user_entity_id" json:"taapUserEntityId"`
	TaapEntityID     int32          `db:"taap_entity_id" json:"taapEntityId"`
	TaprModifiedDate sql.NullTime   `db:"tapr_modified_date" json:"taprModifiedDate"`
	TaprStatus       sql.NullString `db:"tapr_status" json:"taprStatus"`
	TaprComment      sql.NullString `db:"tapr_comment" json:"taprComment"`
	TaprProgressName sql.NullString `db:"tapr_progress_name" json:"taprProgressName"`
}

type MasterAddress struct {
	AddrID              int32          `db:"addr_id" json:"addrId"`
	AddrLine1           sql.NullString `db:"addr_line1" json:"addrLine1"`
	AddrLine2           sql.NullString `db:"addr_line2" json:"addrLine2"`
	AddrPostalCode      sql.NullString `db:"addr_postal_code" json:"addrPostalCode"`
	AddrSpatialLocation sql.NullString `db:"addr_spatial_location" json:"addrSpatialLocation"`
	AddrModifiedDate    sql.NullTime   `db:"addr_modified_date" json:"addrModifiedDate"`
	AddrCityID          int32  `db:"addr_city_id" json:"addrCityId"`
}

type MasterAddressType struct {
	AdtyID           int32          `db:"adty_id" json:"adtyId"`
	AdtyName         sql.NullString `db:"adty_name" json:"adtyName"`
	AdtyModifiedDate sql.NullTime   `db:"adty_modified_date" json:"adtyModifiedDate"`
}

type MasterCategory struct {
	CateID           int32         `db:"cate_id" json:"cateId"`
	CateName         string        `db:"cate_name" json:"cateName"`
	CateCateID       sql.NullInt32 `db:"cate_cate_id" json:"cateCateId"`
	CateModifiedDate time.Time     `db:"cate_modified_date" json:"cateModifiedDate"`
}

type MasterCity struct {
	CityID           int32          `db:"city_id" json:"cityId"`
	CityName         sql.NullString `db:"city_name" json:"cityName"`
	CityModifiedDate sql.NullTime   `db:"city_modified_date" json:"cityModifiedDate"`
	CityProvID       sql.NullInt32  `db:"city_prov_id" json:"cityProvId"`
}

type MasterCountry struct {
	CountryCode         string         `db:"country_code" json:"countryCode"`
	CountryName         sql.NullString `db:"country_name" json:"countryName"`
	CountryModifiedDate sql.NullTime   `db:"country_modified_date" json:"countryModifiedDate"`
}

type MasterEducation struct {
	EduCode string         `db:"edu_code" json:"eduCode"`
	EduName sql.NullString `db:"edu_name" json:"eduName"`
}

type MasterIndustry struct {
	InduCode string         `db:"indu_code" json:"induCode"`
	InduName sql.NullString `db:"indu_name" json:"induName"`
}

type MasterJobRole struct {
	JoroID           int32          `db:"joro_id" json:"joroId"`
	JoroName         sql.NullString `db:"joro_name" json:"joroName"`
	JoroModifiedDate sql.NullTime   `db:"joro_modified_date" json:"joroModifiedDate"`
}

type MasterJobType struct {
	JotyID   int32          `db:"joty_id" json:"jotyId"`
	JotyName sql.NullString `db:"joty_name" json:"jotyName"`
}

type MasterModule struct {
	ModuleName string `db:"module_name" json:"moduleName"`
}

type MasterProvince struct {
	ProvID           int32          `db:"prov_id" json:"provId"`
	ProvCode         sql.NullString `db:"prov_code" json:"provCode"`
	ProvName         sql.NullString `db:"prov_name" json:"provName"`
	ProvModifiedDate sql.NullTime   `db:"prov_modified_date" json:"provModifiedDate"`
	ProvCountryCode  sql.NullString `db:"prov_country_code" json:"provCountryCode"`
}

type MasterRouteAction struct {
	RoacID         int32          `db:"roac_id" json:"roacId"`
	RoacName       sql.NullString `db:"roac_name" json:"roacName"`
	RoacOrderby    sql.NullInt32  `db:"roac_orderby" json:"roacOrderby"`
	RoacDisplay    sql.NullInt32  `db:"roac_display" json:"roacDisplay"`
	RoacModuleName sql.NullString `db:"roac_module_name" json:"roacModuleName"`
}

type MasterSkillTemplate struct {
	SkteID           int32          `db:"skte_id" json:"skteId"`
	SkteSkill        sql.NullString `db:"skte_skill" json:"skteSkill"`
	SkteDescription  sql.NullString `db:"skte_description" json:"skteDescription"`
	SkteWeek         sql.NullInt32  `db:"skte_week" json:"skteWeek"`
	SkteOrderby      sql.NullInt32  `db:"skte_orderby" json:"skteOrderby"`
	SkteModifiedDate sql.NullTime   `db:"skte_modified_date" json:"skteModifiedDate"`
	SktyName         sql.NullString `db:"skty_name" json:"sktyName"`
	SkteSkteID       sql.NullInt32  `db:"skte_skte_id" json:"skteSkteId"`
}

type MasterSkillType struct {
	SktyName string `db:"skty_name" json:"sktyName"`
}

type MasterStatus struct {
	Status             string         `db:"status" json:"status"`
	StatusModifiedDate sql.NullTime   `db:"status_modified_date" json:"statusModifiedDate"`
	StatusModule       sql.NullString `db:"status_module" json:"statusModule"`
}

type MasterWorkingType struct {
	WotyCode string         `db:"woty_code" json:"wotyCode"`
	WotyName sql.NullString `db:"woty_name" json:"wotyName"`
}

type PaymentBank struct {
	BankEntityID     int32          `db:"bank_entity_id" json:"bankEntityId"`
	BankCode         sql.NullString `db:"bank_code" json:"bankCode"`
	BankName         sql.NullString `db:"bank_name" json:"bankName"`
	BankModifiedDate sql.NullTime   `db:"bank_modified_date" json:"bankModifiedDate"`
}

type PaymentFintech struct {
	FintEntityID     int32          `db:"fint_entity_id" json:"fintEntityId"`
	FintCode         sql.NullString `db:"fint_code" json:"fintCode"`
	FintName         sql.NullString `db:"fint_name" json:"fintName"`
	FintModifiedDate sql.NullTime   `db:"fint_modified_date" json:"fintModifiedDate"`
}

type PaymentTransactionPayment struct {
	TrpaID           int32          `db:"trpa_id" json:"trpaId"`
	TrpaCodeNumber   sql.NullString `db:"trpa_code_number" json:"trpaCodeNumber"`
	TrpaOrderNumber  sql.NullString `db:"trpa_order_number" json:"trpaOrderNumber"`
	TrpaDebit        sql.NullString `db:"trpa_debit" json:"trpaDebit"`
	TrpaCredit       sql.NullString `db:"trpa_credit" json:"trpaCredit"`
	TrpaType         sql.NullString `db:"trpa_type" json:"trpaType"`
	TrpaNote         sql.NullString `db:"trpa_note" json:"trpaNote"`
	TrpaModifiedDate sql.NullTime   `db:"trpa_modified_date" json:"trpaModifiedDate"`
	TrpaSourceID     string         `db:"trpa_source_id" json:"trpaSourceId"`
	TrpaTargetID     string         `db:"trpa_target_id" json:"trpaTargetId"`
	TrpaUserEntityID sql.NullInt32  `db:"trpa_user_entity_id" json:"trpaUserEntityId"`
}

type PaymentUsersAccount struct {
	UsacBankEntityID  int32          `db:"usac_bank_entity_id" json:"usacBankEntityId"`
	UsacUserEntityID  int32          `db:"usac_user_entity_id" json:"usacUserEntityId"`
	UsacAccountNumber sql.NullString `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         sql.NullString `db:"usac_saldo" json:"usacSaldo"`
	UsacType          sql.NullString `db:"usac_type" json:"usacType"`
	UsacStartDate     sql.NullTime   `db:"usac_start_date" json:"usacStartDate"`
	UsacEndDate       sql.NullTime   `db:"usac_end_date" json:"usacEndDate"`
	UsacModifiedDate  sql.NullTime   `db:"usac_modified_date" json:"usacModifiedDate"`
	UsacStatus        sql.NullString `db:"usac_status" json:"usacStatus"`
}

type SalesCartItem struct {
	CaitID           int32          `db:"cait_id" json:"caitId"`
	CaitQuantity     int32          `db:"cait_quantity" json:"caitQuantity"`
	CaitUnitPrice    sql.NullString `db:"cait_unit_price" json:"caitUnitPrice"`
	CaitModifiedDate sql.NullString `db:"cait_modified_date" json:"caitModifiedDate"`
	CaitUserEntityID int32          `db:"cait_user_entity_id" json:"caitUserEntityId"`
	CaitProgEntityID int32          `db:"cait_prog_entity_id" json:"caitProgEntityId"`
}

type SalesSalesOrderDetail struct {
	SodeID           int32  `db:"sode_id" json:"sodeId"`
	SodeQty          int32  `db:"sode_qty" json:"sodeQty"`
	SodeUnitPrice    string `db:"sode_unit_price" json:"sodeUnitPrice"`
	SodeUnitDiscount string `db:"sode_unit_discount" json:"sodeUnitDiscount"`
	SodeLineTotal    int32  `db:"sode_line_total" json:"sodeLineTotal"`
	SodeModifiedDate string `db:"sode_modified_date" json:"sodeModifiedDate"`
	SodeSoheID       int32  `db:"sode_sohe_id" json:"sodeSoheId"`
	SodeProgEntityID int32  `db:"sode_prog_entity_id" json:"sodeProgEntityId"`
}

type SalesSalesOrderHeader struct {
	SoheID             int32  `db:"sohe_id" json:"soheId"`
	SoheOrderDate      string `db:"sohe_order_date" json:"soheOrderDate"`
	SoheDueDate        string `db:"sohe_due_date" json:"soheDueDate"`
	SoheShipDate       string `db:"sohe_ship_date" json:"soheShipDate"`
	SoheOrderNumber    string `db:"sohe_order_number" json:"soheOrderNumber"`
	SoheAccountNumber  string `db:"sohe_account_number" json:"soheAccountNumber"`
	SoheTrpaCodeNumber string `db:"sohe_trpa_code_number" json:"soheTrpaCodeNumber"`
	SoheSubtotal       string `db:"sohe_subtotal" json:"soheSubtotal"`
	SoheTax            string `db:"sohe_tax" json:"soheTax"`
	SoheTotalDue       int32  `db:"sohe_total_due" json:"soheTotalDue"`
	SoheLicenseCode    string `db:"sohe_license_code" json:"soheLicenseCode"`
	SoheModifiedDate   string `db:"sohe_modified_date" json:"soheModifiedDate"`
	SoheUserEntityID   int32  `db:"sohe_user_entity_id" json:"soheUserEntityId"`
	SoheStatus         string `db:"sohe_status" json:"soheStatus"`
}

type SalesSpecialOffer struct {
	SpofID           int32          `db:"spof_id" json:"spofId"`
	SpofDescription  string         `db:"spof_description" json:"spofDescription"`
	SpofDiscount     int32          `db:"spof_discount" json:"spofDiscount"`
	SpofType         sql.NullInt32  `db:"spof_type" json:"spofType"`
	SpofStartDate    string         `db:"spof_start_date" json:"spofStartDate"`
	SpofEndDate      string         `db:"spof_end_date" json:"spofEndDate"`
	SpofMinQty       int32          `db:"spof_min_qty" json:"spofMinQty"`
	SpofMaxQty       int32          `db:"spof_max_qty" json:"spofMaxQty"`
	SpofModifiedDate sql.NullString `db:"spof_modified_date" json:"spofModifiedDate"`
	SpofCateID       int32          `db:"spof_cate_id" json:"spofCateId"`
}

type SalesSpecialOfferProgram struct {
	SocoID           int32  `db:"soco_id" json:"socoId"`
	SocoSpofID       int32  `db:"soco_spof_id" json:"socoSpofId"`
	SocoProgEntityID int32  `db:"soco_prog_entity_id" json:"socoProgEntityId"`
	SocoStatus       string `db:"soco_status" json:"socoStatus"`
	SocoModifiedDate string `db:"soco_modified_date" json:"socoModifiedDate"`
}

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
	UserEntityID       int32          `db:"user_entity_id" json:"userEntityId"`
	UserName           sql.NullString `db:"user_name" json:"userName"`
	UserPassword       sql.NullString `db:"user_password" json:"userPassword"`
	UserFirstName      sql.NullString `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString `db:"user_last_name" json:"userLastName"`
	UserBirthDate      sql.NullTime   `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion sql.NullInt32  `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   sql.NullTime   `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          sql.NullString `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    sql.NullInt32  `db:"user_current_role" json:"userCurrentRole"`
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

type UsersUsersMedia struct {
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

type SignUpUser struct {
	User 	UsersUser
	Email   UsersUsersEmail
	Phone 	UsersUsersPhone
}
