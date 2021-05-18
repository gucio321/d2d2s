package d2sskills

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/datautils"
)

const numSkills = 30

// Skills stores skillID:level values
type Skills map[d2senums.SkillID]byte

// New creates a new skills structure
func New() *Skills {
	result := &Skills{}
	*result = make(Skills)

	return result
}

// Load loads skills from bitmuncher given
func (s *Skills) Load(sr *datautils.BitMuncher, class d2senums.CharacterClass) {
	for i := 0; i < numSkills; i++ {
		value := sr.GetByte()
		(*s)[d2senums.SkillID(int(d2senums.GetSkillModifier(class))+i)] = value
	}
}

// Encode encodes skills data into stream writer
func (s *Skills) Encode(sw *d2datautils.StreamWriter, class d2senums.CharacterClass) {
	for i := 0; i < numSkills; i++ {
		sw.PushBytes((*s)[d2senums.SkillID(int(d2senums.GetSkillModifier(class))+i)])
	}
}
