package d2sdifficulty

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/stretchr/testify/assert"
)

func testdata() map[[NumDifficultyBytes]byte]*Difficulty {
	return map[[NumDifficultyBytes]byte]*Difficulty{
		{
			132, 170, 1,
		}: {
			d2enum.DifficultyNormal: &DifficultyLevelStatus{
				4, false, false, false, false, true,
			},
			d2enum.DifficultyNightmare: &DifficultyLevelStatus{
				2, true, false, true, false, true,
			},
			d2enum.DifficultyHell: &DifficultyLevelStatus{
				1, false, false, false, false, false,
			},
		},
		{
			173, 227, 164,
		}: {
			d2enum.DifficultyNormal: &DifficultyLevelStatus{
				5, true, false, true, false, true,
			},
			d2enum.DifficultyNightmare: &DifficultyLevelStatus{
				3, false, false, true, true, true,
			},
			d2enum.DifficultyHell: &DifficultyLevelStatus{
				4, false, false, true, false, true,
			},
		},
	}
}

func Test_Load(t *testing.T) {
	for key, value := range testdata() {
		d := New()
		d.Load(key)

		assert.Equal(t, value, d, "unexpected Difficulty status loaded")
	}
}

func Test_Encode(t *testing.T) {
	for key, value := range testdata() {
		data := value.Encode()

		assert.Equal(t, key, data, "unexpected Difficulty status encoded")
	}
}
