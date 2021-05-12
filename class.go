package d2d2s

// CharacterClass represents class of a character
type CharacterClass byte

// character classes
const (
	CharacterClassAmazon CharacterClass = iota
	CharacterClassSorceress
	CharacterClassNecromancer
	CharacterClassPaladin
	CharacterClassBarbarian
	CharacterClassDruid
	CharacterClassAssassin
)

func (c CharacterClass) String() string {
	lookup := map[CharacterClass]string{
		CharacterClassAmazon:      "Amazon",
		CharacterClassSorceress:   "Sorceress",
		CharacterClassNecromancer: "Necromancer",
		CharacterClassPaladin:     "Paladin",
		CharacterClassBarbarian:   "Barbarian",
		CharacterClassDruid:       "Druid",
		CharacterClassAssassin:    "Assassin",
	}

	s, ok := lookup[c]
	if !ok {
		return "Unknown character"
	}

	return s
}
