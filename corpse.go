package d2d2s

import (
	"errors"

	"github.com/gucio321/d2d2s/datautils"
)

const corpseUnknownBytesCount = 12

type Corpse struct {
	unknown [corpseUnknownBytesCount]byte
	Items   *Items
}

func (c *Corpse) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2)

	if string(id) != "JM" {
		return errors.New("Unexpected header")
	}

	isDead := sr.GetUInt16() == 1
	if !isDead {
		return nil
	}

	unknown := sr.GetBytes(corpseUnknownBytesCount)
	copy(c.unknown[:], unknown[:corpseUnknownBytesCount])

	c.Items = &Items{}
	c.Items.Load(sr)

	return nil
}
