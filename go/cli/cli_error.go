package cli

type CLIError struct {
	err error
}

func NewCLIError(e error) CLIError {
	return CLIError{e}
}

func (ce CLIError) Error() string {
	return ce.err.Error()
}

func (ce CLIError) Unwrap() error {
	return ce.err
}
