package models

import (
	"database/sql"
	"time"
)

type BootcampBatch struct {
	BatchID           int32  `db:"batch_id" json:"batchId"`
	BatchEntityID     int32  `db:"batch_entity_id" json:"batchEntityId"`
	BatchName         string `db:"batch_name" json:"batchName"`
	BatchDescription  string `db:"batch_description" json:"batchDescription"`
	BatchStartDate    string `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate      string `db:"batch_end_date" json:"batchEndDate"`
	BatchReason       string `db:"batch_reason" json:"batchReason"`
	BatchType         string `db:"batch_type" json:"batchType"`
	BatchModifiedDate string `db:"batch_modified_date" json:"batchModifiedDate"`
	BatchStatus       string `db:"batch_status" json:"batchStatus"`
	BatchPicID        int32  `db:"batch_pic_id" json:"batchPicId"`
}

type BootcampBatchTrainee struct {
	BatrID               int32  `db:"batr_id" json:"batrId"`
	BatrStatus           string `db:"batr_status" json:"batrStatus"`
	BatrCertificated     string `db:"batr_certificated" json:"batrCertificated"`
	BatreCertificateLink string `db:"batre_certificate_link" json:"batreCertificateLink"`
	BatrAccessToken      string `db:"batr_access_token" json:"batrAccessToken"`
	BatrAccessGrant      string `db:"batr_access_grant" json:"batrAccessGrant"`
	BatrReview           string `db:"batr_review" json:"batrReview"`
	BatrTotalScore       int32  `db:"batr_total_score" json:"batrTotalScore"`
	BatrModifiedDate     string `db:"batr_modified_date" json:"batrModifiedDate"`
	BatrTraineeEntityID  int32  `db:"batr_trainee_entity_id" json:"batrTraineeEntityId"`
	BatrBatchID          int32  `db:"batr_batch_id" json:"batrBatchId"`
}

type BootcampBatchTraineeEvaluation struct {
	BtevID              int32  `db:"btev_id" json:"btevId"`
	BtevType            string `db:"btev_type" json:"btevType"`
	BtevHeader          string `db:"btev_header" json:"btevHeader"`
	BtevSection         string `db:"btev_section" json:"btevSection"`
	BtevSkill           string `db:"btev_skill" json:"btevSkill"`
	BtevWeek            int32  `db:"btev_week" json:"btevWeek"`
	BtevSkor            int32  `db:"btev_skor" json:"btevSkor"`
	BtevNote            string `db:"btev_note" json:"btevNote"`
	BtevModifiedDate    string `db:"btev_modified_date" json:"btevModifiedDate"`
	BtevBatchID         int32  `db:"btev_batch_id" json:"btevBatchId"`
	BtevTraineeEntityID int32  `db:"btev_trainee_entity_id" json:"btevTraineeEntityId"`
}

type BootcampInstructorProgram struct {
	BatchID           int32  `db:"batch_id" json:"batchId"`
	InproEntityID     int32  `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32  `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate string `db:"inpro_modified_date" json:"inproModifiedDate"`
}

type BootcampProgramApply struct {
	PrapUserEntityID int32  `db:"prap_user_entity_id" json:"prapUserEntityId"`
	PrapProgEntityID int32  `db:"prap_prog_entity_id" json:"prapProgEntityId"`
	PrapTestScore    int32  `db:"prap_test_score" json:"prapTestScore"`
	PrapGpa          int32  `db:"prap_gpa" json:"prapGpa"`
	PrapIqTest       int32  `db:"prap_iq_test" json:"prapIqTest"`
	PrapReview       string `db:"prap_review" json:"prapReview"`
	PrapModifiedDate string `db:"prap_modified_date" json:"prapModifiedDate"`
	PrapStatus       string `db:"prap_status" json:"prapStatus"`
}

type BootcampProgramApplyProgress struct {
	ParogID           int32  `db:"parog_id" json:"parogId"`
	ParogUserEntityID int32  `db:"parog_user_entity_id" json:"parogUserEntityId"`
	ParogProgEntityID int32  `db:"parog_prog_entity_id" json:"parogProgEntityId"`
	ParogActionDate   string `db:"parog_action_date" json:"parogActionDate"`
	ParogModifiedDate string `db:"parog_modified_date" json:"parogModifiedDate"`
	ParogComment      string `db:"parog_comment" json:"parogComment"`
	ParogProgressName string `db:"parog_progress_name" json:"parogProgressName"`
	ParogEmpEntityID  int32  `db:"parog_emp_entity_id" json:"parogEmpEntityId"`
	ParogStatus       string `db:"parog_status" json:"parogStatus"`
}

type CurriculumProgramEntity struct {
	ProgEntityID     int32  `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string `db:"prog_title" json:"progTitle"`
	ProgHeadline     string `db:"prog_headline" json:"progHeadline"`
	ProgType         string `db:"prog_type" json:"progType"`
	ProgLearningType string `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       int32  `db:"prog_rating" json:"progRating"`
	ProgTotalTrainee int32  `db:"prog_total_trainee" json:"progTotalTrainee"`
	ProgModifiedDate string `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string `db:"prog_image" json:"progImage"`
	ProgBestSeller   string `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        int32  `db:"prog_price" json:"progPrice"`
	ProgLanguage     string `db:"prog_language" json:"progLanguage"`
	ProgDuration     int32  `db:"prog_duration" json:"progDuration"`
	ProgDurationType string `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     string `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       int32  `db:"prog_city_id" json:"progCityId"`
	ProgCateID       int32  `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    int32  `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       string `db:"prog_status" json:"progStatus"`
}

type CurriculumProgramEntityDescription struct {
	PredProgEntityID int32  `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning string `db:"pred_item_learning" json:"predItemLearning"`
	PredItemInclude  string `db:"pred_item_include" json:"predItemInclude"`
	PredRequirement  string `db:"pred_requirement" json:"predRequirement"`
	PredDescription  string `db:"pred_description" json:"predDescription"`
	PredTargetLevel  string `db:"pred_target_level" json:"predTargetLevel"`
}

