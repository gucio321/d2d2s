package d2smercenarytype

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

func testdata() map[uint16]*MercenaryType {
	testdata := map[uint16]*MercenaryType{
		10: {
			Class:      MercDesert,
			Difficulty: d2senums.DifficultyNightmare,
			Attribute:  MercAttributeDefensive,
		},
		23: {
			Class:      MercSorcerer,
			Difficulty: d2senums.DifficultyHell,
			Attribute:  MercAttributeLightning,
		},
		27: {
			Class:      MercBarbarian,
			Difficulty: d2senums.DifficultyNightmare,
			Attribute:  MercAttributeBarb2,
		},
	}

	return testdata
}

func Test_Load(t *testing.T) {
	for id, merc := range testdata() {
		r := Load(id)

		assert.Equal(t, merc, r, "unexpected mercenary data returned")
	}
}

func Test_Encode(t *testing.T) {
	for id, merc := range testdata() {
		i := merc.Encode()

		assert.Equal(t, id, i, "unexpected mercenary id returned")
	}
}

func Test_MercenaryType_numAttributes_incorrect_class(t *testing.T) {
	assert.Panics(t, func() { MercClass(20).numAttributes() }, "Unexpected behavior")
}

func Test_GetClassAttributes(t *testing.T) {
	tests := []struct {
		name     string
		class    MercClass
		expected []MercAttribute
	}{
		{"Rogue", MercRogue, []MercAttribute{MercAttributeFireArrow, MercAttributeColdArrow}},
		{"Merc from desert", MercDesert, []MercAttribute{MercAttributeCombat, MercAttributeDefensive, MercAttributeOffensive}},
		{"Sorcerers", MercSorcerer, []MercAttribute{MercAttributeFire, MercAttributeCold, MercAttributeLightning}},
		{"Barbarians", MercBarbarian, []MercAttribute{MercAttributeBarb1, MercAttributeBarb2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, GetClassAttributes(test.class), "Unexpected result")
		})
	}
}

func Test_GetClassAttributes_incorrect_class(t *testing.T) {
	assert.Nil(t, GetClassAttributes(MercClass(20)), "Unexpected value")
}
