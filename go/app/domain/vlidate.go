package domain

type IValidate interface {
	Struct(interface{}) error
}
