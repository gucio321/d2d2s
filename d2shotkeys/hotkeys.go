package d2shotkeys

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/d2senums"
)

const (
	// NumHotkeysBytes is a len of hotkey's binary form
	NumHotkeysBytes = 64

	numHotKeys = 16
)

// Hotkeys represents a skill hotkeys data
type Hotkeys map[byte]d2senums.SkillID

// New creates a new Hotkeys
func New() *Hotkeys {
	result := &Hotkeys{}
	*result = make(Hotkeys)

	return result
}

// Load loads hotkeys
func (h *Hotkeys) Load(data [NumHotkeysBytes]byte) {
	sr := d2datautils.CreateBitMuncher(data[:], 0)
	for i := byte(0); i < numHotKeys; i++ {
		id := sr.GetUInt32()
		(*h)[i] = d2senums.SkillID(id)
	}
}

// Encode encodes hotkeys back into binary form
func (h *Hotkeys) Encode() (result [NumHotkeysBytes]byte) {
	sw := d2datautils.CreateStreamWriter()

	for i := byte(0); i < numHotKeys; i++ {
		sw.PushUint32(uint32((*h)[i]))
	}

	data := sw.GetBytes()
	copy(result[:], data[:NumHotkeysBytes])

	return result
}
