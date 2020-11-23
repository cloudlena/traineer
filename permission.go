package traineer

import (
	"time"
)

// Permission is something a user needs to ask a trainer for permission for.
type Permission struct {
	Entity
	MinMood     float64
	CertainMood float64
	MinInterval time.Duration
	Value       float64
}

// GetPermission retrieves a permission from the DB and returns it.
func GetPermission(id string) (Permission, error) {
	return Permission{}, nil
}