type CurriculumProgramReview struct {
	ProwUserEntityID int32  `db:"prow_user_entity_id" json:"prowUserEntityId"`
	ProwProgEntityID int32  `db:"prow_prog_entity_id" json:"prowProgEntityId"`
	ProwReview       string `db:"prow_review" json:"prowReview"`
	ProwRating       int32  `db:"prow_rating" json:"prowRating"`
	ProwModifiedDate string `db:"prow_modified_date" json:"prowModifiedDate"`
}

type CurriculumSection struct {
	SectID           int32  `db:"sect_id" json:"sectId"`
	SectProgEntityID int32  `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SectTitle        string `db:"sect_title" json:"sectTitle"`
	SectDescription  string `db:"sect_description" json:"sectDescription"`
	SectTotalSection int32  `db:"sect_total_section" json:"sectTotalSection"`
	SectTotalLecture int32  `db:"sect_total_lecture" json:"sectTotalLecture"`
	SectTotalMinute  int32  `db:"sect_total_minute" json:"sectTotalMinute"`
	SectModifiedDate string `db:"sect_modified_date" json:"sectModifiedDate"`
}

type CurriculumSectionDetail struct {
	SecdID           int32  `db:"secd_id" json:"secdId"`
	SecdTitle        string `db:"secd_title" json:"secdTitle"`
	SecdPreview      string `db:"secd_preview" json:"secdPreview"`
	SecdScore        int32  `db:"secd_score" json:"secdScore"`
	SecdNote         string `db:"secd_note" json:"secdNote"`
	SecdMinute       int32  `db:"secd_minute" json:"secdMinute"`
	SecdModifiedDate string `db:"secd_modified_date" json:"secdModifiedDate"`
	SecdSectID       int32  `db:"secd_sect_id" json:"secdSectId"`
}

type CurriculumSectionDetailMaterial struct {
	SedmID           int32  `db:"sedm_id" json:"sedmId"`
	SedmFilename     string `db:"sedm_filename" json:"sedmFilename"`
	SedmFilesize     int32  `db:"sedm_filesize" json:"sedmFilesize"`
	SedmFiletype     string `db:"sedm_filetype" json:"sedmFiletype"`
	SedmFilelink     string `db:"sedm_filelink" json:"sedmFilelink"`
	SedmModifiedDate string `db:"sedm_modified_date" json:"sedmModifiedDate"`
	SedmSecdID       int32  `db:"sedm_secd_id" json:"sedmSecdId"`
}

type HrDepartment struct {
	DeptID           int32  `db:"dept_id" json:"deptId"`
	DeptName         string `db:"dept_name" json:"deptName"`
	DeptModifiedDate string `db:"dept_modified_date" json:"deptModifiedDate"`
}

type HrEmployee struct {
	EmpEntityID       int32         `db:"emp_entity_id" json:"empEntityId"`
	EmpEmpNumber      string        `db:"emp_emp_number" json:"empEmpNumber"`
	EmpNationalID     string        `db:"emp_national_id" json:"empNationalId"`
	EmpBirthDate      string        `db:"emp_birth_date" json:"empBirthDate"`
	EmpMaritalStatus  string        `db:"emp_marital_status" json:"empMaritalStatus"`
	EmpGender         string        `db:"emp_gender" json:"empGender"`
	EmpHireDate       string        `db:"emp_hire_date" json:"empHireDate"`
	EmpSalariedFlag   string        `db:"emp_salaried_flag" json:"empSalariedFlag"`
	EmpVacationHours  sql.NullInt16 `db:"emp_vacation_hours" json:"empVacationHours"`
	EmpSickleaveHours sql.NullInt16 `db:"emp_sickleave_hours" json:"empSickleaveHours"`
	EmpCurrentFlag    sql.NullInt16 `db:"emp_current_flag" json:"empCurrentFlag"`
	EmpModifiedDate   string        `db:"emp_modified_date" json:"empModifiedDate"`
	EmpType           string        `db:"emp_type" json:"empType"`
	EmpJoroID         int32         `db:"emp_joro_id" json:"empJoroId"`
	EmpEmpEntityID    int32         `db:"emp_emp_entity_id" json:"empEmpEntityId"`
}

type HrEmployeeClientContract struct {
	EccoID             int32  `db:"ecco_id" json:"eccoId"`
	EccoEntityID       int32  `db:"ecco_entity_id" json:"eccoEntityId"`
	EccoContractNo     string `db:"ecco_contract_no" json:"eccoContractNo"`
	EccoContractDate   string `db:"ecco_contract_date" json:"eccoContractDate"`
	EccoStartDate      string `db:"ecco_start_date" json:"eccoStartDate"`
	EccoEndDate        string `db:"ecco_end_date" json:"eccoEndDate"`
	EccoNotes          string `db:"ecco_notes" json:"eccoNotes"`
	EccoModifiedDate   string `db:"ecco_modified_date" json:"eccoModifiedDate"`
	EccoMediaLink      string `db:"ecco_media_link" json:"eccoMediaLink"`
	EccoJotyID         int32  `db:"ecco_joty_id" json:"eccoJotyId"`
	EccoAccountManager int32  `db:"ecco_account_manager" json:"eccoAccountManager"`
	EccoClitID         int32  `db:"ecco_clit_id" json:"eccoClitId"`
	EccoStatus         string `db:"ecco_status" json:"eccoStatus"`
}

type HrEmployeeDepartmentHistory struct {
	EdhiID           int32  `db:"edhi_id" json:"edhiId"`
	EdhiEntityID     int32  `db:"edhi_entity_id" json:"edhiEntityId"`
	EdhiStartDate    string `db:"edhi_start_date" json:"edhiStartDate"`
	EdhiEndDate      string `db:"edhi_end_date" json:"edhiEndDate"`
	EdhiModifiedDate string `db:"edhi_modified_date" json:"edhiModifiedDate"`
	EdhiDeptID       int32  `db:"edhi_dept_id" json:"edhiDeptId"`
}

type HrEmployeePayHistory struct {
	EphiEntityID       int32         `db:"ephi_entity_id" json:"ephiEntityId"`
	EphiRateChangeDate time.Time     `db:"ephi_rate_change_date" json:"ephiRateChangeDate"`
	EphiRateSalary     int32         `db:"ephi_rate_salary" json:"ephiRateSalary"`
	EphiPayFrequence   sql.NullInt16 `db:"ephi_pay_frequence" json:"ephiPayFrequence"`
	EphiModifiedDate   string        `db:"ephi_modified_date" json:"ephiModifiedDate"`
}

type JobhireClient struct {
	ClitID           int32  `db:"clit_id" json:"clitId"`
	ClitName         string `db:"clit_name" json:"clitName"`
	ClitAbout        string `db:"clit_about" json:"clitAbout"`
	ClitModifiedDate string `db:"clit_modified_date" json:"clitModifiedDate"`
	ClitAddrID       int32  `db:"clit_addr_id" json:"clitAddrId"`
	ClitEmraID       int32  `db:"clit_emra_id" json:"clitEmraId"`
}

type JobhireEmployeeRange struct {
	EmraID           int32  `db:"emra_id" json:"emraId"`
	EmraRangeMin     int32  `db:"emra_range_min" json:"emraRangeMin"`
	EmraRangeMax     int32  `db:"emra_range_max" json:"emraRangeMax"`
	EmraModifiedDate string `db:"emra_modified_date" json:"emraModifiedDate"`
}

type JobhireJobCategory struct {
	JocaID           int32  `db:"joca_id" json:"jocaId"`
	JocaName         string `db:"joca_name" json:"jocaName"`
	JocaModifiedDate string `db:"joca_modified_date" json:"jocaModifiedDate"`
}

type JobhireJobPhoto struct {
	JophoID           int32  `db:"jopho_id" json:"jophoId"`
	JophoFilename     string `db:"jopho_filename" json:"jophoFilename"`
	JophoFilesize     int32  `db:"jopho_filesize" json:"jophoFilesize"`
	JophoFiletype     string `db:"jopho_filetype" json:"jophoFiletype"`
	JophoModifiedDate string `db:"jopho_modified_date" json:"jophoModifiedDate"`
	JophoEntityID     int32  `db:"jopho_entity_id" json:"jophoEntityId"`
}

type JobhireJobPost struct {
	JopoEntityID       int32  `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoNumber         string `db:"jopo_number" json:"jopoNumber"`
	JopoTitle          string `db:"jopo_title" json:"jopoTitle"`
	JopoStartDate      string `db:"jopo_start_date" json:"jopoStartDate"`
	JopoEndDate        string `db:"jopo_end_date" json:"jopoEndDate"`
	JopoMinSalary      int32  `db:"jopo_min_salary" json:"jopoMinSalary"`
	JopoMaxSalary      int32  `db:"jopo_max_salary" json:"jopoMaxSalary"`
	JopoMinExperience  int32  `db:"jopo_min_experience" json:"jopoMinExperience"`
	JopoMaxExperience  int32  `db:"jopo_max_experience" json:"jopoMaxExperience"`
	JopoPrimarySkill   string `db:"jopo_primary_skill" json:"jopoPrimarySkill"`
	JopoSecondarySkill string `db:"jopo_secondary_skill" json:"jopoSecondarySkill"`
	JopoPublishDate    string `db:"jopo_publish_date" json:"jopoPublishDate"`
	JopoModifiedDate   string `db:"jopo_modified_date" json:"jopoModifiedDate"`
	JopoEmpEntityID    int32  `db:"jopo_emp_entity_id" json:"jopoEmpEntityId"`
	JopoClitID         int32  `db:"jopo_clit_id" json:"jopoClitId"`
	JopoJoroID         int32  `db:"jopo_joro_id" json:"jopoJoroId"`
	JopoJotyID         int32  `db:"jopo_joty_id" json:"jopoJotyId"`
	JopoJocaID         int32  `db:"jopo_joca_id" json:"jopoJocaId"`
	JopoAddrID         int32  `db:"jopo_addr_id" json:"jopoAddrId"`
	JopoWorkCode       string `db:"jopo_work_code" json:"jopoWorkCode"`
	JopoEduCode        string `db:"jopo_edu_code" json:"jopoEduCode"`
	JopoInduCode       string `db:"jopo_indu_code" json:"jopoInduCode"`
	JopoStatus         string `db:"jopo_status" json:"jopoStatus"`
}

