// Code generated by "stringer -type SkillID -linecomment -output skill_id_string.go"; DO NOT EDIT.

package d2senums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SkillAttack-0]
	_ = x[SkillKick-1]
	_ = x[SkillThrow-2]
	_ = x[SkillUnsummon-3]
	_ = x[SkillLHThrow-4]
	_ = x[SkillLHSwing-5]
	_ = x[SkillAmazonMagicArrow-6]
	_ = x[SkillAmazonFireArrow-7]
	_ = x[SkillAmazonInnerSight-8]
	_ = x[SkillAmazonCriticalStrike-9]
	_ = x[SkillAmazonJab-10]
	_ = x[SkillAmazonColdArrow-11]
	_ = x[SkillAmazonMultipleShot-12]
	_ = x[SkillAmazonDodge-13]
	_ = x[SkillAmazonPowerStrike-14]
	_ = x[SkillAmazonPoisonJavelin-15]
	_ = x[SkillAmazonExplodingArrow-16]
	_ = x[SkillAmazonSlowMissile-17]
	_ = x[SkillAmazonAvoid-18]
	_ = x[SkillAmazonImpale-19]
	_ = x[SkillAmazonLightningBolt-20]
	_ = x[SkillAmazonIceArrow-21]
	_ = x[SkillAmazonGuidedArrow-22]
	_ = x[SkillAmazonPenetrate-23]
	_ = x[SkillAmazonChargedStrike-24]
	_ = x[SkillAmazonPlagueJavelin-25]
	_ = x[SkillAmazonStrafe-26]
	_ = x[SkillAmazonImmolationArrow-27]
	_ = x[SkillAmazonDopplezon-28]
	_ = x[SkillAmazonEvade-29]
	_ = x[SkillAmazonFend-30]
	_ = x[SkillAmazonFreezingArrow-31]
	_ = x[SkillAmazonValkyrie-32]
	_ = x[SkillAmazonPierce-33]
	_ = x[SkillAmazonLightningStrike-34]
	_ = x[SkillAmazonLightningFury-35]
	_ = x[SkillSorceressFireBolt-36]
	_ = x[SkillSorceressWarmth-37]
	_ = x[SkillSorceressChargedBolt-38]
	_ = x[SkillSorceressIceBolt-39]
	_ = x[SkillSorceressFrozenArmor-40]
	_ = x[SkillSorceressInferno-41]
	_ = x[SkillSorceressStaticField-42]
	_ = x[SkillSorceressTelekinesis-43]
	_ = x[SkillSorceressFrostNova-44]
	_ = x[SkillSorceressIceBlast-45]
	_ = x[SkillSorceressBlaze-46]
	_ = x[SkillSorceressFireBall-47]
	_ = x[SkillSorceressNova-48]
	_ = x[SkillSorceressLightning-49]
	_ = x[SkillSorceressShiverArmor-50]
	_ = x[SkillSorceressFireWall-51]
	_ = x[SkillSorceressEnchant-52]
	_ = x[SkillSorceressChainLightning-53]
	_ = x[SkillSorceressTeleport-54]
	_ = x[SkillSorceressGlacialSpike-55]
	_ = x[SkillSorceressMeteor-56]
	_ = x[SkillSorceressThunderStorm-57]
	_ = x[SkillSorceressEnergyShield-58]
	_ = x[SkillSorceressBlizzard-59]
	_ = x[SkillSorceressChillingArmor-60]
	_ = x[SkillSorceressFireMastery-61]
	_ = x[SkillSorceressHydra-62]
	_ = x[SkillSorceressLightningMastery-63]
	_ = x[SkillSorceressFrozenOrb-64]
	_ = x[SkillSorceressColdMastery-65]
	_ = x[SkillNecromancerAmplifyDamage-66]
	_ = x[SkillNecromancerTeeth-67]
	_ = x[SkillNecromancerBoneArmor-68]
	_ = x[SkillNecromancerSkeletonMastery-69]
	_ = x[SkillNecromancerRaiseSkeleton-70]
	_ = x[SkillNecromancerDimVision-71]
	_ = x[SkillNecromancerWeaken-72]
	_ = x[SkillNecromancerPoisonDagger-73]
	_ = x[SkillNecromancerCorpseExplosion-74]
	_ = x[SkillNecromancerClyGolem-75]
	_ = x[SkillNecromancerIronMaiden-76]
	_ = x[SkillNecromancerTerror-77]
	_ = x[SkillNecromancerBoneWall-78]
	_ = x[SkillNecromancerGolemMastery-79]
	_ = x[SkillNecromancerRaiseSkeletalMag-80]
	_ = x[SkillNecromancerConfuse-81]
	_ = x[SkillNecromancerLifeTap-82]
	_ = x[SkillNecromancerPoisonExplosion-83]
	_ = x[SkillNecromancerBoneSpear-84]
	_ = x[SkillNecromancerBloodGolem-85]
	_ = x[SkillNecromancerAttract-86]
	_ = x[SkillNecromancerDecrepify-87]
	_ = x[SkillNecromancerBonePrison-88]
	_ = x[SkillNecromancerSummonResist-89]
	_ = x[SkillNecromancerIronGolem-90]
	_ = x[SkillNecromancerLowerResist-91]
	_ = x[SkillNecromancerPoisonNova-92]
	_ = x[SkillNecromancerBoneSpirit-93]
	_ = x[SkillNecromancerFireGolem-94]
	_ = x[SkillNecromancerRevive-95]
	_ = x[SkillPaladinSacrifice-96]
	_ = x[SkillPaladinSmite-97]
	_ = x[SkillPaladinMight-98]
	_ = x[SkillPaladinPrayer-99]
	_ = x[SkillPaladinResistFire-100]
	_ = x[SkillPaladinHolyBolt-101]
	_ = x[SkillPaladinHolyFire-102]
	_ = x[SkillPaladinThorne-103]
	_ = x[SkillPaladinDefiance-104]
	_ = x[SkillPaladinResistCold-105]
	_ = x[SkillPaladinZeal-106]
	_ = x[SkillPaladinCharge-107]
	_ = x[SkillPaladinBlassedAim-108]
	_ = x[SkillPaladinConcentration-109]
	_ = x[SkillPaladinHolyFreeze-110]
	_ = x[SkillPaladinVigor-111]
	_ = x[SkillPaladinConversion-112]
	_ = x[SkillPaladinHolyShield-113]
	_ = x[SkillPaladinHolyShock-114]
	_ = x[SkillPaladinSanctuary-115]
	_ = x[SkillPaladinMeditation-116]
	_ = x[SkillPaladinFOH-117]
	_ = x[SkillPaladinFanaticism-118]
	_ = x[SkillPaladinConviction-119]
	_ = x[SkillPaladinRedemption-120]
	_ = x[SkillPaladinSalvation-121]
	_ = x[SkillBarbarianBash-122]
	_ = x[SkillBarbarianSwordMastery-123]
	_ = x[SkillBarbarianAxeMastery-124]
	_ = x[SkillBarbarianMaceMastery-125]
	_ = x[SkillBarbarianHowl-126]
	_ = x[SkillBarbarianFindPotion-127]
	_ = x[SkillBarbarianLeap-128]
	_ = x[SkillBarbarianDoubleSwing-129]
	_ = x[SkillBarbarianPoleArmMastery-130]
	_ = x[SkillBarbarianThrowingMastery-131]
	_ = x[SkillBarbarianSpearMastery-132]
	_ = x[SkillBarbarianTaunt-133]
	_ = x[SkillBarbarianShout-134]
	_ = x[SkillBarbarianStum-135]
	_ = x[SkillBarbarianDoubleThrow-136]
	_ = x[SkillBarbarianIncreasedStamina-137]
	_ = x[SkillBarbarianFintItems-138]
	_ = x[SkillBarbarianLeapAttack-139]
	_ = x[SkillBarbarianConcentrate-140]
	_ = x[SkillBarbarianIronSkin-141]
	_ = x[SkillBarbarianBattleCry-142]
	_ = x[SkillBarbarianFrenzy-143]
	_ = x[SkillBarbarianIncreasedSpeed-144]
	_ = x[SkillBarbarianBattleOrders-145]
	_ = x[SkillBarbarianGrimWard-146]
	_ = x[SkillBarbarianWhirlwind-147]
	_ = x[SkillBarbarianBerserk-148]
	_ = x[SkillBarbarianNaturalResistance-149]
	_ = x[SkillBarbarianWarCry-150]
	_ = x[SkillBarbarianBattleCommands-151]
	_ = x[SkillScrollIdentify-217]
	_ = x[SkillTomeIdentify-218]
	_ = x[SkillScrollTownPortall-219]
	_ = x[SkillTomeTownPortal-220]
	_ = x[SkillDruidRaven-221]
	_ = x[SkillDruidPlaguePoppy-222]
	_ = x[SkillDruidWearwolf-223]
	_ = x[SkillDruidShapeShifting-224]
	_ = x[SkillDruidFirestorm-225]
	_ = x[SkillDruidOakSage-226]
	_ = x[SkillDruidSummonSpriteWolf-227]
	_ = x[SkillDruidWearbear-228]
	_ = x[SkillDruidMoltenBoulder-229]
	_ = x[SkillDruidArcticBlast-230]
	_ = x[SkillDruidCycleLife-231]
	_ = x[SkillDruidFeralRage-232]
	_ = x[SkillDruidMaul-233]
	_ = x[SkillDruidEruption-234]
	_ = x[SkillDruidCycloneArmor-235]
	_ = x[SkillDruidHeartOfWolverine-236]
	_ = x[SkillDruidSummonFenris-237]
	_ = x[SkillDruidRabies-238]
	_ = x[SkillDruidFireClaws-239]
	_ = x[SkillDruidTwister-240]
	_ = x[SkillDruidVines-241]
	_ = x[SkillDruidHunger-242]
	_ = x[SkillDruidShockWave-243]
	_ = x[SkillDruidVolcano-244]
	_ = x[SkillDruidTornado-245]
	_ = x[SkillDruidBarbsSpirit-246]
	_ = x[SkillDruidSummonGrizzly-247]
	_ = x[SkillDruidFury-248]
	_ = x[SkillDruidArmageddon-249]
	_ = x[SkillDruidHurricane-250]
	_ = x[SkillAssasinFireTrauma-251]
	_ = x[SkillAssasinClawMastery-252]
	_ = x[SkillAssasinPsyhicHammer-253]
	_ = x[SkillAssasinTigerStrike-254]
	_ = x[SkillAssasinDragonTalon-255]
	_ = x[SkillAssasinShockField-256]
	_ = x[SkillAssasinBladeSentinel-257]
	_ = x[SkillAssasinQuickness-258]
	_ = x[SkillAssasinFistsOfFire-259]
	_ = x[SkillAssasinDragonClaw-260]
	_ = x[SkillAssasinChargedBolt-261]
	_ = x[SkillAssasinFireWeak-262]
	_ = x[SkillAssasinWeaponBlock-263]
	_ = x[SkillAssasinShadowsCloak-264]
	_ = x[SkillAssasinCobraStrike-265]
	_ = x[SkillAssasinBladeFury-266]
	_ = x[SkillAssasinFade-267]
	_ = x[SkillAssasinShadowWarrior-268]
	_ = x[SkillAssasinClawsOfThunder-269]
	_ = x[SkillAssasinDragonTail-270]
	_ = x[SkillAssasinLightningSentry-271]
	_ = x[SkillAssasinInfernoSentry-272]
	_ = x[SkillAssasinMindBlast-273]
	_ = x[SkillAssasinBladeOfIce-274]
	_ = x[SkillAssasinDragonFight-275]
	_ = x[SkillAssasinDeathSentry-276]
	_ = x[SkillAssasinBladeShield-277]
	_ = x[SkillAssasinVenom-278]
	_ = x[SkillAssasinShadowMastery-279]
	_ = x[SkillAssasinRoyalStrike-280]
}

