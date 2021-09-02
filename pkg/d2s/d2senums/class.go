package d2senums

//go:generate stringer -type CharacterClass -trimprefix Character -linecomment -output class_string.go

// CharacterClass represents class of a character
type CharacterClass byte

// character classes
const (
	CharacterAmazon CharacterClass = iota
	CharacterSorceress
	CharacterNecromancer
	CharacterPaladin
	CharacterBarbarian
	CharacterDruid
	CharacterAssassin
)
