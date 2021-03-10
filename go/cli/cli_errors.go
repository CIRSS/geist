package cli

type CLIError struct {
	summary string
	err     error
}

func NewCLIError(summary string, err error, prepend bool) CLIError {
	e := CLIError{}
	e.summary = summary
	e.err = err
	if prepend {
		e.summary += ": " + err.Error()
	}
	return e
}

func (e CLIError) Error() string {
	return e.summary
}

func (e CLIError) Unwrap() error {
	return e.err
}
