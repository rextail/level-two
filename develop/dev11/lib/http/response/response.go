package response

const (
	okStatus    = "OK"
	errorStatus = "Error"

	okCode            = "200"
	validationErrCode = "400"
	otherErrCode      = "500"
	usecaseErrCode    = "503"
)

type Response struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Err    string `json:"error,omitempty"`
}

func OK() Response {
	return Response{
		Status: okStatus,
		Code:   okCode,
	}
}

func Error(msg, code string) Response {
	return Response{
		Status: errorStatus,
		Code:   code,
		Err:    msg,
	}
}
