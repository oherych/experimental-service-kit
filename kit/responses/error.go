package responses

type NotFound struct {
	Reason   string
	Internal error
}

func (e NotFound) Error() string {
	return "not found"
}
