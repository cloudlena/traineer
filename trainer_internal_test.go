package traineer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrainerModifyMood(t *testing.T) {
	assert := assert.New(t)

	cases := map[string]struct {
		trainer      Trainer
		modification float64
		expectedMood float64
	}{
		"reduces mood": {
			trainer:      Trainer{PunishmentMultiplier: 1},
			modification: -5,
			expectedMood: -5,
		},
		"reduces mood adjusted": {
			trainer:      Trainer{PunishmentMultiplier: 0.2},
			modification: -5,
			expectedMood: -1,
		},
		"doesn't reduce mood below min": {
			trainer: Trainer{
				PunishmentMultiplier: 1,
				mood:                 TrainerMinMood,
			},
			modification: -5,
			expectedMood: TrainerMinMood,
		},
		"doesn't reduce mood above max": {
			trainer: Trainer{
				PunishmentMultiplier: 1,
				mood:                 TrainerMaxMood,
			},
			modification: 5,
			expectedMood: TrainerMaxMood,
		},
	}

	for tcID, tc := range cases {
		t.Run(tcID, func(t *testing.T) {
			tc.trainer.modifyMood(tc.modification)

			assert.Equal(tc.expectedMood, tc.trainer.mood, tcID)
		})
	}
}
