package d2d2s

import (
	"errors"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/gucio321/d2d2s/itemdata"
)

type MagicAttribute struct {
	ID     uint16
	Name   string
	Values []int64
}

func (i *Items) LoadMagicAttributes(sr *datautils.BitMuncher, n int) error {
	for {
		id := uint16(sr.GetBits(9))
		if id == 0x1ff {
			break
		}

		prop, ok := itemdata.MagicalProperties[uint64(id)]
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

		// attr := MagicAttribute{
		_ = MagicAttribute{
			ID:     id,
			Name:   prop.Name,
			Values: values,
		}

		// i.Items[n].MagicAttributes = append(i.Items[n].MagicAttributes, attr)
	}

	return nil
}
