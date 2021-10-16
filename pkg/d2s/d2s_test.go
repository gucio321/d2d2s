package d2s

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gucio321/d2d2s/internal/testdata"
)

func Test_Load(t *testing.T) {
	data, err := testdata.Testdata.ReadFile("example.d2s")
	if err != nil {
		t.Fatal(err)
	}

	x, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "example", x.Name, "Unexpected name read")
}

func Test_LoadEncode(t *testing.T) {
	data, err := testdata.Testdata.ReadFile("example.d2s")
	if err != nil {
		t.Fatal(err)
	}

	x, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}

	out, err := x.Encode()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, data, out, "Unexpected character data after encoding")
}

func Test_New(t *testing.T) {
	x := New()
	if _, err := x.Encode(); err != nil {
		t.Fatal(err)
	}
}
