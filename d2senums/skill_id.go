package d2senums

import "log"

// NumSkills is a number of skills for particular character class
const NumSkills = 30

//go:generate stringer -type SkillID -linecomment -output skill_id_string.go
//go:generate string2enum -samepkg -type SkillID -linecomment -output skill_id_string2enum.go

// SkillID represents a skill ID
type SkillID int32

// https://user.xmission.com/~trevin/DiabloIIv1.09_Skills.html
const (
	SkillAttack   SkillID = iota // attack
	SkillKick                    // kick
	SkillThrow                   // throw
	SkillUnsummon                // unsummon
	SkillLHThrow                 // left hand throw
	SkillLHSwing                 // left hand swing

	// amazon skills
	SkillAmazonMagicArrow      // magic arrow
	SkillAmazonFireArrow       // fire arrow
	SkillAmazonInnerSight      // inner sight
	SkillAmazonCriticalStrike  // critical strike
	SkillAmazonJab             // jab
	SkillAmazonColdArrow       // cold arrow
	SkillAmazonMultipleShot    // multiple shot
	SkillAmazonDodge           // dodge
	SkillAmazonPowerStrike     // power strike
	SkillAmazonPoisonJavelin   // poison javelin
	SkillAmazonExplodingArrow  // exploding arrow
	SkillAmazonSlowMissile     // slow missile
	SkillAmazonAvoid           // avoid
	SkillAmazonImpale          // impale
	SkillAmazonLightningBolt   // lightning bolt
	SkillAmazonIceArrow        // ice arrow
	SkillAmazonGuidedArrow     // guided arrow
	SkillAmazonPenetrate       // penetrate
	SkillAmazonChargedStrike   // charged strike
	SkillAmazonPlagueJavelin   // plague javelin
	SkillAmazonStrafe          // strafe
	SkillAmazonImmolationArrow // immolation arrow
	SkillAmazonDopplezon       // dopplezon
	SkillAmazonEvade           // evade
	SkillAmazonFend            // fend
	SkillAmazonFreezingArrow   // freezing arrow
	SkillAmazonValkyrie        // valkyrie
	SkillAmazonPierce          // pierce
	SkillAmazonLightningStrike // lightning strike
	SkillAmazonLightningFury   // lightning fury

	// Sorceress Skills
	SkillSorceressFireBolt         // fire bolt
	SkillSorceressWarmth           // warmth
	SkillSorceressChargedBolt      // charged bolt
	SkillSorceressIceBolt          // ice bolt
	SkillSorceressFrozenArmor      // frozen armor
	SkillSorceressInferno          // inferno
	SkillSorceressStaticField      // static field
	SkillSorceressTelekinesis      // telekinesis
	SkillSorceressFrostNova        // frost nova
	SkillSorceressIceBlast         // ice blast
	SkillSorceressBlaze            // blaze
	SkillSorceressFireBall         // fire ball
	SkillSorceressNova             // nova
	SkillSorceressLightning        // lightning
	SkillSorceressShiverArmor      // shiver armor
	SkillSorceressFireWall         // fire wall
	SkillSorceressEnchant          // enchant
	SkillSorceressChainLightning   // chain lightning
	SkillSorceressTeleport         // teleport
	SkillSorceressGlacialSpike     // glacial spike
	SkillSorceressMeteor           // meteor
	SkillSorceressThunderStorm     // thunder storm
	SkillSorceressEnergyShield     // energy shield
	SkillSorceressBlizzard         // blizzard
	SkillSorceressChillingArmor    // chilling armor
	SkillSorceressFireMastery      // fire mastery
	SkillSorceressHydra            // hydra
	SkillSorceressLightningMastery // lightning mastery
	SkillSorceressFrozenOrb        // frozen orb
	SkillSorceressColdMastery      // cold mastery

	// Necromancer Skills
	SkillNecromancerAmplifyDamage    // amplify damage
	SkillNecromancerTeeth            // teeth
	SkillNecromancerBoneArmor        // bone armor
	SkillNecromancerSkeletonMastery  // skeleton mastery
	SkillNecromancerRaiseSkeleton    // raise skeleton
	SkillNecromancerDimVision        // dim vision
	SkillNecromancerWeaken           // weaken
	SkillNecromancerPoisonDagger     // poison dagger
	SkillNecromancerCorpseExplosion  // corpse explosion
	SkillNecromancerClyGolem         // cly golem
	SkillNecromancerIronMaiden       // iron maiden
	SkillNecromancerTerror           // terror
	SkillNecromancerBoneWall         // bone wall
	SkillNecromancerGolemMastery     // golem mastery
	SkillNecromancerRaiseSkeletalMag // raise skeletal mag
	SkillNecromancerConfuse          // confuse
	SkillNecromancerLifeTap          // life tap
	SkillNecromancerPoisonExplosion  // poison explosion
	SkillNecromancerBoneSpear        // bone spear
	SkillNecromancerBloodGolem       // blood golem
	SkillNecromancerAttract          // attract
	SkillNecromancerDecrepify        // decrepify
	SkillNecromancerBonePrison       // bone prison
	SkillNecromancerSummonResist     // summon resist
	SkillNecromancerIronGolem        // iron golem
	SkillNecromancerLowerResist      // lower lesist
	SkillNecromancerPoisonNova       // poison nova
	SkillNecromancerBoneSpirit       // bone spirit
	SkillNecromancerFireGolem        // fire golem
	SkillNecromancerRevive           // revive

	// Paladin Skills
	SkillPaladinSacrifice     // sacrifice
	SkillPaladinSmite         // smite
	SkillPaladinMight         // might
	SkillPaladinPrayer        // prayer
	SkillPaladinResistFire    // resist fire
	SkillPaladinHolyBolt      // holy bolt
	SkillPaladinHolyFire      // holy fire
	SkillPaladinThorne        // thorne
	SkillPaladinDefiance      // defiance
	SkillPaladinResistCold    // resist cold
	SkillPaladinZeal          // zeal
	SkillPaladinCharge        // charge
	SkillPaladinBlassedAim    // blassed aim
	SkillPaladinConcentration // concentration
	SkillPaladinHolyFreeze    // holy freeze
	SkillPaladinVigor         // vigor
	SkillPaladinConversion    // conversion
	SkillPaladinHolyShield    // holy shield
	SkillPaladinHolyShock     // holy shock
	SkillPaladinSanctuary     // sanctuary
	SkillPaladinMeditation    // meditation
	SkillPaladinFOH           // fist of the Heavens
	SkillPaladinFanaticism    // fanaticism
	SkillPaladinConviction    // conviction
	SkillPaladinRedemption    // redemption
	SkillPaladinSalvation     // salvation

	// Barbarian Skills
	SkillBarbarianBash              // bash
	SkillBarbarianSwordMastery      // sword mastery
	SkillBarbarianAxeMastery        // axe mastery
	SkillBarbarianMaceMastery       // mace mastery
	SkillBarbarianHowl              // howl
	SkillBarbarianFindPotion        // find potion
	SkillBarbarianLeap              // leap
	SkillBarbarianDoubleSwing       // double swing
	SkillBarbarianPoleArmMastery    // pole arm mastery
	SkillBarbarianThrowingMastery   // throwing mastery
	SkillBarbarianSpearMastery      // spear mastery
	SkillBarbarianTaunt             // taunt
	SkillBarbarianShout             // shout
	SkillBarbarianStum              // stum
	SkillBarbarianDoubleThrow       // double throw
	SkillBarbarianIncreasedStamina  // increased stamina
	SkillBarbarianFintItems         // find items
	SkillBarbarianLeapAttack        // leap attack
	SkillBarbarianConcentrate       // concentrate
	SkillBarbarianIronSkin          // iron skin
	SkillBarbarianBattleCry         // battle cry
	SkillBarbarianFrenzy            // frenzy
	SkillBarbarianIncreasedSpeed    // increased speed
	SkillBarbarianBattleOrders      // battle orders
	SkillBarbarianGrimWard          // grim ward
	SkillBarbarianWhirlwind         // whirlwind
	SkillBarbarianBerserk           // berserk
	SkillBarbarianNaturalResistance // natural resistance
	SkillBarbarianWarCry            // war cry
	SkillBarbarianBattleCommands    // battle commands

	// common (scrolls/tomes)
	SkillScrollIdentify    SkillID = 217 + iota // scroll of identify
	SkillTomeIdentify                           // tome of identify
	SkillScrollTownPortall                      // scroll of town portal
	SkillTomeTownPortal                         // tome of town portal

	// Druid Skills
	SkillDruidRaven            // raven
	SkillDruidPlaguePoppy      // plague poppy
	SkillDruidWearwolf         // warwolf
	SkillDruidShapeShifting    // shape shifting
	SkillDruidFirestorm        // firestorm
	SkillDruidOakSage          // oak sage
	SkillDruidSummonSpriteWolf // summon sprite wolf
	SkillDruidWearbear         // wearbear
	SkillDruidMoltenBoulder    // molten boulder
	SkillDruidArcticBlast      // arctic blast
	SkillDruidCycleLife        // cycle of life
	SkillDruidFeralRage        // feral rage
	SkillDruidMaul             // maul
	SkillDruidEruption         // eruption
	SkillDruidCycloneArmor     // cyclone armor
	SkillDruidHeartOfWolverine // heart of wolverine
	SkillDruidSummonFenris     // summon fenris
	SkillDruidRabies           // rabies
	SkillDruidFireClaws        // fire claws
	SkillDruidTwister          // twister
	SkillDruidVines            // vines
	SkillDruidHunger           // hunger
	SkillDruidShockWave        // shock wave
	SkillDruidVolcano          // volcano
	SkillDruidTornado          // tornado
	SkillDruidBarbsSpirit      // spirit of barbs
	SkillDruidSummonGrizzly    // summon grizzly
	SkillDruidFury             // fury
	SkillDruidArmageddon       // armageddon
	SkillDruidHurricane        // hurricane

	// Assassin Skills
	SkillAssasinFireTrauma      // fire trauma
	SkillAssasinClawMastery     // claw mastery
	SkillAssasinPsyhicHammer    // psychic hammer
	SkillAssasinTigerStrike     // tiger strike
	SkillAssasinDragonTalon     // dragon talon
	SkillAssasinShockField      // shock field
	SkillAssasinBladeSentinel   // blade sentinel
	SkillAssasinQuickness       // quickness
	SkillAssasinFistsOfFire     // fists of fire
	SkillAssasinDragonClaw      // dragon claw
	SkillAssasinChargedBolt     // charged bolt sentry
	SkillAssasinFireWeak        // weak of fire sentry
	SkillAssasinWeaponBlock     // weapon block
	SkillAssasinShadowsCloak    // cloak of shadows
	SkillAssasinCobraStrike     // cobra strike
	SkillAssasinBladeFury       // blade fury
	SkillAssasinFade            // fade
	SkillAssasinShadowWarrior   // shadow warrior
	SkillAssasinClawsOfThunder  // claws of thunder
	SkillAssasinDragonTail      // dragon tail
	SkillAssasinLightningSentry // lightning sentry
	SkillAssasinInfernoSentry   // inferno sentry
	SkillAssasinMindBlast       // mint blast
	SkillAssasinBladeOfIce      // blade of ice
	SkillAssasinDragonFight     // dragon fight
	SkillAssasinDeathSentry     // death sentry
	SkillAssasinBladeShield     // blade shield
	SkillAssasinVenom           // venom
	SkillAssasinShadowMastery   // shadow mastery
	SkillAssasinRoyalStrike     // royal strike
)

