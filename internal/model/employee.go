package model

import (
	"time"
)

type Employee struct {
	EmpId      int       `json:"empId"`
	Name       string    `json:"name"`
	DOB        time.Time `json:"dob"`
	Department string    `json:"department"`
	Salary     float64   `json:"salary"`
	BossId     int       `json:"bossId"`
}
