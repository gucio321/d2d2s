package d2sskills

import (
	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

const (
	skillsHeaderID = "if"
)

// Skills stores skillID:level values
type Skills struct {
	skillsCount int
	SkillLevels map[d2senums.SkillID]byte
}

// New creates a new skills structure
func New() *Skills {
	result := &Skills{
		skillsCount: d2senums.NumSkills,
		SkillLevels: make(map[d2senums.SkillID]byte),
	}

	return result
}

// SetSkillsCount could be used to parse custom d2s files. It sets
// custom number of character skills.
func (s *Skills) SetSkillsCount(count int) *Skills {
	s.skillsCount = count
	return s
}

// Load loads skills from bitmuncher given
func (s *Skills) Load(sr *datareader.Reader, class d2senums.CharacterClass) error {
	skillsID := sr.GetBytes(len(skillsHeaderID))
	if string(skillsID) != skillsHeaderID {
		return common.ErrUnexpectedHeader
	}

	mod := d2senums.GetSkillModifier(class)

	for skill := 0; skill < s.skillsCount; skill++ {
		skill := skill + int(mod)
		value := sr.GetByte()
		s.SkillLevels[d2senums.SkillID(skill)] = value
	}

	return nil
}

// Encode encodes skills data into stream writer
func (s *Skills) Encode(sw *datautils.StreamWriter, class d2senums.CharacterClass) {
	sw.PushBytes([]byte(skillsHeaderID)...)

	mod := d2senums.GetSkillModifier(class)

	for skill := 0; skill < s.skillsCount; skill++ {
		skill := skill + int(mod)
		sw.PushBytes(s.SkillLevels[d2senums.SkillID(skill)])
	}
}
