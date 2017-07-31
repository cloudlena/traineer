package traineer

// NotFoundError is an error used for things that haven't been found.
type NotFoundError struct {
	msg  string
	code int
}

// Error returns the error string.
func (e NotFoundError) Error() string {
	return e.msg
}

// Errors commonly used throughout the package.
var (
	errNotFound             = NotFoundError{msg: "not found", code: 4040}
	errTrainerNoPunishments = NotFoundError{msg: "trainer doesn't know any punishments", code: 4041}
	errTrainerNoRewards     = NotFoundError{msg: "trainer doesn't know any rewards", code: 4042}
)
