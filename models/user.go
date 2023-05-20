package models

type UserType int

const (
	individual UserType = iota
	legalEntity
)
