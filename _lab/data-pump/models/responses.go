package models

type (
	PingResponse struct {
		Status  string `json:"status"`
		Time    string `json:"time"`
		Latency string `json:"latency"`
	}

	TokenResponse struct {
		Success bool   `json:"success"`
		Token   string `json:"token"`
	}

	FailureResponse struct {
		Code  int    `json:"error_code"`
		Error string `json:"error_message"`
	}
)
