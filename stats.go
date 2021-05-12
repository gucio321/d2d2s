package d2d2s

import (
	"errors"
	"fmt"

	"github.com/gucio321/d2d2s/datautils"
)

const (
	statsHeaderID = "gf"
	statsModifier = 256
)

// Stats represents character stats
type Stats struct {
	Strength,
	Energy,
	Dexterity,
	Vitality,
	UnusedStats,
	UnusedSkillPoints,
	CurrentHP,
	MaxHP,
	CurrentMana,
	MaxMana,
	CurrentStamina,
	MaxStamina,
	Level,
	Experience,
	Gold,
	StashedGold uint64
}

// Load loads hero stats
// nolint:gocyclo // can't reduce (switch-case)
func (s *Stats) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != statsHeaderID {
		return errors.New("unexpected header")
	}

	bm := sr.Copy()

	for {
		i := bm.GetBits(9) // nolint:gomnd // id size
		id := uint64(i)

		// nolint:gomnd // If all 9 bits are set, we've hit the end of the attributes section
		//  at 0x1ff and exit the loop.
		if id == 0x1ff {
			break
		}

		// The attribute value bit length, so we'll know how many bits to read next.
		length, ok := attributeBitMap[id]
		if !ok {
			return fmt.Errorf("unknown attribute id: %d", id)
		}

		// The attribute value.
		// attr := reverseBits(uint64(bm.GetBits(int(length))), length)
		attr := uint64(bm.GetBits(int(length)))

		switch id {
		case strength:
			s.Strength = attr
		case energy:
			s.Energy = attr
		case dexterity:
			s.Dexterity = attr
		case vitality:
			s.Vitality = attr
		case unusedStats:
			s.UnusedStats = attr
		case unusedSkills:
			s.UnusedSkillPoints = attr
		case currentHP:
			s.CurrentHP = attr / statsModifier
		case maxHP:
			s.MaxHP = attr / statsModifier
		case currentMana:
			s.CurrentMana = attr / statsModifier
		case maxMana:
			s.MaxMana = attr / statsModifier
		case currentStamina:
			s.CurrentStamina = attr / statsModifier
		case maxStamina:
			s.MaxStamina = attr / statsModifier
		case level:
			s.Level = attr
		case experience:
			s.Experience = attr
		case gold:
			s.Gold = attr
		case stashedGold:
			s.StashedGold = attr
		}
	}

	// nolint:gomnd // need to equalize offset
	sr.SkipBits((8 * (bm.BitsRead() / 8)) + 8)

	return nil
}

// nolint:gochecknoglobals,gomnd // data variable
var attributeBitMap = map[uint64]uint{
	strength:       10,
	energy:         10,
	dexterity:      10,
	vitality:       10,
	unusedStats:    10,
	unusedSkills:   8,
	currentHP:      21,
	maxHP:          21,
	currentMana:    21,
	maxMana:        21,
	currentStamina: 21,
	maxStamina:     21,
	level:          7,
	experience:     32,
	gold:           25,
	stashedGold:    25,
}

// All attribute ids described.
const (
	strength       = 0
	energy         = 1
	dexterity      = 2
	vitality       = 3
	unusedStats    = 4
	unusedSkills   = 5
	currentHP      = 6
	maxHP          = 7
	currentMana    = 8
	maxMana        = 9
	currentStamina = 10
	maxStamina     = 11
	level          = 12
	experience     = 13
	gold           = 14
	stashedGold    = 15
)
