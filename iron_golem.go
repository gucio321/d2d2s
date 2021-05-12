package d2d2s

import (
	"errors"
	"fmt"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/nokka/d2s"
)

const golemHeaderID = "kf"

type IronGolem struct {
	Item *d2s.Item
}

func (i *IronGolem) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2)
	if string(id) != golemHeaderID {
		return errors.New("unexpected golem header")
	}

	hasGolem := sr.GetByte() == 1

	if !hasGolem {
		fmt.Println("nec but no golem")
		return nil // no golem
	}

	item, err := d2s.ParseItemList(sr, 1)
	if err != nil {
		return err
	}

	i.Item = &item[0]

	return nil
}
