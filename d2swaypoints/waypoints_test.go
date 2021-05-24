package d2swaypoints

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/d2senums"
)

func testdata() map[[NumWaypointsBytes]byte]*Waypoints {
	return map[[NumWaypointsBytes]byte]*Waypoints{
		[NumWaypointsBytes]byte{
			87, 83, 1, 0, 0, 0, 80, 0, 2, 1, 75, 163, 70, 220, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 2, 1, 75, 163, 70, 220, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 2, 1, 75, 163, 70, 220, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		}: &Waypoints{
			d2enum.DifficultyNormal: &[d2senums.NumActs][]bool{
				0: []bool{true, true, false, true, false, false, true, false, true},
				1: []bool{true, false, false, false, true, false, true, false, true},
				2: []bool{true, false, false, false, true, false, false, false, true},
				3: []bool{true, true, false},
				4: []bool{true, true, true, false, false, false, true, false, true},
			},
			d2enum.DifficultyNightmare: &[d2senums.NumActs][]bool{
				0: []bool{true, true, false, true, false, false, true, false, true},
				1: []bool{true, false, false, false, true, false, true, false, true},
				2: []bool{true, false, false, false, true, false, false, false, true},
				3: []bool{true, true, false},
				4: []bool{true, true, true, false, false, false, true, false, true},
			},
			d2enum.DifficultyHell: &[d2senums.NumActs][]bool{
				0: []bool{true, true, false, true, false, false, true, false, true},
				1: []bool{true, false, false, false, true, false, true, false, true},
				2: []bool{true, false, false, false, true, false, false, false, true},
				3: []bool{true, true, false},
				4: []bool{true, true, true, false, false, false, true, false, true},
			},
		},
		[NumWaypointsBytes]byte{
			87, 83, 1, 0, 0, 0, 80, 0, 2, 1, 75, 167, 70, 204, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 2, 1, 1, 195, 124, 126, 53, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2,
			1, 75, 163, 110, 92, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		}: &Waypoints{
			d2enum.DifficultyNormal: &[d2senums.NumActs][]bool{
				0: []bool{true, true, false, true, false, false, true, false, true},
				1: []bool{true, true, false, false, true, false, true, false, true},
				2: []bool{true, false, false, false, true, false, false, false, true},
				3: []bool{true, false, false},
				4: []bool{true, true, true, false, false, false, true, false, false},
			},
			d2enum.DifficultyNightmare: &[d2senums.NumActs][]bool{
				0: []bool{true, false, false, false, false, false, false, false, true},
				1: []bool{true, false, false, false, false, true, true, false, false},
				2: []bool{true, true, true, true, true, false, false, true, true},
				3: []bool{true, true, true},
				4: []bool{true, false, true, false, true, false, true, true, false},
			},
			d2enum.DifficultyHell: &[d2senums.NumActs][]bool{
				0: []bool{true, true, false, true, false, false, true, false, true},
				1: []bool{true, false, false, false, true, false, true, false, true},
				2: []bool{true, true, false, true, true, false, false, false, true},
				3: []bool{true, true, false},
				4: []bool{true, false, true, false, true, false, false, false, false},
			},
		},
	}
}

func Test_Load(t *testing.T) {
	for key, value := range testdata() {
		w := New()
		w.Load(&key)

		assert.Equal(t, value, w, "unexpected WP data loaded")
	}
}

func Test_Encode(t *testing.T) {
	for key, value := range testdata() {
		d := value.Encode()

		assert.Equal(t, key, d, "unexpected WP data encoded")
	}
}
