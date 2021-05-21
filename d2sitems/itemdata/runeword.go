package itemdata

//go:generate stringer -linecomment -type RunewordID -output runeword_string.go

// RunewordID represents a runeword ID
type RunewordID uint16

// runeword IDs TODO: give these constants a bit more descriptive names
const (
	Runeword27 RunewordID = iota + 27 // Ancient's Pledge
	_
	_
	Runeword30 // Beast
	_
	Runeword32 // Black
	_
	Runeword34 // Bone
	Runeword35 // Bramble
	Runeword36 // Brand
	Runeword37 // Breath of the Dying
	_
	Runeword39 // Call to Arms
	Runeword40 // Chains of Honor
	_
	Runeword42 // Chaos
	Runeword43 // Crescent Moon
	_
	_
	Runeword46 // Death
	_
	_
	_
	Runeword51 // Destruction
	Runeword52 // Doom
	Runeword53 // Dragon
	_
	Runeword55 // Dream
	Runeword56 // Duress
	Runeword57 // Edge
	_
	Runeword59 // Enigma
	Runeword60 // Enlightenment
	_
	Runeword62 // Eternity
	Runeword63 // Exile
	Runeword64 // Faith
	Runeword65 // Famine
	_
	Runeword67 // Fortitude
	_
	_
	Runeword70 // Fury
	Runeword71 // Gloom
	_
	Runeword73 // Grief
	Runeword74 // Hand of Justice
	Runeword75 // Harmory
	_
	Runeword77 // Heart of the Oak
	_
	_
	Runeword80 // Holy Thunder
	Runeword81 // Honor
	_
	_
	_
	Runeword85 // Ice
	Runeword86 // Infinity
	_
	Runeword88 // Insight
	_
	_
	Runeword91 // King's Grace
	Runeword92 // Kingslayer
	_
	_
	Runeword95 // Last Wish
	_
	Runeword97 // Lawbringer
	Runeword98 // Leaf
	_
	Runeword100 // Lionheart
	Runeword101 // Lore
	_
	_
	_
	_
	Runeword106 // Malice
	Runeword107 // Melody
	Runeword108 // Memory
	_
	_
	_
	Runeword112 // Myth
	Runeword113 // Nadir
	_
	_
	Runeword116 // Oath
	Runeword117 // Obedience
	_
	_
	Runeword120 // Passion
	_
	_
	Runeword123 // Peace
	Runeword124 // Winter
	_
	_
	_
	Runeword128 // Phoenix
	_
	_
	Runeword131 // Plague
	_
	_
	Runeword134 // Pride
	Runeword135 // Principle
	_
	Runeword137 // Prudence
	_
	_
	_
	Runeword141 // Radiance
	Runeword142 // Rain
	_
	_
	Runeword145 // Rhyme
	Runeword146 // Rift
	Runeword147 // Sanctuary
	_
	_
	_
	Runeword151 // Silence
	_
	Runeword153 // Smoke
	_
	Runeword155 // Spirit
	Runeword156 // Splendor
	_
	Runeword158 // Stealth
	Runeword159 // Steel
	_
	_
	Runeword162 // Stone
	_
	Runeword164 // Strength
	_
	_
	_
	_
	_
	_
	_
	_
	Runeword173 // Treachery
	_
	_
	_
	_
	_
	Runeword179 // Venom
	_
	_
	_
	_
	_
	Runeword185 // Wealth
	_
	Runeword187 // White
	Runeword188 // Wind
	_
	_
	_
	_
	Runeword193 // Wrath
	_
	Runeword195 // Zephyr

	Runeword2718 Runeword = 2718 // Delirium
)
