package util

// Custom error type with error code and message.
type OpError struct {
	Code int
	Msg  string
}

func (e OpError) Error() string {
	return e.Msg
}
