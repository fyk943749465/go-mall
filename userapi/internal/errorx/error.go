package errorx

var ParamsError = New(1101, "parameter error")

type BizError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(code int, msg string) *BizError {
	return &BizError{
		Code: code,
		Msg:  msg,
	}
}

func (e *BizError) Error() string {
	return e.Msg
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *BizError) Data() *ErrorResponse {
	return &ErrorResponse{
		e.Code,
		e.Msg,
	}
}
