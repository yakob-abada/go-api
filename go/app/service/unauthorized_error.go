package service

func NewUnauthorizedError(text string) error {
	return &UnauthorizedError{
		s: text,
	}
}

type UnauthorizedError struct {
	s string
}

func (e *UnauthorizedError) Error() string {
	return e.s
}
