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

// NPC represents npc introduction data (TODO)
type NPC struct {
	Data []byte
}

// Load loads NPC data into NPC structure
func (n *NPC) Load(data [numNPCBytes]byte) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != npcHeaderID {
		return errors.New("unexpected header ID")
	}

	n.Data = sr.GetBytes(49) // nolint:gomnd // TODO: parse to something human-readable

	return nil
}

// Encode encodes NPC data back into byte array
func (n *NPC) Encode() (result [numNPCBytes]byte) {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes([]byte(npcHeaderID)...)
	sw.PushBytes(n.Data...)

	data := sw.GetBytes()
	copy(result[:], data[:numNPCBytes])

	return result
}
