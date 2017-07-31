package traineer

import "time"

// Punishment is something a trainer uses to punish a user.
type Punishment struct {
	mainObject
	MaxMood   float64
	TimeLimit time.Duration
	Value     float64
}

// GetPunishment retrieves a punishment from the DB and returns it.
func GetPunishment(id string) (Punishment, error) {
	return Punishment{}, nil
}