// SkillIDModifier represents an offset for each character class to get a common skill ID
type SkillIDModifier byte

// skillID modifiers
const (
	SkillIDAmazon SkillIDModifier = iota*30 + 6
	SkillIDSorceress
	SkillIDNecromancer
	SkillIDPaladin
	SkillIDBarbarian
)

// skillID modifiers: part 2
const (
	SkillIDDruid SkillIDModifier = iota*30 + 221
	SkillIDAssasin
)

// GetSkillModifier returns a skill id modifier basing on character class
func GetSkillModifier(class CharacterClass) SkillIDModifier {
	lookup := map[CharacterClass]SkillIDModifier{
		CharacterClassAmazon:      SkillIDAmazon,
		CharacterClassSorceress:   SkillIDSorceress,
		CharacterClassNecromancer: SkillIDNecromancer,
		CharacterClassPaladin:     SkillIDPaladin,
		CharacterClassBarbarian:   SkillIDBarbarian,
		CharacterClassDruid:       SkillIDDruid,
		CharacterClassAssassin:    SkillIDAssasin,
	}

	m, ok := lookup[class]
	if !ok {
		log.Panicf("unexpected character class %v", class)
	}

	return m
}

// GetSkillList returns a list of skills for specified character class
func GetSkillList(c CharacterClass) (result [NumSkills]SkillID) {
	mod := GetSkillModifier(c)
	for i := range result {
		result[i] = SkillID(int(mod) + i)
	}

	return result
}
