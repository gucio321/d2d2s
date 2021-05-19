package d2smagicattributes

// MagicalProperty describes a string template, bias, and bit length for a
// particular kind of magical property, such as bonuses to all resistances or
// magic item finding.
type MagicalProperty struct {
	Bits []uint
	Bias uint64
	Name string
}

// GetPropertyData returns data for specified property id
// nolint:funlen // data function - can't reduce
func GetPropertyData(id uint16) *MagicalProperty {
	// nolint:gomnd // TODO
	lookup := map[uint16]MagicalProperty{
		0:  {Bits: []uint{8}, Bias: 32, Name: "+{0} to Strength"},
		1:  {Bits: []uint{7}, Bias: 32, Name: "+{0} to Energy"},
		2:  {Bits: []uint{7}, Bias: 32, Name: "+{0} to Dexterity"},
		3:  {Bits: []uint{7}, Bias: 32, Name: "+{0} to Vitality"},
		7:  {Bits: []uint{9}, Bias: 32, Name: "+{0} to Life"},
		9:  {Bits: []uint{8}, Bias: 32, Name: "+{0} to Mana"},
		11: {Bits: []uint{8}, Bias: 32, Name: "+{0} to Maximum Stamina"},
		16: {Bits: []uint{9}, Name: "+{0}% Enhanced Defense"},

		// It's weird that there's two bit fields here, but it seems like
		// the first bit field enhanced min damage, and the second enchances
		// maxium damage.
		17: {Bits: []uint{9, 9}, Name: "+{0}% Enhanced Damage"},
		19: {Bits: []uint{10}, Name: "+{0} to Attack rating"},
		20: {Bits: []uint{6}, Name: "+{0}% Increased chance of blocking"},
		21: {Bits: []uint{6}, Name: "+{0} to Minimum 1-handed damage"},
		22: {Bits: []uint{7}, Name: "+{0} to Maximum 1-handed damage"},
		23: {Bits: []uint{6}, Name: "+{0} to Minimum 2-handed damage"},
		24: {Bits: []uint{7}, Name: "+{0} to Maximum 2-handed damage"},
		25: {Bits: []uint{8}, Name: "Unknown (Invisible)"}, // damagepercent
		26: {Bits: []uint{8}, Name: "Unknown (Invisible)"}, // manarecovery
		27: {Bits: []uint{8}, Name: "Regenerate Mana {0}%"},
		28: {Bits: []uint{8}, Name: "Heal Stamina {0}%"},
		31: {Bits: []uint{11}, Bias: 10, Name: "+{0} Defense"},
		32: {Bits: []uint{9}, Name: "+{0} vs. Missile"},
		33: {Bits: []uint{8}, Bias: 10, Name: "+{0} vs. Melee"},
		34: {Bits: []uint{6}, Name: "Damage Reduced by {0}"},
		35: {Bits: []uint{6}, Name: "Magic Damage Reduced by {0}"},
		36: {Bits: []uint{8}, Name: "Damage Reduced by {0}%"},
		37: {Bits: []uint{8}, Name: "Magic Resist +{0}%"},
		38: {Bits: []uint{8}, Name: "+{0}% to Maximum Magic Resist"},
		39: {Bits: []uint{8}, Bias: 50, Name: "Fire Resist +{0}%"},
		40: {Bits: []uint{5}, Name: "+{0}% to Maximum Fire Resist"},
		41: {Bits: []uint{8}, Bias: 50, Name: "Lightning Resist +{0}%"},
		42: {Bits: []uint{5}, Name: "+{0}% to Maximum Lightning Resist"},
		43: {Bits: []uint{8}, Bias: 50, Name: "Cold Resist +{0}%"},
		44: {Bits: []uint{5}, Name: "+{0}% to Maximum Cold Resist"},
		45: {Bits: []uint{8}, Bias: 50, Name: "Poison Resist +{0}%"},
		46: {Bits: []uint{5}, Name: "+{0}% to Maximum Poison Resist"},
		48: {Bits: []uint{8, 9}, Name: "Adds {0}-{1} Fire Damage"},
		49: {Bits: []uint{9}, Name: "+{0} to Maximum Fire Damage"},
		50: {Bits: []uint{6, 10}, Name: "Adds {0}-{1} Lightning Damage"},
		52: {Bits: []uint{8, 9}, Name: "Adds {0}-{1} Magic Damage"},
		54: {Bits: []uint{8, 9, 8}, Name: "Adds {0}-{1} Cold Damage"},
		57: {Bits: []uint{10, 10, 9}, Name: "Adds {0}-{1} Poison Damage over {2} Seconds"},
		60: {Bits: []uint{7}, Name: "{0}% Life Stolen Per Hit"},
		62: {Bits: []uint{7}, Name: "{0}% Mana Stolen Per Hit"},
		67: {Bits: []uint{7}, Bias: 30, Name: "Unknown (Invisible)"},  // velocitypercent
		68: {Bits: []uint{7}, Bias: 30, Name: "Unknown (Invisible)"},  // attackrate
		71: {Bits: []uint{8}, Bias: 100, Name: "Unknown (Invisible)"}, // value
		72: {Bits: []uint{9}, Name: "Unknown (Invisible)"},            // durability
		73: {Bits: []uint{8}, Name: "+{0} Maximum Durability"},
		74: {Bits: []uint{6}, Bias: 30, Name: "Replenish Life +{0}"},
		75: {Bits: []uint{7}, Bias: 20, Name: "Increase Maximum Durability {0}%"},
		76: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Life {0}%"},
		77: {Bits: []uint{6}, Bias: 10, Name: "Increase Maximum Mana {0}%"},
		78: {Bits: []uint{7}, Name: "Attacker Takes Damage of {0}"},
		79: {Bits: []uint{9}, Bias: 100, Name: "{0}% Extra Gold from Monsters"},
		80: {Bits: []uint{8}, Bias: 100, Name: "{0}% Better Chance of Getting Magic Items"},
		81: {Bits: []uint{7}, Name: "Knockback"},
		82: {Bits: []uint{9}, Bias: 20, Name: "Unknown (Invisible)"}, // item_timeduration

		// First value is class, second is skill level, but they're printed in reverse
		// e.g. "+3 To Sorceress Skill Levels"
		83: {Bits: []uint{3, 3}, Name: "+{1} to {0} Skill Levels"},
		84: {Bits: []uint{3, 3}, Name: "+{1} to {0} Skill Levels"},

		// TODO: Check if experience gained really have a bias of 50.
		85: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Experience Gained"},
		86: {Bits: []uint{7}, Name: "+{0} Life After Each Kill"},
		87: {Bits: []uint{7}, Name: "Reduces Prices {0}%"},
		88: {Bits: []uint{1}, Name: "Unknown (Invisible)"}, // item_doubleherbduration
		89: {Bits: []uint{4}, Bias: 4, Name: "+{0} to Light Radius"},
		// This property is not displayed on the item, but its effect is to alter
		// the color of the ambient light.
		90: {Bits: []uint{5}, Name: "Ambient light"},
		// After subtracting the bias, this is usually a negative number.
		91: {Bits: []uint{8}, Bias: 100, Name: "Requirements {0}%"},
		92: {Bits: []uint{7}, Name: "Level requirements +{0} (Invisible)"},
		93: {Bits: []uint{7}, Bias: 20, Name: "{0}% Increased Attack Speed"},
		94: {Bits: []uint{7}, Bias: 64, Name: "Unknown (Invisible)"}, // item_levelreqpct
		96: {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Run/Walk"},

		// Number of levels to a certain skill, e.g. +1 To Teleport.
		97: {Bits: []uint{9, 6}, Name: "+{1} To {0}"},

		// NVSTATE Charm attributes. ID 98 only occurs on charms of a special
		// type, called NV state charms, they're basically for visual effects.
		// They're imported charms and does not occur naturally in the game.
		98: {Bits: []uint{8, 1}, Name: "{1}+ to {0} (Visual effect only)"},

		99:  {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Hit Recovery"},
		102: {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Block Rate"},
		105: {Bits: []uint{7}, Bias: 20, Name: "{0}% Faster Cast Rate"},

		// These properties usually applied to class specific items,
		// first value selects the skill, the second determines how many
		// additional skill points are given.
		107: {Bits: []uint{9, 3}, Name: "+{1} To {0}"},
		108: {Bits: []uint{1}, Name: "Rest In Peace"},
		109: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		181: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		182: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		183: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		184: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		185: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		186: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},
		187: {Bits: []uint{9, 5}, Name: "+{1} to spell {0} (char_class Only)"},

		110: {Bits: []uint{8}, Bias: 20, Name: "Poison Length Reduced by {0}%"},
		111: {Bits: []uint{9}, Bias: 20, Name: "Damage +{0}"},
		112: {Bits: []uint{7}, Name: "Hit Causes Monsters to Flee {0}%"},
		113: {Bits: []uint{7}, Name: "Hit Blinds Target +{0}"},
		114: {Bits: []uint{6}, Name: "{0}% Damage Taken Goes to Mana"},
		// The value of the data field is always 1.
		115: {Bits: []uint{1}, Name: "Ignore Target Defense"},
		116: {Bits: []uint{7}, Name: "{0}% Target Defense"},
		// The value of the data field is always 1.
		117: {Bits: []uint{7}, Name: "Prevent Monster Heal"},
		// The value of the data field is always 1.
		118: {Bits: []uint{1}, Name: "Half Freeze Duration"},
		119: {Bits: []uint{9}, Bias: 20, Name: "{0}% Bonus to Attack Rating"},
		120: {Bits: []uint{7}, Bias: 128, Name: "{0} to Monster Defense Per Hit"},
		121: {Bits: []uint{9}, Bias: 20, Name: "+{0}% Damage to Demons"},
		122: {Bits: []uint{9}, Bias: 20, Name: "+{0}% Damage to Undead"},
		123: {Bits: []uint{10}, Bias: 128, Name: "+{0} to Attack Rating against Demons"},
		124: {Bits: []uint{10}, Bias: 128, Name: "+{0} to Attack Rating against Undead"},
		125: {Bits: []uint{1}, Name: "Throwable"},
		// First value is class id, the next one is skill tree
		126: {Bits: []uint{3, 3}, Name: "+{0} to Fire Skills"},
		127: {Bits: []uint{3}, Name: "+{0} to All Skill Levels"},
		128: {Bits: []uint{5}, Name: "Attacker Takes Lightning Damage of {0}"},
		134: {Bits: []uint{5}, Name: "Freezes Target +{0}"},
		135: {Bits: []uint{7}, Name: "{0}% Chance of Open Wounds"},
		136: {Bits: []uint{7}, Name: "{0}% Chance of Crushing Blow"},
		137: {Bits: []uint{7}, Name: "+{0} Kick Damage"},
		138: {Bits: []uint{7}, Name: "+{0} to Mana After Each Kill"},
		139: {Bits: []uint{7}, Name: "+{0} Life after each Demon Kill"},
		// Unknown property, shows up on Swordback Hold Spiked Shield.
		140: {Bits: []uint{7}, Name: "Extra Blood (Invisible)"}, // item_extrablood
		141: {Bits: []uint{7}, Name: "{0}% Deadly Strike"},
		142: {Bits: []uint{7}, Name: "Fire Absorb {0}%"},
		143: {Bits: []uint{7}, Name: "+{0} Fire Absorb"},
		144: {Bits: []uint{7}, Name: "Lightning Absorb {0}%"},
		145: {Bits: []uint{7}, Name: "+{0} Lightning Absorb"},
		146: {Bits: []uint{7}, Name: "Magic Absorb {0}%"},
		147: {Bits: []uint{7}, Name: "+{0} Magic Absorb"},
		148: {Bits: []uint{7}, Name: "Cold Absorb {0}%"},
		149: {Bits: []uint{7}, Name: "+{0} Cold Absorb"},
		150: {Bits: []uint{7}, Name: "Slows Target by {0}%"},
		151: {Bits: []uint{9, 5}, Name: "Level +{1} {0} When Equipped"},
		152: {Bits: []uint{1}, Name: "Indestructible"},
		153: {Bits: []uint{1}, Name: "Cannot Be Frozen"},
		154: {Bits: []uint{7}, Bias: 20, Name: "{0}% Slower Stamina Drain"},
		155: {Bits: []uint{10, 7}, Name: "{0}% Chance to Reanimate Target"},
		156: {Bits: []uint{7}, Name: "Piercing Attack"},
		157: {Bits: []uint{7}, Name: "Fires Magic Arrows"},
		158: {Bits: []uint{7}, Name: "Fires Explosive Arrows or Bolts"},
		159: {Bits: []uint{6}, Name: "+{0} to Minimum Throw Damage"},
		160: {Bits: []uint{7}, Name: "+{0} to Maximum Throw Damage"},
		179: {Bits: []uint{3}, Name: "+{0} to Druid Skill Levels"},
		180: {Bits: []uint{3}, Name: "+{0} to Assassin Skill Levels"},

		// Ok so get this, this is quite complicated, the id 188 is for items with
		// + x to a certain skill tree, e.g. 1 + to defensive auras (paladin).
		//
		// So here's how it works, the field is 19 bits long, here's the bits for
		// the defensive auras skiller.
		//
		// 001             0000000000             011              010
		//  ^                  ^                   ^                ^
		// levels        unknown padding        class id       skill tree offset
		//
		// So in the above example, the first 3 bits 001 are the + levels (1), we'll
		// ignore the padding, the second interesting set of 3 bits is the class id.
		// Refer to the class.go for class ids, but paladin is 011 (3), and the
		// last 3 bits 010 (2) is the offset from the start of the class skill tree.
		// Refer to skills.go to find the different tree offsets. Paladin offset is
		// 9. So remember the last 3 bits 010 (2), that means the skill tree is
		// 9 + 2 = 11, aka the defensive auras tree.
		//
		// When reading the values, remember the bits are read from the right,
		// so the values will be [2 3 1], offset 2, class id 3, 1 + to skills.
		188: {Bits: []uint{3, 13, 3}, Name: "+{2} to {0} Skills ({1} only)"},

		189: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
		190: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
		191: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
		192: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},
		193: {Bits: []uint{10, 9}, Name: "+{0} to {1} Skills (char_class Only)"},

		194: {Bits: []uint{4}, Name: "Adds {0} extra sockets to the item"},

		// Order is spell id, level, % chance.
		195: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When you die"},
		196: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When you die"},
		197: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When you die"},

		// Order is spell id, level, % chance.
		198: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} On Striking"},
		199: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} On Striking"},
		200: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} On Striking"},

		// Order is spell id, level, % chance.
		201: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When Struck"},
		202: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When Struck"},
		203: {Bits: []uint{6, 10, 7}, Name: "{2}% Chance to Cast Level {0} {1} When Struck"},

		// First value selects the spell id, second value is level, third is remaining charges
		// and the last is the total number of charges.
		204: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		205: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		206: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		207: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		208: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		209: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		210: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		211: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		212: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},
		213: {Bits: []uint{6, 10, 8, 8}, Name: "Level {0} {1} ({2}/{3} Charges)"},

		// All values based on character level are stored in eights, so take
		// the number divide by 8 and multiply by the character level and round down.
		// Or, just do (value * 0.125)% per level.
		214: {Bits: []uint{6}, Name: "+{0} to Defense (Based on Character Level)"},
		215: {Bits: []uint{6}, Name: "{0}% Enhanced Defense (Based on Character Level)"},
		216: {Bits: []uint{6}, Name: "+{0} to Life (Based on Character Level)"},
		217: {Bits: []uint{6}, Name: "+{0} to Mana (Based on Character Level)"},
		218: {Bits: []uint{6}, Name: "+{0} to Maximum Damage (Based on Character Level)"},
		219: {Bits: []uint{6}, Name: "{0}% Enhanced Maximum Damage (Based on Character Level)"},
		220: {Bits: []uint{6}, Name: "+{0} to Strength (Based on Character Level)"},
		221: {Bits: []uint{6}, Name: "+{0} to Dexterity (Based on Character Level)"},
		222: {Bits: []uint{6}, Name: "+{0} to Energy (Based on Character Level)"},
		223: {Bits: []uint{6}, Name: "+{0} to Vitality (Based on Character Level)"},
		224: {Bits: []uint{6}, Name: "+{0} to Attack Rating (Based on Character Level)"},
		225: {Bits: []uint{6}, Name: "{0}% Bonus to Attack Rating (Based on Character Level)"},
		226: {Bits: []uint{6}, Name: "+{0} Cold Damage (Based on Character Level)"},
		227: {Bits: []uint{6}, Name: "+{0} Fire Damage (Based on Character Level)"},
		228: {Bits: []uint{6}, Name: "+{0} Lightning Damage (Based on Character Level)"},
		229: {Bits: []uint{6}, Name: "+{0} Poison Damage (Based on Character Level)"},
		230: {Bits: []uint{6}, Name: "Cold Resist +{0}% (Based on Character Level)"},
		231: {Bits: []uint{6}, Name: "Fire Resist +{0}% (Based on Character Level)"},
		232: {Bits: []uint{6}, Name: "Lightning Resist +{0}% (Based on Character Level)"},
		233: {Bits: []uint{6}, Name: "Poison Resist +{0}% (Based on Character Level)"},
		234: {Bits: []uint{6}, Name: "+{0} Cold Absorb (Based on Character Level)"},
		235: {Bits: []uint{6}, Name: "+{0} Fire Absorb (Based on Character Level)"},
		236: {Bits: []uint{6}, Name: "+{0} Lightning Absorb (Based on Character Level)"},
		237: {Bits: []uint{6}, Name: "{0} Poison Absorb (Based on Character Level)"},
		238: {Bits: []uint{5}, Name: "Attacker Takes Damage of {0} (Based on Character Level)"},
		239: {Bits: []uint{6}, Name: "{0}% Extra Gold from Monsters (Based on Character Level)"},
		240: {Bits: []uint{6}, Name: "{0}% Better Chance of Getting Magic Items (Based on Character Level)"},
		241: {Bits: []uint{6}, Name: "Heal Stamina Plus {0}% (Based on Character Level)"},
		242: {Bits: []uint{6}, Name: "+{0} Maxmium Stamina (Based on Character Level)"},
		243: {Bits: []uint{6}, Name: "{0}% Damage to Demons (Based on Character Level)"},
		244: {Bits: []uint{6}, Name: "{0}% Damage to Undead (Based on Character Level)"},
		245: {Bits: []uint{6}, Name: "+{0} to Attack Rating against Demons (Based on Character Level)"},
		246: {Bits: []uint{6}, Name: "+{0} to Attack Rating against Undead (Based on Character Level)"},
		247: {Bits: []uint{6}, Name: "{0}% Chance of Crushing Blow (Based on Character Level)"},
		248: {Bits: []uint{6}, Name: "{0}% Chance of Open Wounds (Based on Character Level)"},
		249: {Bits: []uint{6}, Name: "+{0} Kick Damage (Based on Character Level)"},
		250: {Bits: []uint{6}, Name: "{0}% to Deadly Strike (Based on Character Level)"},

		// The value of the data field is not actually a time period, but a frequency in terms
		// of the number of times durability is repaired over a period of 100 seconds.
		// For example, if the value is 5, then this property repairs 1 durability in 100 / 5 = 20 seconds.
		252: {Bits: []uint{6}, Name: "Repairs 1 Durability in {0} Seconds"},

		// As in the previous property, the value of the data field is a frequency in terms of the number
		// replenished over a period of 100 seconds. For example if the value is 4, then this property
		// replenishes 1 item in 100 / 4 = 25 seconds.
		253: {Bits: []uint{6}, Name: "Replenishes Quantity"},

		// Number of additional items beyond the base limit, for example if the base
		// is 50 and additional is 30, then the total is 50 + 30.
		254: {Bits: []uint{8}, Name: "Increased Stack Size"},

		// IDs 268 - 303 are some weird values that were never used in the actual game.
		// These values change depending on the time of day in the game.
		// The format of the bit fields are the same in all cases, the first 2 bits
		// specifies the the of time when the value is at its maximum.
		//
		// The second and third are respectively the minimum and maximum values of the property.
		// The maximum value at the time specified and the minimum at the opposite.

		305: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Cold"},
		306: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Fire"},
		307: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Lightning"},
		308: {Bits: []uint{8}, Bias: 50, Name: "{0} Pierce Poision"},

		324: {Bits: []uint{6}, Name: "Unknown (Invisible)"}, // item_extra_charges
		329: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Fire Skill Damage"},
		330: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Lightning Skill Damage"},
		331: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Cold Skill Damage"},
		332: {Bits: []uint{9}, Bias: 50, Name: "{0}% To Poison Skill Damage"},
		333: {Bits: []uint{8}, Name: "-{0}% To Enemy Fire Resistance"},
		334: {Bits: []uint{8}, Name: "-{0}% To Enemy Lightning Resistance"},
		335: {Bits: []uint{8}, Name: "-{0}% To Enemy Cold Resistance"},
		336: {Bits: []uint{8}, Name: "-{0}% To Enemy Poison Resistance"},
		356: {Bits: []uint{2}, Name: "Quest Item Difficulty +{0} (Invisible)"},
	}

	m, ok := lookup[id]
	if !ok {
		return nil
	}

	return &m
}
