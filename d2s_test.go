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

	x, err := Unmarshal(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "example", x.Name, "Unexpected name read")
}

// nolint:wsl // just test
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

	x, err := Unmarshal(data)
	if err != nil {
		t.Fatal(err)
	}

	_, err = x.Encode()
	if err != nil {
		t.Fatal(err)
	}

	// nolint:gocritic // will use later
	/*
		// nolint:gosec // only tests
		err = ioutil.WriteFile("/tmp/out.d2s", d, 0o600)
		if err != nil {
			t.Fatal(err)
		}

		_, err = Unmarshal(d)
		if err != nil {
			t.Fatal(err)
		}
	*/

	// nolint:gocritic // will use later
	// m := 16
	// fmt.Println(d[m:])
	// fmt.Println(data[m:len(d)])
	// fmt.Println(len(d[m:]), len(data[m:len(d)]))
	// fmt.Println(bytes.Equal(data[m:len(d)], d[m:]))
	// t.Fail()
}

func Test_New(t *testing.T) {
	x := New()
	if _, err := x.Encode(); err != nil {
		t.Fatal(err)
	}
}
