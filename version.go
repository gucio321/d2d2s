package d2d2s

import "fmt"

type Version int32

/*
   71 is 1.00 through v1.06
   87 is 1.07 or Expansion Set v1.08
   89 is standard game v1.08
   92 is v1.09 (both the standard game and the Expansion Set.)
   96 is v1.10+
*/
func (v Version) String() string {
	lookup := map[Version]string{
		71: "Diablo II v1.00 - v1.06",
		87: "Diablo II v1.07 or Diablo II: Lord of Destruction v1.08",
		89: "Diablo II v1.08",
		92: "Diablo II or Diablo II: Lord of Destruction v.09",
		96: "Diablo II Expansion v.10 or later",
	}

	s, ok := lookup[v]
	if !ok {
		return fmt.Sprintf("Unknown value %d", v)
	}

	return s
}
