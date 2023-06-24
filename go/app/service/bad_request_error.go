package service

func NewBadRequestError(text string) error {
	return &BadRequestError{
		s: text,
	}
}

type BadRequestError struct {
	s string
}

func (e *BadRequestError) Error() string {
	return e.s
}
