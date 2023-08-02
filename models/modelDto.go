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

type ListGroup []struct {
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
	Section          []struct {
		SectID           int32          `db:"sect_id" json:"sectId"`
		SectProgEntityID int32          `db:"sect_prog_entity_id" json:"sectProgEntityId"`
		SectTitle        string         `db:"sect_title" json:"sectTitle"`
		SectDescription  sql.NullString `db:"sect_description" json:"sectDescription"`
		SectTotalSection int32          `db:"sect_total_section" json:"sectTotalSection"`
		SectTotalLecture int32          `db:"sect_total_lecture" json:"sectTotalLecture"`
		SectTotalMinute  int32          `db:"sect_total_minute" json:"sectTotalMinute"`
		SectModifiedDate time.Time      `db:"sect_modified_date" json:"sectModifiedDate"`
	}
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

type CreateGroupDto struct {
	CreateProgEntityDto
	Group []CreateSectionDto
}
