package d2sstats

import (
	"testing"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/stretchr/testify/assert"
)

func testdata() map[*Stats][]byte {
	return map[*Stats][]byte{
		{
			Strength:          5,
			Energy:            29,
			Dexterity:         105,
			Vitality:          208,
			UnusedStats:       0,
			UnusedSkillPoints: 5,
			CurrentHP:         200,
			MaxHP:             200,
			CurrentMana:       500,
			MaxMana:           512,
			CurrentStamina:    1000,
			MaxStamina:        2000,
			Level:             40,
			Experience:        2048,
			Gold:              20000,
			StashedGold:       100000,
		}: {
			103, 102, 0, 10, 8, 208, 129, 128, 52, 6, 64, 83, 160, 192, 0, 0, 50, 56, 0, 128, 12, 16,
			0, 208, 135, 4, 0, 0, 66, 1, 0, 250, 88, 0, 0, 125, 24, 160, 26, 0, 32, 0, 0, 56, 0, 113, 2, 240, 0, 212, 48, 192, 127,
		},
	}
}

func Test_Load(t *testing.T) {
	td := testdata()

	for key, value := range td {
		sr := datautils.CreateBitMuncher(value, 0)
		l := New()

		err := l.Load(sr)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, key, l, "unexpected stats loaded")
	}
}

func Test_Encode(t *testing.T) {
	td := testdata()

	for key, value := range td {
		data, err := key.Encode()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, value, data, "unexpected data after encoding stats")
	}
}