type JobhireJobPostDesc struct {
	JopoEntityID       int32  `db:"jopo_entity_id" json:"jopoEntityId"`
	JopoDescription    string `db:"jopo_description" json:"jopoDescription"`
	JopoResponsibility string `db:"jopo_responsibility" json:"jopoResponsibility"`
	JopoTarget         string `db:"jopo_target" json:"jopoTarget"`
	JopoBenefit        string `db:"jopo_benefit" json:"jopoBenefit"`
}

type JobhireTalentApply struct {
	TaapUserEntityID int32  `db:"taap_user_entity_id" json:"taapUserEntityId"`
	TaapEntityID     int32  `db:"taap_entity_id" json:"taapEntityId"`
	TaapIntro        string `db:"taap_intro" json:"taapIntro"`
	TaapScoring      int32  `db:"taap_scoring" json:"taapScoring"`
	TaapModifiedDate string `db:"taap_modified_date" json:"taapModifiedDate"`
	TaapStatus       string `db:"taap_status" json:"taapStatus"`
}

type JobhireTalentApplyProgress struct {
	TaprID           int32  `db:"tapr_id" json:"taprId"`
	TaapUserEntityID int32  `db:"taap_user_entity_id" json:"taapUserEntityId"`
	TaapEntityID     int32  `db:"taap_entity_id" json:"taapEntityId"`
	TaprModifiedDate string `db:"tapr_modified_date" json:"taprModifiedDate"`
	TaprStatus       string `db:"tapr_status" json:"taprStatus"`
	TaprComment      string `db:"tapr_comment" json:"taprComment"`
	TaprProgressName string `db:"tapr_progress_name" json:"taprProgressName"`
}

