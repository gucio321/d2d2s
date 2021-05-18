package d2senums

//go:generate stringer -type SkillID -linecomment -output skill_id_string.go
//go:generate string2enum -type SkillID -linecomment -output skill_id_string2enum.go

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
)

// SkillIDModifier represents an offset for each character class to get a common skill ID
type SkillIDModifier byte

// skillID modifiers
const (
	SkillIDAmazon SkillIDModifier = iota*30 + 6
	SkillIDSorceress
	SkillIDNecromancer
)
