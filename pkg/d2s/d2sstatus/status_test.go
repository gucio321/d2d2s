package d2sstatus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testdata() map[byte]*Status {
	return map[byte]*Status{
		255: { // averything true
			Unknown0:  true,
			Unknown1:  true,
			Hardcore:  true,
			Died:      true,
			Unknown4:  true,
			Expansion: true,
			Ladder:    true,
			Unknown7:  true,
		},
		104: {
			Unknown0:  false,
			Unknown1:  false,
			Hardcore:  false,
			Died:      true,
			Unknown4:  false,
			Expansion: true,
			Ladder:    true,
			Unknown7:  false,
		},
	}
}

func Test_Load(t *testing.T) {
	td := testdata()
	for key, value := range td {
		s := New()
		s.Load(key)

		assert.Equal(t, value, s, "unexpected status structure loaded")
	}
}

func Test_Encode(t *testing.T) {
	td := testdata()
	for key, value := range td {
		d := value.Encode()

		assert.Equal(t, key, d, "unexpected status encoded")
	}
}
