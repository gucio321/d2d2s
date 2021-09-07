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
