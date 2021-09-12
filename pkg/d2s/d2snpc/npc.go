package d2snpc

import (
	"errors"

	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

const (
	// NumNPCBytes is a number of bytes for NPC data
	NumNPCBytes = 51

	npcHeaderID = "w4"
)

// New creates a new npc struct
func New() *NPCs {
	result := &NPCs{
		Introductions: make(map[d2senums.DifficultyType]*Introduction),
	}

	for d := d2senums.DifficultyNormal; d <= d2senums.DifficultyHell; d++ {
		result.Introductions[d] = newIntroduction()
	}

	return result
}

// NPCs represents npc introduction data (TODO)
type NPCs struct {
	Data          []byte
	unknown       byte
	Introductions map[d2senums.DifficultyType]*Introduction
}

// Load loads NPCs data into NPCs structure
func (n *NPCs) Load(data [NumNPCBytes]byte) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	id := sr.GetBytes(len(npcHeaderID))
	if string(id) != npcHeaderID {
		return errors.New("unexpected header ID")
	}

	n.unknown = sr.GetByte()

	for d := d2senums.DifficultyNormal; d <= d2senums.DifficultyHell; d++ {
		data := sr.GetBytes(8)
		var idata [8]byte
		copy(idata[:], data)
		n.Introductions[d].Decode(idata)
	}
	n.Data = sr.GetBytes(24) // nolint:gomnd // TODO: parse to something human-readable

	return nil
}

// Encode encodes NPCs data back into byte array
func (n *NPCs) Encode() (result [NumNPCBytes]byte) {
	sw := datautils.CreateStreamWriter()

	sw.PushBytes([]byte(npcHeaderID)...)
	sw.PushBytes(n.Data...)

	data := sw.GetBytes()
	copy(result[:], data[:NumNPCBytes])

	return result
}

type NPC struct {
	Act  int
	Name d2senums.NPCID
}
