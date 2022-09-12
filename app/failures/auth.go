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

type EmailNotVerifiedFailure struct {
	M string
	E error
}

func (e *EmailNotVerifiedFailure) Error() string {
	return e.M
}

type SendEmailVerificationFailure struct {
	M string
	E error
}

func (e *SendEmailVerificationFailure) Error() string {
	return e.M
}
