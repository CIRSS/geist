package geist

type Error interface {
	Error() string
	Unwrap() error
}

type errorBase struct {
	summary string
	err     error
}

func NewGeistError(summary string, err error, prepend bool) Error {
	ge := errorBase{}
	ge.summary = summary
	ge.err = err
	if prepend {
		ge.summary += ": " + err.Error()
	}
	return &ge
}

func (e *errorBase) Error() string {
	return e.summary
}

func (e *errorBase) Unwrap() error {
	return e.err
}
