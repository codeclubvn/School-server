package config

var (
	StatusActive   = "on"
	StatusInactive = "off"
)

type Role string

const (
	USER    Role = "USER"
	TEACHER Role = "TEACHER"
)
