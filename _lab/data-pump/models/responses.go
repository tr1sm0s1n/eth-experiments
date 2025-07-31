package models

type (
	PingResponse struct {
		Status  string `json:"status"`
		Time    string `json:"time"`
		Latency string `json:"latency"`
	}

	SuccessResponse struct {
		Success bool   `json:"success"`
		Data    string `json:"data,omitempty"`
	}

	FailureResponse struct {
		Code  int    `json:"error_code"`
		Error string `json:"error_message"`
	}
)
