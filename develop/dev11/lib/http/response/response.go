package response

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

const (
	OtherErrCode   = "500"
	UsecaseErrCode = "503"

	okCode            = "200"
	validationErrCode = "400"

	okStatus    = "OK"
	errorStatus = "Error"
)

type Response struct {
	Status string `json:"status"`
	Code   string `json:"code"`
	Result any    `json:"result,omitempty"`
	Err    string `json:"error,omitempty"`
}

func OK() Response {
	return Response{
		Status: okStatus,
		Code:   okCode,
	}
}

func Result(res any) Response {
	return Response{
		Status: okStatus,
		Code:   okCode,
		Result: res,
	}

}

func Error(msg, code string) Response {
	return Response{
		Status: errorStatus,
		Code:   code,
		Err:    msg,
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		errMsgs = append(errMsgs, fmt.Sprintf(`field %s is a required field`, err.Field()))
	}

	return Response{
		Status: errorStatus,
		Code:   validationErrCode,
		Err:    strings.Join(errMsgs, ", "),
	}
}
