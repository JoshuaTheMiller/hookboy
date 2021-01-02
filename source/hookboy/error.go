package hookboy

type hookboyError struct {
	description   string
	internalError error
}

func (pe hookboyError) Error() string {
	return pe.description
}

func (pe hookboyError) InternalError() error {
	return pe.internalError
}

func (pe hookboyError) Description() string {
	return pe.description
}

// WrapError returns a new HookboyError, which satisfies the Error interface
func WrapError(err error, desc string) Error {
	return hookboyError{
		description:   desc,
		internalError: err,
	}
}

// NewError returns a new HookboyError, which satisfies the Error interface
func NewError(desc string) Error {
	return hookboyError{
		description: desc,
	}
}

// Error are errors returned by Hookboy
type Error interface {
	Error() string
	Description() string
	InternalError() error
}
