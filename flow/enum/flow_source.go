package enum

import (
	"database/sql/driver"
	"errors"
	"github.com/goccy/go-json"
)

type FlowSource struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

var (
	FlowSourceCandidate = FlowSource{Code: 1, Name: "候选人"}
	FlowSourceCustomer  = FlowSource{Code: 2, Name: "客户"}
	FlowSourceInterview = FlowSource{Code: 3, Name: "面试"}
)

// change to use map
var flowSourceMap = map[int]FlowSource{
	1: FlowSourceCandidate,
	2: FlowSourceCustomer,
	3: FlowSourceInterview,
}

func FlowSourceFromValue(code int) FlowSource {
	return flowSourceMap[code]
}

func (f *FlowSource) UnmarshalJSON(data []byte) error {
	var code int
	if err := json.Unmarshal(data, &code); err != nil {
		return err
	}
	flowSource := FlowSourceFromValue(code)
	*f = flowSource
	return nil
}

func (f *FlowSource) Scan(value interface{}) error {
	code, ok := value.(int64)
	if !ok {
		return errors.New("invalid type for ResourceType")
	}
	*f = FlowSourceFromValue(int(code))
	return nil
}

func (f FlowSource) Value() (driver.Value, error) {
	return f.Code, nil
}
