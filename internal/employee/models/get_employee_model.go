package models

type (
	GetEmpByIdRequest struct {
		EmpId string
	}

	GetEmpByIdResponse struct {
		Response interface{} `json:"response"`
	}
)
