package responses

type NotFound struct {
	Reason string
}

func (e NotFound) Error() string {
	return "not found"
}
