package d2sstats

import (
	"errors"
	"fmt"

	"github.com/gucio321/d2d2s/internal/datareader"
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
	UnusedSkillPoints uint64
	CurrentHP,
	MaxHP,
	CurrentMana,
	MaxMana,
	CurrentStamina,
	MaxStamina float64
	Level,
	Experience,
	Gold,
	StashedGold uint64

	// user can also specify self-definied stat ids
	ExtraStats map[StatID]uint64
	// this map should contain pairs of StatID:bit length of stat value
	userStatIdMap map[StatID]int
}

// New creates a new stats list
func New() *Stats {
	result := &Stats{
		ExtraStats: make(map[StatID]uint64),
	}
	return result
}

// UserStatMap allows to specify user-definied map of stat ids
// These ids will be saved inside of ExtraStats map
// values in m should be of types:
// key - StatID
// value - int (how many bits does value in d2s file use to be saved)
// NOTE: if existing value was added (e.g. user want to modify len of stat value)
// this stat will be still saved in corseponding field in Stats structure (not ExtraStats).
func (s *Stats) UserStatMap(m map[StatID]int) *Stats {
	s.userStatIdMap = m
	return s
}

// Load loads hero stats
func (s *Stats) Load(sr *datareader.Reader) error {
	id := sr.GetBytes(numHeaderBytes)
	if string(id) != statsHeaderID {
		return common.ErrUnexpectedHeader
	}

	for {
		id := StatID(sr.GetBits16(statIDLen))

		// If all 9 bits are set, we've hit the end of the attributes section
		//  at 0x1ff and exit the loop.
		if id == statEndMark {
			break
		}

		// The attribute value bit length, so we'll know how many bits to read next.
		length, err := id.GetStatLen(s.userStatIdMap)
		if err != nil {
			return fmt.Errorf("error reading stat id: %w", err)
		}

		// The attribute value.
		attr := sr.GetBits64(length)

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
			value := statMap[id].(*float64)
			*value = float64(attr) / statsModifier
		case Strength, Energy, Dexterity, Vitality, UnusedStats,
			UnusedSkills, Level, Experience, Gold, StashedGold:
			value := statMap[id].(*uint64)
			*value = attr
		default: // check extra stats map
			if _, exists := s.userStatIdMap[id]; exists {
				s.ExtraStats[id] = attr
			}
		}
	}

	sr.Align()

	return nil
}

// Encode encodes stats back into a bytes slice
func (s *Stats) Encode() ([]byte, error) {
	sw := datautils.CreateStreamWriter()
	sw.PushBytes([]byte(statsHeaderID)...)

	statMap := map[StatID]uint64{
		Strength:       s.Strength,
		Energy:         s.Energy,
		Dexterity:      s.Dexterity,
		Vitality:       s.Vitality,
		UnusedStats:    s.UnusedStats,
		UnusedSkills:   s.UnusedSkillPoints,
		CurrentHP:      uint64(s.CurrentHP * statsModifier),
		MaxHP:          uint64(s.MaxHP * statsModifier),
		CurrentMana:    uint64(s.CurrentMana * statsModifier),
		MaxMana:        uint64(s.MaxMana * statsModifier),
		CurrentStamina: uint64(s.CurrentStamina * statsModifier),
		MaxStamina:     uint64(s.MaxStamina * statsModifier),
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

		l, err := i.GetStatLen(s.userStatIdMap)
		if err != nil {
			return nil, err
		}

		if statMap[i] == 0 {
			continue
		}

		sw.PushBits16(uint16(i), statIDLen)
		sw.PushBits64(statMap[i], l)
	}

	if s.ExtraStats != nil {
		for key, value := range s.ExtraStats {
			l, err := key.GetStatLen(s.userStatIdMap)
			if err != nil {
				return nil, errors.New("Adding user-definied stat: custom stat value len wasn't found. Did you forgot to call UserStatMap()?")
			}

			sw.PushBits16(uint16(key), statIDLen)
			sw.PushBits64(value, l)
		}
	}

	sw.PushBits16(statEndMark, statIDLen)
	sw.Align()

	return sw.GetBytes(), nil
}

// it may be easier to debug, if each stat id have its own human-readable name
//go:generate stringer -type StatID

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
func (i StatID) GetStatLen(extra map[StatID]int) (int, error) {
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

	if s, ok := extra[i]; ok {
		return s, nil
	} else if s, ok := attributeBitMap[i]; ok {
		return s, nil
	}

	return 0, ErrIncorrectStatID
}
