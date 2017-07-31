package traineer

// Reward is something a trainer uses to reward a user.
type Reward struct {
	mainObject
	MinMood float64
	Value   float64
}

// GetReward retrieves a reward from the DB and returns it.
func GetReward(id string) (Reward, error) {
	return Reward{}, nil
}
