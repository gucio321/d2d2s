package itemdata

//go:generate stringer -linecomment -type MagicalSuffix -output magical_suffix_string.go

// MagicalSuffix represents a magical suffix
type MagicalSuffix uint16

// magical suffixes
const (
	MagicalSuffixNone                MagicalSuffix = iota //
	MagicalSuffixProtection                               // Protection
	MagicalSuffixAbsorption                               // Absorption
	MagicalSuffixLife                                     // Life
	MagicalSuffixEmpty                                    // Empty(nothing?)
	MagicalSuffixWarding                                  // Warding
	MagicalSuffixtheSentinel                              // theSentinel
	MagicalSuffixGuarding                                 // Guarding
	MagicalSuffixNegation                                 // Negation
	MagicalSuffixPiercing                                 // Piercing
	MagicalSuffixBashing                                  // Bashing
	MagicalSuffixPuncturing                               // Puncturing
	MagicalSuffixThorns                                   // Thorns
	MagicalSuffixSpikes                                   // Spikes
	MagicalSuffixFleadiness                               // Fleadiness
	MagicalSuffixAlacrity                                 // Alacrity
	MagicalSuffixSwiitness                                // Swiitness
	MagicalSuffixQuickness                                // Quickness
	MagicalSuffixBlocking                                 // Blocking
	MagicalSuffixDeilecting                               // Deilecting
	MagicalSuffixtheApprentice                            // theApprentice
	MagicalSuffixtheMagus                                 // theMagus
	MagicalSuffixFrost                                    // Frost
	MagicalSuffixtheGlacier                               // theGlacier
	MagicalSuffixWarmth                                   // Warmth
	MagicalSuffixFlame                                    // Flame
	MagicalSuffixFire                                     // Fire
	MagicalSuffixBurning                                  // Burning
	MagicalSuffixShook                                    // Shook
	MagicalSuffixLightning                                // Lightning
	MagicalSuffixThunder                                  // Thunder
	MagicalSuffixShock                                    // Shock
	MagicalSuffixCraftsmanship                            // Craftsmanship
	MagicalSuffixQuality                                  // Quality
	MagicalSuffixMaiming                                  // Maiming
	MagicalSuffixSlaying                                  // Slaying
	MagicalSuffixGore                                     // Gore
	MagicalSuffixCarnage                                  // Carnage
	MagicalSuffixSlaughter                                // Slaughter
	MagicalSuffixWorth                                    // Worth
	MagicalSuffixMeasure                                  // Measure
	MagicalSuffixExcellence                               // Excellence
	MagicalSuffixPetlctmance                              // Petlctmance
	MagicalSuffixBlight                                   // Blight
	MagicalSuffixVenom                                    // Venom
	MagicalSuffixPestilence                               // Pestilence
	MagicalSuffixDextetity                                // Dextetity
	MagicalSuffixSkill                                    // Skill
	MagicalSuffixAccuracy                                 // Accuracy
	MagicalSuffixPrecision                                // Precision
	MagicalSuffixPetlection                               // Petlection
	MagicalSuffixBalance                                  // Balance
	MagicalSuffixStability                                // Stability
	MagicalSuffixRegenetation                             // Regenetation
	MagicalSuffixRegrowth                                 // Regrowth
	MagicalSuffixVileness                                 // Vileness
	MagicalSuffixGreed                                    // Greed
	MagicalSuffixWealth                                   // Wealth
	MagicalSuffixChance                                   // Chance
	MagicalSuffixFortune                                  // Fortune
	MagicalSuffixEnergy                                   // Energy
	MagicalSuffixtheMind                                  // theMind
	MagicalSuffixBrilliance                               // Brilliance
	MagicalSuffixSorcery                                  // Sorcery
	MagicalSuffixWizardry                                 // Wizardry
	MagicalSuffixtheBeat                                  // theBeat
	MagicalSuffixLight                                    // Light
	MagicalSuffixRadiance                                 // Radiance
	MagicalSuffixtheSun                                   // theSun
	MagicalSuffixtheJackal                                // theJackal
	MagicalSuffixtheFox                                   // theFox
	MagicalSuffixtheWolf                                  // theWolf
	MagicalSuffixtheTiget                                 // theTiget
	MagicalSuffixtheMammoth                               // theMammoth
	MagicalSuffixtheColosuss                              // theColosuss
	MagicalSuffixtheLeech                                 // theLeech
	MagicalSuffixtheLocust                                // theLocust
	MagicalSuffixtheBat                                   // theBat
	MagicalSuffixtheVampire                               // theVampire
	MagicalSuffixDefiance                                 // Defiance
	MagicalSuffixAmelioration                             // Amelioration
	MagicalSuffixRemedy                                   // Remedy
	MagicalSuffixSimplicity                               // Simplicity
	MagicalSuffixEase                                     // Ease
	MagicalSuffixStrength                                 // Strength
	MagicalSuffixMight                                    // Might
	MagicalSuffixtheOx                                    // theOx
	MagicalSuffixtheGiant                                 // theGiant
	MagicalSuffixtheTitan                                 // theTitan
	MagicalSuffixPacing                                   // Pacing
	MagicalSuffixHaste                                    // Haste
	MagicalSuffixSpeed                                    // Speed
	MagicalSuffixHealth                                   // Health
	MagicalSuffixLifeEverlasting                          // LifeEverlasting
	MagicalSuffixAnima                                    // Anima
	MagicalSuffixCoolness                                 // Coolness
	MagicalSuffixIncombustibility                         // Incombustibility
	MagicalSuffixAmianthus                                // Amianthus
	MagicalSuffixFireQuenching                            // FireQuenching
	MagicalSuffixFaith                                    // Faith
	MagicalSuffixResistance                               // Resistance
	MagicalSuffixInsulation                               // Insulation
	MagicalSuffixGrounding                                // Grounding
	MagicalSuffixtheDynamo                                // theDynamo
	MagicalSuffixStoicism                                 // Stoicism
	MagicalSuffixWarming                                  // Warming
	MagicalSuffixThawing                                  // Thawing
	MagicalSuffixtheDunes                                 // theDunes
	MagicalSuffixtheSirocco                               // theSirocco
	MagicalSuffixDesire                                   // Desire
	MagicalSuffixRazors                                   // Razors
	MagicalSuffixSwords                                   // Swords
	MagicalSuffixMalice                                   // Malice
	MagicalSuffixReadiness                                // Readiness
	MagicalSuffixSwiftness                                // Swiftness
	MagicalSuffixFewer                                    // Fewer
	MagicalSuffixDeflecting                               // Deflecting
	MagicalSuffixtheIcicle                                // theIcicle
	MagicalSuffixWinter                                   // Winter
	MagicalSuffixFrigidity                                // Frigidity
	MagicalSuffixIncineration                             // Incineration
	MagicalSuffixPassion                                  // Passion
	MagicalSuffixStorms                                   // Storms
	MagicalSuffixEnnui                                    // Ennui
	MagicalSuffixDamage                                   // Damage
	MagicalSuffixButchery                                 // Butchery
	MagicalSuffixEvisceration                             // Evisceration
	MagicalSuffixIre                                      // Ire
	MagicalSuffixWrath                                    // Wrath
	MagicalSuffixPerformance                              // Performance
	MagicalSuffixTranscendence                            // Transcendence
	MagicalSuffixJoyfulness                               // Joyfulness
	MagicalSuffixBliss                                    // Bliss
	MagicalSuffixAnthrax                                  // Anthrax
	MagicalSuffixEnvy                                     // Envy
	MagicalSuffixDexterity                                // Dexterity
	MagicalSuffixPerfection                               // Perfection
	MagicalSuffixNirvana                                  // Nirvana
	MagicalSuffixDaring                                   // Daring
	MagicalSuffixEquilibrium                              // Equilibrium
	MagicalSuffixTruth                                    // Truth
	MagicalSuffixRegeneration                             // Regeneration
	MagicalSuffixRevivification                           // Revivification
	MagicalSuffixHonor                                    // Honor
	MagicalSuffixAvarice                                  // Avarice
	MagicalSuffixLuck                                     // Luck
	MagicalSuffixGoodLuck                                 // GoodLuck
	MagicalSuffixProsperity                               // Prosperity
	MagicalSuffixEnlightenment                            // Enlightenment
	MagicalSuffixKnowledge                                // Knowledge
	MagicalSuffixtheBear                                  // theBear
	MagicalSuffixtheTiger                                 // theTiger
	MagicalSuffixtheSquid                                 // theSquid
	MagicalSuffixtheWhale                                 // theWhale
	MagicalSuffixSubstinence                              // Substinence
	MagicalSuffixVita                                     // Vita
	MagicalSuffixSpirit                                   // Spirit
	MagicalSuffixHope                                     // Hope
	MagicalSuffixtheLamprey                               // theLamprey
	MagicalSuffixtheWraith                                // theWraith
	MagicalSuffixFreedom                                  // Freedom
	MagicalSuffixAtlus                                    // Atlus
	MagicalSuffixtheUs                                    // theUs
	MagicalSuffixVirility                                 // Virility
	MagicalSuffixTraveling                                // Traveling
	MagicalSuffixAcceleration                             // Acceleration
	MagicalSuffixInertia                                  // Inertia
	MagicalSuffixSelfRepair                               // Self-Repair
	MagicalSuffixFastRepair                               // FastRepair
	MagicalSuffixAges                                     // Ages
	MagicalSuffixHeplenishing                             // Heplenishing
	MagicalSuffixPropagation                              // Propagation
	MagicalSuffixtheKraken                                // theKraken
	MagicalSuffixMemory                                   // Memory
	MagicalSuffixtheElephant                              // theElephant
	MagicalSuffixPower                                    // Power
	MagicalSuffixGrace                                    // Grace
	MagicalSuffixGraceAndPower                            // GraceandPower
	MagicalSuffixtheYeti                                  // theYeti
	MagicalSuffixthePhoenix                               // thePhoenix
	MagicalSuffixtheEfreeti                               // theEfreeti
	MagicalSuffixtheCobra                                 // theCobra
	MagicalSuffixtheElements                              // theElements
	MagicalSuffixFirebolts                                // Firebolts
	MagicalSuffixChargedShield                            // ChargedShield
	MagicalSuffixIcebolt                                  // Icebolt
	MagicalSuffixFrozenArmor                              // FrozenArmor
	MagicalSuffixStaticField                              // StaticField
	MagicalSuffixTelekinesis                              // Telekinesis
	MagicalSuffixFrostShield                              // FrostShield
	MagicalSuffixIceBlast                                 // IceBlast
	MagicalSuffixBlaze                                    // Blaze
	MagicalSuffixFireBall                                 // FireBall
	MagicalSuffixNova                                     // Nova
	MagicalSuffixNovaShield                               // NovaShield
	MagicalSuffixShiverArmor                              // ShiverArmor
	MagicalSuffixFireWall                                 // FireWall
	MagicalSuffixEnchant                                  // Enchant
	MagicalSuffixChainLightning                           // ChainLightning
	MagicalSuffixTeleportShield                           // TeleportShield
	MagicalSuffixGlacialSpike                             // GlacialSpike
	MagicalSuffixMeteor                                   // Meteor
	MagicalSuffixThunderStorm                             // ThunderStorm
	MagicalSuffixEnergyShield                             // EnergyShield
	MagicalSuffixBlizzard                                 // Blizzard
	MagicalSuffixChillingArmor                            // ChillingArmor
	MagicalSuffixHydraShield                              // HydraShield
	MagicalSuffixFrozenler                                // Frozenler
	MagicalSuffixDawn                                     // Dawn
	MagicalSuffixSunlight                                 // Sunlight
	MagicalSuffixMagicArrows                              // MagicArrows
	MagicalSuffixFireArrows                               // FireArrows
	MagicalSuffixlnnerSight                               // lnnerSight
	MagicalSuffixInnerSight                               // InnerSight
	MagicalSuffixJabbing                                  // Jabbing
	MagicalSuffixColdArrows                               // ColdArrows
	MagicalSuffixMultipleShot                             // MultipleShot
	MagicalSuffixPowerStrike                              // PowerStrike
	MagicalSuffixPoisonJab                                // PoisonJab
	MagicalSuffixExplodingArrows                          // ExplodingArrows
	MagicalSuffixSlowMissiles                             // SlowMissiles
	MagicalSuffixlmpalingStrike                           // lmpalingStrike
	MagicalSuffixLightningJavelin                         // LightningJavelin
	MagicalSuffixIceArrows                                // IceArrows
	MagicalSuffixGuidedArrows                             // GuidedArrows
	MagicalSuffixChargedStrike                            // ChargedStrike
	MagicalSuffixPlagueJab                                // PlagueJab
	MagicalSuffixlmmolatingArrows                         // lmmolatingArrows
	MagicalSuffixFending                                  // Fending
	MagicalSuffixFreezingArrows                           // FreezingArrows
	MagicalSuffixLightningStrike                          // LightningStrike
	MagicalSuffixLightningFury                            // LightningFury
	MagicalSuffixFireBolts                                // FireBolts
	MagicalSuffixChargedBolts                             // ChargedBolts
	MagicalSuffixIceBolts                                 // IceBolts
	MagicalSuffixFrostNovas                               // FrostNovas
	MagicalSuffixIceBlasts                                // IceBlasts
	MagicalSuffixBlazing                                  // Blazing
	MagicalSuffixFireBalls                                // FireBalls
	MagicalSuffixNovas                                    // Novas
	MagicalSuffixFireWalls                                // FireWalls
	MagicalSuffixEnchantment                              // Enchantment
	MagicalSuffixTeleportation                            // Teleportation
	MagicalSuffixGlacialSpikes                            // GlacialSpikes
	MagicalSuffixMeteors                                  // Meteors
	MagicalSuffixBlizzards                                // Blizzards
	MagicalSuffixHydras                                   // Hydras
	MagicalSuffixFrozenOrbs                               // FrozenOrbs
	MagicalSuffixAmplifyDamage                            // AmplifyDamage
	MagicalSuffixTeeth                                    // Teeth
	MagicalSuffixBoneArmor                                // BoneArmor
	MagicalSuffixRaiseSkeletons                           // RaiseSkeletons
	MagicalSuffixDimVision                                // DimVision
	MagicalSuffixWeaken                                   // Weaken
	MagicalSuffixPoisonDagger                             // PoisonDagger
	MagicalSuffixCorpseExplosions                         // CorpseExplosions
	MagicalSuffixClayGolemSummoning                       // ClayGolemSummoning
	MagicalSuffixIronMaiden                               // IronMaiden
	MagicalSuffixTerror                                   // Terror
	MagicalSuffixBoneWalls                                // BoneWalls
	MagicalSuffixRaiseSkeletalMages                       // RaiseSkeletalMages
	MagicalSuffixConfusion                                // Confusion
	MagicalSuffixLifeTap                                  // LifeTap
	MagicalSuffixPoisonExplosion                          // PoisonExplosion
	MagicalSuffixBoneSpears                               // BoneSpears
	MagicalSuffixBloodGolemSummoning                      // BloodGolemSummoning
	MagicalSuffixAttraction                               // Attraction
	MagicalSuffixDecrepification                          // Decrepification
	MagicalSuffixBoneImprisonment                         // BoneImprisonment
	MagicalSuffixIronGolemCreation                        // IronGolemCreation
	MagicalSuffixLowerResistance                          // LowerResistance
	MagicalSuffixLowerFlesistance                         // LowerFlesistance
	MagicalSuffixPoisonNovas                              // PoisonNovas
	MagicalSuffixBoneSpirits                              // BoneSpirits
	MagicalSuffixFireGolemSummoning                       // FireGolemSummoning
	MagicalSuffixSacrifice                                // Sacrifice
	MagicalSuffixHolyBolts                                // HolyBolts
	MagicalSuffixZeal                                     // Zeal
	MagicalSuffixVengeance                                // Vengeance
	MagicalSuffixBlessedHammers                           // BlessedHammers
	MagicalSuffixConversion                               // Conversion
	MagicalSuffixFistOfTheHeavens                         // FistoftheHeavens
	MagicalSuffixHowling                                  // Howling
	MagicalSuffixPotionFinding                            // PotionFinding
	MagicalSuffixTaunting                                 // Taunting
	MagicalSuffixShouting                                 // Shouting
	MagicalSuffixStunning                                 // Stunning
	MagicalSuffixItemFinding                              // ItemFinding
	MagicalSuffixConcentration                            // Concentration
	MagicalSuffixBattleCry                                // BattleCry
	MagicalSuffixBattleOrders                             // BattleOrders
	MagicalSuffixGrimWard                                 // GrimWard
	MagicalSuffixWarCry                                   // WarCry
	MagicalSuffixBattleCommand                            // BattleCommand
	MagicalSuffixFirestorms                               // Firestorms
	MagicalSuffixMoltenBoulders                           // MoltenBoulders
	MagicalSuffixEruption                                 // Eruption
	MagicalSuffixCycloneArmor                             // CycloneArmor
	MagicalSuffixTwister                                  // Twister
	MagicalSuffixVolcano                                  // Volcano
	MagicalSuffixTornado                                  // Tornado
	MagicalSuffixArmageddon                               // Armageddon
	MagicalSuffixHurricane                                // Hurricane
	MagicalSuffixDamageAmplification                      // DamageAmplification
)

