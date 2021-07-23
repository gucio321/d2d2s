package d2sdifficulty

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	// NumDifficultyBytes is a length of binary form of `Difficulty` structure
	NumDifficultyBytes = 3

	actBitsCount = 3
)

// New creates a new Difficulty
func New() *Difficulty {
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
func (d *Difficulty) Load(data [NumDifficultyBytes]byte) {
	sr := d2datautils.CreateBitMuncher(data[:], 0)
	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		data := sr.GetByte()

		(*d)[i].Load(data)
	}
}

// Encode encodes difficulty status back into a byte slice
func (d *Difficulty) Encode() (result [NumDifficultyBytes]byte) {
	sw := d2datautils.CreateStreamWriter()
	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		sw.PushBytes((*d)[i].Encode())
	}

	data := sw.GetBytes()

	copy(result[:], data[:NumDifficultyBytes])

	return result
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

// Load loads byte into DifficultyLevelStatus structure
func (d *DifficultyLevelStatus) Load(data byte) {
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
	sw.PushBits(d.Act, actBitsCount)
	sw.PushBit(d.unknown3)
	sw.PushBit(d.unknown4)
	sw.PushBit(d.unknown5)
	sw.PushBit(d.unknown6)
	sw.PushBit(d.Active)

	return sw.GetBytes()[0]
}
