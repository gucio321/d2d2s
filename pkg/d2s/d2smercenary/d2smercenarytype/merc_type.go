package d2smercenarytype

import (
	"fmt"

	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

const (
	numDifficulties = 3
)

// MercenaryType represents a mercenary type data
type MercenaryType struct {
	Class      MercClass
	Attribute  MercAttribute
	Difficulty d2senums.DifficultyType
}

// Load loads mercenary type data into a structure
func Load(id uint16) *MercenaryType {
	result := &MercenaryType{}

	var attribute int

search:
	for class := MercRogue; class <= MercBarbarian; class++ {
		n := class.numAttributes()
		for d := d2senums.DifficultyNormal; d <= d2senums.DifficultyHell; d++ {
			for a := 0; a < n; a++ {
				if id == 0 {
					result.Class = class
					result.Difficulty = d

					result.Attribute = MercAttribute(attribute + a)

					break search
				}
				id--
			}
		}
		attribute += n
	}

	return result
}

// Encode encodes merc type data into an uint16
func (m *MercenaryType) Encode() (result uint16) {
	attr := 0

	for class := MercRogue; class < m.Class; class++ {
		n := class.numAttributes()
		attr += n
		result += uint16(n * numDifficulties)
	}

	for d := d2senums.DifficultyNormal; d < m.Difficulty; d++ {
		result += uint16(m.Class.numAttributes())
	}

	result += uint16(int(m.Attribute) - attr)

	return result
}

//go:generate stringer -linecomment -type MercClass -output mercenary_class_string.go

// MercClass represents a merc class (e.g. Rogue Scout or Barbarian)
type MercClass byte

// mercenary classes
const (
	MercRogue     MercClass = iota // Rogue Scout
	MercDesert                     // Desert Mercenary
	MercSorcerer                   // Eastern Sorcerer
	MercBarbarian                  // Barbarian
)

func (m MercClass) numAttributes() int {
	// nolint:gomnd // data function
	lookup := map[MercClass]int{
		MercRogue:     2,
		MercDesert:    3,
		MercSorcerer:  3,
		MercBarbarian: 2,
	}

	result, ok := lookup[m]
	if !ok {
		panic("d2d2s: d2smercenarytype.(MercType).numAttributes() (internal): unexpected input value")
	}

	return result
}

//go:generate stringer -linecomment -type MercAttribute -output mercenary_attribute_string.go

// MercAttribute represents a mercenary attribute (like fire arrows or defensive auras)
type MercAttribute byte

// merc attributes
const (
	MercAttributeFireArrow MercAttribute = iota // Fire Arrow
	MercAttributeColdArrow                      // Cold Arrow

	MercAttributeCombat    // Combat
	MercAttributeDefensive // Devensive
	MercAttributeOffensive // Offensive

	MercAttributeFire      // Fire Spells
	MercAttributeCold      // Cold Spells
	MercAttributeLightning // Lightning Spells

	MercAttributeBarb1 //
	MercAttributeBarb2 //
)

// GetClassAttributes returns a slice of attributes for specified merc class
func GetClassAttributes(class MercClass) []MercAttribute {
	attrIdx := [...]MercAttribute{MercAttributeFireArrow, MercAttributeCombat, MercAttributeFire, MercAttributeBarb1, 10}

	if int(class) > len(attrIdx) {
		return nil
	}

	result := make([]MercAttribute, 0)

	for i := attrIdx[int(class)]; i < attrIdx[int(class)+1]; i++ {
		result = append(result, i)
	}

	return result
}

func (m *MercenaryType) String() string {
	return fmt.Sprintf("%s - %s (%s)", m.Class, m.Attribute, m.Difficulty)
}
