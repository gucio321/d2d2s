package d2d2s

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/gucio321/d2d2s/datautils"
)

const golemHeaderID = "kf"

// IronGolem represents an iron golem
type IronGolem struct {
	Item *Item
}

// Load loads a golem's data
func (i *IronGolem) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != golemHeaderID {
		return errors.New("unexpected golem header")
	}

	hasGolem := sr.GetByte() == 1

	if !hasGolem {
		return nil // no golem
	}

	item := &Items{}
	err := item.LoadList(sr, 1)
	if err != nil {
		return err
	}

	i.Item = &(*item)[0]

	return nil
}

func (i *IronGolem) Encode(sw *d2datautils.StreamWriter) {
	sw.PushBytes([]byte(golemHeaderID)...)
	hasGolem := i.Item != nil
	if !hasGolem {
		sw.PushBytes(0)
		return
	}
	sw.PushBytes(1)

	sw.PushBytes(i.Item.Encode()...)
}
