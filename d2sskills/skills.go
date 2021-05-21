package d2sskills

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	NumSkillBytes  = d2senums.NumSkills + 2
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
func (s *Skills) Load(data [NumSkillBytes]byte, class d2senums.CharacterClass) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	skillsID := sr.GetBytes(2) // nolint:gomnd // skills header
	if string(skillsID) != skillsHeaderID {
		return errors.New("unexpected skills section header")
	}

	skills := d2senums.GetSkillList(class)

	for _, skill := range skills {
		value := sr.GetByte()
		(*s)[skill] = value
	}

	return nil
}

// Encode encodes skills data into stream writer
func (s *Skills) Encode(class d2senums.CharacterClass) (result [NumSkillBytes]byte) {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes([]byte(skillsHeaderID)...)

	skills := d2senums.GetSkillList(class)

	for _, skill := range skills {
		sw.PushBytes((*s)[skill])
	}

	data := sw.GetBytes()

	copy(result[:], data[:NumSkillBytes])

	return result
}
