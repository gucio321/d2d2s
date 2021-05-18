package d2d2s

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	testFile, fileErr := os.Open("testdata/example.d2s")

	if fileErr != nil {
		t.Error("cannot open test data file")
		return
	}

	data := make([]byte, 0)
	buf := make([]byte, 16)

	for {
		numRead, err := testFile.Read(buf)

		data = append(data, buf[:numRead]...)

		if err != nil {
			break
		}
	}

	x, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "example", x.Name, "Unexpected name read")
}

func Test_LoadEncode(t *testing.T) {
	testFile, fileErr := os.Open("testdata/example.d2s")

	if fileErr != nil {
		t.Error("cannot open test data file")
		return
	}

	data := make([]byte, 0)
	buf := make([]byte, 16)

	for {
		numRead, err := testFile.Read(buf)

		data = append(data, buf[:numRead]...)

		if err != nil {
			break
		}
	}

	x, err := Load(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = x.Encode()
	if err != nil {
		t.Fatal(err)
	}
}

func Test_New(t *testing.T) {
	x := New()
	if _, err := x.Encode(); err != nil {
		t.Fatal(err)
	}
}
