package d2sskills

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	numSkills      = 30
	skillsHeaderID = "if"
)

// Skills stores skillID:level values
type Skills map[d2senums.SkillID]byte

// New creates a new skills structure
func New() *Skills {
	result := &Skills{}
	*result = make(Skills)

	return result
}

// Load loads skills from bitmuncher given
func (s *Skills) Load(sr *datautils.BitMuncher, class d2senums.CharacterClass) error {
	skillsID := sr.GetBytes(2) // nolint:gomnd // skills header
	if string(skillsID) != skillsHeaderID {
		return errors.New("unexpected skills section header")
	}

	for i := 0; i < numSkills; i++ {
		value := sr.GetByte()
		(*s)[d2senums.SkillID(int(d2senums.GetSkillModifier(class))+i)] = value
	}

	return nil
}

// Encode encodes skills data into stream writer
func (s *Skills) Encode(sw *d2datautils.StreamWriter, class d2senums.CharacterClass) {
	sw.PushBytes([]byte(skillsHeaderID)...)

	for i := 0; i < numSkills; i++ {
		sw.PushBytes((*s)[d2senums.SkillID(int(d2senums.GetSkillModifier(class))+i)])
	}
}
