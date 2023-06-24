package service

func NewUnprocessableEntityError(text string) error {
	return &UnprocessableEntityError{
		s: text,
	}
}

type UnprocessableEntityError struct {
	s string
}

func (e *UnprocessableEntityError) Error() string {
	return e.s
}
