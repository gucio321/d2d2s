package d2d2s

import (
	"errors"
	"fmt"

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
	fmt.Println(len(*c.Items))

	return nil
}
