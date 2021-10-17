package d2sstatus

import (
	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
)

// Status represents character status
type Status struct {
	Unknown0,
	Unknown1,
	Hardcore,
	Died, // it is true, if you hav died in some point in past
	Unknown4,
	Expansion,
	Ladder, // uncertain
	Unknown7 bool
}

// New creates a new status structure
func New() *Status {
	result := &Status{}
	return result
}

// Load loads data into status structure
func (s *Status) Load(data byte) {
	bm := datareader.NewReader([]byte{data})
	s.Unknown0 = bm.GetBit()
	s.Unknown1 = bm.GetBit()
	s.Hardcore = bm.GetBit()
	s.Died = bm.GetBit()
	s.Unknown4 = bm.GetBit()
	s.Expansion = bm.GetBit()
	s.Ladder = bm.GetBit()
	s.Unknown7 = bm.GetBit()
}

// Encode encodes status back into a byte data
func (s *Status) Encode() (result byte) {
	sw := datautils.CreateStreamWriter()
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
