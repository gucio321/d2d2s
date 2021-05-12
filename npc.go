package d2d2s

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	numNPCBytes = 51
	npcHeaderID = "w4"
)

type NPC struct {
	Data []byte
}

func (n *NPC) Load(data [numNPCBytes]byte) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	id := sr.GetBytes(2)
	if string(id) != npcHeaderID {
		return errors.New("unexpected header ID")
	}

	n.Data = sr.GetBytes(49)

	return nil
}

func (n *NPC) Encode() (result [numNPCBytes]byte) {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes([]byte(npcHeaderID)...)
	sw.PushBytes(n.Data...)

	data := sw.GetBytes()
	copy(result[:], data[:numNPCBytes])

	return result
}
