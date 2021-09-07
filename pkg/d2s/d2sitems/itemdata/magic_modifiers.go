package itemdata

//go:generate stringer -linecomment -type MagicalPrefix -output magical_prefix_string.go

// MagicalPrefix represents a magical prefix
type MagicalPrefix uint16

const (
	MagicalPrefixStrong        MagicalPrefix = iota // Strong
	MagicalPrefixGlorious                           // Glorious
	MagicalPrefixBlessed                            // Blessed
	MagicalPrefixSaintly                            // Saintly
	MagicalPrefixHoly                               // Holy
	MagicalPrefixDevious                            // Devious
	MagicalPrefixFortified                          // Fortified
	MagicalPrefixJagged                             // Jagged
	MagicalPrefixDeadly                             // Deadly
	MagicalPrefixVicious                            // Vicious
	MagicalPrefixBrutal                             // Brutal
	MagicalPrefixMassive                            // Massive
	MagicalPrefixSavage                             // Savage
	MagicalPrefixMerciless                          // Merciless
	MagicalPrefixVulpine                            // Vulpine
	MagicalPrefixTireless                           // Tireless
	MagicalPrefixRugged                             // Rugged
	MagicalPrefixBronze                             // Bronze
	MagicalPrefixIron                               // Iron
	MagicalPrefixSteel                              // Steel
	MagicalPrefixSilver                             // Silver
	MagicalPrefixGold                               // Gold
	MagicalPrefixPlatinum                           // Platinum
	MagicalPrefixMeteoric                           // Meteoric
	MagicalPrefixSharp                              // Sharp
	MagicalPrefixFine                               // Fine
	MagicalPrefixWarriors                           // Warrior's
	MagicalPrefixSoldiers                           // Soldier's
	MagicalPrefixKnights                            // Knight's
	MagicalPrefixLords                              // Lord's
	MagicalPrefixKings                              // King's
	MagicalPrefixHowling                            // Howling
	MagicalPrefixFortuitous                         // Fortuitous
	MagicalPrefixGlimmering                         // Glimmering
	MagicalPrefixGlowing                            // Glowing
	MagicalPrefixLizards                            // Lizard's
	MagicalPrefixSnakes                             // Snake's
	MagicalPrefixSerpents                           // Serpent's
	MagicalPrefixDrakes                             // Drake's
	MagicalPrefixDragons                            // Dragon's
	MagicalPrefixWyrms                              // Wyrm's
	MagicalPrefixPrismatic                          // Prismatic
	MagicalPrefixAzure                              // Azure
	MagicalPrefixLapis                              // Lapis
	MagicalPrefixCobalt                             // Cobalt
	MagicalPrefixSapphire                           // Sapphire
	MagicalPrefixCrimson                            // Crimson
	MagicalPrefixBurgundy                           // Burgundy
	MagicalPrefixGarnet                             // Garnet
	MagicalPrefixRuby                               // Ruby
	MagicalPrefixOcher                              // Ocher
	MagicalPrefixTangerine                          // Tangerine
	MagicalPrefixCoral                              // Coral
	MagicalPrefixAmber                              // Amber
	MagicalPrefixBeryl                              // Beryl
	MagicalPrefixJade                               // Jade
	MagicalPrefixViridian                           // Viridian
	MagicalPrefixEmerald                            // Emerald
	MagicalPrefixFletchers                          // Fletcher's
	MagicalPrefixArchers                            // Archer's
	MagicalPrefixMonks                              // Monk's
	MagicalPrefixPriests                            // Priest's
	MagicalPrefixSummoners                          // Summoner's
	MagicalPrefixNecromancers                       // Necromancer's
	MagicalPrefixAngels                             // Angel's
	MagicalPrefixArchAngels                         // Arch-Angel's
	MagicalPrefixSlayers                            // Slayer's
	MagicalPrefixBerserkers                         // Berserker's
	MagicalPrefixTriumphant                         // Triumphant
	MagicalPrefixStout                              // Stout
	MagicalPrefixBurly                              // Burly
	MagicalPrefixStalwart                           // Stalwart
	MagicalPrefixBlanched                           // Blanched
	MagicalPrefixEburin                             // Eburin
	MagicalPrefixBone                               // Bone
	MagicalPrefixIvory                              // Ivory
	MagicalPrefixSturdy                             // Sturdy
	MagicalPrefixGodly                              // Godly
	MagicalPrefixBlank                              // Blank
	MagicalPrefixNull                               // Null
	MagicalPrefixAntimagic                          // Antimagic
	MagicalPrefixRed                                // Red
	MagicalPrefixSanguinary                         // Sanguinary
	MagicalPrefixBloody                             // Bloody
	MagicalPrefixScarlet                            // Scarlet
	MagicalPrefixForked                             // Forked
	MagicalPrefixSerrated                           // Serrated
	MagicalPrefixCarbuncle                          // Carbuncle
	MagicalPrefixCarmine                            // Carmine
	MagicalPrefixVermillion                         // Vermillion
	MagicalPrefixFerocious                          // Ferocious
	MagicalPrefixCruel                              // Cruel
	MagicalPrefixCinnabar                           // Cinnabar
	MagicalPrefixRusty                              // Rusty
	MagicalPrefixRealgar                            // Realgar
	MagicalPrefixDun                                // Dun
	MagicalPrefixBrown                              // Brown
	MagicalPrefixVigorous                           // Vigorous
	MagicalPrefixChestnut                           // Chestnut
	MagicalPrefixMaroon                             // Maroon
	MagicalPrefixStrange                            // Strange
	MagicalPrefixWeird                              // Weird
	MagicalPrefixNickel                             // Nickel
	MagicalPrefixTin                                // Tin
	MagicalPrefixArgent                             // Argent
	MagicalPrefixMasters                            // Master's
	MagicalPrefixGrandmasters                       // Grandmaster's
	MagicalPrefixBright                             // Bright
	MagicalPrefixScreaming                          // Screaming
	MagicalPrefixWailing                            // Wailing
	MagicalPrefixLucky                              // Lucky
	MagicalPrefixFelicitous                         // Felicitous
	MagicalPrefixGreatWyrms                         // GreatWyrm's
	MagicalPrefixBahamuts                           // Bahamut's
	MagicalPrefixZircon                             // Zircon
	MagicalPrefixJacinth                            // Jacinth
	MagicalPrefixTurquoise                          // Turquoise
	MagicalPrefixShimmering                         // Shimmering
	MagicalPrefixRainbow                            // Rainbow
	MagicalPrefixScintillating                      // Scintillating
	MagicalPrefixChromatic                          // Chromatic
	MagicalPrefixLapisLazuli                        // LapisLazuli
	MagicalPrefixRusset                             // Russet
	MagicalPrefixCamphor                            // Camphor
	MagicalPrefixAmbergris                          // Ambergris
	MagicalPrefixVictorious                         // Victorious
	MagicalPrefixAureolin                           // Aureolin
	MagicalPrefixMechanists                         // Mechanist's
	MagicalPrefixArtificers                         // Artificer's
	MagicalPrefixJewelers                           // Jeweler's
	MagicalPrefixAssamic                            // Assamic
	MagicalPrefixArcadian                           // Arcadian
	MagicalPrefixUnearthly                          // Unearthly
	MagicalPrefixAstral                             // Astral
	MagicalPrefixElysian                            // Elysian
	MagicalPrefixCelestial                          // Celestial
	MagicalPrefixDiamond                            // Diamond
	MagicalPrefixAcrobats                           // Acrobat's
	MagicalPrefixHarpoonists                        // Harpoonist's
	MagicalPrefixBowyers                            // Bowyer's
	MagicalPrefixGymnasts                           // Gymnast's
	MagicalPrefixAthletes                           // Athlete's
	MagicalPrefixSpearmaidens                       // Spearmaiden's
	MagicalPrefixLancers                            // Lancer's
	MagicalPrefixBurning                            // Burning
	MagicalPrefixSparking                           // Sparking
	MagicalPrefixChilling                           // Chilling
	MagicalPrefixBlazing                            // Blazing
	MagicalPrefixVolcanic                           // Volcanic
	MagicalPrefixCharged                            // Charged
	MagicalPrefixPowered                            // Powered
	MagicalPrefixFreezing                           // Freezing
	MagicalPrefixGlacial                            // Glacial
	MagicalPrefixHexing                             // Hexing
	MagicalPrefixFungal                             // Fungal
	MagicalPrefixGraverobbers                       // Graverobber's
	MagicalPrefixBlighting                          // Blighting
	MagicalPrefixAccursed                           // Accursed
	MagicalPrefixNoxious                            // Noxious
	MagicalPrefixVenomous                           // Venomous
	MagicalPrefixVodoun                             // Vodoun
	MagicalPrefixGolemlords                         // Golemlord's
	MagicalPrefixLionBranded                        // LionBranded
	MagicalPrefixCaptains                           // Captain's
	MagicalPrefixPreservers                         // Preserver's
	MagicalPrefixHawkBranded                        // HawkBranded
	MagicalPrefixRoseBranded                        // RoseBranded
	MagicalPrefixCommanders                         // Commander's
	MagicalPrefixMarshals                           // Marshal's
	MagicalPrefixWarders                            // Warder's
	MagicalPrefixGuardians                          // Guardian's
	MagicalPrefixExperts                            // Expert's
	MagicalPrefixFanatic                            // Fanatic
	MagicalPrefixSounding                           // Sounding
	MagicalPrefixVeterans                           // Veteran's
	MagicalPrefixRaging                             // Raging
	MagicalPrefixFurious                            // Furious
	MagicalPrefixResonant                           // Resonant
	MagicalPrefixEchoing                            // Echoing
	MagicalPrefixTrainers                           // Trainer's
	MagicalPrefixSpiritual                          // Spiritual
	MagicalPrefixNatures                            // Nature's
	MagicalPrefixCaretakers                         // Caretaker's
	MagicalPrefixKeepers                            // Keeper's
	MagicalPrefixFeral                              // Feral
	MagicalPrefixCommunal                           // Communal
	MagicalPrefixTerras                             // Terra's
	MagicalPrefixGaeas                              // Gaea's
	MagicalPrefixEntrapping                         // Entrapping
	MagicalPrefixMentalists                         // Mentalist's
	MagicalPrefixShogukushas                        // Shogukusha's
	MagicalPrefixTricksters                         // Trickster's
	MagicalPrefixCunning                            // Cunning
	MagicalPrefixPsychic                            // Psychic
	MagicalPrefixShadow                             // Shadow
	MagicalPrefixSenseis                            // Sensei's
	MagicalPrefixKenshis                            // Kenshi's
	MagicalPrefixMiocene                            // Miocene
	MagicalPrefixOligocene                          // Oligocene
	MagicalPrefixEocene                             // Eocene
	MagicalPrefixPaleocene                          // Paleocene
	MagicalPrefixKnaves                             // Knave's
	MagicalPrefixJacks                              // Jack's
	MagicalPrefixJesters                            // Jester's
	MagicalPrefixJokers                             // Joker's
	MagicalPrefixTrump                              // Trump
	MagicalPrefixLoud                               // Loud
	MagicalPrefixCalling                            // Calling
	MagicalPrefixYelling                            // Yelling
	MagicalPrefixShouting                           // Shouting
	MagicalPrefixParadox                            // Paradox
	MagicalPrefixRobineye                           // Robineye
	MagicalPrefixSparroweye                         // Sparroweye
	MagicalPrefixFalconeye                          // Falconeye
	MagicalPrefixHawkeye                            // Hawkeye
	MagicalPrefixEagleeye                           // Eagleeye
	MagicalPrefixVisionary                          // Visionary
	MagicalPrefixMnemonic                           // Mnemonic
	MagicalPrefixSnowflake                          // Snowflake
	MagicalPrefixShivering                          // Shivering
	MagicalPrefixBoreal                             // Boreal
	MagicalPrefixHibernal                           // Hibernal
	MagicalPrefixEmber                              // Ember
	MagicalPrefixSmoldering                         // Smoldering
	MagicalPrefixSmoking                            // Smoking
	MagicalPrefixFlaming                            // Flaming
	MagicalPrefixScorching                          // Scorching
	MagicalPrefixStatic                             // Static
	MagicalPrefixBuzzing                            // Buzzing
	MagicalPrefixArcing                             // Arcing
	MagicalPrefixShocking                           // Shocking
	MagicalPrefixSeptic                             // Septic
	MagicalPrefixEnvenomed                          // Envenomed
	MagicalPrefixCorosive                           // Corosive
	MagicalPrefixToxic                              // Toxic
	MagicalPrefixPestilent                          // Pestilent
	MagicalPrefixMaidens                            // Maiden's
	MagicalPrefixValkyries                          // Valkyrie's
	MagicalPrefixShamans                            // Shaman's
	MagicalPrefixHierophants                        // Hierophant's
	MagicalPrefixMagekillers                        // Magekiller's
	MagicalPrefixWitchHunters                       // Witch-hunter's
	MagicalPrefixCompact                            // Compact
	MagicalPrefixThin                               // Thin
	MagicalPrefixDense                              // Dense
	MagicalPrefixConsecrated                        // Consecrated
	MagicalPrefixPure                               // Pure
	MagicalPrefixSacred                             // Sacred
	MagicalPrefixHallowed                           // Hallowed
	MagicalPrefixDivine                             // Divine
	MagicalPrefixPearl                              // Pearl
)

