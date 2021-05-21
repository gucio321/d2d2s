package d2sstats

import (
	"errors"
	"fmt"

	"github.com/gucio321/d2d2s/datautils"
)

const (
	numHeaderBytes = 2
	statsHeaderID  = "gf"
	statsModifier  = 256
	statIDLen      = 9
	statEndMark    = 1<<statIDLen - 1 // all 9 bits set
)

// Stats represents character stats
type Stats struct {
	Strength,
	Energy,
	Dexterity,
	Vitality,
	UnusedStats,
	UnusedSkillPoints uint32
	CurrentHP,
	MaxHP,
	CurrentMana,
	MaxMana,
	CurrentStamina,
	MaxStamina float32
	Level,
	Experience,
	Gold,
	StashedGold uint32
}

// New creates a new stats list
func New() *Stats {
	result := &Stats{}
	return result
}

// Load loads hero stats
func (s *Stats) Load(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(numHeaderBytes)
	if string(id) != statsHeaderID {
		return errors.New("unexpected header")
	}

	bm := sr.Copy()

	for {
		id := statID(bm.GetBits(statIDLen))

		// If all 9 bits are set, we've hit the end of the attributes section
		//  at 0x1ff and exit the loop.
		if id == statEndMark {
			break
		}

		// The attribute value bit length, so we'll know how many bits to read next.
		length, err := id.getStatLen()
		if err != nil {
			return fmt.Errorf("error reading stat id: %w", err)
		}

		// The attribute value.
		attr := bm.GetBits(length)

		// this map connects statID with appropriate value from stats structure
		statMap := map[statID]interface{}{
			strength:       &s.Strength,
			energy:         &s.Energy,
			dexterity:      &s.Dexterity,
			vitality:       &s.Vitality,
			unusedStats:    &s.UnusedStats,
			unusedSkills:   &s.UnusedSkillPoints,
			currentHP:      &s.CurrentHP,
			maxHP:          &s.MaxHP,
			currentMana:    &s.CurrentMana,
			maxMana:        &s.MaxMana,
			currentStamina: &s.CurrentStamina,
			maxStamina:     &s.MaxStamina,
			level:          &s.Level,
			experience:     &s.Experience,
			gold:           &s.Gold,
			stashedGold:    &s.StashedGold,
		}

		// if id is HP/mana/stamina, make it float, by default uint32
		switch id {
		case currentHP, maxHP, currentMana, maxMana, currentStamina, maxStamina:
			value := statMap[id].(*float32)
			*value = float32(attr) / statsModifier
		default:
			value := statMap[id].(*uint32)
			*value = attr
		}
	}

	// nolint:gomnd // need to equalize offset
	sr.SkipBits((8 * (bm.BitsRead() / 8)) + 8)

	return nil
}

// Encode encodes stats back into a bytes slice
func (s *Stats) Encode() ([]byte, error) {
	sw := datautils.CreateStreamWriter()
	sw.PushBytes([]byte(statsHeaderID)...)

	statMap := map[statID]uint32{
		strength:       s.Strength,
		energy:         s.Energy,
		dexterity:      s.Dexterity,
		vitality:       s.Vitality,
		unusedStats:    s.UnusedStats,
		unusedSkills:   s.UnusedSkillPoints,
		currentHP:      uint32(s.CurrentHP * statsModifier),
		maxHP:          uint32(s.MaxHP * statsModifier),
		currentMana:    uint32(s.CurrentMana * statsModifier),
		maxMana:        uint32(s.MaxMana * statsModifier),
		currentStamina: uint32(s.CurrentStamina * statsModifier),
		maxStamina:     uint32(s.MaxStamina * statsModifier),
		level:          s.Level,
		experience:     s.Experience,
		gold:           s.Gold,
		stashedGold:    s.StashedGold,
	}

	for i := strength; i <= stashedGold; i++ {
		// if unused stats or skill points are 0 - than don't push dhem
		switch i {
		case unusedStats, unusedSkills:
			if statMap[i] == 0 {
				continue
			}
		}

		l, err := i.getStatLen()
		if err != nil {
			return nil, err
		}

		sw.PushBits16(uint16(i), statIDLen)
		sw.PushBits32(statMap[i], l)
	}

	sw.PushBits16(statEndMark, statIDLen)
	sw.Align()

	return sw.GetBytes(), nil
}

type statID uint16

// All attribute ids described.
const (
	strength statID = iota
	energy
	dexterity
	vitality
	unusedStats
	unusedSkills
	currentHP
	maxHP
	currentMana
	maxMana
	currentStamina
	maxStamina
	level
	experience
	gold
	stashedGold
)

func (i statID) getStatLen() (int, error) {
	// nolint:gomnd // constant values
	attributeBitMap := map[statID]int{
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

	s, ok := attributeBitMap[i]
	if !ok {
		return 0, errors.New("unexpected attribute ID")
	}

	return s, nil
}
