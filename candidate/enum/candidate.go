package enum

import (
	"database/sql/driver"
	"errors"
	"github.com/goccy/go-json"
)

type JobStatus struct {
	Code        int    `json:"value"`
	Description string `json:"description"`
}

// Scan get from DB
func (r *JobStatus) Scan(value interface{}) error {
	intVal, ok := value.(int)
	if !ok {
		return errors.New("invalid type for ResourceType")
	}
	*r = JobStatusFromValue(intVal)
	return nil
}

// Value save in DB
func (r *JobStatus) Value() (driver.Value, error) {
	return r.Value, nil
}

// UnmarshalJSON from request
func (r *JobStatus) UnmarshalJSON(data []byte) error {
	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}
	resourceType := JobStatusFromValue(code)
	*r = resourceType
	return nil
}

var (
	NONE   = JobStatus{Code: 0, Description: "在职-暂不考虑"}
	ONJOB1 = JobStatus{Code: 1, Description: "在职-考虑机会"}
	ONJOB2 = JobStatus{Code: 2, Description: "在职-月内到岗"}
	OUTJOB = JobStatus{Code: 3, Description: "离职-随时到岗"}
)

var jobStatusByValue = map[int]JobStatus{
	NONE.Code:   NONE,
	ONJOB1.Code: ONJOB1,
	ONJOB2.Code: ONJOB2,
	OUTJOB.Code: OUTJOB,
}

func JobStatusFromValue(code int) JobStatus {
	status, _ := jobStatusByValue[code]
	return status
}
