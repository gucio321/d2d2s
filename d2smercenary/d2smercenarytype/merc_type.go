package d2smercenarytype

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	numDifficulties   = 3
	numMercSkillsAct1 = 2
	numMercSkillsAct2 = 2
	numMercSkillsAct3 = 2
	numMercSkillsAct5 = 2
)

type MercenaryType struct {
	Class      MercClass
	Attribute  MercAttribute
	Difficulty d2enum.DifficultyType
}

func Load(id uint16) *MercenaryType {
	result := &MercenaryType{}

	numAttributes := map[MercClass]int{
		MercRogue:     2,
		MercDesert:    3,
		MercSorceror:  3,
		MercBarbarian: 2,
	}

	var attribute int = 0

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

func (m *MercenaryType) Encode() (result uint16) {
	numAttributes := map[MercClass]int{
		MercRogue:     2,
		MercDesert:    3,
		MercSorceror:  3,
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
	fmt.Println(result)

	return result
}

//go:generate stringer -linecomment -type MercClass -output mercenary_class_string.go

type MercClass byte

const (
	MercRogue     MercClass = iota // Rogue Scout
	MercDesert                     // Desert Mercenary
	MercSorceror                   // Eastern Sorceror
	MercBarbarian                  // Barbarian
)

//go:generate stringer -linecomment -type MercAttribute -output mercenary_attribute_string.go

type MercAttribute byte

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
