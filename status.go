package d2d2s

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
