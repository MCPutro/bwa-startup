package response

type New struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	//ErrorMessage string      `json:"error_message,omitempty"`
	ErrorDetail interface{} `json:"error_detail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
