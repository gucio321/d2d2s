package d2smagicattributes

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
)

const (
	idLen         = 9
	endOfListMark = 1<<idLen - 1
)

// MagicAttributes represents a list of magic attributes
type MagicAttributes []MagicAttribute

// Load loads magic attributes for n item
func (m *MagicAttributes) Load(sr *datareader.Reader) error {
	for {
		id := sr.GetBits16(idLen)
		if id == endOfListMark {
			break
		}

		prop := GetPropertyData(id)
		if prop == nil {
			return errors.New("unexpected magical property")
		}

		var values []int64

		for _, bitLength := range prop.Bits {
			val := sr.GetBits32(int(bitLength))
			if prop.Bias != 0 {
				val -= uint32(prop.Bias)
			}

			values = append(values, int64(val))
		}

		attr := MagicAttribute{
			ID:     id,
			Name:   prop.Name,
			Values: values,
		}

		*m = append(*m, attr)
	}

	return nil
}

// Encode encodes magical attributes back into byte slice
func (m *MagicAttributes) Encode(sw *datautils.StreamWriter) (err error) {
	for _, a := range *m {
		sw.PushBits16(a.ID, idLen)

		prop := GetPropertyData(a.ID)
		if prop == nil {
			return errors.New("unexpected magical property")
		}

		for n, bitLength := range prop.Bits {
			val := a.Values[n]
			if prop.Bias != 0 {
				val += int64(prop.Bias)
			}

			sw.PushBits16(uint16(val), int(bitLength))
		}
	}

	sw.PushBits16(endOfListMark, idLen)

	return nil
}

// MagicAttribute represents magic attributes data
type MagicAttribute struct {
	ID     uint16
	Name   string
	Values []int64
}

func (m *MagicAttribute) String() string {
	result := m.Name

	for i, a := range m.Values {
		result = strings.ReplaceAll(result, "{"+strconv.Itoa(i)+"}", strconv.Itoa(int(a)))
	}

	return result
}
