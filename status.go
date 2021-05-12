package d2d2s

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

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
	s.Unknown0 = ((data >> 0) & 1) > 0
	s.Unknown1 = ((data >> 1) & 1) > 0
	s.Hardcore = ((data >> 2) & 1) > 0
	s.Died = ((data >> 3) & 1) > 0
	s.Unknown4 = ((data >> 4) & 1) > 0
	s.Expansion = ((data >> 5) & 1) > 0
	s.Ladder = ((data >> 6) & 1) > 0
	s.Unknown7 = ((data >> 7) & 1) > 0
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
