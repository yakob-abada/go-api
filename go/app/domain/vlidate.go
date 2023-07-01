package domain

type Validate interface {
	Struct(interface{}) error
}
