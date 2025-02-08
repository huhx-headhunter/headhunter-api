package req

import "github.com/huhx/common-go/times"

type FlowAddRequest struct {
	CandidateId int64                `json:"candidateId"`
	CustomerId  int64                `json:"customerId"`
	JobId       int64                `json:"jobId"`
	Stage       string               `json:"stage"`
	StageStatus string               `json:"stageStatus"`
	StateStart  times.LocalDateTime  `json:"stateStart"`
	StateEnd    *times.LocalDateTime `json:"stateEnd"`
	Remark      string               `json:"remark"`
}
