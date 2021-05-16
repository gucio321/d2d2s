package enums

//go:generate stringer -type CharacterClass -trimprefix CharacterClass -linecomment -output class_string.go

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
