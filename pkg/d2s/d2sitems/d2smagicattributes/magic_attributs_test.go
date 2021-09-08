package d2smagicattributes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/internal/datautils"
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

func testdata() map[*MagicAttributes][]byte {
	return map[*MagicAttributes][]byte{
		{
			{333, "-{0}% To Enemy Fire Resistance", []int64{20}},
			{20, "+{0}% Increased chance of blocking", []int64{60}},
		}: {
			77, 41, 40, 240, 255, 1,
		},
	}
}

func Test_Load(t *testing.T) {
	for key, value := range testdata() {
		sr := datautils.CreateBitMuncher(value, 0)
		m := &MagicAttributes{}

		if err := m.Load(sr); err != nil {
			t.Error(err)
		}

		assert.Equal(t, key, m, "unexpected magicall attributes recceived")
	}
}

func Test_Encode(t *testing.T) {
	for key, value := range testdata() {
		sw := datautils.CreateStreamWriter()

		if err := key.Encode(sw); err != nil {
			t.Error(err)
		}

		sw.Align()

		assert.Equal(t, value, sw.GetBytes(), "Unexpected data encoded")
	}
}
