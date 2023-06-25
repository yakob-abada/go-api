package entity

import "time"

type Session struct {
	Id       int8      `json:"id"`
	Time     time.Time `json:"time"`
	Name     string    `json:"name"`
	Duration int8      `json:"duration"`
	IsFull   bool      `json:"is_full"`
}
