package traineer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

// Trainer is a personal trainer or a coach.
type Trainer struct {
	MainObject
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
		t.triggerScenario()

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
		return errors.Wrap(errNotFound, "trainer doesn't know confession")
	}

	c, err := GetConfession(confessionID)
	if err != nil {
		return errors.Wrap(err, "error getting confession")
	}

	t.modifyMood(c.Value)

	return nil
}

// AskPermission asks the trainer for permission to do something.
func (t Trainer) AskPermission(permissionID string) (bool, error) {
	if !sliceContains(t.Permissions, permissionID) {
		return false, errors.Wrap(errNotFound, "trainer doesn't know permission")
	}

	p, err := GetPermission(permissionID)
	if err != nil {
		return false, errors.Wrap(err, "error getting permission")
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
		return Punishment{}, errors.Wrap(err, "error getting punishment")
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
			return Reward{}, errors.Wrap(err, "error getting reward")
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
		t.mood = t.mood + (val * t.RewardMultiplier)
	} else {
		t.mood = t.mood + (val * t.PunishmentMultiplier)
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
		return errors.Wrap(err, "error getting scenario")
	}

	for _, st := range s.Steps {
		fmt.Println(st.Title)
		fmt.Println(st.Description)
	}

	t.modifyMood(s.Reward)

	return nil
}
