package d2sstats

import (
	"errors"
	"fmt"

	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
)

// ErrIncorrectStatID is returned by GetStatLen when stat id is unknown
var ErrIncorrectStatID = errors.New("incorrect stat identifier")

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
		return common.ErrUnexpectedHeader
	}

	bm := sr.Copy()

	for {
		id := StatID(bm.GetBits(statIDLen))

		// If all 9 bits are set, we've hit the end of the attributes section
		//  at 0x1ff and exit the loop.
		if id == statEndMark {
			break
		}

		// The attribute value bit length, so we'll know how many bits to read next.
		length, err := id.GetStatLen()
		if err != nil {
			return fmt.Errorf("error reading stat id: %w", err)
		}

		// The attribute value.
		attr := bm.GetBits(length)

		// this map connects StatID with appropriate value from stats structure
		statMap := map[StatID]interface{}{
			Strength:       &s.Strength,
			Energy:         &s.Energy,
			Dexterity:      &s.Dexterity,
			Vitality:       &s.Vitality,
			UnusedStats:    &s.UnusedStats,
			UnusedSkills:   &s.UnusedSkillPoints,
			CurrentHP:      &s.CurrentHP,
			MaxHP:          &s.MaxHP,
			CurrentMana:    &s.CurrentMana,
			MaxMana:        &s.MaxMana,
			CurrentStamina: &s.CurrentStamina,
			MaxStamina:     &s.MaxStamina,
			Level:          &s.Level,
			Experience:     &s.Experience,
			Gold:           &s.Gold,
			StashedGold:    &s.StashedGold,
		}

		// if id is HP/mana/stamina, make it float, by default uint32
		switch id {
		case CurrentHP, MaxHP, CurrentMana, MaxMana, CurrentStamina, MaxStamina:
			value := statMap[id].(*float32)
			*value = float32(attr) / statsModifier
		default:
			value := statMap[id].(*uint32)
			*value = attr
		}
	}

	// need to equalize offsets
	const byteSize = 8

	sr.SkipBits((byteSize * (bm.BitsRead() / byteSize)) + byteSize)

	return nil
}

// Encode encodes stats back into a bytes slice
func (s *Stats) Encode() ([]byte, error) {
	sw := datautils.CreateStreamWriter()
	sw.PushBytes([]byte(statsHeaderID)...)

	statMap := map[StatID]uint32{
		Strength:       s.Strength,
		Energy:         s.Energy,
		Dexterity:      s.Dexterity,
		Vitality:       s.Vitality,
		UnusedStats:    s.UnusedStats,
		UnusedSkills:   s.UnusedSkillPoints,
		CurrentHP:      uint32(s.CurrentHP * statsModifier),
		MaxHP:          uint32(s.MaxHP * statsModifier),
		CurrentMana:    uint32(s.CurrentMana * statsModifier),
		MaxMana:        uint32(s.MaxMana * statsModifier),
		CurrentStamina: uint32(s.CurrentStamina * statsModifier),
		MaxStamina:     uint32(s.MaxStamina * statsModifier),
		Level:          s.Level,
		Experience:     s.Experience,
		Gold:           s.Gold,
		StashedGold:    s.StashedGold,
	}

	for i := Strength; i <= StashedGold; i++ {
		// if unused stats or skill points are 0 - than don't push dhem
		switch i {
		case UnusedStats, UnusedSkills:
			if statMap[i] == 0 {
				continue
			}
		}

		l, err := i.GetStatLen()
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

// StatID represents a stat id
type StatID uint16

// All attribute ids described.
const (
	Strength StatID = iota
	Energy
	Dexterity
	Vitality
	UnusedStats
	UnusedSkills
	CurrentHP
	MaxHP
	CurrentMana
	MaxMana
	CurrentStamina
	MaxStamina
	Level
	Experience
	Gold
	StashedGold
)

// GetStatLen returns length of stat id data
func (i StatID) GetStatLen() (int, error) {
	// nolint:gomnd // data function
	attributeBitMap := map[StatID]int{
		Strength:       10,
		Energy:         10,
		Dexterity:      10,
		Vitality:       10,
		UnusedStats:    10,
		UnusedSkills:   8,
		CurrentHP:      21,
		MaxHP:          21,
		CurrentMana:    21,
		MaxMana:        21,
		CurrentStamina: 21,
		MaxStamina:     21,
		Level:          7,
		Experience:     32,
		Gold:           25,
		StashedGold:    25,
	}

	s, ok := attributeBitMap[i]
	if !ok {
		return 0, ErrIncorrectStatID
	}

	return s, nil
}
