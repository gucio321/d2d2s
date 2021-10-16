package d2sstats

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/internal/datautils"
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
		{
			Strength:          220,
			Energy:            324,
			Dexterity:         245,
			Vitality:          408,
			UnusedStats:       0,
			UnusedSkillPoints: 0,
			CurrentHP:         5000,
			MaxHP:             6000,
			CurrentMana:       512,
			MaxMana:           1512,
			CurrentStamina:    1024,
			MaxStamina:        8000,
			Level:             99,
			Experience:        112848,
			Gold:              200000,
			StashedGold:       1000000,
		}: {
			103, 102, 0, 184, 9, 64, 148, 128, 122, 6, 96,
			102, 0, 0, 113, 30, 0, 128, 187, 8, 0, 0, 68, 2, 0, 244, 162, 0, 0, 128, 44,
			0, 0, 250, 12, 198, 13, 160, 113, 3, 0, 28, 0, 53, 12, 120, 0, 36, 244, 224, 63,
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

func Test_GetStatLen_incorrect_id(t *testing.T) {
	_, err := StatID(200).GetStatLen()
	assert.Equal(t, ErrIncorrectStatID, err, "Unexpected behavior")
}
