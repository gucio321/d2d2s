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
		id := uint64(sr.GetBits(9)) // nolint:gomnd // id size (bitfield)
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
			ID:     uint16(id),
			Name:   prop.Name,
			Values: values,
		}

		// i.Items[n].MagicAttributes = append(i.Items[n].MagicAttributes, attr)
		*m = append(*m, attr)
	}

	return nil
}
