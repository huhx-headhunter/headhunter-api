package resp

import "github.com/huhx/common-go/times"

type FlowItemResp struct {
	Id            int64  `json:"id"`
	CandidateId   int64  `json:"candidateId"`
	CandidateName string `json:"candidateName"`
	JobId         int64  `json:"jobId"`
	JobName       string `json:"jobName"`
	Stage         string `json:"stage"`
	StageStatus   string `json:"stageStatus"`
}

type FlowDetailResp struct {
	Id            int64                `json:"id"`
	CandidateId   int64                `json:"candidateId"`
	CandidateName string               `json:"candidateName"`
	CustomerId    int64                `json:"customerId"`
	CustomerName  int64                `json:"customerName"`
	JobId         int64                `json:"jobId"`
	JobName       string               `json:"jobName"`
	Stage         string               `json:"stage"`
	StageStatus   string               `json:"stageStatus"`
	StateStart    times.LocalDateTime  `json:"stateStart"`
	StateEnd      *times.LocalDateTime `json:"stateEnd"`
	Remark        string               `json:"remark"`
}
