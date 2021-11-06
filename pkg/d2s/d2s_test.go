package d2s

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/internal/testdata"
)

func Test_Load(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		charname string
	}{
		{"empty character created by d2d2s.New()", "example.d2s", "example"},
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
	tests := []struct {
		name     string
		filename string
	}{
		{"empty character created by d2d2s.New()", "example.d2s"},
		{"New character created by Diablo II: Lord of Destruction v1.14d", "newGameChar.d2s"},
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

			out, err := x.Encode()
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
