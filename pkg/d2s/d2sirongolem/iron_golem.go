package d2sirongolem

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
)

const golemHeaderID = "kf"

// New creates a new IronGolem
func New() *IronGolem {
	result := &IronGolem{
		HasGolem: false,
		Item:     &d2sitems.Item{},
	}

	return result
}

// IronGolem represents an iron golem
type IronGolem struct {
	HasGolem bool
	Item     *d2sitems.Item
}

// Load loads a golem's data
func (i *IronGolem) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != golemHeaderID {
		return errors.New("unexpected golem header")
	}

	i.HasGolem = sr.GetByte() == 1

	if !i.HasGolem {
		return nil // no golem
	}

	item := &d2sitems.Items{}
	if err := item.LoadList(sr, 1); err != nil {
		return fmt.Errorf("error lading golem item: %w", err)
	}

	i.Item = (*item)[0]

	return nil
}

// Encode encodes iron golem data into a byte slice
func (i *IronGolem) Encode(sw *d2datautils.StreamWriter) {
	sw.PushBytes([]byte(golemHeaderID)...)

	if !i.HasGolem {
		sw.PushBytes(0)
		return
	}

	sw.PushBytes(1)

	sw.PushBytes(i.Item.Encode()...)
}
