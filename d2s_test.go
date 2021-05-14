package d2d2s

import (
	"io/ioutil"
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
	x.Name = "kejcam"

	d, err := x.Encode()
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile("/tmp/maciej1.d2s", d, 0o644)
	if err != nil {
		t.Fatal(err)
	}

	_, err = Unmarshal(d)
	if err != nil {
		t.Fatal(err)
	}

	/*m := 500
	fmt.Println(d[m:])
	fmt.Println(data[m:len(d)])
	fmt.Println(len(d[m:]), len(data[m:len(d)]))
	fmt.Println(bytes.Equal(data[m:len(d)], d[m:]))
	*/
	t.Fail()
}
