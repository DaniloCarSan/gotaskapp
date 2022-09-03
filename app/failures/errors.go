package fail

type SignInFailure struct {
	M string
	E error
}

func (e *SignInFailure) Error() string {
	return e.M
}

type GenerateJwtTokenFailure struct {
	M string
	E error
}

func (e *GenerateJwtTokenFailure) Error() string {
	return e.M
}
