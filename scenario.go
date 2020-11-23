package traineer

import (
	"time"
)

// Scenario is a happening triggered by a trainer which may consist of multiple steps.
type Scenario struct {
	Entity
	MinMood    float64
	MaxMood    float64
	Commonness float64
	Reward     float64
	Punishment float64
	TimeLimit  time.Duration
	Steps      []ScenarioStep
}

// GetScenario retrieves a scenario from the DB and returns it.
func GetScenario(id string) (Scenario, error) {
	return Scenario{}, nil
}

// ScenarioStep is a step of a scenario.
type ScenarioStep struct {
	Title        string
	Description  string
	Reward       float64
	Punishment   float64
	BodyTemplate string
	Options      []ScenarioStepOption
}

// ScenarioStepOption leads to another step or terminates a scenario.
type ScenarioStepOption struct {
	ID       string
	Label    string
	NextStep string
}