const (
	_SkillID_name_0 = "attackkickthrowunsummonleft hand throwleft hand swingmagic arrowfire arrowinner sightcritical strikejabcold arrowmultiple shotdodgepower strikepoison javelinexploding arrowslow missileavoidimpalelightning boltice arrowguided arrowpenetratecharged strikeplague javelinstrafeimmolation arrowdopplezonevadefendfreezing arrowvalkyriepiercelightning strikelightning furyfire boltwarmthcharged boltice boltfrozen armorinfernostatic fieldtelekinesisfrost novaice blastblazefire ballnovalightningshiver armorfire wallenchantchain lightningteleportglacial spikemeteorthunder stormenergy shieldblizzardchilling armorfire masteryhydralightning masteryfrozen orbcold masteryamplify damageteethbone armorskeleton masteryraise skeletondim visionweakenpoison daggercorpse explosioncly golemiron maidenterrorbone wallgolem masteryraise skeletal magconfuselife tappoison explosionbone spearblood golemattractdecrepifybone prisonsummon resistiron golemlower lesistpoison novabone spiritfire golemrevivesacrificesmitemightprayerresist fireholy boltholy firethornedefianceresist coldzealchargeblassed aimconcentrationholy freezevigorconversionholy shieldholy shocksanctuarymeditationfist of the Heavensfanaticismconvictionredemptionsalvationbashsword masteryaxe masterymace masteryhowlfind potionleapdouble swingpole arm masterythrowing masteryspear masterytauntshoutstumdouble throwincreased staminafind itemsleap attackconcentrateiron skinbattle cryfrenzyincreased speedbattle ordersgrim wardwhirlwindberserknatural resistancewar crybattle commands"
	_SkillID_name_1 = "scroll of identifytome of identifyscroll of town portaltome of town portalravenplague poppywarwolfshape shiftingfirestormoak sagesummon sprite wolfwearbearmolten boulderarctic blastcycle of lifeferal ragemauleruptioncyclone armorheart of wolverinesummon fenrisrabiesfire clawstwistervineshungershock wavevolcanotornadospirit of barbssummon grizzlyfuryarmageddonhurricanefire traumaclaw masterypsychic hammertiger strikedragon talonshock fieldblade sentinelquicknessfists of firedragon clawcharged bolt sentryweak of fire sentryweapon blockcloak of shadowscobra strikeblade furyfadeshadow warriorclaws of thunderdragon taillightning sentryinferno sentrymint blastblade of icedragon fightdeath sentryblade shieldvenomshadow masteryroyal strike"
)

