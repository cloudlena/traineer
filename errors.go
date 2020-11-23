package traineer

// NotFoundError is an error used for things that haven't been found.
type NotFoundError struct {
	msg string
}

// Error returns the error string.
func (e NotFoundError) Error() string {
	return e.msg
}

// UnprocessableDataError is an error used for invalid data input.
type UnprocessableDataError struct {
	msg string
}

// Error returns the error string.
func (e UnprocessableDataError) Error() string {
	return e.msg
}

// Errors commonly used throughout the package.
var (
	errNotFound             = NotFoundError{msg: "not found"}
	errTrainerNoPunishments = NotFoundError{msg: "trainer doesn't know any punishments"}
	errTrainerNoRewards     = NotFoundError{msg: "trainer doesn't know any rewards"}
	errInvalid              = UnprocessableDataError{msg: "invalid data"}
)
