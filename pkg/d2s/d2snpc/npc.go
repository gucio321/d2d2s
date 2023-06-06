package d2snpc

import (
	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
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
	sr := datareader.NewReader(data[:])

	id := sr.GetBytes(len(npcHeaderID))
	if string(id) != npcHeaderID {
		return common.ErrUnexpectedHeader
	}

	n.Data = sr.GetBytes(49) //nolint:gomnd // TODO: parse to something human-readable

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
