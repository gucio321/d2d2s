package d2d2s

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/gucio321/d2d2s/datautils"
)

const (
	actBitsCount = 3
)

// NewDifficulty creates a new Difficulty
func NewDifficulty() *Difficulty {
	result := &Difficulty{}
	*result = make(Difficulty)

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		(*result)[i] = &DifficultyLevelStatus{}
	}

	return result
}

// Difficulty represents a status of all difficulty levels
type Difficulty map[d2enum.DifficultyType]*DifficultyLevelStatus

// Load loads difficulty status
func (d *Difficulty) Load(sr *datautils.BitMuncher) {
	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		data := sr.GetByte()

		(*d)[i].Unmarshal(data)
	}
}

// Encode encodes difficulty status back into a byte slice
func (d *Difficulty) Encode(sw *d2datautils.StreamWriter) {
	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		sw.PushBytes((*d)[i].Encode())
	}
}

// DifficultyLevelStatus represents a status of difficulty level
type DifficultyLevelStatus struct {
	Act byte
	unknown3,
	unknown4,
	unknown5,
	unknown6 bool
	Active bool
}

// Unmarshal loads byte into DifficultyLevelStatus structure
func (d *DifficultyLevelStatus) Unmarshal(data byte) {
	bm := d2datautils.CreateBitMuncher([]byte{data}, 0)
	d.Act = byte(bm.GetBits(actBitsCount))
	d.unknown3 = bm.GetBit() == 1
	d.unknown4 = bm.GetBit() == 1
	d.unknown5 = bm.GetBit() == 1
	d.unknown6 = bm.GetBit() == 1
	d.Active = bm.GetBit() == 1
}

// Encode encodes a difficulty level status back into byte data
func (d *DifficultyLevelStatus) Encode() (result byte) {
	sw := d2datautils.CreateStreamWriter()
	sw.PushBits(d.Act, 3)
	sw.PushBit(d.unknown3)
	sw.PushBit(d.unknown4)
	sw.PushBit(d.unknown5)
	sw.PushBit(d.unknown6)
	sw.PushBit(d.Active)

	return sw.GetBytes()[0]
}
