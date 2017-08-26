package traineer_test

import (
	"testing"

	. "github.com/mastertinner/traineer"
	"github.com/stretchr/testify/assert"
)

func TestTrainerAskPermission(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]struct {
		trainer        Trainer
		permission     Permission
		expectedResult bool
	}{
		"permits if certain": {
			trainer:        Trainer{},
			permission:     Permission{MinMood: -5, CertainMood: 0},
			expectedResult: true,
		},
		"doesn't permit if mood is below min": {
			trainer:        Trainer{},
			permission:     Permission{MinMood: 5, CertainMood: 10},
			expectedResult: false,
		},
	}

	for tcID, tc := range cases {
		t.Run(tcID, func(t *testing.T) {
			result := tc.trainer.AskPermission(tc.permission)
			assert.Equal(tc.expectedResult, result, tcID)
		})
	}
}

func TestTrainerReward(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]struct {
		trainer            Trainer
		value              float64
		expectedRewardName string
	}{
		"returns reward": {
			trainer: Trainer{
				Rewards: []Reward{
					{Name: "test-reward"},
				},
			},
			expectedRewardName: "test-reward",
		},
		"returns correct reward": {
			trainer: Trainer{
				Rewards: []Reward{
					{
						Name:  "test-reward1",
						Value: 1,
					},
					{
						Name:  "test-reward5",
						Value: 5,
					},
				},
			},
			value:              5,
			expectedRewardName: "test-reward5",
		},
	}

	for tcID, tc := range cases {
		t.Run(tcID, func(t *testing.T) {
			result, err := tc.trainer.Reward(tc.value)
			assert.NoError(err, tcID)

			assert.Equal(tc.expectedRewardName, result.Name, tcID)
		})
	}
}
