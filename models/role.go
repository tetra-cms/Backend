package models

type Role string

const (
	RoleUser     Role = "USER"
	RoleEmployee Role = "EMPLOYEE"
	RoleAdmin    Role = "ADMIN"
)