func getMagicalPrefixMap() map[uint16]MagicalPrefix {
	return map[uint16]MagicalPrefix{
		2:   MagicalPrefixSturdy,
		3:   MagicalPrefixStrong,
		4:   MagicalPrefixGlorious,
		5:   MagicalPrefixBlessed,
		6:   MagicalPrefixSaintly,
		7:   MagicalPrefixHoly,
		8:   MagicalPrefixDevious,
		9:   MagicalPrefixFortified,
		13:  MagicalPrefixJagged,
		14:  MagicalPrefixDeadly,
		15:  MagicalPrefixVicious,
		16:  MagicalPrefixBrutal,
		17:  MagicalPrefixMassive,
		18:  MagicalPrefixSavage,
		19:  MagicalPrefixMerciless,
		20:  MagicalPrefixVulpine,
		25:  MagicalPrefixTireless,
		26:  MagicalPrefixRugged,
		27:  MagicalPrefixBronze,
		28:  MagicalPrefixIron,
		29:  MagicalPrefixSteel,
		30:  MagicalPrefixSilver,
		32:  MagicalPrefixGold,
		33:  MagicalPrefixPlatinum,
		34:  MagicalPrefixMeteoric,
		35:  MagicalPrefixSharp,
		36:  MagicalPrefixFine,
		37:  MagicalPrefixWarriors,
		38:  MagicalPrefixSoldiers,
		39:  MagicalPrefixKnights,
		40:  MagicalPrefixLords,
		41:  MagicalPrefixKings,
		42:  MagicalPrefixHowling,
		43:  MagicalPrefixFortuitous,
		49:  MagicalPrefixGlimmering,
		50:  MagicalPrefixGlowing,
		53:  MagicalPrefixLizards,
		55:  MagicalPrefixSnakes,
		56:  MagicalPrefixSerpents,
		57:  MagicalPrefixSerpents,
		58:  MagicalPrefixDrakes,
		59:  MagicalPrefixDragons,
		60:  MagicalPrefixDragons,
		61:  MagicalPrefixWyrms,
		64:  MagicalPrefixPrismatic,
		65:  MagicalPrefixPrismatic,
		66:  MagicalPrefixAzure,
		67:  MagicalPrefixLapis,
		68:  MagicalPrefixLapis,
		69:  MagicalPrefixCobalt,
		70:  MagicalPrefixCobalt,
		72:  MagicalPrefixSapphire,
		75:  MagicalPrefixCrimson,
		76:  MagicalPrefixBurgundy,
		77:  MagicalPrefixBurgundy,
		78:  MagicalPrefixGarnet,
		79:  MagicalPrefixGarnet,
		81:  MagicalPrefixRuby,
		84:  MagicalPrefixOcher,
		85:  MagicalPrefixTangerine,
		86:  MagicalPrefixTangerine,
		87:  MagicalPrefixCoral,
		88:  MagicalPrefixCoral,
		90:  MagicalPrefixAmber,
		93:  MagicalPrefixBeryl,
		94:  MagicalPrefixJade,
		95:  MagicalPrefixJade,
		96:  MagicalPrefixViridian,
		97:  MagicalPrefixViridian,
		99:  MagicalPrefixEmerald,
		101: MagicalPrefixFletchers,
		102: MagicalPrefixArchers,
		103: MagicalPrefixArchers,
		104: MagicalPrefixMonks,
		105: MagicalPrefixPriests,
		106: MagicalPrefixPriests,
		107: MagicalPrefixSummoners,
		108: MagicalPrefixNecromancers,
		109: MagicalPrefixNecromancers,
		110: MagicalPrefixAngels,
		111: MagicalPrefixArchAngels,
		112: MagicalPrefixArchAngels,
		113: MagicalPrefixSlayers,
		114: MagicalPrefixBerserkers,
		115: MagicalPrefixBerserkers,
		118: MagicalPrefixTriumphant,
		119: MagicalPrefixStout,
		120: MagicalPrefixStout,
		121: MagicalPrefixStout,
		122: MagicalPrefixBurly,
		123: MagicalPrefixBurly,
		124: MagicalPrefixBurly,
		125: MagicalPrefixStalwart,
		126: MagicalPrefixStalwart,
		127: MagicalPrefixStalwart,
		128: MagicalPrefixStout,
		129: MagicalPrefixStout,
		130: MagicalPrefixStout,
		131: MagicalPrefixBurly,
		132: MagicalPrefixBurly,
		133: MagicalPrefixStalwart,
		134: MagicalPrefixStalwart,
		135: MagicalPrefixStout,
		136: MagicalPrefixStout,
		137: MagicalPrefixBurly,
		138: MagicalPrefixStalwart,
		139: MagicalPrefixBlanched,
		140: MagicalPrefixEburin,
		141: MagicalPrefixBone,
		142: MagicalPrefixIvory,
		143: MagicalPrefixSturdy,
		144: MagicalPrefixSturdy,
		145: MagicalPrefixStrong,
		146: MagicalPrefixGlorious,
		147: MagicalPrefixBlessed,
		148: MagicalPrefixSaintly,
		149: MagicalPrefixHoly,
		150: MagicalPrefixGodly,
		151: MagicalPrefixDevious,
		152: MagicalPrefixBlank,
		153: MagicalPrefixNull,
		154: MagicalPrefixAntimagic,
		155: MagicalPrefixRed,
		156: MagicalPrefixRed,
		157: MagicalPrefixSanguinary,
		158: MagicalPrefixSanguinary,
		159: MagicalPrefixBloody,
		160: MagicalPrefixRed,
		161: MagicalPrefixSanguinary,
		162: MagicalPrefixBloody,
		163: MagicalPrefixRed,
		164: MagicalPrefixSanguinary,
		165: MagicalPrefixBloody,
		166: MagicalPrefixScarlet,
		167: MagicalPrefixCrimson,
		168: MagicalPrefixJagged,
		169: MagicalPrefixJagged,
		170: MagicalPrefixJagged,
		171: MagicalPrefixForked,
		172: MagicalPrefixForked,
		173: MagicalPrefixSerrated,
		174: MagicalPrefixSerrated,
		175: MagicalPrefixJagged,
		176: MagicalPrefixJagged,
		177: MagicalPrefixForked,
		178: MagicalPrefixForked,
		179: MagicalPrefixSerrated,
		180: MagicalPrefixJagged,
		181: MagicalPrefixForked,
		182: MagicalPrefixSerrated,
		183: MagicalPrefixCarbuncle,
		184: MagicalPrefixCarmine,
		185: MagicalPrefixVermillion,
		186: MagicalPrefixJagged,
		187: MagicalPrefixDeadly,
		188: MagicalPrefixVicious,
		189: MagicalPrefixBrutal,
		190: MagicalPrefixMassive,
		191: MagicalPrefixSavage,
		192: MagicalPrefixMerciless,
		193: MagicalPrefixFerocious,
		194: MagicalPrefixCruel,
		195: MagicalPrefixCinnabar,
		196: MagicalPrefixRusty,
		197: MagicalPrefixRealgar,
		198: MagicalPrefixRuby,
		199: MagicalPrefixVulpine,
		200: MagicalPrefixDun,
		201: MagicalPrefixTireless,
		202: MagicalPrefixTireless,
		203: MagicalPrefixBrown,
		204: MagicalPrefixRugged,
		205: MagicalPrefixRugged,
		206: MagicalPrefixRugged,
		207: MagicalPrefixRugged,
		208: MagicalPrefixRugged,
		209: MagicalPrefixRugged,
		210: MagicalPrefixRugged,
		211: MagicalPrefixRugged,
		212: MagicalPrefixRugged,
		213: MagicalPrefixRugged,
		214: MagicalPrefixRugged,
		215: MagicalPrefixVigorous,
		216: MagicalPrefixChestnut,
		217: MagicalPrefixMaroon,
		218: MagicalPrefixBronze,
		219: MagicalPrefixBronze,
		220: MagicalPrefixBronze,
		221: MagicalPrefixIron,
		222: MagicalPrefixIron,
		223: MagicalPrefixIron,
		224: MagicalPrefixSteel,
		225: MagicalPrefixSteel,
		226: MagicalPrefixSteel,
		227: MagicalPrefixBronze,
		228: MagicalPrefixBronze,
		229: MagicalPrefixBronze,
		230: MagicalPrefixIron,
		231: MagicalPrefixIron,
		232: MagicalPrefixSteel,
		233: MagicalPrefixSteel,
		234: MagicalPrefixBronze,
		235: MagicalPrefixBronze,
		236: MagicalPrefixIron,
		237: MagicalPrefixSteel,
		238: MagicalPrefixBronze,
		239: MagicalPrefixIron,
		240: MagicalPrefixSteel,
		241: MagicalPrefixSilver,
		242: MagicalPrefixGold,
		243: MagicalPrefixPlatinum,
		244: MagicalPrefixMeteoric,
		245: MagicalPrefixStrange,
		246: MagicalPrefixWeird,
		247: MagicalPrefixNickel,
		248: MagicalPrefixTin,
		249: MagicalPrefixSilver,
		250: MagicalPrefixArgent,
		251: MagicalPrefixFine,
		252: MagicalPrefixFine,
		253: MagicalPrefixSharp,
		254: MagicalPrefixFine,
		255: MagicalPrefixSharp,
		256: MagicalPrefixFine,
		257: MagicalPrefixSharp,
		258: MagicalPrefixFine,
		259: MagicalPrefixWarriors,
		260: MagicalPrefixSoldiers,
		261: MagicalPrefixKnights,
		262: MagicalPrefixLords,
		263: MagicalPrefixKings,
		264: MagicalPrefixMasters,
		265: MagicalPrefixGrandmasters,
		266: MagicalPrefixGlimmering,
		267: MagicalPrefixGlowing,
		268: MagicalPrefixBright,
		269: MagicalPrefixScreaming,
		270: MagicalPrefixHowling,
		271: MagicalPrefixWailing,
		272: MagicalPrefixScreaming,
		273: MagicalPrefixHowling,
		274: MagicalPrefixWailing,
		275: MagicalPrefixLucky,
		276: MagicalPrefixLucky,
		277: MagicalPrefixLucky,
		278: MagicalPrefixLucky,
		279: MagicalPrefixLucky,
		280: MagicalPrefixLucky,
		281: MagicalPrefixFelicitous,
		282: MagicalPrefixFortuitous,
		283: MagicalPrefixEmerald,
		284: MagicalPrefixLizards,
		285: MagicalPrefixLizards,
		286: MagicalPrefixLizards,
		287: MagicalPrefixSnakes,
		288: MagicalPrefixSnakes,
		289: MagicalPrefixSnakes,
		290: MagicalPrefixSerpents,
		291: MagicalPrefixSerpents,
		292: MagicalPrefixSerpents,
		293: MagicalPrefixLizards,
		294: MagicalPrefixLizards,
		295: MagicalPrefixLizards,
		296: MagicalPrefixSnakes,
		297: MagicalPrefixSnakes,
		298: MagicalPrefixSerpents,
		299: MagicalPrefixSerpents,
		300: MagicalPrefixLizards,
		301: MagicalPrefixLizards,
		302: MagicalPrefixSnakes,
		303: MagicalPrefixSerpents,
		304: MagicalPrefixLizards,
		305: MagicalPrefixSnakes,
		306: MagicalPrefixSerpents,
		307: MagicalPrefixSerpents,
		308: MagicalPrefixDrakes,
		309: MagicalPrefixDragons,
		310: MagicalPrefixDragons,
		311: MagicalPrefixWyrms,
		312: MagicalPrefixGreatWyrms,
		313: MagicalPrefixBahamuts,
		314: MagicalPrefixZircon,
		315: MagicalPrefixJacinth,
		316: MagicalPrefixTurquoise,
		317: MagicalPrefixShimmering,
		318: MagicalPrefixShimmering,
		319: MagicalPrefixShimmering,
		320: MagicalPrefixShimmering,
		321: MagicalPrefixShimmering,
		322: MagicalPrefixShimmering,
		323: MagicalPrefixShimmering,
		324: MagicalPrefixRainbow,
		325: MagicalPrefixScintillating,
		326: MagicalPrefixPrismatic,
		327: MagicalPrefixChromatic,
		328: MagicalPrefixShimmering,
		329: MagicalPrefixRainbow,
		330: MagicalPrefixScintillating,
		331: MagicalPrefixPrismatic,
		332: MagicalPrefixChromatic,
		333: MagicalPrefixShimmering,
		334: MagicalPrefixRainbow,
		335: MagicalPrefixScintillating,
		336: MagicalPrefixShimmering,
		337: MagicalPrefixScintillating,
		338: MagicalPrefixAzure,
		339: MagicalPrefixLapis,
		340: MagicalPrefixCobalt,
		341: MagicalPrefixSapphire,
		342: MagicalPrefixAzure,
		343: MagicalPrefixLapis,
		344: MagicalPrefixCobalt,
		345: MagicalPrefixSapphire,
		346: MagicalPrefixAzure,
		347: MagicalPrefixLapis,
		348: MagicalPrefixCobalt,
		349: MagicalPrefixSapphire,
		350: MagicalPrefixAzure,
		351: MagicalPrefixLapis,
		352: MagicalPrefixLapis,
		353: MagicalPrefixCobalt,
		354: MagicalPrefixCobalt,
		355: MagicalPrefixSapphire,
		356: MagicalPrefixLapisLazuli,
		357: MagicalPrefixSapphire,
		358: MagicalPrefixCrimson,
		359: MagicalPrefixRusset,
		360: MagicalPrefixGarnet,
		361: MagicalPrefixRuby,
		362: MagicalPrefixCrimson,
		363: MagicalPrefixRusset,
		364: MagicalPrefixGarnet,
		365: MagicalPrefixRuby,
		366: MagicalPrefixCrimson,
		367: MagicalPrefixRusset,
		368: MagicalPrefixGarnet,
		369: MagicalPrefixRuby,
		370: MagicalPrefixRusset,
		371: MagicalPrefixRusset,
		372: MagicalPrefixGarnet,
		373: MagicalPrefixGarnet,
		374: MagicalPrefixRuby,
		375: MagicalPrefixGarnet,
		376: MagicalPrefixRuby,
		377: MagicalPrefixTangerine,
		378: MagicalPrefixOcher,
		379: MagicalPrefixCoral,
		380: MagicalPrefixAmber,
		381: MagicalPrefixTangerine,
		382: MagicalPrefixOcher,
		383: MagicalPrefixCoral,
		384: MagicalPrefixAmber,
		385: MagicalPrefixTangerine,
		386: MagicalPrefixOcher,
		387: MagicalPrefixCoral,
		388: MagicalPrefixAmber,
		389: MagicalPrefixTangerine,
		390: MagicalPrefixOcher,
		391: MagicalPrefixOcher,
		392: MagicalPrefixCoral,
		393: MagicalPrefixCoral,
		394: MagicalPrefixAmber,
		395: MagicalPrefixCamphor,
		396: MagicalPrefixAmbergris,
		397: MagicalPrefixBeryl,
		398: MagicalPrefixViridian,
		399: MagicalPrefixJade,
		400: MagicalPrefixEmerald,
		401: MagicalPrefixBeryl,
		402: MagicalPrefixViridian,
		403: MagicalPrefixJade,
		404: MagicalPrefixEmerald,
		405: MagicalPrefixBeryl,
		406: MagicalPrefixViridian,
		407: MagicalPrefixJade,
		408: MagicalPrefixEmerald,
		409: MagicalPrefixBeryl,
		410: MagicalPrefixViridian,
		411: MagicalPrefixViridian,
		412: MagicalPrefixJade,
		413: MagicalPrefixJade,
		414: MagicalPrefixEmerald,
		415: MagicalPrefixBeryl,
		416: MagicalPrefixJade,
		417: MagicalPrefixTriumphant,
		418: MagicalPrefixVictorious,
		419: MagicalPrefixAureolin,
		420: MagicalPrefixMechanists,
		421: MagicalPrefixArtificers,
		422: MagicalPrefixJewelers,
		423: MagicalPrefixAssamic,
		424: MagicalPrefixArcadian,
		425: MagicalPrefixUnearthly,
		426: MagicalPrefixAstral,
		427: MagicalPrefixElysian,
		428: MagicalPrefixCelestial,
		429: MagicalPrefixDiamond,
		430: MagicalPrefixFletchers,
		431: MagicalPrefixAcrobats,
		432: MagicalPrefixHarpoonists,
		433: MagicalPrefixFletchers,
		434: MagicalPrefixBowyers,
		435: MagicalPrefixArchers,
		436: MagicalPrefixAcrobats,
		437: MagicalPrefixGymnasts,
		438: MagicalPrefixAthletes,
		439: MagicalPrefixHarpoonists,
		440: MagicalPrefixSpearmaidens,
		441: MagicalPrefixLancers,
		442: MagicalPrefixBurning,
		443: MagicalPrefixSparking,
		444: MagicalPrefixChilling,
		445: MagicalPrefixBurning,
		446: MagicalPrefixBlazing,
		447: MagicalPrefixVolcanic,
		448: MagicalPrefixSparking,
		449: MagicalPrefixCharged,
		450: MagicalPrefixPowered,
		451: MagicalPrefixChilling,
		452: MagicalPrefixFreezing,
		453: MagicalPrefixGlacial,
		454: MagicalPrefixHexing,
		455: MagicalPrefixFungal,
		456: MagicalPrefixGraverobbers,
		457: MagicalPrefixHexing,
		458: MagicalPrefixBlighting,
		459: MagicalPrefixAccursed,
		460: MagicalPrefixFungal,
		461: MagicalPrefixNoxious,
		462: MagicalPrefixVenomous,
		463: MagicalPrefixGraverobbers,
		464: MagicalPrefixVodoun,
		465: MagicalPrefixGolemlords,
		466: MagicalPrefixLionBranded,
		467: MagicalPrefixCaptains,
		468: MagicalPrefixPreservers,
		469: MagicalPrefixLionBranded,
		470: MagicalPrefixHawkBranded,
		471: MagicalPrefixRoseBranded,
		472: MagicalPrefixCaptains,
		473: MagicalPrefixCommanders,
		474: MagicalPrefixMarshals,
		475: MagicalPrefixPreservers,
		476: MagicalPrefixWarders,
		477: MagicalPrefixGuardians,
		478: MagicalPrefixExperts,
		479: MagicalPrefixFanatic,
		480: MagicalPrefixSounding,
		481: MagicalPrefixExperts,
		482: MagicalPrefixVeterans,
		483: MagicalPrefixMasters,
		484: MagicalPrefixFanatic,
		485: MagicalPrefixRaging,
		486: MagicalPrefixFurious,
		487: MagicalPrefixSounding,
		488: MagicalPrefixResonant,
		489: MagicalPrefixEchoing,
		490: MagicalPrefixTrainers,
		491: MagicalPrefixSpiritual,
		492: MagicalPrefixNatures,
		493: MagicalPrefixTrainers,
		494: MagicalPrefixCaretakers,
		495: MagicalPrefixKeepers,
		496: MagicalPrefixSpiritual,
		497: MagicalPrefixFeral,
		498: MagicalPrefixCommunal,
		499: MagicalPrefixNatures,
		500: MagicalPrefixTerras,
		501: MagicalPrefixGaeas,
		502: MagicalPrefixEntrapping,
		503: MagicalPrefixMentalists,
		504: MagicalPrefixShogukushas,
		505: MagicalPrefixEntrapping,
		506: MagicalPrefixTricksters,
		507: MagicalPrefixCunning,
		508: MagicalPrefixMentalists,
		509: MagicalPrefixPsychic,
		510: MagicalPrefixShadow,
		511: MagicalPrefixShogukushas,
		512: MagicalPrefixSenseis,
		513: MagicalPrefixKenshis,
		514: MagicalPrefixMiocene,
		515: MagicalPrefixMiocene,
		516: MagicalPrefixOligocene,
		517: MagicalPrefixOligocene,
		518: MagicalPrefixEocene,
		519: MagicalPrefixEocene,
		520: MagicalPrefixPaleocene,
		521: MagicalPrefixPaleocene,
		522: MagicalPrefixKnaves,
		523: MagicalPrefixJacks,
		524: MagicalPrefixJesters,
		525: MagicalPrefixJokers,
		526: MagicalPrefixTrump,
		527: MagicalPrefixLoud,
		528: MagicalPrefixCalling,
		529: MagicalPrefixYelling,
		530: MagicalPrefixShouting,
		531: MagicalPrefixScreaming,
		532: MagicalPrefixParadox,
		533: MagicalPrefixParadox,
		534: MagicalPrefixRobineye,
		535: MagicalPrefixSparroweye,
		536: MagicalPrefixFalconeye,
		537: MagicalPrefixHawkeye,
		538: MagicalPrefixEagleeye,
		539: MagicalPrefixVisionary,
		540: MagicalPrefixMnemonic,
		541: MagicalPrefixSnowflake,
		542: MagicalPrefixShivering,
		543: MagicalPrefixBoreal,
		544: MagicalPrefixHibernal,
		545: MagicalPrefixEmber,
		546: MagicalPrefixSmoldering,
		547: MagicalPrefixSmoking,
		548: MagicalPrefixFlaming,
		549: MagicalPrefixScorching,
		550: MagicalPrefixStatic,
		551: MagicalPrefixGlowing,
		552: MagicalPrefixBuzzing,
		553: MagicalPrefixArcing,
		554: MagicalPrefixShocking,
		555: MagicalPrefixSeptic,
		556: MagicalPrefixEnvenomed,
		557: MagicalPrefixCorosive,
		558: MagicalPrefixToxic,
		559: MagicalPrefixPestilent,
		560: MagicalPrefixMaidens,
		561: MagicalPrefixValkyries,
		562: MagicalPrefixMaidens,
		563: MagicalPrefixValkyries,
		564: MagicalPrefixMonks,
		565: MagicalPrefixPriests,
		566: MagicalPrefixMonks,
		567: MagicalPrefixPriests,
		568: MagicalPrefixMonks,
		569: MagicalPrefixPriests,
		570: MagicalPrefixSummoners,
		571: MagicalPrefixNecromancers,
		572: MagicalPrefixSummoners,
		573: MagicalPrefixNecromancers,
		574: MagicalPrefixAngels,
		575: MagicalPrefixArchAngels,
		576: MagicalPrefixAngels,
		577: MagicalPrefixArchAngels,
		578: MagicalPrefixSlayers,
		579: MagicalPrefixBerserkers,
		580: MagicalPrefixSlayers,
		581: MagicalPrefixBerserkers,
		582: MagicalPrefixSlayers,
		583: MagicalPrefixBerserkers,
		584: MagicalPrefixShamans,
		585: MagicalPrefixHierophants,
		586: MagicalPrefixShamans,
		587: MagicalPrefixHierophants,
		588: MagicalPrefixMagekillers,
		589: MagicalPrefixWitchHunters,
		590: MagicalPrefixMagekillers,
		591: MagicalPrefixWitchHunters,
		592: MagicalPrefixCompact,
		593: MagicalPrefixThin,
		594: MagicalPrefixDense,
		595: MagicalPrefixConsecrated,
		596: MagicalPrefixPure,
		597: MagicalPrefixSacred,
		598: MagicalPrefixHallowed,
		599: MagicalPrefixDivine,
		600: MagicalPrefixPearl,
		601: MagicalPrefixCrimson,
		602: MagicalPrefixRed,
		603: MagicalPrefixSanguinary,
		604: MagicalPrefixBloody,
		605: MagicalPrefixRed,
		606: MagicalPrefixSanguinary,
		607: MagicalPrefixRed,
		608: MagicalPrefixJagged,
		609: MagicalPrefixForked,
		610: MagicalPrefixSerrated,
		611: MagicalPrefixJagged,
		612: MagicalPrefixForked,
		613: MagicalPrefixJagged,
		614: MagicalPrefixSnowflake,
		615: MagicalPrefixShivering,
		616: MagicalPrefixBoreal,
		617: MagicalPrefixHibernal,
		618: MagicalPrefixSnowflake,
		619: MagicalPrefixShivering,
		620: MagicalPrefixBoreal,
		621: MagicalPrefixHibernal,
		622: MagicalPrefixSnowflake,
		623: MagicalPrefixShivering,
		624: MagicalPrefixBoreal,
		625: MagicalPrefixHibernal,
		626: MagicalPrefixEmber,
		627: MagicalPrefixSmoldering,
		628: MagicalPrefixSmoking,
		629: MagicalPrefixFlaming,
		630: MagicalPrefixEmber,
		631: MagicalPrefixSmoldering,
		632: MagicalPrefixSmoking,
		633: MagicalPrefixFlaming,
		634: MagicalPrefixEmber,
		635: MagicalPrefixSmoldering,
		636: MagicalPrefixSmoking,
		637: MagicalPrefixFlaming,
		638: MagicalPrefixStatic,
		639: MagicalPrefixGlowing,
		640: MagicalPrefixArcing,
		641: MagicalPrefixShocking,
		642: MagicalPrefixStatic,
		643: MagicalPrefixGlowing,
		644: MagicalPrefixArcing,
		645: MagicalPrefixShocking,
		646: MagicalPrefixStatic,
		647: MagicalPrefixGlowing,
		648: MagicalPrefixArcing,
		649: MagicalPrefixShocking,
		650: MagicalPrefixSeptic,
		651: MagicalPrefixEnvenomed,
		652: MagicalPrefixToxic,
		653: MagicalPrefixPestilent,
		654: MagicalPrefixSeptic,
		655: MagicalPrefixEnvenomed,
		656: MagicalPrefixToxic,
		657: MagicalPrefixPestilent,
		658: MagicalPrefixSeptic,
		659: MagicalPrefixEnvenomed,
		660: MagicalPrefixToxic,
		661: MagicalPrefixPestilent,
		662: MagicalPrefixTireless,
		663: MagicalPrefixLizards,
		664: MagicalPrefixAzure,
		665: MagicalPrefixCrimson,
		666: MagicalPrefixTangerine,
		667: MagicalPrefixBeryl,
		668: MagicalPrefixGodly,
		669: MagicalPrefixCruel,
	}
}

