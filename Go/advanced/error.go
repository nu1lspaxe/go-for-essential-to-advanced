package advanced

type InternalError struct {
	msg string
}

func (ie InternalError) Error() string {
	return ie.msg
}

func New(msg string) error {
	return InternalError{msg: msg}
}
