package datareader

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Read(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
		{"one index - 0", []byte{0}},
		{"one index - non-zero", []byte{20}},
		{"byte slice", []byte{20, 15, 228, 189, 72}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			a := assert.New(tt)
			p := make([]byte, len(test.data))

			if test.data == nil {
				p = nil
			}

			n, err := NewReader(test.data).Read(p)
			a.Nil(err, "Error occurred")
			a.Equal(len(test.data), n, "Unexpected bytes count")
			a.Equal(test.data, p, "unexpected result")
		})
	}
}

func Test_Reade_eof(t *testing.T) {
	_, err := NewReader([]byte{}).Read([]byte{0})
	assert.Equal(t, io.EOF, err, "Unexpected behavior")
}

func Test_SkipBits(t *testing.T) {
	r := NewReader([]byte{5})
	r.SkipBits(5)
	assert.Equal(t, uint64(5), r.bitPosition, "unexpected position")
}

func Test_SkipByte(t *testing.T) {
	r := NewReader([]byte{5})
	r.SkipByte()
	assert.Equal(t, uint64(8), r.bitPosition, "unexpected position")
}

func Test_SkipBytes(t *testing.T) {
	tests := []struct {
		name  string
		data  []byte
		count uint64
	}{
		{"skip 0 bytes with empty data", []byte{}, 0},
		{"skip 0 bytes", []byte{5, 8}, 0},
		{"skip 5 bytes", []byte{5, 8, 29, 48, 79, 102}, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			r := NewReader(test.data)
			r.SkipBytes(test.count)
			assert.Equal(tt, test.count*8, r.bitPosition, "Unexpected bit position")
		})
	}
}

func Test_GetBit(t *testing.T) {
	tests := []struct {
		name     string
		source   []byte
		expected bool
	}{
		{"True", []byte{1}, true},
		{"False", []byte{0}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, NewReader(test.source).GetBit(), "Unexpected value")
		})
	}
}

func Test_GetBits(t *testing.T) {
	tests := []struct {
		name     string
		source   []byte
		count    int
		expected byte
	}{
		{"0, bits count = 0", []byte{0}, 0, 0},
		{"0, bits count = 5", []byte{0}, 5, 0},
		{"7 - variant 1", []byte{23}, 3, 7},
		{"7 - variant 2", []byte{151}, 3, 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, NewReader(test.source).GetBits(test.count), "Unexpected result")
		})
	}
}

func Test_GetBits_incorrect_bit_count(t *testing.T) {
	assert.Panics(t, func() { NewReader([]byte{0}).GetBits(20) }, "Unexpected behavior")
}

func Test_GetUint16(t *testing.T) {
	tests := []struct {
		name     string
		source   []byte
		expected uint16
	}{
		{"Example number", []byte{224, 252}, 64736},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, NewReader(test.source).GetUint16(), "unexpected result")
		})
	}
}

func Test_GetInt16(t *testing.T) {
	tests := []struct {
		name     string
		source   []byte
		expected int16
	}{
		{"Example number (-800)", []byte{224, 252}, -800},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			assert.Equal(tt, test.expected, NewReader(test.source).GetInt16(), "unexpected result")
		})
	}
}
