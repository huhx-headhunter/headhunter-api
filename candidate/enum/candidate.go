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

var (
	JobStatusNone       = JobStatus{Code: 0, Description: "在职-暂不考虑"}
	JobStatusNoneOnJob1 = JobStatus{Code: 1, Description: "在职-考虑机会"}
	JobStatusNoneOnJob2 = JobStatus{Code: 2, Description: "在职-月内到岗"}
	JobStatusOutJob     = JobStatus{Code: 3, Description: "离职-随时到岗"}
)

var jobStatusByValue = map[int]JobStatus{
	JobStatusNone.Code:       JobStatusNone,
	JobStatusNoneOnJob1.Code: JobStatusNoneOnJob1,
	JobStatusNoneOnJob2.Code: JobStatusNoneOnJob2,
	JobStatusOutJob.Code:     JobStatusOutJob,
}

func JobStatusFromValue(code int) JobStatus {
	return jobStatusByValue[code]
}

func (r *JobStatus) Scan(value interface{}) error {
	code, ok := value.(int64)
	if !ok {
		return errors.New("invalid type for JobStatus")
	}
	*r = JobStatusFromValue(int(code))
	return nil
}

func (r JobStatus) Value() (driver.Value, error) {
	return r.Code, nil
}

func (r *JobStatus) UnmarshalJSON(data []byte) error {
	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}
	resourceType := JobStatusFromValue(code)
	*r = resourceType
	return nil
}
