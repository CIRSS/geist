package geist

type Error interface {
	Error() string
	Unwrap() error
}

type errorBase struct {
	summary string
	err     error
}

func NewGeistError(summary string, err error) Error {
	ge := errorBase{}
	ge.summary = summary
	ge.err = err
	return &ge
}

func (e *errorBase) Error() string {
	return e.summary
}

func (e *errorBase) Unwrap() error {
	return e.err
}
