package d2s

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/internal/testdata"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sstats"
)

func Test_Load(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		charname string
	}{
		// {"empty character created by d2d2s.New()", "example.d2s", "example"},
		{"New character created by Diablo II: Lord of Destruction v1.14d", "newGameChar.d2s", "emptychar"},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			data, err := testdata.Testdata.ReadFile(test.filename)
			if err != nil {
				tt.Fatal(err)
			}

			x, err := Load(data)
			if err != nil {
				tt.Fatal(err)
			}

			assert.Equal(tt, test.charname, x.Name, "Unexpected name read")
		})
	}
}

func Test_LoadEncode(t *testing.T) {
	mxl := New()
	mxl.Stats.UserStatMap(
		map[d2sstats.StatID]int{
			d2sstats.Strength:       11,
			d2sstats.Energy:         11,
			d2sstats.Dexterity:      11,
			d2sstats.Vitality:       11,
			d2sstats.UnusedStats:    11,
			d2sstats.UnusedSkills:   8,
			d2sstats.CurrentHP:      21,
			d2sstats.MaxHP:          21,
			d2sstats.CurrentMana:    21,
			d2sstats.MaxMana:        21,
			d2sstats.CurrentStamina: 21,
			d2sstats.MaxStamina:     21,
			d2sstats.Level:          8,
			d2sstats.Experience:     32,
			d2sstats.Gold:           22,
			d2sstats.StashedGold:    25,
			88:                      48,
		},
	)
	mxl.Skills.SetSkillsCount(95)
	mxl.Items.IgnoreErrors()

	tests := []struct {
		name     string
		filename string
		base     *D2S
	}{
		// {"empty character created by d2d2s.New()", "example.d2s"},
		{"New character created by Diablo II: Lord of Destruction v1.14d", "newGameChar.d2s", New()},
		{"MedianXL char", "maxsocket.d2s", mxl},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			data, err := testdata.Testdata.ReadFile(test.filename)
			if err != nil {
				tt.Fatal(err)
			}

			test.base.Load(data)
			if err != nil {
				tt.Fatal(err)
			}

			out, err := test.base.Encode()
			if err != nil {
				tt.Fatal(err)
			}

			assert.Equal(tt, data, out, "Unexpected character data after encoding")
		})
	}
}

func Test_New(t *testing.T) {
	x := New()
	if _, err := x.Encode(); err != nil {
		t.Fatal(err)
	}
}
