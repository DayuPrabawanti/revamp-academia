package models

import "time"

type CreateBatchDto struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	BatchEntityID     int32     `db:"batch_entity_id" json:"batchEntityId"`
	BatchName         string    `db:"batch_name" json:"batchName"`
	BatchStartDate    time.Time `db:"batch_start_date" json:"batchStartDate"`
	BatchEndDate      time.Time `db:"batch_end_date" json:"batchEndDate"`
	BatchStatus       string    `db:"batch_status" json:"batchStatus"`
	BatchModifiedDate time.Time `db:"batch_modified_date" json:"batchModifiedDate"`
}

type CreateInstructorProgramsDto struct {
	BatchID           int32     `db:"batch_id" json:"batchId"`
	InproEntityID     int32     `db:"inpro_entity_id" json:"inproEntityId"`
	InproEmpEntityID  int32     `db:"inpro_emp_entity_id" json:"inproEmpEntityId"`
	InproModifiedDate time.Time `db:"inpro_modified_date" json:"inproModifiedDate"`
}

type CreateBatchTraineeDto struct {
	BatrModifiedDate    time.Time `db:"batr_modified_date" json:"batrModifiedDate"`
	BatrTraineeEntityID int32     `db:"batr_trainee_entity_id" json:"batrTraineeEntityId"`
	BatrBatchID         int32     `db:"batr_batch_id" json:"batrBatchId"`
}

type CreateBatchInstructorTraineeDto struct {
	CreateBatchDto
	Instructor []CreateInstructorProgramsDto
	Trainee    []CreateBatchTraineeDto
}
