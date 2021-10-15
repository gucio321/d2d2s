package d2sskills

import (
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

const (
	// NumSkillBytes is a total number of encoded skills structure
	NumSkillBytes = d2senums.NumSkills + 2

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

	skillsID := sr.GetBytes(len(skillsHeaderID))
	if string(skillsID) != skillsHeaderID {
		return common.ErrUnexpectedHeader
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
	sw := datautils.CreateStreamWriter()

	sw.PushBytes([]byte(skillsHeaderID)...)

	skills := d2senums.GetSkillList(class)

	for _, skill := range skills {
		sw.PushBytes((*s)[skill])
	}

	data := sw.GetBytes()

	copy(result[:], data[:NumSkillBytes])

	return result
}
