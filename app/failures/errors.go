package fail

type GenerateJwtTokenFailure struct {
	M string
	E error
}

func (e *GenerateJwtTokenFailure) Error() string {
	return e.M
}

type PasswordToHashFailure struct {
	M string
	E error
}

func (e *PasswordToHashFailure) Error() string {
	return e.M
}
