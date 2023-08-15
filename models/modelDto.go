package models

import (
	"database/sql"
	"time"
)

type CreateProgEntityDto struct {
	ProgEntityID     int32     `db:"prog_entity_id" json:"progEntityId"`
	ProgTitle        string    `db:"prog_title" json:"progTitle"`
	ProgHeadline     string    `db:"prog_headline" json:"progHeadline"`
	ProgType         string    `db:"prog_type" json:"progType"`
	ProgLearningType string    `db:"prog_learning_type" json:"progLearningType"`
	ProgRating       string    `db:"prog_rating" json:"progRating"`
	ProgTotalTraniee int32     `db:"prog_total_trainee" json:"progTotalTrainee"`
	ProgModifiedDate time.Time `db:"prog_modified_date" json:"progModifiedDate"`
	ProgImage        string    `db:"prog_image" json:"progImage"`
	ProgBestSeller   string    `db:"prog_best_seller" json:"progBestSeller"`
	ProgPrice        int32     `db:"prog_price" json:"progPrice"`
	ProgLanguage     string    `db:"prog_language" json:"progLanguage"`
	ProgDuration     int32     `db:"prog_duration" json:"progDuration"`
	ProgDurationType string    `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill     string    `db:"prog_tag_skill" json:"progTagSkill"`
	ProgCityID       int32     `db:"prog_city_id" json:"progCityId"`
	ProgCateID       int32     `db:"prog_cate_id" json:"progCateId"`
	ProgCreatedBy    int32     `db:"prog_created_by" json:"progCreatedBy"`
	ProgStatus       string    `db:"prog_status" json:"progStatus"`
}

type CreateProgramEntityDescDto struct {
	PredProgEntityID int32          `db:"pred_prog_entity_id" json:"predProgEntityId"`
	PredItemLearning sql.NullString `db:"pred_item_learning" json:"predItemLearning"`
	PredDescription  sql.NullString `db:"pred_description" json:"predDescription"`
	PredTargetLevel  sql.NullString `db:"pred_target_level" json:"predTargetLevel"`
}

type CreateSectionDto struct {
	SectID           int32          `db:"sect_id" json:"sectId"`
	SectProgEntityID int32          `db:"sect_prog_entity_id" json:"sectProgEntityId"`
	SectTitle        string         `db:"sect_title" json:"sectTitle"`
	SectDescription  sql.NullString `db:"sect_description" json:"sectDescription"`
	SectTotalSection int32          `db:"sect_total_section" json:"sectTotalSection"`
	SectTotalLecture int32          `db:"sect_total_lecture" json:"sectTotalLecture"`
	SectTotalMinute  int32          `db:"sect_total_minute" json:"sectTotalMinute"`
	SectModifiedDate time.Time      `db:"sect_modified_date" json:"sectModifiedDate"`
}

type CreateMasterCategoryDto struct {
	CateID           int32         `db:"cate_id" json:"cateId"`
	CateName         string        `db:"cate_name" json:"cateName"`
	CateCateID       sql.NullInt32 `db:"cate_cate_id" json:"cateCateId"`
	CateModifiedDate sql.NullTime  `db:"cate_modified_date" json:"cateModifiedDate"`
}

type CreateSectionDetailDto struct {
	SecdID           int32          `db:"secd_id" json:"secdId"`
	SecdTitle        string         `db:"secd_title" json:"secdTitle"`
	SecdPreview      string         `db:"secd_preview" json:"secdPreview"`
	SecdScore        int32          `db:"secd_score" json:"secdScore"`
	SecdNote         sql.NullString `db:"secd_note" json:"secdNote"`
	SecdMinute       int32          `db:"secd_minute" json:"secdMinute"`
	SecdModifiedDate time.Time      `db:"secd_modified_date" json:"secdModifiedDate"`
	SecdSectID       int32          `db:"secd_sect_id" json:"secdSectId"`
}

type CreateGroup []struct {
	ProgTitle         string `db:"prog_title" json:"progTitle"`
	ProgHeadline      string `db:"prog_headline" json:"progHeadline"`
	ProgType          string `db:"prog_type" json:"progType"`
	ProgLearningType  string `db:"prog_learning_type" json:"progLearningType"`
	ProgImage         string `db:"prog_image" json:"progImage"`
	ProgPrice         int32  `db:"prog_price" json:"progPrice"`
	ProgLanguage      string `db:"prog_language" json:"progLanguage"`
	ProgDuration      int32  `db:"prog_duration" json:"progDuration"`
	ProgDurationType  string `db:"prog_duration_type" json:"progDurationType"`
	ProgTagSkill      string `db:"prog_tag_skill" json:"progTagSkill"`
	ProgramEntityDesc []struct {
		PredDescription sql.NullString `db:"pred_description" json:"predDescription"`
	}
	MasterCategory []struct {
		CateName string `db:"cate_name" json:"cateName"`
	}
	Section []struct {
		SectTitle        string         `db:"sect_title" json:"sectTitle"`
		SectDescription  sql.NullString `db:"sect_description" json:"sectDescription"`
		SectTotalSection int32          `db:"sect_total_section" json:"sectTotalSection"`
		SectTotalLecture int32          `db:"sect_total_lecture" json:"sectTotalLecture"`
		SectTotalMinute  int32          `db:"sect_total_minute" json:"sectTotalMinute"`
		SectModifiedDate time.Time      `db:"sect_modified_date" json:"sectModifiedDate"`
	}
	SubSection []struct {
		SecdScore int32 `db:"secd_score" json:"secdScore"`
	}
}
