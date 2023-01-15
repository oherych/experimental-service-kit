package kit

type NotFound struct {
	Reason   string
	Internal error
}

func (e NotFound) Error() string {
	return "not found"
}

type WrongParameter struct {
	Name     string
	Internal error
}

func (e WrongParameter) Error() string {
	return "wrong parameter " + e.Name
}
