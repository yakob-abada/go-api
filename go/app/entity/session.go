package entity

type Session struct {
	Id       int8   `json:"id"`
	Time     string `json:"time"`
	Name     string `json:"name"`
	Duration int8   `json:"duration"`
	IsFull   bool   `json:"is_full"`
}
