package traineer

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Trainer is a personal trainer or a coach.
type Trainer struct {
	Entity
	Active               bool
	RewardMultiplier     float64
	PunishmentMultiplier float64
	ScenarioRate         int
	Scenarios            []string
	Rewards              []string
	Punishments          []string
	Permissions          []string
	Confessions          []string
	mood                 float64
}

// Init initiates a trainer and triggers scenarios on a regular basis.
func (t *Trainer) Init() {
	t.Active = true

	go func() {
		err := t.triggerScenario()
		if err != nil {
			log.Fatal(fmt.Errorf("error triggering scenario: %w", err))
		}

		time.Sleep(time.Duration(t.ScenarioRate) * time.Second)
	}()
}

// Mood returns the trainer's current mood.
func (t Trainer) Mood() float64 {
	return t.mood
}

// ConfessTo confesses something to the trainer.
func (t *Trainer) ConfessTo(confessionID string) error {
	if !sliceContains(t.Confessions, confessionID) {
		return fmt.Errorf("trainer doesn't know confession: %w", errNotFound)
	}

	c, err := GetConfession(confessionID)
	if err != nil {
		return fmt.Errorf("error getting confession: %w", err)
	}

	t.modifyMood(c.Value)

	return nil
}

// AskPermission asks the trainer for permission to do something.
func (t Trainer) AskPermission(permissionID string) (bool, error) {
	if !sliceContains(t.Permissions, permissionID) {
		return false, fmt.Errorf("trainer doesn't know permission: %w", errNotFound)
	}

	p, err := GetPermission(permissionID)
	if err != nil {
		return false, fmt.Errorf("error getting permission: %w", err)
	}

	if t.mood < p.MinMood {
		return false, nil
	}

	if t.mood < p.CertainMood {
		pb := (t.mood - p.MinMood) / (p.CertainMood - p.MinMood)

		rand.Seed(time.Now().UnixNano())
		r := rand.Float64()

		if r <= pb {
			return true, nil
		}
	}

	return true, nil
}

// GetPunished imposes a punishment on the user which they must fulfill.
func (t Trainer) GetPunished(val float64) (Punishment, error) {
	if len(t.Punishments) == 0 {
		return Punishment{}, errTrainerNoPunishments
	}

	p, err := GetPunishment(t.Punishments[0])
	if err != nil {
		return Punishment{}, fmt.Errorf("error getting punishment: %w", err)
	}

	return p, nil
}

// Reward treats a user with a reward.
func (t *Trainer) Reward(val float64) (Reward, error) {
	if len(t.Rewards) == 0 {
		return Reward{}, errTrainerNoRewards
	}

	rewards := make([]Reward, 0, len(t.Rewards))
	for _, r := range t.Rewards {
		rew, err := GetReward(r)
		if err != nil {
			return Reward{}, fmt.Errorf("error getting reward: %w", err)
		}

		rewards = append(rewards, rew)
	}

	smallestDeviation := TrainerMaxMood - TrainerMinMood
	for _, r := range rewards {
		currentDeviation := r.Value - val
		if currentDeviation*currentDeviation < smallestDeviation*smallestDeviation {
			smallestDeviation = currentDeviation
		}
	}
	var possibleRewards []Reward
	for _, r := range rewards {
		if r.Value == smallestDeviation {
			possibleRewards = append(possibleRewards, r)
		}
	}
	rand.Seed(time.Now().UnixNano())
	chosen := rand.Intn(len(possibleRewards))
	r := possibleRewards[chosen]

	t.modifyMood(r.Value)

	return r, nil
}

// modifyMood modifies a trainer's mood adjusted by the multiplier.
func (t *Trainer) modifyMood(val float64) {
	if val > 0 {
		t.mood += val * t.RewardMultiplier
	} else {
		t.mood += val * t.PunishmentMultiplier
	}

	if t.mood > TrainerMaxMood {
		t.mood = TrainerMaxMood
	}
	if t.mood < TrainerMinMood {
		t.mood = TrainerMinMood
	}
}

// triggerScenario triggers a random scenario of the trainer.
func (t *Trainer) triggerScenario() error {
	s, err := GetScenario(t.Scenarios[0])
	if err != nil {
		return fmt.Errorf("error getting scenario: %w", err)
	}

	for _, st := range s.Steps {
		fmt.Println(st.Title)
		fmt.Println(st.Description)
	}

	t.modifyMood(s.Reward)

	return nil
}
