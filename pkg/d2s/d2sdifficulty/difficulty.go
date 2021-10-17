package d2sdifficulty

import (
	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
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

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		(*result)[i] = &DifficultyLevelStatus{}
	}

	return result
}

// Difficulty represents a status of all difficulty levels
type Difficulty map[d2senums.DifficultyType]*DifficultyLevelStatus

// Load loads difficulty status
func (d *Difficulty) Load(data [NumDifficultyBytes]byte) {
	sr := datareader.NewReader(data[:])
	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		data := sr.GetByte()

		(*d)[i].Load(data)
	}
}

// Encode encodes difficulty status back into a byte slice
func (d *Difficulty) Encode() (result [NumDifficultyBytes]byte) {
	sw := datautils.CreateStreamWriter()
	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
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
	bm := datareader.NewReader([]byte{data})
	d.Act = byte(bm.GetBits(actBitsCount))
	d.unknown3 = bm.GetBit()
	d.unknown4 = bm.GetBit()
	d.unknown5 = bm.GetBit()
	d.unknown6 = bm.GetBit()
	d.Active = bm.GetBit()
}

// Encode encodes a difficulty level status back into byte data
func (d *DifficultyLevelStatus) Encode() (result byte) {
	sw := datautils.CreateStreamWriter()
	sw.PushBits(d.Act, actBitsCount)
	sw.PushBit(d.unknown3)
	sw.PushBit(d.unknown4)
	sw.PushBit(d.unknown5)
	sw.PushBit(d.unknown6)
	sw.PushBit(d.Active)

	return sw.GetBytes()[0]
}

// SetAct sets valid difficultyLevelStatus.Act value
func (d *DifficultyLevelStatus) SetAct(act int) {
	if act > 0 && act < d2senums.NumActs {
		d.Act = byte(act)
	}
}
