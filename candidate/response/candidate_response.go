package response

import "github.com/huhx/headhunter-api/candidate/enum"

type CandidateResponse struct {
	Id         int64               `json:"id"`
	Avatar     string              `json:"avatar"`
	Name       string              `json:"name"`
	Age        int                 `json:"age"`
	Address    string              `json:"address"`
	Level      enum.EducationLevel `json:"level"` // 0: 学士 1: 硕士 2: 博士
	Educations []EducationItem     `json:"educations"`
	Works      []WorkItem          `json:"works"`
}

type EducationItem struct {
	School string `json:"school"`
	Major  string `json:"major"`
}

type WorkItem struct {
	Company  string `json:"company"`
	JobTitle string `json:"jobTitle"`
}
