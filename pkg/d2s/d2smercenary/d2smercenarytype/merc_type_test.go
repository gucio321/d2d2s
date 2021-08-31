package d2smercenarytype

import (
	"testing"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	testdata := map[uint16]*MercenaryType{
		10: {
			Class:      MercDesert,
			Difficulty: d2enum.DifficultyNightmare,
			Attribute:  MercAttributeDefensive,
		},
		23: {
			Class:      MercSorcerer,
			Difficulty: d2enum.DifficultyHell,
			Attribute:  MercAttributeLightning,
		},
		27: {
			Class:      MercBarbarian,
			Difficulty: d2enum.DifficultyNightmare,
			Attribute:  MercAttributeBarb2,
		},
	}

	for id, merc := range testdata {
		r := Load(id)

		assert.Equal(t, merc, r, "unexpected mercenary data returned")
	}
}

func Test_Encode(t *testing.T) {
	testdata := map[uint16]*MercenaryType{
		10: {
			Class:      MercDesert,
			Difficulty: d2enum.DifficultyNightmare,
			Attribute:  MercAttributeDefensive,
		},
		23: {
			Class:      MercSorcerer,
			Difficulty: d2enum.DifficultyHell,
			Attribute:  MercAttributeLightning,
		},
		27: {
			Class:      MercBarbarian,
			Difficulty: d2enum.DifficultyNightmare,
			Attribute:  MercAttributeBarb2,
		},
	}

	for id, merc := range testdata {
		i := merc.Encode()

		assert.Equal(t, id, i, "unexpected mercenary id returned")
	}
}
