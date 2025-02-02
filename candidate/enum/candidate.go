package enum

// JobStatus job status
type JobStatus struct {
	Value       int    `json:"value"`
	Description string `json:"description"`
}

var (
	NONE   = JobStatus{Value: 0, Description: "在职-暂不考虑"}
	ONJOB1 = JobStatus{Value: 1, Description: "在职-考虑机会"}
	ONJOB2 = JobStatus{Value: 2, Description: "在职-月内到岗"}
	OUTJOB = JobStatus{Value: 3, Description: "离职-随时到岗"}
)

var jobStatusByValue = map[int]JobStatus{
	NONE.Value:   NONE,
	ONJOB1.Value: ONJOB1,
	ONJOB2.Value: ONJOB2,
	OUTJOB.Value: OUTJOB,
}

func JobStatusFromValue(code int) JobStatus {
	status, _ := jobStatusByValue[code]
	return status
}
