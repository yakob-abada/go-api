package service

func NewNotFoundError(text string) error {
	return &NotFoundError{
		s: text,
	}
}

type NotFoundError struct {
	s string
}

func (e *NotFoundError) Error() string {
	return e.s
}