func GetMagicalPrefix(id uint16) (result MagicalPrefix, found bool) {
	lookup := getMagicalPrefixMap()

	result, found = lookup[id]
	return result, found
}

func (m MagicalPrefix) GetID() uint16 {
	lookup := getMagicalPrefixMap()

	for key, value := range lookup {
		if value == m {
			return key
		}
	}

	// should not be reached
	panic("d2d2s: (MagicalPrefix).GetID: unknown id")
}

//go:generate stringer -linecomment -type MagicalSuffix -output magical_suffix_string.go

// MagicalSuffix represents a magical suffix
type MagicalSuffix uint16

// magical suffixes - TODO: constant names
const (
	MagicalSuffix1   MagicalSuffix = iota + 1 //   Health
	MagicalSuffix2                            //   Protection
	MagicalSuffix3                            //   Absorption
	MagicalSuffix4                            //   Life
	MagicalSuffix5                            //   (nothing?)
	MagicalSuffix6                            //   Warding
	MagicalSuffix7                            //   the Sentinel
	MagicalSuffix8                            //   Guarding
	MagicalSuffix9                            //   Negation
	MagicalSuffix10                           //  (nothing?)
	MagicalSuffix11                           //  Piercing
	MagicalSuffix12                           //  Bashing
	MagicalSuffix13                           //  Puncturing
	MagicalSuffix14                           //  Thorns
	MagicalSuffix15                           //  Spikes
	MagicalSuffix16                           //  Fleadiness
	MagicalSuffix17                           //  Alacrity
	MagicalSuffix18                           //  Swiitness
	MagicalSuffix19                           //  Quickness
	MagicalSuffix20                           //  Blocking
	MagicalSuffix21                           //  Deilecting
	MagicalSuffix22                           //  the Apprentice
	MagicalSuffix23                           //  the Magus
	MagicalSuffix24                           //  Frost
	MagicalSuffix25                           //  the Glacier
	MagicalSuffix26                           //  Frost
	MagicalSuffix27                           //  Warmth
	MagicalSuffix28                           //  Flame
	MagicalSuffix29                           //  Fire
	MagicalSuffix30                           //  Burning
	MagicalSuffix31                           //  Flame
	MagicalSuffix32                           //  Shook
	MagicalSuffix33                           //  Lightning
	MagicalSuffix34                           //  Thunder
	MagicalSuffix35                           //  Shock
	MagicalSuffix36                           //  Craftsmanship
	MagicalSuffix37                           //  Quality
	MagicalSuffix38                           //  Maiming
	MagicalSuffix39                           //  Slaying
	MagicalSuffix40                           //  Gore
	MagicalSuffix41                           //  Carnage
	MagicalSuffix42                           //  Slaughter
	MagicalSuffix43                           //  Maiming
	MagicalSuffix44                           //  Worth
	MagicalSuffix45                           //  Measure
	MagicalSuffix46                           //  Excellence
	MagicalSuffix47                           //  Petlctmance
	MagicalSuffix48                           //  Measure
	MagicalSuffix49                           //  Blight
	MagicalSuffix50                           //  Venom
	MagicalSuffix51                           //  Pestilence
	MagicalSuffix52                           //  Blight
	MagicalSuffix53                           //  Dextetity
	MagicalSuffix54                           //  Dextetity
	MagicalSuffix55                           //  Skill
	MagicalSuffix56                           //  Skill
	MagicalSuffix57                           //  Accuracy
	MagicalSuffix58                           //  Precision
	MagicalSuffix59                           //  Precision
	MagicalSuffix60                           //  Petlection
	MagicalSuffix61                           //  Balance
	MagicalSuffix62                           //  Stability
	MagicalSuffix63                           //  (nothing?)
	MagicalSuffix64                           //  Regenetation
	MagicalSuffix65                           //  Regenetation
	MagicalSuffix66                           //  Regenetation
	MagicalSuffix67                           //  Regrowth
	MagicalSuffix68                           //  Regrowth
	MagicalSuffix69                           //  Vileness
	MagicalSuffix70                           //  (nothing?)
	MagicalSuffix71                           //  Greed
	MagicalSuffix72                           //  Wealth
	MagicalSuffix73                           //  Chance
	MagicalSuffix74                           //  Fortune
	MagicalSuffix75                           //  Energy
	MagicalSuffix76                           //  Energy
	MagicalSuffix77                           //  the Mind
	MagicalSuffix78                           //  Brilliance
	MagicalSuffix79                           //  Sorcery
	MagicalSuffix80                           //  Wizardry
	MagicalSuffix81                           //  the Beat
	MagicalSuffix82                           //  Light
	MagicalSuffix83                           //  Radiance
	MagicalSuffix84                           //  the Sun
	MagicalSuffix85                           //  Life
	MagicalSuffix86                           //  the Jackal
	MagicalSuffix87                           //  the Fox
	MagicalSuffix88                           //  the Wolf
	MagicalSuffix89                           //  the Wolf
	MagicalSuffix90                           //  the Tiget
	MagicalSuffix91                           //  the Mammoth
	MagicalSuffix92                           //  the Mammoth
	MagicalSuffix93                           //  the Colosuss
	MagicalSuffix94                           //  the Leech
	MagicalSuffix95                           //  the Locust
	MagicalSuffix96                           //  the Bat
	MagicalSuffix97                           //  the Vampire
	MagicalSuffix98                           //  Defiance
	MagicalSuffix99                           //  Amelioration
	MagicalSuffix100                          // Remedy
	MagicalSuffix101                          // (nothing?)
	MagicalSuffix102                          // Simplicity
	MagicalSuffix103                          // Ease
	MagicalSuffix104                          // (nothing?)
	MagicalSuffix105                          // Strength
	MagicalSuffix106                          // Might
	MagicalSuffix107                          // the Ox
	MagicalSuffix108                          // the Ox
	MagicalSuffix109                          // the Giant
	MagicalSuffix110                          // the Giant
	MagicalSuffix111                          // the Titan
	MagicalSuffix112                          // Pacing
	MagicalSuffix113                          // Haste
	MagicalSuffix114                          // Speed
	MagicalSuffix115                          // Health
	MagicalSuffix116                          // Protection
	MagicalSuffix117                          // Absorption
	MagicalSuffix118                          // Life
	MagicalSuffix119                          // Life Everlasting
	MagicalSuffix120                          // Protection
	MagicalSuffix121                          // Absorption
	MagicalSuffix122                          // Life
	MagicalSuffix123                          // Anima
	MagicalSuffix124                          // Warding
	MagicalSuffix125                          // the Sentinel
	MagicalSuffix126                          // Guarding
	MagicalSuffix127                          // Negation
	MagicalSuffix128                          // the Sentinel
	MagicalSuffix129                          // Guarding
	MagicalSuffix130                          // Negation
	MagicalSuffix131                          // Coolness
	MagicalSuffix132                          // Incombustibility
	MagicalSuffix133                          // Amianthus
	MagicalSuffix134                          // Fire Quenching
	MagicalSuffix135                          // Coolness
	MagicalSuffix136                          // Incombustibility
	MagicalSuffix137                          // Amianthus
	MagicalSuffix138                          // Fire Quenching
	MagicalSuffix139                          // Faith
	MagicalSuffix140                          // Resistance
	MagicalSuffix141                          // Insulation
	MagicalSuffix142                          // Grounding
	MagicalSuffix143                          // the Dynamo
	MagicalSuffix144                          // Resistance
	MagicalSuffix145                          // Insulation
	MagicalSuffix146                          // Grounding
	MagicalSuffix147                          // the Dynamo
	MagicalSuffix148                          // Stoicism
	MagicalSuffix149                          // Warming
	MagicalSuffix150                          // Thawing
	MagicalSuffix151                          // the Dunes
	MagicalSuffix152                          // the Sirocco
	MagicalSuffix153                          // Warming
	MagicalSuffix154                          // Thawing
	MagicalSuffix155                          // the Dunes
	MagicalSuffix156                          // the Sirocco
	MagicalSuffix157                          // Desire
	MagicalSuffix158                          // Piercing
	MagicalSuffix159                          // Bashing
	MagicalSuffix160                          // Puncturing
	MagicalSuffix161                          // Thorns
	MagicalSuffix162                          // Spikes
	MagicalSuffix163                          // Razors
	MagicalSuffix164                          // Swords
	MagicalSuffix165                          // Malice
	MagicalSuffix166                          // Readiness
	MagicalSuffix167                          // Alacrity
	MagicalSuffix168                          // Swiftness
	MagicalSuffix169                          // Quickness
	MagicalSuffix170                          // Alacrity
	MagicalSuffix171                          // Fewer
	MagicalSuffix172                          // Blocking
	MagicalSuffix173                          // Deflecting
	MagicalSuffix174                          // the Apprentice
	MagicalSuffix175                          // the Magus
	MagicalSuffix176                          // Frost
	MagicalSuffix177                          // the Icicle
	MagicalSuffix178                          // the Glacier
	MagicalSuffix179                          // Winter
	MagicalSuffix180                          // Frost
	MagicalSuffix181                          // Frigidity
	MagicalSuffix182                          // Warmth
	MagicalSuffix183                          // Flame
	MagicalSuffix184                          // Fire
	MagicalSuffix185                          // Burning
	MagicalSuffix186                          // Incineration
	MagicalSuffix187                          // Flame
	MagicalSuffix188                          // Passion
	MagicalSuffix189                          // Shock
	MagicalSuffix190                          // Lightning
	MagicalSuffix191                          // Thunder
	MagicalSuffix192                          // Storms
	MagicalSuffix193                          // Shock
	MagicalSuffix194                          // Ennui
	MagicalSuffix195                          // Craftsmanship
	MagicalSuffix196                          // Quality
	MagicalSuffix197                          // Maiming
	MagicalSuffix198                          // Slaying
	MagicalSuffix199                          // Gore
	MagicalSuffix200                          // Damage
	MagicalSuffix201                          // Slaughter
	MagicalSuffix202                          // Butchery
	MagicalSuffix203                          // Evisceration
	MagicalSuffix204                          // Maiming
	MagicalSuffix205                          // Craftsmanship
	MagicalSuffix206                          // Craftsmanship
	MagicalSuffix207                          // Craftsmanship
	MagicalSuffix208                          // Quality
	MagicalSuffix209                          // Quality
	MagicalSuffix210                          // Maiming
	MagicalSuffix211                          // Maiming
	MagicalSuffix212                          // Craftsmanship
	MagicalSuffix213                          // Craftsmanship
	MagicalSuffix214                          // Quality
	MagicalSuffix215                          // Quality
	MagicalSuffix216                          // Maiming
	MagicalSuffix217                          // Craftsmanship
	MagicalSuffix218                          // Quality
	MagicalSuffix219                          // Maiming
	MagicalSuffix220                          // Ire
	MagicalSuffix221                          // Wrath
	MagicalSuffix222                          // Damage
	MagicalSuffix223                          // Worth
	MagicalSuffix224                          // Measure
	MagicalSuffix225                          // Excellence
	MagicalSuffix226                          // Performance
	MagicalSuffix227                          // Transcendence
	MagicalSuffix228                          // Worth
	MagicalSuffix229                          // Measure
	MagicalSuffix230                          // Excellence
	MagicalSuffix231                          // Performance
	MagicalSuffix232                          // Joyfulness
	MagicalSuffix233                          // Bliss
	MagicalSuffix234                          // Blight
	MagicalSuffix235                          // Venom
	MagicalSuffix236                          // Pestilence
	MagicalSuffix237                          // Anthrax
	MagicalSuffix238                          // Blight
	MagicalSuffix239                          // Envy
	MagicalSuffix240                          // Dexterity
	MagicalSuffix241                          // Skill
	MagicalSuffix242                          // Accuracy
	MagicalSuffix243                          // Precision
	MagicalSuffix244                          // Perfection
	MagicalSuffix245                          // Nirvana
	MagicalSuffix246                          // Dexterity
	MagicalSuffix247                          // Skill
	MagicalSuffix248                          // Accuracy
	MagicalSuffix249                          // Precision
	MagicalSuffix250                          // Perfection
	MagicalSuffix251                          // Dexterity
	MagicalSuffix252                          // Skill
	MagicalSuffix253                          // Accuracy
	MagicalSuffix254                          // Precision
	MagicalSuffix255                          // Dexterity
	MagicalSuffix256                          // Dexterity
	MagicalSuffix257                          // Dexterity
	MagicalSuffix258                          // Dexterity
	MagicalSuffix259                          // Dexterity
	MagicalSuffix260                          // Dexterity
	MagicalSuffix261                          // Daring
	MagicalSuffix262                          // Balance
	MagicalSuffix263                          // Equilibrium
	MagicalSuffix264                          // Stability
	MagicalSuffix265                          // Balance
	MagicalSuffix266                          // Balance
	MagicalSuffix267                          // Balance
	MagicalSuffix268                          // Truth
	MagicalSuffix269                          // Regeneration
	MagicalSuffix270                          // Regeneration
	MagicalSuffix271                          // Regeneration
	MagicalSuffix272                          // Regrowth
	MagicalSuffix273                          // Regrowth
	MagicalSuffix274                          // Revivification
	MagicalSuffix275                          // Honor
	MagicalSuffix276                          // Vileness
	MagicalSuffix277                          // Greed
	MagicalSuffix278                          // Wealth
	MagicalSuffix279                          // Greed
	MagicalSuffix280                          // Greed
	MagicalSuffix281                          // Greed
	MagicalSuffix282                          // Greed
	MagicalSuffix283                          // Greed
	MagicalSuffix284                          // Greed
	MagicalSuffix285                          // Avarice
	MagicalSuffix286                          // Chance
	MagicalSuffix287                          // Fortune
	MagicalSuffix288                          // Fortune
	MagicalSuffix289                          // Luck
	MagicalSuffix290                          // Fortune
	MagicalSuffix291                          // Good Luck
	MagicalSuffix292                          // Prosperity
	MagicalSuffix293                          // Energy
	MagicalSuffix294                          // the Mind
	MagicalSuffix295                          // Brilliance
	MagicalSuffix296                          // Sorcery
	MagicalSuffix297                          // Wizardry
	MagicalSuffix298                          // Enlightenment
	MagicalSuffix299                          // Energy
	MagicalSuffix300                          // the Mind
	MagicalSuffix301                          // Brilliance
	MagicalSuffix302                          // Sorcery
	MagicalSuffix303                          // Wizardry
	MagicalSuffix304                          // Energy
	MagicalSuffix305                          // the Mind
	MagicalSuffix306                          // Brilliance
	MagicalSuffix307                          // Sorcery
	MagicalSuffix308                          // Knowledge
	MagicalSuffix309                          // the Bear
	MagicalSuffix310                          // Light
	MagicalSuffix311                          // Radiance
	MagicalSuffix312                          // the Sun
	MagicalSuffix313                          // the Jackal
	MagicalSuffix314                          // the Fox
	MagicalSuffix315                          // the Wolf
	MagicalSuffix316                          // the Tiger
	MagicalSuffix317                          // the Mammoth
	MagicalSuffix318                          // the Colosuss
	MagicalSuffix319                          // the Squid
	MagicalSuffix320                          // the Whale
	MagicalSuffix321                          // the Jackal
	MagicalSuffix322                          // the Fox
	MagicalSuffix323                          // the Wolf
	MagicalSuffix324                          // the Tiger
	MagicalSuffix325                          // the Mammoth
	MagicalSuffix326                          // the Colosuss
	MagicalSuffix327                          // the Jackal
	MagicalSuffix328                          // the Fox
	MagicalSuffix329                          // the Wolf
	MagicalSuffix330                          // the Tiger
	MagicalSuffix331                          // the Mammoth
	MagicalSuffix332                          // Life
	MagicalSuffix333                          // Life
	MagicalSuffix334                          // Life
	MagicalSuffix335                          // Substinence
	MagicalSuffix336                          // Substinence
	MagicalSuffix337                          // Substinence
	MagicalSuffix338                          // Vita
	MagicalSuffix339                          // Vita
	MagicalSuffix340                          // Vita
	MagicalSuffix341                          // Life
	MagicalSuffix342                          // Life
	MagicalSuffix343                          // Substinence
	MagicalSuffix344                          // Substinence
	MagicalSuffix345                          // Vita
	MagicalSuffix346                          // Vita
	MagicalSuffix347                          // Life
	MagicalSuffix348                          // Substinence
	MagicalSuffix349                          // Vita
	MagicalSuffix350                          // Spirit
	MagicalSuffix351                          // Hope
	MagicalSuffix352                          // the Leech
	MagicalSuffix353                          // the Locust
	MagicalSuffix354                          // the Lamprey
	MagicalSuffix355                          // the Leech
	MagicalSuffix356                          // the Locust
	MagicalSuffix357                          // the Lamprey
	MagicalSuffix358                          // the Leech
	MagicalSuffix359                          // the Bat
	MagicalSuffix360                          // the Wraith
	MagicalSuffix361                          // the Vampire
	MagicalSuffix362                          // the Bat
	MagicalSuffix363                          // the Wraith
	MagicalSuffix364                          // the Vampire
	MagicalSuffix365                          // the Bat
	MagicalSuffix366                          // Defiance
	MagicalSuffix367                          // Amelioration
	MagicalSuffix368                          // Remedy
	MagicalSuffix369                          // Simplicity
	MagicalSuffix370                          // Ease
	MagicalSuffix371                          // Freedom
	MagicalSuffix372                          // Strength
	MagicalSuffix373                          // Might
	MagicalSuffix374                          // the Ox
	MagicalSuffix375                          // the Giant
	MagicalSuffix376                          // the Titan
	MagicalSuffix377                          // Atlus
	MagicalSuffix378                          // Strength
	MagicalSuffix379                          // Might
	MagicalSuffix380                          // the Us
	MagicalSuffix381                          // the Giant
	MagicalSuffix382                          // the Titan
	MagicalSuffix383                          // Strength
	MagicalSuffix384                          // Might
	MagicalSuffix385                          // the Us
	MagicalSuffix386                          // the Giant
	MagicalSuffix387                          // Strength
	MagicalSuffix388                          // Strength
	MagicalSuffix389                          // Strength
	MagicalSuffix390                          // Strength
	MagicalSuffix391                          // Strength
	MagicalSuffix392                          // Strength
	MagicalSuffix393                          // Virility
	MagicalSuffix394                          // Pacing
	MagicalSuffix395                          // Haste
	MagicalSuffix396                          // Speed
	MagicalSuffix397                          // Traveling
	MagicalSuffix398                          // Acceleration
	MagicalSuffix399                          // Inertia
	MagicalSuffix400                          // Inertia
	MagicalSuffix401                          // Inertia
	MagicalSuffix402                          // Self-Repair
	MagicalSuffix403                          // Fast Repair
	MagicalSuffix404                          // Ages
	MagicalSuffix405                          // Heplenishing
	MagicalSuffix406                          // Propagation
	MagicalSuffix407                          // the Kraken
	MagicalSuffix408                          // Memory
	MagicalSuffix409                          // the Elephant
	MagicalSuffix410                          // Power
	MagicalSuffix411                          // Grace
	MagicalSuffix412                          // Grace and Power
	MagicalSuffix413                          // the Yeti
	MagicalSuffix414                          // the Phoenix
	MagicalSuffix415                          // the Efreeti
	MagicalSuffix416                          // the Cobra
	MagicalSuffix417                          // the Elements
	MagicalSuffix418                          // Firebolts
	MagicalSuffix419                          // Firebolts
	MagicalSuffix420                          // Firebolts
	MagicalSuffix421                          // Charged Shield
	MagicalSuffix422                          // Charged Shield
	MagicalSuffix423                          // Charged Shield
	MagicalSuffix424                          // Icebolt
	MagicalSuffix425                          // Frozen Armor
	MagicalSuffix426                          // Static Field
	MagicalSuffix427                          // Telekinesis
	MagicalSuffix428                          // Frost Shield
	MagicalSuffix429                          // Ice Blast
	MagicalSuffix430                          // Blaze
	MagicalSuffix431                          // Fire Ball
	MagicalSuffix432                          // Nova
	MagicalSuffix433                          // Nova
	MagicalSuffix434                          // Nova Shield
	MagicalSuffix435                          // Nova Shield
	MagicalSuffix436                          // Nova Shield
	MagicalSuffix437                          // Lightning
	MagicalSuffix438                          // Lightning
	MagicalSuffix439                          // Shiver Armor
	MagicalSuffix440                          // Fire Wall
	MagicalSuffix441                          // Enchant
	MagicalSuffix442                          // Chain Lightning
	MagicalSuffix443                          // Chain Lightning
	MagicalSuffix444                          // Chain Lightning
	MagicalSuffix445                          // Teleport Shield
	MagicalSuffix446                          // Teleport Shield
	MagicalSuffix447                          // Teleport Shield
	MagicalSuffix448                          // Glacial Spike
	MagicalSuffix449                          // Meteor
	MagicalSuffix450                          // Thunder Storm
	MagicalSuffix451                          // Energy Shield
	MagicalSuffix452                          // Blizzard
	MagicalSuffix453                          // Chilling Armor
	MagicalSuffix454                          // Hydra Shield
	MagicalSuffix455                          // Frozen ler
	MagicalSuffix456                          // Dawn
	MagicalSuffix457                          // Sunlight
	MagicalSuffix458                          // Magic Arrows
	MagicalSuffix459                          // Magic Arrows
	MagicalSuffix460                          // Fire Arrows
	MagicalSuffix461                          // Fire Arrows
	MagicalSuffix462                          // lnner Sight
	MagicalSuffix463                          // Inner Sight
	MagicalSuffix464                          // Jabbing
	MagicalSuffix465                          // Jabbing
	MagicalSuffix466                          // Cold Arrows
	MagicalSuffix467                          // Cold Arrows
	MagicalSuffix468                          // Multiple Shot
	MagicalSuffix469                          // Multiple Shot
	MagicalSuffix470                          // Power Strike
	MagicalSuffix471                          // Power Strike
	MagicalSuffix472                          // Poison Jab
	MagicalSuffix473                          // Poison Jab
	MagicalSuffix474                          // Exploding Arrows
	MagicalSuffix475                          // Exploding Arrows
	MagicalSuffix476                          // Slow Missiles
	MagicalSuffix477                          // Slow Missiles
	MagicalSuffix478                          // lmpaling Strike
	MagicalSuffix479                          // lmpaling Strike
	MagicalSuffix480                          // Lightning Javelin
	MagicalSuffix481                          // Lightning Javelin
	MagicalSuffix482                          // Ice Arrows
	MagicalSuffix483                          // Ice Arrows
	MagicalSuffix484                          // Guided Arrows
	MagicalSuffix485                          // Guided Arrows
	MagicalSuffix486                          // Charged Strike
	MagicalSuffix487                          // Charged Strike
	MagicalSuffix488                          // Plague Jab
	MagicalSuffix489                          // Plague Jab
	MagicalSuffix490                          // lmmolating Arrows
	MagicalSuffix491                          // lmmolating Arrows
	MagicalSuffix492                          // Fending
	MagicalSuffix493                          // Fending
	MagicalSuffix494                          // Freezing Arrows
	MagicalSuffix495                          // Freezing Arrows
	MagicalSuffix496                          // Lightning Strike
	MagicalSuffix497                          // Lightning Strike
	MagicalSuffix498                          // Lightning Fury
	MagicalSuffix499                          // Lightning Fury
	MagicalSuffix500                          // Fire Bolts
	MagicalSuffix501                          // Fire Bolts
	MagicalSuffix502                          // Charged Bolts
	MagicalSuffix503                          // Charged Bolts
	MagicalSuffix504                          // Ice Bolts
	MagicalSuffix505                          // Ice Bolts
	MagicalSuffix506                          // Frozen Armor
	MagicalSuffix507                          // Frozen Armor
	MagicalSuffix508                          // Static Field
	MagicalSuffix509                          // Static Field
	MagicalSuffix510                          // Telekinesis
	MagicalSuffix511                          // Telekinesis
	MagicalSuffix512                          // Frost Novas
	MagicalSuffix513                          // Frost Novas
	MagicalSuffix514                          // Ice Blasts
	MagicalSuffix515                          // Ice Blasts
	MagicalSuffix516                          // Blazing
	MagicalSuffix517                          // Blazing
	MagicalSuffix518                          // Fire Balls
	MagicalSuffix519                          // Fire Balls
	MagicalSuffix520                          // Novas
	MagicalSuffix521                          // Novas
	MagicalSuffix522                          // Lightning
	MagicalSuffix523                          // Lightning
	MagicalSuffix524                          // Shiver Armor
	MagicalSuffix525                          // Shiver Armor
	MagicalSuffix526                          // Fire Walls
	MagicalSuffix527                          // Fire Walls
	MagicalSuffix528                          // Enchantment
	MagicalSuffix529                          // Enchantment
	MagicalSuffix530                          // Chain Lightning
	MagicalSuffix531                          // Chain Lightning
	MagicalSuffix532                          // Teleportation
	MagicalSuffix533                          // Teleportation
	MagicalSuffix534                          // Glacial Spikes
	MagicalSuffix535                          // Glacial Spikes
	MagicalSuffix536                          // Meteors
	MagicalSuffix537                          // Meteors
	MagicalSuffix538                          // Thunder Storm
	MagicalSuffix539                          // Thunder Storm
	MagicalSuffix540                          // Energy Shield
	MagicalSuffix541                          // Energy Shield
	MagicalSuffix542                          // Blizzards
	MagicalSuffix543                          // Blizzards
	MagicalSuffix544                          // Chilling Armor
	MagicalSuffix545                          // Chilling Armor
	MagicalSuffix546                          // Hydras
	MagicalSuffix547                          // Hydras
	MagicalSuffix548                          // Frozen Orbs
	MagicalSuffix549                          // Frozen Orbs
	MagicalSuffix550                          // Amplify Damage
	MagicalSuffix551                          // Amplify Damage
	MagicalSuffix552                          // Teeth
	MagicalSuffix553                          // Teeth
	MagicalSuffix554                          // Bone Armor
	MagicalSuffix555                          // Bone Armor
	MagicalSuffix556                          // Raise Skeletons
	MagicalSuffix557                          // Raise Skeletons
	MagicalSuffix558                          // Dim Vision
	MagicalSuffix559                          // Dim Vision
	MagicalSuffix560                          // Weaken
	MagicalSuffix561                          // Weaken
	MagicalSuffix562                          // Poison Dagger
	MagicalSuffix563                          // Poison Dagger
	MagicalSuffix564                          // Corpse Explosions
	MagicalSuffix565                          // Corpse Explosions
	MagicalSuffix566                          // Clay Golem Summoning
	MagicalSuffix567                          // Clay Golem Summoning
	MagicalSuffix568                          // Iron Maiden
	MagicalSuffix569                          // Iron Maiden
	MagicalSuffix570                          // Terror
	MagicalSuffix571                          // Terror
	MagicalSuffix572                          // Bone Walls
	MagicalSuffix573                          // Bone Walls
	MagicalSuffix574                          // Raise Skeletal Mages
	MagicalSuffix575                          // Raise Skeletal Mages
	MagicalSuffix576                          // Confusion
	MagicalSuffix577                          // Confusion
	MagicalSuffix578                          // Life Tap
	MagicalSuffix579                          // Life Tap
	MagicalSuffix580                          // Poison Explosion
	MagicalSuffix581                          // Poison Explosion
	MagicalSuffix582                          // Bone Spears
	MagicalSuffix583                          // Bone Spears
	MagicalSuffix584                          // Blood Golem Summoning
	MagicalSuffix585                          // Blood Golem Summoning
	MagicalSuffix586                          // Attraction
	MagicalSuffix587                          // Attraction
	MagicalSuffix588                          // Decrepification
	MagicalSuffix589                          // Decrepification
	MagicalSuffix590                          // Bone Imprisonment
	MagicalSuffix591                          // Bone Imprisonment
	MagicalSuffix592                          // Iron Golem Creation
	MagicalSuffix593                          // Iron Golem Creation
	MagicalSuffix594                          // Lower Resistance
	MagicalSuffix595                          // Lower Flesistance
	MagicalSuffix596                          // Poison Novas
	MagicalSuffix597                          // Poison Novas
	MagicalSuffix598                          // Bone Spirits
	MagicalSuffix599                          // Bone Spirits
	MagicalSuffix600                          // Fire Golem Summoning
	MagicalSuffix601                          // Fire Golem Summoning
	MagicalSuffix602                          // Revivification
	MagicalSuffix603                          // Revivification
	MagicalSuffix604                          // Sacrifice
	MagicalSuffix605                          // Sacrifice
	MagicalSuffix606                          // Holy Bolts
	MagicalSuffix607                          // Holy Bolts
	MagicalSuffix608                          // Zeal
	MagicalSuffix609                          // Zeal
	MagicalSuffix610                          // Vengeance
	MagicalSuffix611                          // Vengeance
	MagicalSuffix612                          // Blessed Hammers
	MagicalSuffix613                          // Blessed Hammers
	MagicalSuffix614                          // Conversion
	MagicalSuffix615                          // Conversion
	MagicalSuffix616                          // Fist of the Heavens
	MagicalSuffix617                          // Fist of the Heavens
	MagicalSuffix618                          // Bashing
	MagicalSuffix619                          // Bashing
	MagicalSuffix620                          // Howling
	MagicalSuffix621                          // Howling
	MagicalSuffix622                          // Potion Finding
	MagicalSuffix623                          // Potion Finding
	MagicalSuffix624                          // Taunting
	MagicalSuffix625                          // Taunting
	MagicalSuffix626                          // Shouting
	MagicalSuffix627                          // Shouting
	MagicalSuffix628                          // Stunning
	MagicalSuffix629                          // Stunning
	MagicalSuffix630                          // Item Finding
	MagicalSuffix631                          // Item Finding
	MagicalSuffix632                          // Concentration
	MagicalSuffix633                          // Concentration
	MagicalSuffix634                          // Battle Cry
	MagicalSuffix635                          // Battle Cry
	MagicalSuffix636                          // Battle Orders
	MagicalSuffix637                          // Battle Orders
	MagicalSuffix638                          // Grim Ward
	MagicalSuffix639                          // Grim Ward
	MagicalSuffix640                          // War Cry
	MagicalSuffix641                          // War Cry
	MagicalSuffix642                          // Battle Command
	MagicalSuffix643                          // Battle Command
	MagicalSuffix644                          // Firestorms
	MagicalSuffix645                          // Firestorms
	MagicalSuffix646                          // Molten Boulders
	MagicalSuffix647                          // Molten Boulders
	MagicalSuffix648                          // Eruption
	MagicalSuffix649                          // Eruption
	MagicalSuffix650                          // Cyclone Armor
	MagicalSuffix651                          // Cyclone Armor
	MagicalSuffix652                          // Twister
	MagicalSuffix653                          // Twister
	MagicalSuffix654                          // Volcano
	MagicalSuffix655                          // Volcano
	MagicalSuffix656                          // Tornado
	MagicalSuffix657                          // Tornado
	MagicalSuffix658                          // Armageddon
	MagicalSuffix659                          // Armageddon
	MagicalSuffix660                          // Hurricane
	MagicalSuffix661                          // Hurricane
	MagicalSuffix662                          // Damage Amplification
	MagicalSuffix663                          // the Icicle
	MagicalSuffix664                          // the Glacier
	MagicalSuffix665                          // Fire
	MagicalSuffix666                          // Burning
	MagicalSuffix667                          // Lightning
	MagicalSuffix668                          // Thunder
	MagicalSuffix669                          // Daring
	MagicalSuffix670                          // Daring
	MagicalSuffix671                          // Knowledge
	MagicalSuffix672                          // Knowledge
	MagicalSuffix673                          // Virility
	MagicalSuffix674                          // Virility
	MagicalSuffix675                          // Readiness
	MagicalSuffix676                          // Craftsmanship
	MagicalSuffix677                          // Quality
	MagicalSuffix678                          // Maiming
	MagicalSuffix679                          // Craftsmanship
	MagicalSuffix680                          // Quality
	MagicalSuffix681                          // Craftsmanship
	MagicalSuffix682                          // Blight
	MagicalSuffix683                          // Venom
	MagicalSuffix684                          // Pestilence
	MagicalSuffix685                          // Anthrax
	MagicalSuffix686                          // Blight
	MagicalSuffix687                          // Venom
	MagicalSuffix688                          // Pestilence
	MagicalSuffix689                          // Anthrax
	MagicalSuffix690                          // Blight
	MagicalSuffix691                          // Venom
	MagicalSuffix692                          // Pestilence
	MagicalSuffix693                          // Anthrax
	MagicalSuffix694                          // Frost
	MagicalSuffix695                          // the Icicle
	MagicalSuffix696                          // the Glacier
	MagicalSuffix697                          // Winter
	MagicalSuffix698                          // Frost
	MagicalSuffix699                          // the Icicle
	MagicalSuffix700                          // the Glacier
	MagicalSuffix701                          // Winter
	MagicalSuffix702                          // Frost
	MagicalSuffix703                          // the Icicle
	MagicalSuffix704                          // the Glacier
	MagicalSuffix705                          // Winter
	MagicalSuffix706                          // Flame
	MagicalSuffix707                          // Fire
	MagicalSuffix708                          // Burning
	MagicalSuffix709                          // Incineration
	MagicalSuffix710                          // Flame
	MagicalSuffix711                          // Fire
	MagicalSuffix712                          // Burning
	MagicalSuffix713                          // Incineration
	MagicalSuffix714                          // Flame
	MagicalSuffix715                          // Fire
	MagicalSuffix716                          // Burning
	MagicalSuffix717                          // Incineration
	MagicalSuffix718                          // Shock
	MagicalSuffix719                          // Lightning
	MagicalSuffix720                          // Thunder
	MagicalSuffix721                          // Storms
	MagicalSuffix722                          // Shock
	MagicalSuffix723                          // Lightning
	MagicalSuffix724                          // Thunder
	MagicalSuffix725                          // Storms
	MagicalSuffix726                          // Shock
	MagicalSuffix727                          // Lightning
	MagicalSuffix728                          // Thunder
	MagicalSuffix729                          // Storms
	MagicalSuffix730                          // Dexterity
	MagicalSuffix731                          // Dexterity
	MagicalSuffix732                          // Strength
	MagicalSuffix733                          // Strength
	MagicalSuffix734                          // Thorns
	MagicalSuffix735                          // Frost
	MagicalSuffix736                          // Flame
	MagicalSuffix737                          // Blight
	MagicalSuffix738                          // Shock
	MagicalSuffix739                          // Regeneration
	MagicalSuffix740                          // Energy
	MagicalSuffix741                          // Light
	MagicalSuffix742                          // the Leech
	MagicalSuffix743                          // the Locust
	MagicalSuffix744                          // the Lamprey
	MagicalSuffix745                          // the Bat
	MagicalSuffix746                          // the Wraith
	MagicalSuffix747                          // the Vampire
)
