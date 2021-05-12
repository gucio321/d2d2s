package d2d2s

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/datautils"
)

// Status represents character status
type Status struct {
	Unknown0,
	Unknown1,
	Hardcore,
	Died, // it is true, if you hav died in some point in past
	Unknown4,
	Expansion,
	Ladder,
	Unknown7 bool
}

// Unmarshal loads data into status structure
func (s *Status) Unmarshal(data byte) {
	bm := datautils.CreateBitMuncher([]byte{data}, 0)
	s.Unknown0 = bm.GetBit() == 1
	s.Unknown1 = bm.GetBit() == 1
	s.Hardcore = bm.GetBit() == 1
	s.Died = bm.GetBit() == 1
	s.Unknown4 = bm.GetBit() == 1
	s.Expansion = bm.GetBit() == 1
	s.Ladder = bm.GetBit() == 1
	s.Unknown7 = bm.GetBit() == 1
}

// Encode encodes status back into a byte data
func (s *Status) Encode() (result byte) {
	sw := d2datautils.CreateStreamWriter()
	sw.PushBit(s.Unknown0)
	sw.PushBit(s.Unknown1)
	sw.PushBit(s.Hardcore)
	sw.PushBit(s.Died)
	sw.PushBit(s.Unknown4)
	sw.PushBit(s.Expansion)
	sw.PushBit(s.Ladder)
	sw.PushBit(s.Unknown7)

	return sw.GetBytes()[0]
}
