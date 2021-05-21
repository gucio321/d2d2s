package itemdata

//go:generate stringer -linecomment -type RunewordID -output runeword_string.go

// RunewordID represents a runeword ID
type RunewordID uint16

// runeword IDs TODO: give these constants a bit more descriptive names
const (
	RunewordAncientsPladge RunewordID = iota + 27 // Ancient's Pledge
	_
	_
	RunewordBeast // Beast
	_
	RunewordBlack // Black
	_
	RunewordBone        // Bone
	RunewordBramble     // Bramble
	RunewordBrand       // Brand
	RunewordDyingBreath // Breath of the Dying
	_
	RunewordCTA    // Call to Arms
	RunewordChains // Chains of Honor
	_
	RunewordChaos        // Chaos
	RunewordCrescentMoon // Crescent Moon
	_
	_
	RunewordDeath // Death
	_
	_
	_
	RunewordDestruction // Destruction
	RunewordDoom        // Doom
	RunewordDragon      // Dragon
	_
	RunewordDream  // Dream
	RunewordDuress // Duress
	RunewordEdge   // Edge
	_
	RunewordEnigma        // Enigma
	RunewordEnlightenment // Enlightenment
	_
	RunewordEternity // Eternity
	RunewordExile    // Exile
	RunewordFaith    // Faith
	RunewordFamine   // Famine
	_
	RunewordFortitude // Fortitude
	_
	_
	RunewordFury  // Fury
	RunewordGloom // Gloom
	_
	RunewordGrief       // Grief
	RunewordJusticeHand // Hand of Justice
	RunewordHarmory     // Harmory
	_
	RunewordHOTO // Heart of the Oak
	_
	_
	RunewordHolyThunder // Holy Thunder
	RunewordHonor       // Honor
	_
	_
	_
	RunewordIce      // Ice
	RunewordInfinity // Infinity
	_
	RunewordInsight // Insight
	_
	_
	RunewordKingsGrace // King's Grace
	RunewordKingslayer // Kingslayer
	_
	_
	RunewordLastWish // Last Wish
	_
	RunewordLawbringer // Lawbringer
	RunewordLeaf       // Leaf
	_
	RunewordLionheart // Lionheart
	RunewordLore      // Lore
	_
	_
	_
	_
	RunewordMalice // Malice
	RunewordMelody // Melody
	RunewordMemory // Memory
	_
	_
	_
	RunewordMyth  // Myth
	RunewordNadir // Nadir
	_
	_
	RunewordOath      // Oath
	RunewordObedience // Obedience
	_
	_
	RunewordPassion // Passion
	_
	_
	RunewordPeace  // Peace
	RunewordWinter // Winter
	_
	_
	_
	RunewordPhoenix // Phoenix
	_
	_
	RunewordPlague // Plague
	_
	_
	RunewordPride     // Pride
	RunewordPrinciple // Principle
	_
	RunewordPrudence // Prudence
	_
	_
	_
	RunewordRadiance // Radiance
	RunewordRain     // Rain
	_
	_
	RunewordRhyme     // Rhyme
	RunewordRift      // Rift
	RunewordSanctuary // Sanctuary
	_
	_
	_
	RunewordSilence // Silence
	_
	RunewordSmoke // Smoke
	_
	RunewordSpirit   // Spirit
	RunewordSplendor // Splendor
	_
	RunewordStealth // Stealth
	RunewordSteel   // Steel
	_
	_
	RunewordStone // Stone
	_
	RunewordStrength // Strength
	_
	_
	_
	_
	_
	_
	_
	_
	RunewordTreachery // Treachery
	_
	_
	_
	_
	_
	RunewordVenom // Venom
	_
	_
	_
	_
	_
	RunewordWealth // Wealth
	_
	RunewordWhite // White
	RunewordWind  // Wind
	_
	_
	_
	_
	RunewordWrath // Wrath
	_
	RunewordZephyr // Zephyr

	RunewordDelirium RunewordID = 2718 // Delirium
)