type MasterAddress struct {
	AddrID              int32  `db:"addr_id" json:"addrId"`
	AddrLine1           string `db:"addr_line1" json:"addrLine1"`
	AddrLine2           string `db:"addr_line2" json:"addrLine2"`
	AddrPostalCode      string `db:"addr_postal_code" json:"addrPostalCode"`
	AddrSpatialLocation string `db:"addr_spatial_location" json:"addrSpatialLocation"`
	AddrModifiedDate    string `db:"addr_modified_date" json:"addrModifiedDate"`
	AddrCityID          int32  `db:"addr_city_id" json:"addrCityId"`
}

type MasterAddressType struct {
	AdtyID           int32  `db:"adty_id" json:"adtyId"`
	AdtyName         string `db:"adty_name" json:"adtyName"`
	AdtyModifiedDate string `db:"adty_modified_date" json:"adtyModifiedDate"`
}

type MasterCategory struct {
	CateID           int32  `db:"cate_id" json:"cateId"`
	CateName         string `db:"cate_name" json:"cateName"`
	CateCateID       int32  `db:"cate_cate_id" json:"cateCateId"`
	CateModifiedDate string `db:"cate_modified_date" json:"cateModifiedDate"`
}

type MasterCity struct {
	CityID           int32  `db:"city_id" json:"cityId"`
	CityName         string `db:"city_name" json:"cityName"`
	CityModifiedDate string `db:"city_modified_date" json:"cityModifiedDate"`
	CityProvID       int32  `db:"city_prov_id" json:"cityProvId"`
}

