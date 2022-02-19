package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ApiResponse(code int, status string, message string, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Status:  status,
		Message: message,
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}
