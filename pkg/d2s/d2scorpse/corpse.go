package d2scorpse

import (
	"fmt"

	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
)

const (
	corpseUnknownBytesCount = 12
	corpseHeaderID          = "JM"
)

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
func (c *Corpse) Load(sr *datareader.Reader) error {
	id := sr.GetBytes(len(corpseHeaderID))

	if string(id) != corpseHeaderID {
		return common.ErrUnexpectedHeader
	}

	if isDead := sr.GetUint16() == 1; !isDead {
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

	if isDead := c.Items != nil; !isDead {
		sw.PushUint16(0)
		return nil
	}

	sw.PushUint16(1)
	sw.PushBytes(c.unknown[:]...)
	sw.PushBytes(c.Items.Encode()...)

	return nil
}
