package model

type Employee struct {
	EmployeeId int     `json:"empId"`
	Name       string  `json:"name"`
	DOB        string  `json:"dob"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	BossId     int     `json:"bossId"`
}
