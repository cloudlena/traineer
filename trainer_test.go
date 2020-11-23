package traineer_test

import (
	"testing"

	. "github.com/mastertinner/traineer"
	"github.com/matryer/is"
)

func TestTrainerAskPermission(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it             string
		trainer        Trainer
		permission     Permission
		expectedResult bool
	}{
		{
			it:             "permits if certain",
			trainer:        Trainer{},
			permission:     Permission{MinMood: -5, CertainMood: 0},
			expectedResult: true,
		},
		{
			it:             "doesn't permit if mood is below min",
			trainer:        Trainer{},
			permission:     Permission{MinMood: 5, CertainMood: 10},
			expectedResult: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result, err := tc.trainer.AskPermission("permission-id")
			is.NoErr(err)
			is.Equal(result, tc.expectedResult)
		})
	}
}

func TestTrainerReward(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it                 string
		trainer            Trainer
		value              float64
		expectedRewardName string
	}{
		{
			it: "returns reward",
			trainer: Trainer{
				Rewards: []string{"test-reward"},
			},
			expectedRewardName: "test-reward",
		},
		{
			it: "returns correct reward",
			trainer: Trainer{
				Rewards: []string{
					"test-reward1",
					"test-reward5",
				},
			},
			value:              5,
			expectedRewardName: "test-reward5",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			result, err := tc.trainer.Reward(tc.value)
			is.NoErr(err)

			is.Equal(result.Name, tc.expectedRewardName)
		})
	}
}
