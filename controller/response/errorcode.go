package response

const (
	SUCCESS        = 200
	INVALID_PARAMS = 400
	ERROR          = 500
	FILENOTEXIST   = 404
)

type ErrorCode struct {
	code int
}

func (e *ErrorCode) Error() string {
	return GetMsg(e.code)
}

func (e *ErrorCode) Code() int {
	return e.code
}

func NewError(code int) error {
	return &ErrorCode{code}
}
