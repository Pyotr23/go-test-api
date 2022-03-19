package models

import "github.com/google/uuid"

type User struct {
	Id         uuid.UUID `json:"id"`
	FirstName  string    `json:"firstName"`
	Surname    string    `json:"surname"`
	MiddleName string    `json:"middleName"`
	FIO        string    `json:"fio"`
	Sex        Sex       `json:"sex"`
	Age        byte      `json:"age"`
}

type Sex byte

const (
	Male Sex = iota
	Female
)
