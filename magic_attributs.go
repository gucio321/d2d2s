package d2d2s

import (
	"errors"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/gucio321/d2d2s/itemdata"
)

type MagicAttributes []MagicAttribute

// MagicAttribute represents magic attributes data
type MagicAttribute struct {
	ID     uint16
	Name   string
	Values []int64
}

// LoadMagicAttributes loads magic attributes for n item
func (m *MagicAttributes) Load(sr *datautils.BitMuncher) error {
	for {
		id := uint16(sr.GetBits(9)) // nolint:gomnd // id size (bitfield)
		if id == 0x1ff {            // nolint:gomnd // reached next section
			break
		}

		prop, ok := itemdata.MagicalProperties[id]
		if !ok {
			return errors.New("unexpected magical property")
		}

		var values []int64

		for _, bitLength := range prop.Bits {
			val := sr.GetBits(int(bitLength))
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

		// i.Items[n].MagicAttributes = append(i.Items[n].MagicAttributes, attr)
		*m = append(*m, attr)
	}

	return nil
}

func (m *MagicAttributes) Encode(sw *datautils.StreamWriter) (err error) {
	for _, a := range *m {
		sw.PushBits16(a.ID, 9)

		prop, ok := itemdata.MagicalProperties[a.ID]
		if !ok {
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

	sw.PushBits16(0x1ff, 9)
	return nil
}
