package d2snpc

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
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
	result := &NPCs{}

	return result
}

// NPCs represents npc introduction data (TODO)
type NPCs struct {
	Data          []byte
	Introductions map[d2enum.DifficultyType]*Introduction
}

// Load loads NPCs data into NPCs structure
func (n *NPCs) Load(data [NumNPCBytes]byte) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	id := sr.GetBytes(len(npcHeaderID))
	if string(id) != npcHeaderID {
		return errors.New("unexpected header ID")
	}

	n.Data = sr.GetBytes(49) // nolint:gomnd // TODO: parse to something human-readable

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

type Introduction struct {
	Unknown1 byte
	NPCs     []NPC
}

type NPC struct {
	Act  int
	Name d2senums.NPCID
}
