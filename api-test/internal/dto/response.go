package dto

// Standard case for every API
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
}

func ErrorResponse(code int, msg string, errs interface{}) APIResponse {
	return APIResponse{
		Code:    code,
		Message: msg,
		Errors:  errs,
	}
}
