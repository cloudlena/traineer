package traineer

import (
	"strings"
	"time"
)

// Reward is something a trainer uses to reward a user.
type Reward struct {
	Entity
	MinMood   float64
	TimeLimit time.Duration
	Value     float64
}

// NewReward creates a new reward in the DB.
func NewReward(r Reward) (Reward, error) {
	return r, nil
}

// Validate validates a reward.
func (r Reward) Validate() error {
	if len(strings.TrimSpace(r.Name)) < 2 {
		return errInvalid
	}

	return nil
}

// GetReward retrieves a reward from the DB and returns it.
func GetReward(id string) (Reward, error) {
	return Reward{}, nil
}
