package model

var (
	UserStatusActive = "active"
	UserStatusFreeze = "freeze"
	UserStatusDelete = "delete"
	UserStatusAll    = []string{UserStatusActive, UserStatusFreeze, UserStatusDelete}
)
