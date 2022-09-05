package fail

type DatabaseConnectFailure struct {
	M string
	E error
}

func (e *DatabaseConnectFailure) Error() string {
	return e.M
}

type SqlInsertFailure struct {
	M string
	E error
}

func (e *SqlInsertFailure) Error() string {
	return e.M
}

type SqlSelectFailure struct {
	M string
	E error
}

func (e *SqlSelectFailure) Error() string {
	return e.M
}

type SqlSelectNotFoundFailure struct {
	M string
	E error
}

func (e *SqlSelectNotFoundFailure) Error() string {
	return e.M
}

type SqlUpdateFailure struct {
	M string
	E error
}

func (e *SqlUpdateFailure) Error() string {
	return e.M
}

type SqlDeleteFailure struct {
	M string
	E error
}

func (e *SqlDeleteFailure) Error() string {
	return e.M
}

type GetLastInsertIdFailure struct {
	M string
	E error
}

func (e *GetLastInsertIdFailure) Error() string {
	return e.M
}
