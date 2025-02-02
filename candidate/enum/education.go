package enum

// EducationLevel education level
type EducationLevel struct {
	Value       int
	Description string
}

var (
	BACHELOR = EducationLevel{Value: 0, Description: "学士"}
	MASTER   = EducationLevel{Value: 1, Description: "硕士"}
	DOCTORAL = EducationLevel{Value: 2, Description: "博士"}
)

var educationLevelStatusByValue = map[int]EducationLevel{
	BACHELOR.Value: BACHELOR,
	MASTER.Value:   MASTER,
	DOCTORAL.Value: DOCTORAL,
}

func EducationLevelFromValue(code int) EducationLevel {
	status, _ := educationLevelStatusByValue[code]
	return status
}
