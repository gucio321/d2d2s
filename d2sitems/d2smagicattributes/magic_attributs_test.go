package d2smagicattributes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testdataString() map[string]*MagicAttribute {
	return map[string]*MagicAttribute{
		"2 * 8 != 7": {
			Name:   "{0} * {1} != {2}",
			Values: []int64{2, 8, 7},
		},
		"Adds 20-40 Poison Damage over 3 Seconds": {
			Name:   "Adds {2}-{0} Poison Damage over {1} Seconds",
			Values: []int64{40, 3, 20},
		},
	}
}

func Test_String(t *testing.T) {
	for key, value := range testdataString() {
		s := value.String()

		assert.Equal(t, key, s, fmt.Sprintf("unexpected name received: %s, expected %s", s, key))
	}
}
