package d2d2s

import (
	"errors"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/nokka/d2s"
)

const golemHeaderID = "kf"

// IronGolem represents an iron golem
type IronGolem struct {
	Item *d2s.Item
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

	i.Item = &item.Items[0]

	return nil
}
