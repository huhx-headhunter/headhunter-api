package enum

import (
	"database/sql/driver"
	"errors"
	"github.com/goccy/go-json"
)

type EducationLevel struct {
	Code        int
	Description string
}

var (
	EducationLevelBachelor = EducationLevel{Code: 0, Description: "学士"}
	EducationLevelMaster   = EducationLevel{Code: 1, Description: "硕士"}
	EducationLevelDoctor   = EducationLevel{Code: 2, Description: "博士"}
)

var educationLevelStatusByValue = map[int]EducationLevel{
	EducationLevelBachelor.Code: EducationLevelBachelor,
	EducationLevelMaster.Code:   EducationLevelMaster,
	EducationLevelDoctor.Code:   EducationLevelDoctor,
}

func EducationLevelFromValue(code int) EducationLevel {
	return educationLevelStatusByValue[code]
}

func (f *EducationLevel) UnmarshalJSON(data []byte) error {
	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}
	flowSource := EducationLevelFromValue(code)
	*f = flowSource
	return nil
}

func (f *EducationLevel) Scan(value interface{}) error {
	code, ok := value.(int64)
	if !ok {
		return errors.New("invalid type for EducationLevel")
	}
	*f = EducationLevelFromValue(int(code))
	return nil
}

func (f EducationLevel) Value() (driver.Value, error) {
	return f.Code, nil
}