type MasterCountry struct {
	CountryCode         string `db:"country_code" json:"countryCode"`
	CountryName         string `db:"country_name" json:"countryName"`
	CountryModifiedDate string `db:"country_modified_date" json:"countryModifiedDate"`
}

type MasterEducation struct {
	EduCode string `db:"edu_code" json:"eduCode"`
	EduName string `db:"edu_name" json:"eduName"`
}

type MasterIndustry struct {
	InduCode string `db:"indu_code" json:"induCode"`
	InduName string `db:"indu_name" json:"induName"`
}

type MasterJobRole struct {
	JoroID           int32  `db:"joro_id" json:"joroId"`
	JoroName         string `db:"joro_name" json:"joroName"`
	JoroModifiedDate string `db:"joro_modified_date" json:"joroModifiedDate"`
}

type MasterJobType struct {
	JotyID   int32  `db:"joty_id" json:"jotyId"`
	JotyName string `db:"joty_name" json:"jotyName"`
}

type MasterModule struct {
	ModuleName string `db:"module_name" json:"moduleName"`
}

type MasterProvince struct {
	ProvID           int32  `db:"prov_id" json:"provId"`
	ProvCode         string `db:"prov_code" json:"provCode"`
	ProvName         string `db:"prov_name" json:"provName"`
	ProvModifiedDate string `db:"prov_modified_date" json:"provModifiedDate"`
	ProvCountryCode  string `db:"prov_country_code" json:"provCountryCode"`
}

