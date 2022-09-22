package model

type Department struct {
	Deptid  int    `json:"depTDID"`
	Name    string `json:"depName" validate:"required,len=4"`
	Descdep string `json:"descDep" validate:"required,len=10"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}
