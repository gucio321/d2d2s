package d2sdifficulty

import (
	"testing"

	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
	"github.com/stretchr/testify/assert"
)

func testdata() map[[NumDifficultyBytes]byte]*Difficulty {
	return map[[NumDifficultyBytes]byte]*Difficulty{
		{
			132, 170, 1,
		}: {
			d2senums.DifficultyNormal: &DifficultyLevelStatus{
				4, false, false, false, false, true,
			},
			d2senums.DifficultyNightmare: &DifficultyLevelStatus{
				2, true, false, true, false, true,
			},
			d2senums.DifficultyHell: &DifficultyLevelStatus{
				1, false, false, false, false, false,
			},
		},
		{
			173, 227, 164,
		}: {
			d2senums.DifficultyNormal: &DifficultyLevelStatus{
				5, true, false, true, false, true,
			},
			d2senums.DifficultyNightmare: &DifficultyLevelStatus{
				3, false, false, true, true, true,
			},
			d2senums.DifficultyHell: &DifficultyLevelStatus{
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
