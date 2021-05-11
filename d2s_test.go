package d2d2s

import (
	"fmt"
	"os"
	"testing"
)

func Test_Marshal(t *testing.T) {
	testFile, fileErr := os.Open("testdata/testdata.d2s")

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

	fmt.Println(x)
	fmt.Println(x.Items)
	t.Fail()
}
