package d2d2s

import (
	"errors"

	"github.com/gucio321/d2d2s/datautils"
)

const (
	numNPCBytes = 51
	npcHeaderID = "w4"
)

type NPC struct{}

func (n *NPC) Load(data [numNPCBytes]byte) error {
	sr := datautils.CreateStreamReader(data[:])

	id, err := sr.ReadBytes(2)
	if err != nil {
		return err
	}

	if string(id) != npcHeaderID {
		return errors.New("unexpected header ID")
	}

	return nil
}
