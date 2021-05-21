package d2smercenarytype

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	numDifficulties = 3
)

// MercenaryType represents a mercenary type data
type MercenaryType struct {
	Class      MercClass
	Attribute  MercAttribute
	Difficulty d2enum.DifficultyType
}

// Load loads mercenary type data into a structure
func Load(id uint16) *MercenaryType {
	result := &MercenaryType{}

	// nolint:gomnd // constant values
	numAttributes := map[MercClass]int{
		MercRogue:     2,
		MercDesert:    3,
		MercSorcerer:  3,
		MercBarbarian: 2,
	}

	var attribute int

search:
	for class := MercRogue; class <= MercBarbarian; class++ {
		n := numAttributes[class]
		for d := d2enum.DifficultyNormal; d <= d2enum.DifficultyHell; d++ {
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
	// nolint:gomnd // constant values
	numAttributes := map[MercClass]int{
		MercRogue:     2,
		MercDesert:    3,
		MercSorcerer:  3,
		MercBarbarian: 2,
	}

	attr := 0

	for class := MercRogue; class < m.Class; class++ {
		attr += numAttributes[class]
		result += uint16(numAttributes[class] * numDifficulties)
	}

	for d := d2enum.DifficultyNormal; d < m.Difficulty; d++ {
		result += uint16(numAttributes[m.Class])
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

//go:generate stringer -linecomment -type MercAttribute -output mercenary_attribute_string.go

// MercAttribute represents a mercenary attribute (like fire arrows or defensive auras)
type MercAttribute byte

// merc attributes
const (
	MercAttributeFireArrow MercAttribute = iota // Fire Arrow
	MercAttributeColdArrow                      // ColdArrow

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
	attrIdx := [...]MercAttribute{MercAttributeFireArrow, MercAttributeCombat, MercAttributeFire, MercAttributeBarb1, 12}

	if int(class) > len(attrIdx) {
		return nil
	}

	result := make([]MercAttribute, 0)

	for i := attrIdx[int(class)]; i < attrIdx[int(class)+1]; i++ {
		result = append(result, i)
	}

	return result
}
