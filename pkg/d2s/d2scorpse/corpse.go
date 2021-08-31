package d2scorpse

import (
	"errors"
	"fmt"

	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
)

const corpseUnknownBytesCount = 12

// New creates a new corpse data
func New() *Corpse {
	result := &Corpse{}

	return result
}

// Corpse represents a ... corpse ?!
type Corpse struct {
	unknown [corpseUnknownBytesCount]byte
	Items   *d2sitems.Items
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

	c.Items = &d2sitems.Items{}

	numItems, err := c.Items.LoadHeader(sr)
	if err != nil {
		return fmt.Errorf("error loading items header: %w", err)
	}

	if err := c.Items.LoadList(sr, numItems); err != nil {
		return fmt.Errorf("error loading corpse items list: %w", err)
	}

	return nil
}

// Encode encodes corpse data into a byte slice
func (c *Corpse) Encode(sw *datautils.StreamWriter) (err error) {
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
