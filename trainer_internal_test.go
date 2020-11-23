package traineer

import (
	"testing"

	"github.com/matryer/is"
)

func TestTrainerModifyMood(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it           string
		trainer      Trainer
		modification float64
		expectedMood float64
	}{
		{
			it:           "reduces mood",
			trainer:      Trainer{PunishmentMultiplier: 1},
			modification: -5,
			expectedMood: -5,
		},
		{
			it:           "reduces mood adjusted",
			trainer:      Trainer{PunishmentMultiplier: 0.2},
			modification: -5,
			expectedMood: -1,
		},
		{
			it: "doesn't reduce mood below min",
			trainer: Trainer{
				PunishmentMultiplier: 1,
				mood:                 TrainerMinMood,
			},
			modification: -5,
			expectedMood: TrainerMinMood,
		},
		{
			it: "doesn't reduce mood above max",
			trainer: Trainer{
				PunishmentMultiplier: 1,
				mood:                 TrainerMaxMood,
			},
			modification: 5,
			expectedMood: TrainerMaxMood,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			tc.trainer.modifyMood(tc.modification)

			is.Equal(tc.trainer.mood, tc.expectedMood)
		})
	}
}
