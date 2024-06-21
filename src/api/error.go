package api

type ApiError struct {
	ErrorCode int         `json:"code"`
	Details   interface{} `json:"details,omitempty"`
}
