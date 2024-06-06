package errors

type Error struct {
	Desc string
	Err  error
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Description() string {
	return e.Desc
}

func NewError(desc string, err error) Error {
	return Error{
		Desc: desc,
		Err:  err,
	}
}
