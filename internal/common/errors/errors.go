package errors

type Error struct {
	s string
}

func newError(s string) error {
	return &Error{s: s}
}

func (e *Error) Error() string {
	return e.s
}

func StatusCode(s string) int {
	code := statusMessage[s]
	if code < 100 {
		code = statusMessage[statusInternalServerError]
	}
	return code
}
