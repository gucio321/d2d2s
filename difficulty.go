package d2d2s

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	actBitsCount = 3
)

// Difficulty represents a status of all difficulty levels
type Difficulty map[d2enum.DifficultyType]*DifficultyLevelStatus

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
