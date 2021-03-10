package geist

// type Error interface {
// 	Error() string
// 	Unwrap() error
// }

type GeistError struct {
	summary string
	err     error
}

func NewGeistError(summary string, err error, prepend bool) GeistError {
	ge := GeistError{}
	ge.summary = summary
	ge.err = err
	if prepend {
		ge.summary += ": " + err.Error()
	}
	return ge
}

func (e GeistError) Error() string {
	return e.summary
}

func (e GeistError) Unwrap() error {
	return e.err
}