type MasterRouteAction struct {
	RoacID         int32  `db:"roac_id" json:"roacId"`
	RoacName       string `db:"roac_name" json:"roacName"`
	RoacOrderby    int32  `db:"roac_orderby" json:"roacOrderby"`
	RoacDisplay    int32  `db:"roac_display" json:"roacDisplay"`
	RoacModuleName string `db:"roac_module_name" json:"roacModuleName"`
}

type MasterSkillTemplate struct {
	SkteID           int32  `db:"skte_id" json:"skteId"`
	SkteSkill        string `db:"skte_skill" json:"skteSkill"`
	SkteDescription  string `db:"skte_description" json:"skteDescription"`
	SkteWeek         int32  `db:"skte_week" json:"skteWeek"`
	SkteOrderby      int32  `db:"skte_orderby" json:"skteOrderby"`
	SkteModifiedDate string `db:"skte_modified_date" json:"skteModifiedDate"`
	SktyName         string `db:"skty_name" json:"sktyName"`
	SkteSkteID       int32  `db:"skte_skte_id" json:"skteSkteId"`
}

type MasterSkillType struct {
	SktyName string `db:"skty_name" json:"sktyName"`
}

type MasterStatus struct {
	Status             string `db:"status" json:"status"`
	StatusModifiedDate string `db:"status_modified_date" json:"statusModifiedDate"`
	StatusModule       string `db:"status_module" json:"statusModule"`
}

type MasterWorkingType struct {
	WotyCode string `db:"woty_code" json:"wotyCode"`
	WotyName string `db:"woty_name" json:"wotyName"`
}

type PaymentBank struct {
	BankEntityID     int32  `db:"bank_entity_id" json:"bankEntityId"`
	BankCode         string `db:"bank_code" json:"bankCode"`
	BankName         string `db:"bank_name" json:"bankName"`
	BankModifiedDate string `db:"bank_modified_date" json:"bankModifiedDate"`
}

type PaymentFintech struct {
	FintEntityID     int32  `db:"fint_entity_id" json:"fintEntityId"`
	FintCode         string `db:"fint_code" json:"fintCode"`
	FintName         string `db:"fint_name" json:"fintName"`
	FintModifiedDate string `db:"fint_modified_date" json:"fintModifiedDate"`
}

type PaymentTransactionPayment struct {
	TrpaID           int32  `db:"trpa_id" json:"trpaId"`
	TrpaCodeNumber   string `db:"trpa_code_number" json:"trpaCodeNumber"`
	TrpaOrderNumber  string `db:"trpa_order_number" json:"trpaOrderNumber"`
	TrpaDebit        string `db:"trpa_debit" json:"trpaDebit"`
	TrpaCredit       string `db:"trpa_credit" json:"trpaCredit"`
	TrpaType         string `db:"trpa_type" json:"trpaType"`
	TrpaNote         string `db:"trpa_note" json:"trpaNote"`
	TrpaModifiedDate string `db:"trpa_modified_date" json:"trpaModifiedDate"`
	TrpaSourceID     string `db:"trpa_source_id" json:"trpaSourceId"`
	TrpaTargetID     string `db:"trpa_target_id" json:"trpaTargetId"`
	TrpaUserEntityID int32  `db:"trpa_user_entity_id" json:"trpaUserEntityId"`
}

