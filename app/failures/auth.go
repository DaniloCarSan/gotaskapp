package fail

type SignInFailure struct {
	M string
	E error
}

func (e *SignInFailure) Error() string {
	return e.M
}

type SignUpFailure struct {
	M string
	E error
}

func (e *SignUpFailure) Error() string {
	return e.M
}
