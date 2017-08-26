package traineer

// Confession is something a user confesses to a trainer.
type Confession struct {
	mainObject
	Value float64
}

// NewConfession creates a new confession in the DB.
func NewConfession(c Confession) (Confession, error) {
	return c, nil
}

// GetConfession retrieves a confession from the DB and returns it.
func GetConfession(id string) (Confession, error) {
	return Confession{}, nil
}
