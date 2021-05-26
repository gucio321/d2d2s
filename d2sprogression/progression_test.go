package d2sprogression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testdata() map[byte]*Progression {
	return map[byte]*Progression{
		2: {
			DifficultyLevel: 0,
			Act:             2,
		},
		5: {
			DifficultyLevel: 0,
			Act:             5,
		},
		6: {
			DifficultyLevel: 1,
			Act:             1,
		},
		8: {
			DifficultyLevel: 1,
			Act:             3,
		},
		11: {
			DifficultyLevel: 2,
			Act:             1,
		},
	}
}

func Test_Load(t *testing.T) {
	for key, value := range testdata() {
		p := New()
		p.Load(key)

		assert.Equal(t, value, p, "unexpected progression loaded")
	}
}

func Test_Encode(t *testing.T) {
	for key, value := range testdata() {
		data := value.Encode()

		assert.Equal(t, key, data, "unexpected data encoded")
	}
}
