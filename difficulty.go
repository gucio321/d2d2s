package d2d2s

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	actBitsCount = 3
	actMask      = 1<<actBitsCount - 1
)

type Difficulty map[d2enum.DifficultyType]*DifficultyLevelStatus

type DifficultyLevelStatus struct {
	act byte
	unknown3,
	unknown4,
	unknown5,
	unknown6 bool
	active bool
}

func (d *DifficultyLevelStatus) Unmarshal(data byte) {
	bm := d2datautils.CreateBitMuncher([]byte{data}, 0)
	d.act = byte(bm.GetBits(3))
	d.unknown3 = bm.GetBit() == 1
	d.unknown4 = bm.GetBit() == 1
	d.unknown5 = bm.GetBit() == 1
	d.unknown6 = bm.GetBit() == 1
	d.active = bm.GetBit() == 1
}

func (d *DifficultyLevelStatus) Encode() (result byte) {
	sw := d2datautils.CreateStreamWriter()
	sw.PushBits(d.act, 3)
	sw.PushBit(d.unknown3)
	sw.PushBit(d.unknown4)
	sw.PushBit(d.unknown5)
	sw.PushBit(d.unknown6)
	sw.PushBit(d.active)

	return sw.GetBytes()[0]
}
