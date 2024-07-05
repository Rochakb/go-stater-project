package model

import "database/sql"

type Employee struct {
	EmployeeId int           `json:"empId"`
	Name       string        `json:"name"`
	DOB        string        `json:"dob"`
	Department string        `json:"department"`
	Salary     float64       `json:"salary"`
	BossId     sql.NullInt64 `json:"bossId"`
}
