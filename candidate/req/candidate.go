package req

import (
	"github.com/huhx/common-go/times"
)

type CandidateAddRequest struct {
	Avatar       string                 `json:"avatar"`
	Name         string                 `json:"name"`
	JobStatus    int                    `json:"jobStatus"`
	FirstJobDate times.LocalDate        `json:"firstJobDate"`
	Birthday     times.LocalDate        `json:"birthday"`
	Address      string                 `json:"address"`
	Email        string                 `json:"email"`
	Phone        string                 `json:"phone"`
	Wechat       string                 `json:"wechat"`
	Tags         []*TagAddRequest       `json:"tags"`
	Educations   []*EducationAddRequest `json:"educations"`
	Resumes      []*ResumeAddRequest    `json:"resumes"`
	Works        []*WorkAddRequest      `json:"works"`
	Skills       []*SkillAddRequest     `json:"skills"`
}

type TagAddRequest struct {
	Id          *int64 `json:"id" example:"1800783867820785152"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type EducationAddRequest struct {
	School    string           `json:"school"`
	Major     string           `json:"major"`
	Level     int              `json:"level"` // 0: 学士 1: 硕士 2: 博士
	StartTime times.LocalDate  `json:"startTime"`
	EndTime   *times.LocalDate `json:"endTime"`
}

type ResumeAddRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type WorkAddRequest struct {
	JobTitle  string            `json:"jobTitle"`
	Company   CompanyAddRequest `json:"company"`
	City      string            `json:"city"`
	StartTime times.LocalDate   `json:"startTime"`
	EndTime   times.LocalDate   `json:"endTime"`
	Content   string            `json:"content"`
	Summary   string            `json:"summary"`
}

type SkillAddRequest struct {
	Content string `json:"content"`
}

type CompanyAddRequest struct {
	Id          *int64 `json:"id" example:"1800783867820785152"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
