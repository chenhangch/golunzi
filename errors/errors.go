package errors

type WithCode struct {
	err error
	code int
	cause error
	*stack
}