type PaymentUsersAccount struct {
	UsacBankEntityID  int32  `db:"usac_bank_entity_id" json:"usacBankEntityId"`
	UsacUserEntityID  int32  `db:"usac_user_entity_id" json:"usacUserEntityId"`
	UsacAccountNumber string `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         string `db:"usac_saldo" json:"usacSaldo"`
	UsacType          string `db:"usac_type" json:"usacType"`
	UsacStartDate     string `db:"usac_start_date" json:"usacStartDate"`
	UsacEndDate       string `db:"usac_end_date" json:"usacEndDate"`
	UsacModifiedDate  string `db:"usac_modified_date" json:"usacModifiedDate"`
	UsacStatus        string `db:"usac_status" json:"usacStatus"`
}

type SalesCartItem struct {
	CaitID           int32  `db:"cait_id" json:"caitId"`
	CaitQuantity     int32  `db:"cait_quantity" json:"caitQuantity"`
	CaitUnitPrice    string `db:"cait_unit_price" json:"caitUnitPrice"`
	CaitModifiedDate string `db:"cait_modified_date" json:"caitModifiedDate"`
	CaitUserEntityID int32  `db:"cait_user_entity_id" json:"caitUserEntityId"`
	CaitProgEntityID int32  `db:"cait_prog_entity_id" json:"caitProgEntityId"`
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
	PontyCode         string `db:"ponty_code" json:"pontyCode"`
	PontyModifiedDate string `db:"ponty_modified_date" json:"pontyModifiedDate"`
}

type UsersRole struct {
	RoleID           int32  `db:"role_id" json:"roleId"`
	RoleName         string `db:"role_name" json:"roleName"`
	RoleType         string `db:"role_type" json:"roleType"`
	RoleModifiedDate string `db:"role_modified_date" json:"roleModifiedDate"`
}

type UsersUser struct {
	UserEntityID       int32  `db:"user_entity_id" json:"userEntityId"`
	UserName           string `db:"user_name" json:"userName"`
	UserPassword       string `db:"user_password" json:"userPassword"`
	UserFirstName      string `db:"user_first_name" json:"userFirstName"`
	UserLastName       string `db:"user_last_name" json:"userLastName"`
	UserBirthDate      string `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion int32  `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    string `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   string `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          string `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    int32  `db:"user_current_role" json:"userCurrentRole"`
}

type UsersUsersAddress struct {
	EtadAddrID       int32  `db:"etad_addr_id" json:"etadAddrId"`
	EtadModifiedDate string `db:"etad_modified_date" json:"etadModifiedDate"`
	EtadEntityID     int32  `db:"etad_entity_id" json:"etadEntityId"`
	EtadAdtyID       int32  `db:"etad_adty_id" json:"etadAdtyId"`
}

type UsersUsersEducation struct {
	UsduID           int32  `db:"usdu_id" json:"usduId"`
	UsduEntityID     int32  `db:"usdu_entity_id" json:"usduEntityId"`
	UsduSchool       string `db:"usdu_school" json:"usduSchool"`
	UsduDegree       string `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy   string `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduGraduateYear string `db:"usdu_graduate_year" json:"usduGraduateYear"`
	UsduStartDate    string `db:"usdu_start_date" json:"usduStartDate"`
	UsduEndDate      string `db:"usdu_end_date" json:"usduEndDate"`
	UsduGrade        string `db:"usdu_grade" json:"usduGrade"`
	UsduActivities   string `db:"usdu_activities" json:"usduActivities"`
	UsduDescription  string `db:"usdu_description" json:"usduDescription"`
	UsduModifiedDate string `db:"usdu_modified_date" json:"usduModifiedDate"`
}