// magical suffixes - TODO: constant names
//nolint:funlen // cannot reduce
func getSuffixMap() map[uint16]MagicalSuffix {
	return map[uint16]MagicalSuffix{
		0:   MagicalSuffixNone,
		1:   MagicalSuffixHealth,              //   Health
		2:   MagicalSuffixProtection,          // Protection
		3:   MagicalSuffixAbsorption,          // Absorption
		4:   MagicalSuffixLife,                // Life
		5:   MagicalSuffixEmpty,               // Empty (nothing?)
		6:   MagicalSuffixWarding,             // Warding
		7:   MagicalSuffixtheSentinel,         // the Sentinel
		8:   MagicalSuffixGuarding,            // Guarding
		9:   MagicalSuffixNegation,            // Negation
		10:  MagicalSuffixEmpty,               // Empty (nothing?)
		11:  MagicalSuffixPiercing,            // Piercing
		12:  MagicalSuffixBashing,             // Bashing
		13:  MagicalSuffixPuncturing,          // Puncturing
		14:  MagicalSuffixThorns,              // Thorns
		15:  MagicalSuffixSpikes,              // Spikes
		16:  MagicalSuffixFleadiness,          // Fleadiness
		17:  MagicalSuffixAlacrity,            // Alacrity
		18:  MagicalSuffixSwiitness,           // Swiitness
		19:  MagicalSuffixQuickness,           // Quickness
		20:  MagicalSuffixBlocking,            // Blocking
		21:  MagicalSuffixDeilecting,          // Deilecting
		22:  MagicalSuffixtheApprentice,       // the Apprentice
		23:  MagicalSuffixtheMagus,            // the Magus
		24:  MagicalSuffixFrost,               // Frost
		25:  MagicalSuffixtheGlacier,          // the Glacier
		26:  MagicalSuffixFrost,               // Frost
		27:  MagicalSuffixWarmth,              // Warmth
		28:  MagicalSuffixFlame,               // Flame
		29:  MagicalSuffixFire,                // Fire
		30:  MagicalSuffixBurning,             // Burning
		31:  MagicalSuffixFlame,               // Flame
		32:  MagicalSuffixShook,               // Shook
		33:  MagicalSuffixLightning,           // Lightning
		34:  MagicalSuffixThunder,             // Thunder
		35:  MagicalSuffixShock,               // Shock
		36:  MagicalSuffixCraftsmanship,       // Craftsmanship
		37:  MagicalSuffixQuality,             // Quality
		38:  MagicalSuffixMaiming,             // Maiming
		39:  MagicalSuffixSlaying,             // Slaying
		40:  MagicalSuffixGore,                // Gore
		41:  MagicalSuffixCarnage,             // Carnage
		42:  MagicalSuffixSlaughter,           // Slaughter
		43:  MagicalSuffixMaiming,             // Maiming
		44:  MagicalSuffixWorth,               // Worth
		45:  MagicalSuffixMeasure,             // Measure
		46:  MagicalSuffixExcellence,          // Excellence
		47:  MagicalSuffixPetlctmance,         // Petlctmance
		48:  MagicalSuffixMeasure,             // Measure
		49:  MagicalSuffixBlight,              // Blight
		50:  MagicalSuffixVenom,               // Venom
		51:  MagicalSuffixPestilence,          // Pestilence
		52:  MagicalSuffixBlight,              // Blight
		53:  MagicalSuffixDextetity,           // Dextetity
		54:  MagicalSuffixDextetity,           // Dextetity
		55:  MagicalSuffixSkill,               // Skill
		56:  MagicalSuffixSkill,               // Skill
		57:  MagicalSuffixAccuracy,            // Accuracy
		58:  MagicalSuffixPrecision,           // Precision
		59:  MagicalSuffixPrecision,           // Precision
		60:  MagicalSuffixPetlection,          // Petlection
		61:  MagicalSuffixBalance,             // Balance
		62:  MagicalSuffixStability,           // Stability
		63:  MagicalSuffixEmpty,               // Empty (nothing?)
		64:  MagicalSuffixRegenetation,        // Regenetation
		65:  MagicalSuffixRegenetation,        // Regenetation
		66:  MagicalSuffixRegenetation,        // Regenetation
		67:  MagicalSuffixRegrowth,            // Regrowth
		68:  MagicalSuffixRegrowth,            // Regrowth
		69:  MagicalSuffixVileness,            // Vileness
		70:  MagicalSuffixEmpty,               // Empty (nothing?)
		71:  MagicalSuffixGreed,               // Greed
		72:  MagicalSuffixWealth,              // Wealth
		73:  MagicalSuffixChance,              // Chance
		74:  MagicalSuffixFortune,             // Fortune
		75:  MagicalSuffixEnergy,              // Energy
		76:  MagicalSuffixEnergy,              // Energy
		77:  MagicalSuffixtheMind,             // the Mind
		78:  MagicalSuffixBrilliance,          // Brilliance
		79:  MagicalSuffixSorcery,             // Sorcery
		80:  MagicalSuffixWizardry,            // Wizardry
		81:  MagicalSuffixtheBeat,             // the Beat
		82:  MagicalSuffixLight,               // Light
		83:  MagicalSuffixRadiance,            // Radiance
		84:  MagicalSuffixtheSun,              // the Sun
		85:  MagicalSuffixLife,                // Life
		86:  MagicalSuffixtheJackal,           // the Jackal
		87:  MagicalSuffixtheFox,              // the Fox
		88:  MagicalSuffixtheWolf,             // the Wolf
		89:  MagicalSuffixtheWolf,             // the Wolf
		90:  MagicalSuffixtheTiget,            // the Tiget
		91:  MagicalSuffixtheMammoth,          // the Mammoth
		92:  MagicalSuffixtheMammoth,          // the Mammoth
		93:  MagicalSuffixtheColosuss,         // the Colosuss
		94:  MagicalSuffixtheLeech,            // the Leech
		95:  MagicalSuffixtheLocust,           // the Locust
		96:  MagicalSuffixtheBat,              // the Bat
		97:  MagicalSuffixtheVampire,          // the Vampire
		98:  MagicalSuffixDefiance,            // Defiance
		99:  MagicalSuffixAmelioration,        // Amelioration
		100: MagicalSuffixRemedy,              // Remedy
		101: MagicalSuffixEmpty,               // Empty (nothing?)
		102: MagicalSuffixSimplicity,          // Simplicity
		103: MagicalSuffixEase,                // Ease
		104: MagicalSuffixEmpty,               // Empty (nothing?)
		105: MagicalSuffixStrength,            // Strength
		106: MagicalSuffixMight,               // Might
		107: MagicalSuffixtheOx,               // the Ox
		108: MagicalSuffixtheOx,               // the Ox
		109: MagicalSuffixtheGiant,            // the Giant
		110: MagicalSuffixtheGiant,            // the Giant
		111: MagicalSuffixtheTitan,            // the Titan
		112: MagicalSuffixPacing,              // Pacing
		113: MagicalSuffixHaste,               // Haste
		114: MagicalSuffixSpeed,               // Speed
		115: MagicalSuffixHealth,              // Health
		116: MagicalSuffixProtection,          // Protection
		117: MagicalSuffixAbsorption,          // Absorption
		118: MagicalSuffixLife,                // Life
		119: MagicalSuffixLifeEverlasting,     // Life Everlasting
		120: MagicalSuffixProtection,          // Protection
		121: MagicalSuffixAbsorption,          // Absorption
		122: MagicalSuffixLife,                // Life
		123: MagicalSuffixAnima,               // Anima
		124: MagicalSuffixWarding,             // Warding
		125: MagicalSuffixtheSentinel,         // the Sentinel
		126: MagicalSuffixGuarding,            // Guarding
		127: MagicalSuffixNegation,            // Negation
		128: MagicalSuffixtheSentinel,         // the Sentinel
		129: MagicalSuffixGuarding,            // Guarding
		130: MagicalSuffixNegation,            // Negation
		131: MagicalSuffixCoolness,            // Coolness
		132: MagicalSuffixIncombustibility,    // Incombustibility
		133: MagicalSuffixAmianthus,           // Amianthus
		134: MagicalSuffixFireQuenching,       // Fire Quenching
		135: MagicalSuffixCoolness,            // Coolness
		136: MagicalSuffixIncombustibility,    // Incombustibility
		137: MagicalSuffixAmianthus,           // Amianthus
		138: MagicalSuffixFireQuenching,       // Fire Quenching
		139: MagicalSuffixFaith,               // Faith
		140: MagicalSuffixResistance,          // Resistance
		141: MagicalSuffixInsulation,          // Insulation
		142: MagicalSuffixGrounding,           // Grounding
		143: MagicalSuffixtheDynamo,           // the Dynamo
		144: MagicalSuffixResistance,          // Resistance
		145: MagicalSuffixInsulation,          // Insulation
		146: MagicalSuffixGrounding,           // Grounding
		147: MagicalSuffixtheDynamo,           // the Dynamo
		148: MagicalSuffixStoicism,            // Stoicism
		149: MagicalSuffixWarming,             // Warming
		150: MagicalSuffixThawing,             // Thawing
		151: MagicalSuffixtheDunes,            // the Dunes
		152: MagicalSuffixtheSirocco,          // the Sirocco
		153: MagicalSuffixWarming,             // Warming
		154: MagicalSuffixThawing,             // Thawing
		155: MagicalSuffixtheDunes,            // the Dunes
		156: MagicalSuffixtheSirocco,          // the Sirocco
		157: MagicalSuffixDesire,              // Desire
		158: MagicalSuffixPiercing,            // Piercing
		159: MagicalSuffixBashing,             // Bashing
		160: MagicalSuffixPuncturing,          // Puncturing
		161: MagicalSuffixThorns,              // Thorns
		162: MagicalSuffixSpikes,              // Spikes
		163: MagicalSuffixRazors,              // Razors
		164: MagicalSuffixSwords,              // Swords
		165: MagicalSuffixMalice,              // Malice
		166: MagicalSuffixReadiness,           // Readiness
		167: MagicalSuffixAlacrity,            // Alacrity
		168: MagicalSuffixSwiftness,           // Swiftness
		169: MagicalSuffixQuickness,           // Quickness
		170: MagicalSuffixAlacrity,            // Alacrity
		171: MagicalSuffixFewer,               // Fewer
		172: MagicalSuffixBlocking,            // Blocking
		173: MagicalSuffixDeflecting,          // Deflecting
		174: MagicalSuffixtheApprentice,       // the Apprentice
		175: MagicalSuffixtheMagus,            // the Magus
		176: MagicalSuffixFrost,               // Frost
		177: MagicalSuffixtheIcicle,           // the Icicle
		178: MagicalSuffixtheGlacier,          // the Glacier
		179: MagicalSuffixWinter,              // Winter
		180: MagicalSuffixFrost,               // Frost
		181: MagicalSuffixFrigidity,           // Frigidity
		182: MagicalSuffixWarmth,              // Warmth
		183: MagicalSuffixFlame,               // Flame
		184: MagicalSuffixFire,                // Fire
		185: MagicalSuffixBurning,             // Burning
		186: MagicalSuffixIncineration,        // Incineration
		187: MagicalSuffixFlame,               // Flame
		188: MagicalSuffixPassion,             // Passion
		189: MagicalSuffixShock,               // Shock
		190: MagicalSuffixLightning,           // Lightning
		191: MagicalSuffixThunder,             // Thunder
		192: MagicalSuffixStorms,              // Storms
		193: MagicalSuffixShock,               // Shock
		194: MagicalSuffixEnnui,               // Ennui
		195: MagicalSuffixCraftsmanship,       // Craftsmanship
		196: MagicalSuffixQuality,             // Quality
		197: MagicalSuffixMaiming,             // Maiming
		198: MagicalSuffixSlaying,             // Slaying
		199: MagicalSuffixGore,                // Gore
		200: MagicalSuffixDamage,              // Damage
		201: MagicalSuffixSlaughter,           // Slaughter
		202: MagicalSuffixButchery,            // Butchery
		203: MagicalSuffixEvisceration,        // Evisceration
		204: MagicalSuffixMaiming,             // Maiming
		205: MagicalSuffixCraftsmanship,       // Craftsmanship
		206: MagicalSuffixCraftsmanship,       // Craftsmanship
		207: MagicalSuffixCraftsmanship,       // Craftsmanship
		208: MagicalSuffixQuality,             // Quality
		209: MagicalSuffixQuality,             // Quality
		210: MagicalSuffixMaiming,             // Maiming
		211: MagicalSuffixMaiming,             // Maiming
		212: MagicalSuffixCraftsmanship,       // Craftsmanship
		213: MagicalSuffixCraftsmanship,       // Craftsmanship
		214: MagicalSuffixQuality,             // Quality
		215: MagicalSuffixQuality,             // Quality
		216: MagicalSuffixMaiming,             // Maiming
		217: MagicalSuffixCraftsmanship,       // Craftsmanship
		218: MagicalSuffixQuality,             // Quality
		219: MagicalSuffixMaiming,             // Maiming
		220: MagicalSuffixIre,                 // Ire
		221: MagicalSuffixWrath,               // Wrath
		222: MagicalSuffixDamage,              // Damage
		223: MagicalSuffixWorth,               // Worth
		224: MagicalSuffixMeasure,             // Measure
		225: MagicalSuffixExcellence,          // Excellence
		226: MagicalSuffixPerformance,         // Performance
		227: MagicalSuffixTranscendence,       // Transcendence
		228: MagicalSuffixWorth,               // Worth
		229: MagicalSuffixMeasure,             // Measure
		230: MagicalSuffixExcellence,          // Excellence
		231: MagicalSuffixPerformance,         // Performance
		232: MagicalSuffixJoyfulness,          // Joyfulness
		233: MagicalSuffixBliss,               // Bliss
		234: MagicalSuffixBlight,              // Blight
		235: MagicalSuffixVenom,               // Venom
		236: MagicalSuffixPestilence,          // Pestilence
		237: MagicalSuffixAnthrax,             // Anthrax
		238: MagicalSuffixBlight,              // Blight
		239: MagicalSuffixEnvy,                // Envy
		240: MagicalSuffixDexterity,           // Dexterity
		241: MagicalSuffixSkill,               // Skill
		242: MagicalSuffixAccuracy,            // Accuracy
		243: MagicalSuffixPrecision,           // Precision
		244: MagicalSuffixPerfection,          // Perfection
		245: MagicalSuffixNirvana,             // Nirvana
		246: MagicalSuffixDexterity,           // Dexterity
		247: MagicalSuffixSkill,               // Skill
		248: MagicalSuffixAccuracy,            // Accuracy
		249: MagicalSuffixPrecision,           // Precision
		250: MagicalSuffixPerfection,          // Perfection
		251: MagicalSuffixDexterity,           // Dexterity
		252: MagicalSuffixSkill,               // Skill
		253: MagicalSuffixAccuracy,            // Accuracy
		254: MagicalSuffixPrecision,           // Precision
		255: MagicalSuffixDexterity,           // Dexterity
		256: MagicalSuffixDexterity,           // Dexterity
		257: MagicalSuffixDexterity,           // Dexterity
		258: MagicalSuffixDexterity,           // Dexterity
		259: MagicalSuffixDexterity,           // Dexterity
		260: MagicalSuffixDexterity,           // Dexterity
		261: MagicalSuffixDaring,              // Daring
		262: MagicalSuffixBalance,             // Balance
		263: MagicalSuffixEquilibrium,         // Equilibrium
		264: MagicalSuffixStability,           // Stability
		265: MagicalSuffixBalance,             // Balance
		266: MagicalSuffixBalance,             // Balance
		267: MagicalSuffixBalance,             // Balance
		268: MagicalSuffixTruth,               // Truth
		269: MagicalSuffixRegeneration,        // Regeneration
		270: MagicalSuffixRegeneration,        // Regeneration
		271: MagicalSuffixRegeneration,        // Regeneration
		272: MagicalSuffixRegrowth,            // Regrowth
		273: MagicalSuffixRegrowth,            // Regrowth
		274: MagicalSuffixRevivification,      // Revivification
		275: MagicalSuffixHonor,               // Honor
		276: MagicalSuffixVileness,            // Vileness
		277: MagicalSuffixGreed,               // Greed
		278: MagicalSuffixWealth,              // Wealth
		279: MagicalSuffixGreed,               // Greed
		280: MagicalSuffixGreed,               // Greed
		281: MagicalSuffixGreed,               // Greed
		282: MagicalSuffixGreed,               // Greed
		283: MagicalSuffixGreed,               // Greed
		284: MagicalSuffixGreed,               // Greed
		285: MagicalSuffixAvarice,             // Avarice
		286: MagicalSuffixChance,              // Chance
		287: MagicalSuffixFortune,             // Fortune
		288: MagicalSuffixFortune,             // Fortune
		289: MagicalSuffixLuck,                // Luck
		290: MagicalSuffixFortune,             // Fortune
		291: MagicalSuffixGoodLuck,            // Good Luck
		292: MagicalSuffixProsperity,          // Prosperity
		293: MagicalSuffixEnergy,              // Energy
		294: MagicalSuffixtheMind,             // the Mind
		295: MagicalSuffixBrilliance,          // Brilliance
		296: MagicalSuffixSorcery,             // Sorcery
		297: MagicalSuffixWizardry,            // Wizardry
		298: MagicalSuffixEnlightenment,       // Enlightenment
		299: MagicalSuffixEnergy,              // Energy
		300: MagicalSuffixtheMind,             // the Mind
		301: MagicalSuffixBrilliance,          // Brilliance
		302: MagicalSuffixSorcery,             // Sorcery
		303: MagicalSuffixWizardry,            // Wizardry
		304: MagicalSuffixEnergy,              // Energy
		305: MagicalSuffixtheMind,             // the Mind
		306: MagicalSuffixBrilliance,          // Brilliance
		307: MagicalSuffixSorcery,             // Sorcery
		308: MagicalSuffixKnowledge,           // Knowledge
		309: MagicalSuffixtheBear,             // the Bear
		310: MagicalSuffixLight,               // Light
		311: MagicalSuffixRadiance,            // Radiance
		312: MagicalSuffixtheSun,              // the Sun
		313: MagicalSuffixtheJackal,           // the Jackal
		314: MagicalSuffixtheFox,              // the Fox
		315: MagicalSuffixtheWolf,             // the Wolf
		316: MagicalSuffixtheTiger,            // the Tiger
		317: MagicalSuffixtheMammoth,          // the Mammoth
		318: MagicalSuffixtheColosuss,         // the Colosuss
		319: MagicalSuffixtheSquid,            // the Squid
		320: MagicalSuffixtheWhale,            // the Whale
		321: MagicalSuffixtheJackal,           // the Jackal
		322: MagicalSuffixtheFox,              // the Fox
		323: MagicalSuffixtheWolf,             // the Wolf
		324: MagicalSuffixtheTiger,            // the Tiger
		325: MagicalSuffixtheMammoth,          // the Mammoth
		326: MagicalSuffixtheColosuss,         // the Colosuss
		327: MagicalSuffixtheJackal,           // the Jackal
		328: MagicalSuffixtheFox,              // the Fox
		329: MagicalSuffixtheWolf,             // the Wolf
		330: MagicalSuffixtheTiger,            // the Tiger
		331: MagicalSuffixtheMammoth,          // the Mammoth
		332: MagicalSuffixLife,                // Life
		333: MagicalSuffixLife,                // Life
		334: MagicalSuffixLife,                // Life
		335: MagicalSuffixSubstinence,         // Substinence
		336: MagicalSuffixSubstinence,         // Substinence
		337: MagicalSuffixSubstinence,         // Substinence
		338: MagicalSuffixVita,                // Vita
		339: MagicalSuffixVita,                // Vita
		340: MagicalSuffixVita,                // Vita
		341: MagicalSuffixLife,                // Life
		342: MagicalSuffixLife,                // Life
		343: MagicalSuffixSubstinence,         // Substinence
		344: MagicalSuffixSubstinence,         // Substinence
		345: MagicalSuffixVita,                // Vita
		346: MagicalSuffixVita,                // Vita
		347: MagicalSuffixLife,                // Life
		348: MagicalSuffixSubstinence,         // Substinence
		349: MagicalSuffixVita,                // Vita
		350: MagicalSuffixSpirit,              // Spirit
		351: MagicalSuffixHope,                // Hope
		352: MagicalSuffixtheLeech,            // the Leech
		353: MagicalSuffixtheLocust,           // the Locust
		354: MagicalSuffixtheLamprey,          // the Lamprey
		355: MagicalSuffixtheLeech,            // the Leech
		356: MagicalSuffixtheLocust,           // the Locust
		357: MagicalSuffixtheLamprey,          // the Lamprey
		358: MagicalSuffixtheLeech,            // the Leech
		359: MagicalSuffixtheBat,              // the Bat
		360: MagicalSuffixtheWraith,           // the Wraith
		361: MagicalSuffixtheVampire,          // the Vampire
		362: MagicalSuffixtheBat,              // the Bat
		363: MagicalSuffixtheWraith,           // the Wraith
		364: MagicalSuffixtheVampire,          // the Vampire
		365: MagicalSuffixtheBat,              // the Bat
		366: MagicalSuffixDefiance,            // Defiance
		367: MagicalSuffixAmelioration,        // Amelioration
		368: MagicalSuffixRemedy,              // Remedy
		369: MagicalSuffixSimplicity,          // Simplicity
		370: MagicalSuffixEase,                // Ease
		371: MagicalSuffixFreedom,             // Freedom
		372: MagicalSuffixStrength,            // Strength
		373: MagicalSuffixMight,               // Might
		374: MagicalSuffixtheOx,               // the Ox
		375: MagicalSuffixtheGiant,            // the Giant
		376: MagicalSuffixtheTitan,            // the Titan
		377: MagicalSuffixAtlus,               // Atlus
		378: MagicalSuffixStrength,            // Strength
		379: MagicalSuffixMight,               // Might
		380: MagicalSuffixtheUs,               // the Us
		381: MagicalSuffixtheGiant,            // the Giant
		382: MagicalSuffixtheTitan,            // the Titan
		383: MagicalSuffixStrength,            // Strength
		384: MagicalSuffixMight,               // Might
		385: MagicalSuffixtheUs,               // the Us
		386: MagicalSuffixtheGiant,            // the Giant
		387: MagicalSuffixStrength,            // Strength
		388: MagicalSuffixStrength,            // Strength
		389: MagicalSuffixStrength,            // Strength
		390: MagicalSuffixStrength,            // Strength
		391: MagicalSuffixStrength,            // Strength
		392: MagicalSuffixStrength,            // Strength
		393: MagicalSuffixVirility,            // Virility
		394: MagicalSuffixPacing,              // Pacing
		395: MagicalSuffixHaste,               // Haste
		396: MagicalSuffixSpeed,               // Speed
		397: MagicalSuffixTraveling,           // Traveling
		398: MagicalSuffixAcceleration,        // Acceleration
		399: MagicalSuffixInertia,             // Inertia
		400: MagicalSuffixInertia,             // Inertia
		401: MagicalSuffixInertia,             // Inertia
		402: MagicalSuffixSelfRepair,          // Self-Repair
		403: MagicalSuffixFastRepair,          // Fast Repair
		404: MagicalSuffixAges,                // Ages
		405: MagicalSuffixHeplenishing,        // Heplenishing
		406: MagicalSuffixPropagation,         // Propagation
		407: MagicalSuffixtheKraken,           // the Kraken
		408: MagicalSuffixMemory,              // Memory
		409: MagicalSuffixtheElephant,         // the Elephant
		410: MagicalSuffixPower,               // Power
		411: MagicalSuffixGrace,               // Grace
		412: MagicalSuffixGraceAndPower,       // Grace and Power
		413: MagicalSuffixtheYeti,             // the Yeti
		414: MagicalSuffixthePhoenix,          // the Phoenix
		415: MagicalSuffixtheEfreeti,          // the Efreeti
		416: MagicalSuffixtheCobra,            // the Cobra
		417: MagicalSuffixtheElements,         // the Elements
		418: MagicalSuffixFirebolts,           // Firebolts
		419: MagicalSuffixFirebolts,           // Firebolts
		420: MagicalSuffixFirebolts,           // Firebolts
		421: MagicalSuffixChargedShield,       // Charged Shield
		422: MagicalSuffixChargedShield,       // Charged Shield
		423: MagicalSuffixChargedShield,       // Charged Shield
		424: MagicalSuffixIcebolt,             // Icebolt
		425: MagicalSuffixFrozenArmor,         // Frozen Armor
		426: MagicalSuffixStaticField,         // Static Field
		427: MagicalSuffixTelekinesis,         // Telekinesis
		428: MagicalSuffixFrostShield,         // Frost Shield
		429: MagicalSuffixIceBlast,            // Ice Blast
		430: MagicalSuffixBlaze,               // Blaze
		431: MagicalSuffixFireBall,            // Fire Ball
		432: MagicalSuffixNova,                // Nova
		433: MagicalSuffixNova,                // Nova
		434: MagicalSuffixNovaShield,          // Nova Shield
		435: MagicalSuffixNovaShield,          // Nova Shield
		436: MagicalSuffixNovaShield,          // Nova Shield
		437: MagicalSuffixLightning,           // Lightning
		438: MagicalSuffixLightning,           // Lightning
		439: MagicalSuffixShiverArmor,         // Shiver Armor
		440: MagicalSuffixFireWall,            // Fire Wall
		441: MagicalSuffixEnchant,             // Enchant
		442: MagicalSuffixChainLightning,      // Chain Lightning
		443: MagicalSuffixChainLightning,      // Chain Lightning
		444: MagicalSuffixChainLightning,      // Chain Lightning
		445: MagicalSuffixTeleportShield,      // Teleport Shield
		446: MagicalSuffixTeleportShield,      // Teleport Shield
		447: MagicalSuffixTeleportShield,      // Teleport Shield
		448: MagicalSuffixGlacialSpike,        // Glacial Spike
		449: MagicalSuffixMeteor,              // Meteor
		450: MagicalSuffixThunderStorm,        // Thunder Storm
		451: MagicalSuffixEnergyShield,        // Energy Shield
		452: MagicalSuffixBlizzard,            // Blizzard
		453: MagicalSuffixChillingArmor,       // Chilling Armor
		454: MagicalSuffixHydraShield,         // Hydra Shield
		455: MagicalSuffixFrozenler,           // Frozen ler
		456: MagicalSuffixDawn,                // Dawn
		457: MagicalSuffixSunlight,            // Sunlight
		458: MagicalSuffixMagicArrows,         // Magic Arrows
		459: MagicalSuffixMagicArrows,         // Magic Arrows
		460: MagicalSuffixFireArrows,          // Fire Arrows
		461: MagicalSuffixFireArrows,          // Fire Arrows
		462: MagicalSuffixlnnerSight,          // lnner Sight
		463: MagicalSuffixInnerSight,          // Inner Sight
		464: MagicalSuffixJabbing,             // Jabbing
		465: MagicalSuffixJabbing,             // Jabbing
		466: MagicalSuffixColdArrows,          // Cold Arrows
		467: MagicalSuffixColdArrows,          // Cold Arrows
		468: MagicalSuffixMultipleShot,        // Multiple Shot
		469: MagicalSuffixMultipleShot,        // Multiple Shot
		470: MagicalSuffixPowerStrike,         // Power Strike
		471: MagicalSuffixPowerStrike,         // Power Strike
		472: MagicalSuffixPoisonJab,           // Poison Jab
		473: MagicalSuffixPoisonJab,           // Poison Jab
		474: MagicalSuffixExplodingArrows,     // Exploding Arrows
		475: MagicalSuffixExplodingArrows,     // Exploding Arrows
		476: MagicalSuffixSlowMissiles,        // Slow Missiles
		477: MagicalSuffixSlowMissiles,        // Slow Missiles
		478: MagicalSuffixlmpalingStrike,      // lmpaling Strike
		479: MagicalSuffixlmpalingStrike,      // lmpaling Strike
		480: MagicalSuffixLightningJavelin,    // Lightning Javelin
		481: MagicalSuffixLightningJavelin,    // Lightning Javelin
		482: MagicalSuffixIceArrows,           // Ice Arrows
		483: MagicalSuffixIceArrows,           // Ice Arrows
		484: MagicalSuffixGuidedArrows,        // Guided Arrows
		485: MagicalSuffixGuidedArrows,        // Guided Arrows
		486: MagicalSuffixChargedStrike,       // Charged Strike
		487: MagicalSuffixChargedStrike,       // Charged Strike
		488: MagicalSuffixPlagueJab,           // Plague Jab
		489: MagicalSuffixPlagueJab,           // Plague Jab
		490: MagicalSuffixlmmolatingArrows,    // lmmolating Arrows
		491: MagicalSuffixlmmolatingArrows,    // lmmolating Arrows
		492: MagicalSuffixFending,             // Fending
		493: MagicalSuffixFending,             // Fending
		494: MagicalSuffixFreezingArrows,      // Freezing Arrows
		495: MagicalSuffixFreezingArrows,      // Freezing Arrows
		496: MagicalSuffixLightningStrike,     // Lightning Strike
		497: MagicalSuffixLightningStrike,     // Lightning Strike
		498: MagicalSuffixLightningFury,       // Lightning Fury
		499: MagicalSuffixLightningFury,       // Lightning Fury
		500: MagicalSuffixFireBolts,           // Fire Bolts
		501: MagicalSuffixFireBolts,           // Fire Bolts
		502: MagicalSuffixChargedBolts,        // Charged Bolts
		503: MagicalSuffixChargedBolts,        // Charged Bolts
		504: MagicalSuffixIceBolts,            // Ice Bolts
		505: MagicalSuffixIceBolts,            // Ice Bolts
		506: MagicalSuffixFrozenArmor,         // Frozen Armor
		507: MagicalSuffixFrozenArmor,         // Frozen Armor
		508: MagicalSuffixStaticField,         // Static Field
		509: MagicalSuffixStaticField,         // Static Field
		510: MagicalSuffixTelekinesis,         // Telekinesis
		511: MagicalSuffixTelekinesis,         // Telekinesis
		512: MagicalSuffixFrostNovas,          // Frost Novas
		513: MagicalSuffixFrostNovas,          // Frost Novas
		514: MagicalSuffixIceBlasts,           // Ice Blasts
		515: MagicalSuffixIceBlasts,           // Ice Blasts
		516: MagicalSuffixBlazing,             // Blazing
		517: MagicalSuffixBlazing,             // Blazing
		518: MagicalSuffixFireBalls,           // Fire Balls
		519: MagicalSuffixFireBalls,           // Fire Balls
		520: MagicalSuffixNovas,               // Novas
		521: MagicalSuffixNovas,               // Novas
		522: MagicalSuffixLightning,           // Lightning
		523: MagicalSuffixLightning,           // Lightning
		524: MagicalSuffixShiverArmor,         // Shiver Armor
		525: MagicalSuffixShiverArmor,         // Shiver Armor
		526: MagicalSuffixFireWalls,           // Fire Walls
		527: MagicalSuffixFireWalls,           // Fire Walls
		528: MagicalSuffixEnchantment,         // Enchantment
		529: MagicalSuffixEnchantment,         // Enchantment
		530: MagicalSuffixChainLightning,      // Chain Lightning
		531: MagicalSuffixChainLightning,      // Chain Lightning
		532: MagicalSuffixTeleportation,       // Teleportation
		533: MagicalSuffixTeleportation,       // Teleportation
		534: MagicalSuffixGlacialSpikes,       // Glacial Spikes
		535: MagicalSuffixGlacialSpikes,       // Glacial Spikes
		536: MagicalSuffixMeteors,             // Meteors
		537: MagicalSuffixMeteors,             // Meteors
		538: MagicalSuffixThunderStorm,        // Thunder Storm
		539: MagicalSuffixThunderStorm,        // Thunder Storm
		540: MagicalSuffixEnergyShield,        // Energy Shield
		541: MagicalSuffixEnergyShield,        // Energy Shield
		542: MagicalSuffixBlizzards,           // Blizzards
		543: MagicalSuffixBlizzards,           // Blizzards
		544: MagicalSuffixChillingArmor,       // Chilling Armor
		545: MagicalSuffixChillingArmor,       // Chilling Armor
		546: MagicalSuffixHydras,              // Hydras
		547: MagicalSuffixHydras,              // Hydras
		548: MagicalSuffixFrozenOrbs,          // Frozen Orbs
		549: MagicalSuffixFrozenOrbs,          // Frozen Orbs
		550: MagicalSuffixAmplifyDamage,       // Amplify Damage
		551: MagicalSuffixAmplifyDamage,       // Amplify Damage
		552: MagicalSuffixTeeth,               // Teeth
		553: MagicalSuffixTeeth,               // Teeth
		554: MagicalSuffixBoneArmor,           // Bone Armor
		555: MagicalSuffixBoneArmor,           // Bone Armor
		556: MagicalSuffixRaiseSkeletons,      // Raise Skeletons
		557: MagicalSuffixRaiseSkeletons,      // Raise Skeletons
		558: MagicalSuffixDimVision,           // Dim Vision
		559: MagicalSuffixDimVision,           // Dim Vision
		560: MagicalSuffixWeaken,              // Weaken
		561: MagicalSuffixWeaken,              // Weaken
		562: MagicalSuffixPoisonDagger,        // Poison Dagger
		563: MagicalSuffixPoisonDagger,        // Poison Dagger
		564: MagicalSuffixCorpseExplosions,    // Corpse Explosions
		565: MagicalSuffixCorpseExplosions,    // Corpse Explosions
		566: MagicalSuffixClayGolemSummoning,  // Clay Golem Summoning
		567: MagicalSuffixClayGolemSummoning,  // Clay Golem Summoning
		568: MagicalSuffixIronMaiden,          // Iron Maiden
		569: MagicalSuffixIronMaiden,          // Iron Maiden
		570: MagicalSuffixTerror,              // Terror
		571: MagicalSuffixTerror,              // Terror
		572: MagicalSuffixBoneWalls,           // Bone Walls
		573: MagicalSuffixBoneWalls,           // Bone Walls
		574: MagicalSuffixRaiseSkeletalMages,  // Raise Skeletal Mages
		575: MagicalSuffixRaiseSkeletalMages,  // Raise Skeletal Mages
		576: MagicalSuffixConfusion,           // Confusion
		577: MagicalSuffixConfusion,           // Confusion
		578: MagicalSuffixLifeTap,             // Life Tap
		579: MagicalSuffixLifeTap,             // Life Tap
		580: MagicalSuffixPoisonExplosion,     // Poison Explosion
		581: MagicalSuffixPoisonExplosion,     // Poison Explosion
		582: MagicalSuffixBoneSpears,          // Bone Spears
		583: MagicalSuffixBoneSpears,          // Bone Spears
		584: MagicalSuffixBloodGolemSummoning, // Blood Golem Summoning
		585: MagicalSuffixBloodGolemSummoning, // Blood Golem Summoning
		586: MagicalSuffixAttraction,          // Attraction
		587: MagicalSuffixAttraction,          // Attraction
		588: MagicalSuffixDecrepification,     // Decrepification
		589: MagicalSuffixDecrepification,     // Decrepification
		590: MagicalSuffixBoneImprisonment,    // Bone Imprisonment
		591: MagicalSuffixBoneImprisonment,    // Bone Imprisonment
		592: MagicalSuffixIronGolemCreation,   // Iron Golem Creation
		593: MagicalSuffixIronGolemCreation,   // Iron Golem Creation
		594: MagicalSuffixLowerResistance,     // Lower Resistance
		595: MagicalSuffixLowerFlesistance,    // Lower Flesistance
		596: MagicalSuffixPoisonNovas,         // Poison Novas
		597: MagicalSuffixPoisonNovas,         // Poison Novas
		598: MagicalSuffixBoneSpirits,         // Bone Spirits
		599: MagicalSuffixBoneSpirits,         // Bone Spirits
		600: MagicalSuffixFireGolemSummoning,  // Fire Golem Summoning
		601: MagicalSuffixFireGolemSummoning,  // Fire Golem Summoning
		602: MagicalSuffixRevivification,      // Revivification
		603: MagicalSuffixRevivification,      // Revivification
		604: MagicalSuffixSacrifice,           // Sacrifice
		605: MagicalSuffixSacrifice,           // Sacrifice
		606: MagicalSuffixHolyBolts,           // Holy Bolts
		607: MagicalSuffixHolyBolts,           // Holy Bolts
		608: MagicalSuffixZeal,                // Zeal
		609: MagicalSuffixZeal,                // Zeal
		610: MagicalSuffixVengeance,           // Vengeance
		611: MagicalSuffixVengeance,           // Vengeance
		612: MagicalSuffixBlessedHammers,      // Blessed Hammers
		613: MagicalSuffixBlessedHammers,      // Blessed Hammers
		614: MagicalSuffixConversion,          // Conversion
		615: MagicalSuffixConversion,          // Conversion
		616: MagicalSuffixFistOfTheHeavens,    // Fist of the Heavens
		617: MagicalSuffixFistOfTheHeavens,    // Fist of the Heavens
		618: MagicalSuffixBashing,             // Bashing
		619: MagicalSuffixBashing,             // Bashing
		620: MagicalSuffixHowling,             // Howling
		621: MagicalSuffixHowling,             // Howling
		622: MagicalSuffixPotionFinding,       // Potion Finding
		623: MagicalSuffixPotionFinding,       // Potion Finding
		624: MagicalSuffixTaunting,            // Taunting
		625: MagicalSuffixTaunting,            // Taunting
		626: MagicalSuffixShouting,            // Shouting
		627: MagicalSuffixShouting,            // Shouting
		628: MagicalSuffixStunning,            // Stunning
		629: MagicalSuffixStunning,            // Stunning
		630: MagicalSuffixItemFinding,         // Item Finding
		631: MagicalSuffixItemFinding,         // Item Finding
		632: MagicalSuffixConcentration,       // Concentration
		633: MagicalSuffixConcentration,       // Concentration
		634: MagicalSuffixBattleCry,           // Battle Cry
		635: MagicalSuffixBattleCry,           // Battle Cry
		636: MagicalSuffixBattleOrders,        // Battle Orders
		637: MagicalSuffixBattleOrders,        // Battle Orders
		638: MagicalSuffixGrimWard,            // Grim Ward
		639: MagicalSuffixGrimWard,            // Grim Ward
		640: MagicalSuffixWarCry,              // War Cry
		641: MagicalSuffixWarCry,              // War Cry
		642: MagicalSuffixBattleCommand,       // Battle Command
		643: MagicalSuffixBattleCommand,       // Battle Command
		644: MagicalSuffixFirestorms,          // Firestorms
		645: MagicalSuffixFirestorms,          // Firestorms
		646: MagicalSuffixMoltenBoulders,      // Molten Boulders
		647: MagicalSuffixMoltenBoulders,      // Molten Boulders
		648: MagicalSuffixEruption,            // Eruption
		649: MagicalSuffixEruption,            // Eruption
		650: MagicalSuffixCycloneArmor,        // Cyclone Armor
		651: MagicalSuffixCycloneArmor,        // Cyclone Armor
		652: MagicalSuffixTwister,             // Twister
		653: MagicalSuffixTwister,             // Twister
		654: MagicalSuffixVolcano,             // Volcano
		655: MagicalSuffixVolcano,             // Volcano
		656: MagicalSuffixTornado,             // Tornado
		657: MagicalSuffixTornado,             // Tornado
		658: MagicalSuffixArmageddon,          // Armageddon
		659: MagicalSuffixArmageddon,          // Armageddon
		660: MagicalSuffixHurricane,           // Hurricane
		661: MagicalSuffixHurricane,           // Hurricane
		662: MagicalSuffixDamageAmplification, // Damage Amplification
		663: MagicalSuffixtheIcicle,           // the Icicle
		664: MagicalSuffixtheGlacier,          // the Glacier
		665: MagicalSuffixFire,                // Fire
		666: MagicalSuffixBurning,             // Burning
		667: MagicalSuffixLightning,           // Lightning
		668: MagicalSuffixThunder,             // Thunder
		669: MagicalSuffixDaring,              // Daring
		670: MagicalSuffixDaring,              // Daring
		671: MagicalSuffixKnowledge,           // Knowledge
		672: MagicalSuffixKnowledge,           // Knowledge
		673: MagicalSuffixVirility,            // Virility
		674: MagicalSuffixVirility,            // Virility
		675: MagicalSuffixReadiness,           // Readiness
		676: MagicalSuffixCraftsmanship,       // Craftsmanship
		677: MagicalSuffixQuality,             // Quality
		678: MagicalSuffixMaiming,             // Maiming
		679: MagicalSuffixCraftsmanship,       // Craftsmanship
		680: MagicalSuffixQuality,             // Quality
		681: MagicalSuffixCraftsmanship,       // Craftsmanship
		682: MagicalSuffixBlight,              // Blight
		683: MagicalSuffixVenom,               // Venom
		684: MagicalSuffixPestilence,          // Pestilence
		685: MagicalSuffixAnthrax,             // Anthrax
		686: MagicalSuffixBlight,              // Blight
		687: MagicalSuffixVenom,               // Venom
		688: MagicalSuffixPestilence,          // Pestilence
		689: MagicalSuffixAnthrax,             // Anthrax
		690: MagicalSuffixBlight,              // Blight
		691: MagicalSuffixVenom,               // Venom
		692: MagicalSuffixPestilence,          // Pestilence
		693: MagicalSuffixAnthrax,             // Anthrax
		694: MagicalSuffixFrost,               // Frost
		695: MagicalSuffixtheIcicle,           // the Icicle
		696: MagicalSuffixtheGlacier,          // the Glacier
		697: MagicalSuffixWinter,              // Winter
		698: MagicalSuffixFrost,               // Frost
		699: MagicalSuffixtheIcicle,           // the Icicle
		700: MagicalSuffixtheGlacier,          // the Glacier
		701: MagicalSuffixWinter,              // Winter
		702: MagicalSuffixFrost,               // Frost
		703: MagicalSuffixtheIcicle,           // the Icicle
		704: MagicalSuffixtheGlacier,          // the Glacier
		705: MagicalSuffixWinter,              // Winter
		706: MagicalSuffixFlame,               // Flame
		707: MagicalSuffixFire,                // Fire
		708: MagicalSuffixBurning,             // Burning
		709: MagicalSuffixIncineration,        // Incineration
		710: MagicalSuffixFlame,               // Flame
		711: MagicalSuffixFire,                // Fire
		712: MagicalSuffixBurning,             // Burning
		713: MagicalSuffixIncineration,        // Incineration
		714: MagicalSuffixFlame,               // Flame
		715: MagicalSuffixFire,                // Fire
		716: MagicalSuffixBurning,             // Burning
		717: MagicalSuffixIncineration,        // Incineration
		718: MagicalSuffixShock,               // Shock
		719: MagicalSuffixLightning,           // Lightning
		720: MagicalSuffixThunder,             // Thunder
		721: MagicalSuffixStorms,              // Storms
		722: MagicalSuffixShock,               // Shock
		723: MagicalSuffixLightning,           // Lightning
		724: MagicalSuffixThunder,             // Thunder
		725: MagicalSuffixStorms,              // Storms
		726: MagicalSuffixShock,               // Shock
		727: MagicalSuffixLightning,           // Lightning
		728: MagicalSuffixThunder,             // Thunder
		729: MagicalSuffixStorms,              // Storms
		730: MagicalSuffixDexterity,           // Dexterity
		731: MagicalSuffixDexterity,           // Dexterity
		732: MagicalSuffixStrength,            // Strength
		733: MagicalSuffixStrength,            // Strength
		734: MagicalSuffixThorns,              // Thorns
		735: MagicalSuffixFrost,               // Frost
		736: MagicalSuffixFlame,               // Flame
		737: MagicalSuffixBlight,              // Blight
		738: MagicalSuffixShock,               // Shock
		739: MagicalSuffixRegeneration,        // Regeneration
		740: MagicalSuffixEnergy,              // Energy
		741: MagicalSuffixLight,               // Light
		742: MagicalSuffixtheLeech,            // the Leech
		743: MagicalSuffixtheLocust,           // the Locust
		744: MagicalSuffixtheLamprey,          // the Lamprey
		745: MagicalSuffixtheBat,              // the Bat
		746: MagicalSuffixtheWraith,           // the Wraith
		747: MagicalSuffixtheVampire,          // the Vampire
	}
}

// GetMagicalSuffix returns magical suffix basing on its id
func GetMagicalSuffix(id uint16) (result MagicalSuffix, found bool) {
	lookup := getSuffixMap()
	result, found = lookup[id]

	return result, found
}

// GetID returns magical suffix ID
func (m MagicalSuffix) GetID() uint16 {
	lookup := getSuffixMap()

	for key, value := range lookup {
		if value == m {
			return key
		}
	}

	// should not be reached
	panic("d2d2s: (MagicalSuffix).GetID: unknown id")
}
