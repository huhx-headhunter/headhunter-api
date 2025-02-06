package resp

import (
	"github.com/huhx-headhunter/headhunter-api/candidate/enum"
	"github.com/huhx/common-go/times"
)

type CandidateDetailResponse struct {
	Id           int64             `json:"id"`
	Name         string            `json:"name"`
	Avatar       string            `json:"avatar"`
	Birthday     times.LocalDate   `json:"birthday"`
	JobStatus    enum.JobStatus    `json:"jobStatus"`
	FirstJobDate times.LocalDate   `json:"firstJobDate"`
	Address      string            `json:"address"`
	Email        string            `json:"email"`
	Phone        string            `json:"phone"`
	Wechat       string            `json:"wechat"`
	Tags         []TagDetail       `json:"tags"`
	Educations   []EducationDetail `json:"educations"`
	Resumes      []ResumeDetail    `json:"resumes"`
	Works        []WorkDetail      `gorm:"foreignKey:CandidateId" json:"works"`
	Skills       []SkillDetail     `json:"skills"`
}

type TagDetail struct {
	Id          int64  `json:"id" example:"1800783867820785152"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type EducationDetail struct {
	Id        int64               `json:"id" example:"1800783867820785152"`
	School    string              `json:"school"`
	Major     string              `json:"major"`
	Level     enum.EducationLevel `json:"level"` // 0: 学士 1: 硕士 2: 博士
	StartTime times.LocalDate     `json:"startTime"`
	EndTime   *times.LocalDate    `json:"endTime"`
}

type ResumeDetail struct {
	Id          int64  `json:"id" example:"1800783867820785152"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type WorkDetail struct {
	Id        int64           `json:"id" example:"1800783867820785152"`
	JobTitle  string          `json:"jobTitle"`
	Company   CompanyDetail   `json:"company"`
	City      string          `json:"city"`
	StartTime times.LocalDate `json:"startTime"`
	EndTime   times.LocalDate `json:"endTime"`
	Content   string          `json:"content"`
	Summary   string          `json:"summary"`
}

type CompanyDetail struct {
	Id          int64  `json:"id" example:"1800783867820785152"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type SkillDetail struct {
	Id      int64  `json:"id" example:"1800783867820785152"`
	Content string `json:"content"`
}