var (
	_SkillID_index_0 = [...]uint16{0, 6, 10, 15, 23, 38, 53, 64, 74, 85, 100, 103, 113, 126, 131, 143, 157, 172, 184, 189, 195, 209, 218, 230, 239, 253, 267, 273, 289, 298, 303, 307, 321, 329, 335, 351, 365, 374, 380, 392, 400, 412, 419, 431, 442, 452, 461, 466, 475, 479, 488, 500, 509, 516, 531, 539, 552, 558, 571, 584, 592, 606, 618, 623, 640, 650, 662, 676, 681, 691, 707, 721, 731, 737, 750, 766, 775, 786, 792, 801, 814, 832, 839, 847, 863, 873, 884, 891, 900, 911, 924, 934, 946, 957, 968, 978, 984, 993, 998, 1003, 1009, 1020, 1029, 1038, 1044, 1052, 1063, 1067, 1073, 1084, 1097, 1108, 1113, 1123, 1134, 1144, 1153, 1163, 1182, 1192, 1202, 1212, 1221, 1225, 1238, 1249, 1261, 1265, 1276, 1280, 1292, 1308, 1324, 1337, 1342, 1347, 1351, 1363, 1380, 1390, 1401, 1412, 1421, 1431, 1437, 1452, 1465, 1474, 1483, 1490, 1508, 1515, 1530}
	_SkillID_index_1 = [...]uint16{0, 18, 34, 55, 74, 79, 91, 98, 112, 121, 129, 147, 155, 169, 181, 194, 204, 208, 216, 229, 247, 260, 266, 276, 283, 288, 294, 304, 311, 318, 333, 347, 351, 361, 370, 381, 393, 407, 419, 431, 442, 456, 465, 478, 489, 508, 527, 539, 555, 567, 577, 581, 595, 611, 622, 638, 652, 662, 674, 686, 698, 710, 715, 729, 741}
)

func (i SkillID) String() string {
	switch {
	case 0 <= i && i <= 151:
		return _SkillID_name_0[_SkillID_index_0[i]:_SkillID_index_0[i+1]]
	case 217 <= i && i <= 280:
		i -= 217
		return _SkillID_name_1[_SkillID_index_1[i]:_SkillID_index_1[i+1]]
	default:
		return "SkillID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
