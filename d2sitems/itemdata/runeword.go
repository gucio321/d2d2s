package itemdata

//go:generate stringer -linecomment -type RunewordID -output runeword_string.go

// RunewordID represents a runeword ID
type RunewordID uint16

// runeword IDs TODO: give these constants a bit more descriptive names
const (
	Runeword27   RunewordID = iota + 27 // Ancient's Pledge
	Runeword30                          // Beast
	Runeword32                          // Black
	Runeword34                          // Bone
	Runeword35                          // Bramble
	Runeword36                          // Brand
	Runeword37                          // Breath of the Dying
	Runeword39                          // Call to Arms
	Runeword40                          // Chains of Honor
	Runeword42                          // Chaos
	Runeword43                          // Crescent Moon
	Runeword46                          // Death
	Runeword51                          // Destruction
	Runeword52                          // Doom
	Runeword53                          // Dragon
	Runeword55                          // Dream
	Runeword56                          // Duress
	Runeword57                          // Edge
	Runeword59                          // Enigma
	Runeword60                          // Enlightenment
	Runeword62                          // Eternity
	Runeword63                          // Exile
	Runeword64                          // Faith
	Runeword65                          // Famine
	Runeword67                          // Fortitude
	Runeword70                          // Fury
	Runeword71                          // Gloom
	Runeword73                          // Grief
	Runeword74                          // Hand of Justice
	Runeword75                          // Harmory
	Runeword77                          // Heart of the Oak
	Runeword80                          // Holy Thunder
	Runeword81                          // Honor
	Runeword85                          // Ice
	Runeword86                          // Infinity
	Runeword88                          // Insight
	Runeword91                          // King's Grace
	Runeword92                          // Kingslayer
	Runeword95                          // Last Wish
	Runeword97                          // Lawbringer
	Runeword98                          // Leaf
	Runeword100                         // Lionheart
	Runeword101                         // Lore
	Runeword106                         // Malice
	Runeword107                         // Melody
	Runeword108                         // Memory
	Runeword112                         // Myth
	Runeword113                         // Nadir
	Runeword116                         // Oath
	Runeword117                         // Obedience
	Runeword120                         // Passion
	Runeword123                         // Peace
	Runeword124                         // Winter
	Runeword128                         // Phoenix
	Runeword131                         // Plague
	Runeword134                         // Pride
	Runeword135                         // Principle
	Runeword137                         // Prudence
	Runeword141                         // Radiance
	Runeword142                         // Rain
	Runeword145                         // Rhyme
	Runeword146                         // Rift
	Runeword147                         // Sanctuary
	Runeword151                         // Silence
	Runeword153                         // Smoke
	Runeword155                         // Spirit
	Runeword156                         // Splendor
	Runeword158                         // Stealth
	Runeword159                         // Steel
	Runeword162                         // Stone
	Runeword164                         // Strength
	Runeword173                         // Treachery
	Runeword179                         // Venom
	Runeword185                         // Wealth
	Runeword187                         // White
	Runeword188                         // Wind
	Runeword193                         // Wrath
	Runeword195                         // Zephyr
	Runeword2718                        // Delirium
)