type UsersUsersEmail struct {
	PmailEntityID     int32  `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32  `db:"pmail_id" json:"pmailId"`
	PmailAddress      string `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate string `db:"pmail_modified_date" json:"pmailModifiedDate"`
}

type UsersUsersExperience struct {
	UsexID              int32  `db:"usex_id" json:"usexId"`
	UsexEntityID        int32  `db:"usex_entity_id" json:"usexEntityId"`
	UsexTitle           string `db:"usex_title" json:"usexTitle"`
	UsexProfileHeadline string `db:"usex_profile_headline" json:"usexProfileHeadline"`
	UsexEmploymentType  string `db:"usex_employment_type" json:"usexEmploymentType"`
	UsexCompanyName     string `db:"usex_company_name" json:"usexCompanyName"`
	UsexIsCurrent       string `db:"usex_is_current" json:"usexIsCurrent"`
	UsexStartDate       string `db:"usex_start_date" json:"usexStartDate"`
	UsexEndDate         string `db:"usex_end_date" json:"usexEndDate"`
	UsexIndustry        string `db:"usex_industry" json:"usexIndustry"`
	UsexDescription     string `db:"usex_description" json:"usexDescription"`
	UsexExperienceType  string `db:"usex_experience_type" json:"usexExperienceType"`
	UsexCityID          int32  `db:"usex_city_id" json:"usexCityId"`
}

type UsersUsersExperiencesSkill struct {
	UeskUsexID int32 `db:"uesk_usex_id" json:"ueskUsexId"`
	UeskUskiID int32 `db:"uesk_uski_id" json:"ueskUskiId"`
}

type UsersUsersLicense struct {
	UsliID           int32  `db:"usli_id" json:"usliId"`
	UsliLicenseCode  string `db:"usli_license_code" json:"usliLicenseCode"`
	UsliModifiedDate string `db:"usli_modified_date" json:"usliModifiedDate"`
	UsliStatus       string `db:"usli_status" json:"usliStatus"`
	UsliEntityID     int32  `db:"usli_entity_id" json:"usliEntityId"`
}

type UsersUsersMedium struct {
	UsmeID           int32  `db:"usme_id" json:"usmeId"`
	UsmeEntityID     int32  `db:"usme_entity_id" json:"usmeEntityId"`
	UsmeFileLink     string `db:"usme_file_link" json:"usmeFileLink"`
	UsmeFilename     string `db:"usme_filename" json:"usmeFilename"`
	UsmeFilesize     int32  `db:"usme_filesize" json:"usmeFilesize"`
	UsmeFiletype     string `db:"usme_filetype" json:"usmeFiletype"`
	UsmeNote         string `db:"usme_note" json:"usmeNote"`
	UsmeModifiedDate string `db:"usme_modified_date" json:"usmeModifiedDate"`
}

type UsersUsersPhone struct {
	UspoEntityID     int32  `db:"uspo_entity_id" json:"uspoEntityId"`
	UspoNumber       string `db:"uspo_number" json:"uspoNumber"`
	UspoModifiedDate string `db:"uspo_modified_date" json:"uspoModifiedDate"`
	UspoPontyCode    string `db:"uspo_ponty_code" json:"uspoPontyCode"`
}

type UsersUsersRole struct {
	UsroEntityID     int32  `db:"usro_entity_id" json:"usroEntityId"`
	UsroRoleID       int32  `db:"usro_role_id" json:"usroRoleId"`
	UsroModifiedDate string `db:"usro_modified_date" json:"usroModifiedDate"`
}

type UsersUsersSkill struct {
	UskiID           int32  `db:"uski_id" json:"uskiId"`
	UskiEntityID     int32  `db:"uski_entity_id" json:"uskiEntityId"`
	UskiModifiedDate string `db:"uski_modified_date" json:"uskiModifiedDate"`
	UskiSktyName     string `db:"uski_skty_name" json:"uskiSktyName"`
}
