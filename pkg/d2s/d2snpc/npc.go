package d2snpc

import (
	"errors"

	"github.com/gucio321/d2d2s/internal/datautils"
)

const (
	// NumNPCBytes is a number of bytes for NPC data
	NumNPCBytes = 51

	npcHeaderID = "w4"
)

// New creates a new npc struct
func New() *NPC {
	result := &NPC{}

	return result
}

// NPC represents npc introduction data (TODO)
type NPC struct {
	Data []byte
}

// Load loads NPC data into NPC structure
func (n *NPC) Load(data [NumNPCBytes]byte) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != npcHeaderID {
		return errors.New("unexpected header ID")
	}

	n.Data = sr.GetBytes(49) // nolint:gomnd // TODO: parse to something human-readable

	return nil
}

// Encode encodes NPC data back into byte array
func (n *NPC) Encode() (result [NumNPCBytes]byte) {
	sw := datautils.CreateStreamWriter()

	sw.PushBytes([]byte(npcHeaderID)...)
	sw.PushBytes(n.Data...)

	data := sw.GetBytes()
	copy(result[:], data[:NumNPCBytes])

	return result
}
