package d2d2s

import (
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
	unknown5 bool
	active bool
}

func (d *DifficultyLevelStatus) Unmarshal(data byte) {
	d.act = data & actMask
	d.unknown3 = ((data >> 3) & 1) > 0
	d.unknown4 = ((data >> 4) & 1) > 0
	d.unknown5 = ((data >> 5) & 1) > 0
	d.active = ((data >> 6) & 1) > 0
}
