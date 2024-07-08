package models

type (
	GetEmpByIdRequest struct {
		EmpId int
	}

	GetEmpByIdResponse struct {
		Response interface{} `json:"response"`
	}

	CreateEmpResponse struct {
		Response interface{} `json:"response"`
	}

	FilterRequest struct {
		BossId int
	}

	FilterResponse struct {
		Response interface{} `json:"response"`
	}

	DeleteEmpRequest struct {
		EmpId int
	}

	DeleteEmpResponse struct {
		Response interface{} `json:"response"`
	}
)
