package service

func NewInternalServerError(text string) error {
	return &InternalServerError{
		s: text,
	}
}

type InternalServerError struct {
	s string
}

func (e *InternalServerError) Error() string {
	return e.s
}
