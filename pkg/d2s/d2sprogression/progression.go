package d2sprogression

import "github.com/gucio321/d2d2s/pkg/d2s/d2senums"

// New creates a new Progression structure
func New() *Progression {
	result := &Progression{
		DifficultyLevel: d2senums.DifficultyNormal,
		Act:             1,
	}

	return result
}

// Progression represents a character progress in the game
// in fact, the progression data byte is increased every time
// when you cplete an act. In Diablo II: Lord of Destruction
// it doesn't get increased, when you kill Diablo, but is
// increased by 2 when you kill Baal
/*
from: https://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml
Character progression.  This number tells (sort of) how many acts you have completed from all
difficulty levels.  It appears to be incremented when you kill the final demon in an act -- i.e.,
Andarial, Duriel, Mephisto, and Diablo / Baal.  There's a catch to that last one: in an
Expansion game, the value is not incremented after killing Diablo, but is incremented by 2 after killing Baal.
(The reason is unknown.)  So it skips the values 4, 9, and 14.
*/
type Progression struct {
	DifficultyLevel d2senums.DifficultyType
	Act             int
}

// Load loads progression
func (p *Progression) Load(data byte) {
	d := 0

	for data > d2senums.NumActs {
		data -= d2senums.NumActs
		d++
	}

	p.DifficultyLevel = d2senums.DifficultyType(d)

	p.Act = int(data)
}

// Encode encodes progression int a binary form
func (p *Progression) Encode() (data byte) {
	data += byte(p.DifficultyLevel * d2senums.NumActs)
	data += byte(p.Act)

	return data
}
