package d2d2s

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/datautils"
)

const corpseUnknownBytesCount = 12

// Corpse represents a ... corpse ?!
type Corpse struct {
	unknown [corpseUnknownBytesCount]byte
	Items   *Items
}

// Load loads corpse
func (c *Corpse) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2) // nolint:gomnd // header

	if string(id) != "JM" {
		return errors.New("unexpected header")
	}

	isDead := sr.GetUInt16() == 1
	if !isDead {
		return nil
	}

	unknown := sr.GetBytes(corpseUnknownBytesCount)
	copy(c.unknown[:], unknown[:corpseUnknownBytesCount])

	c.Items = &Items{}

	numItems, err := c.Items.LoadHeader(sr)
	if err != nil {
		return err
	}

	if err := c.Items.LoadList(sr, numItems); err != nil {
		return err
	}

	return nil
}

// Encode encodes corpse data into a byte slice
func (c *Corpse) Encode(sw *d2datautils.StreamWriter) (err error) {
	sw.PushBytes([]byte("JM")...)

	isDead := c.Items != nil

	if !isDead {
		sw.PushUint16(0)
		return nil
	}

	sw.PushUint16(1)
	sw.PushBytes(c.unknown[:]...)
	sw.PushBytes(c.Items.Encode()...)

	return nil
}
