package itemdata

//go:generate stringer -linecomment -type UniqueID -output unique_string.go

// UniqueID represents an uinque item's id
type UniqueID uint16

// unique IDs TODO: name these constants
const (
	UniqueTheGnasher           UniqueID = iota // The Gnasher
	UniqueDeathspade                           // Deathspade
	UniqueBladebone                            // Bladebone
	UniqueSkullSplitter                        // Skull splitter
	UniqueRakescar                             // Rakescar
	UniqueAxeOfFechmar                         // Axe of Fechmar
	UniqueGoreshovel                           // Goreshovel
	UniqueTheChiefthan                         // The Chiefthan
	UniqueBrainhew                             // Brainhew
	UniqueHumongous                            // Humongous
	UniqueTorchOfIros                          // Torch of Iros
	UniqueMaelstorm                            // Maelstorm
	UniqueGravenspine                          // Gravenspine
	UniqueUmesLament                           // Umes Lament
	UniqueFelloak                              // Felloak
	UniqueKnellStriker                         // Knell Striker
	UniqueRusthandle                           // Rusthandle
	UniqueStormeye                             // Stormeye
	UniqueStoutnail                            // Stoutnail
	UniqueCrushflange                          // Crushflange
	UniqueBloodrise                            // Bloodrise
	UniqueTheGeneralsTanDoLiGa                 // The Generals Tan Do Li Ga
	UniqueIronstone                            // Ironstone
	UniqueBonesnap                             // Bonesnap
	UniqueSteeldriver                          // Steeldriver
	UniqueRixotsKeen                           // Rixot's Keen
	UniqueBloodCrescent                        // Blood Crescent
	UniqueSkewerOfKrintiz                      // Skewer of Krintiz
	UniqueGleamscythe                          // Gleamscythe
	UniqueAzurewrath1                          // Azurewrath
	UniqueGriswoldsEdge                        // Griswold's Edge
	UniqueHellplague                           // Hellplague
	UniqueCulwensPoint                         // Culwens Point
	UniqueShadowfang                           // Shadowfang
	UniqueSoulflay                             // Soulflay
	UniqueKinemilsAwl                          // Kinemils Awl
	UniqueBlacktongue                          // Blacktongue
	UniqueRipsaw                               // Ripsaw
	UniqueThePatriarch                         // The Patriarch
	UniqueGull                                 // Gull
	UniqueTheDiggler                           // The Diggler
	UniqueTheJadeTanDo                         // The Jade Tan Do
	UniqueSpectralShard                        // Spectral Shard
	UniqueTheDragonChang                       // The Dragon Chang
	UniqueRazortine                            // Razortine
	UniqueBloodthief                           // Bloodthief
	UniqueLanceOfYaggai                        // Lance of Yaggai
	UniqueTheTannrGorerod                      // The Tannr Gorerod
	UniqueDimoaksHew                           // Dimoaks Hew
	UniqueSteelgoad                            // Steelgoad
	UniqueSoulHarvest                          // Soul Harvest
	UniqueTheBattlebranch                      // The Battlebranch
	UniqueWoestave                             // Woestave
	UniqueTheGrimReaper                        // The Grim Reaper
	UniqueBaneAsh                              // Bane Ash
	UniqueSerpentLord                          // Serpent Lord
	UniqueSpireOfLazarus                       // Spire of Lazarus
	UniqueTheSalamander                        // The Salamander
	UniqueTheIronJangBong                      // The Iron Jang Bong
	UniquePluckeye                             // Pluckeye
	UniqueWitherstring                         // Witherstring
	UniqueRavenClaw                            // Raven Claw
	UniqueRoguesBow                            // Rogue's Bow
	UniqueStormstrike                          // Stormstrike
	UniqueWizendraw                            // Wizendraw
	UniqueHellclap                             // Hellclap
	UniqueBlastbark                            // Blastbark
	UniqueLeadcrow                             // Leadcrow
	UniqueIchorsting                           // Ichorsting
	UniqueHellcast                             // Hellcast
	UniqueDoomslinger                          // Doomslinger
	UniqueBigginsBonnet                        // Biggin's Bonnet
	UniqueTarnhelm                             // Tarnhelm
	UniqueCoifOfGlory                          // Coif of Glory
	UniqueDuskdeep                             // Duskdeep
	UniqueWormskull                            // Wormskull
	UniqueHowltusk                             // Howltusk
	UniqueUndeadCrown                          // Undead Crown
	UniqueTheFaceOfHorror                      // The Face of Horror
	UniqueGreyform                             // Greyform
	UniqueBlinkbatsForm                        // Blinkbat's Form
	UniqueTheCenturion                         // The Centurion
	UniqueTwitchthroe                          // Twitchthroe
	UniqueDarkglow                             // Darkglow
	UniqueHawkmail                             // Hawkmail
	UniqueSparkingMail                         // Sparking Mail
	UniqueVenomWard                            // Venom Ward
	UniqueIceblink                             // Iceblink
	UniqueBoneflesh                            // Boneflesh
	UniqueRockfleece                           // Rockfleece
	UniqueRattlecage                           // Rattlecage
	UniqueGoldskin                             // Goldskin
	UniqueVictorsSilk                          // Victors Silk
	UniqueHeavenlyGarb                         // Heavenly Garb
	UniquePeltaLunata                          // Pelta Lunata
	UniqueUmbralDisk                           // Umbral Disk
	UniqueStormguild                           // Stormguild
	UniqueWallOfTheEyeless                     // Wall of the Eyeless
	UniqueSwordbackHold                        // Swordback Hold
	UniqueSteelclash                           // Steelclash
	UniqueBverritKeep                          // Bverrit Keep
	UniqueTheWard                              // The Ward
	UniqueTheHandOfBroc                        // The Hand of Broc
	UniqueBloodfist                            // Bloodfist
	UniqueChanceGuards                         // Chance Guards
	UniqueMagefist                             // Magefist
	UniqueFrostburn                            // Frostburn
	UniqueHotspur                              // Hotspur
	UniqueGorefoot                             // Gorefoot
	UniqueTreadsOfCthon                        // Treads of Cthon
	UniqueGoblinToe                            // Goblin Toe
	UniqueTearhaunch                           // Tearhaunch
	UniqueLenymo                               // Lenymo
	UniqueSnakecord                            // Snakecord
	UniqueNightsmoke                           // Nightsmoke
	UniqueGoldwrap                             // Goldwrap
	UniqueBladebuckle                          // Bladebuckle
	UniqueNokozanRelic                         // Nokozan Relic
	UniqueTheEyeOfEtlich                       // The Eye of Etlich
	UniqueTheMahimOakCurio                     // The Mahim-Oak Curio
	UniqueNagelring                            // Nagelring
	UniqueManaldHeal                           // Manald Heal
	UniqueTheStoneofJordan                     // The Stone of Jordan
	UniqueAmuletOfTheViper                     // Amulet of the Viper
	UniqueStaffOfKings                         // Staff of Kings
	UniqueHoradricStaff                        // Horadric Staff
	UniqueHellForgeHammer                      // Hell Forge Hammer
	UniqueKhalimsFlail                         // Khalim's Flail
	UniqueSuperKhalimsFlail                    // Super Khalim's Flail
	UniqueColdkill                             // Coldkill
	UniqueButchersPupil                        // Butcher's Pupil
	UniqueIslestrike                           // Islestrike
	UniquePompesWrath                          // Pompe's Wrath
	UniqueGuardianNaga                         // Guardian Naga
	UniqueWarlordsTrust                        // Warlord's Trust
	UniqueSpellsteel                           // Spellsteel
	UniqueStormrider                           // Stormrider
	UniqueBoneslayerBlade                      // Boneslayer Blade
	UniqueTheMinataur                          // The Minataur
	UniqueSuicideBranch                        // Suicide Branch
	UniqueCarinShard                           // Carin Shard
	UniqueArmOfKingLeoric                      // Arm of King Leoric
	UniqueBlackhandKey                         // Blackhand Key
	UniqueDarkClanCrusher                      // Dark Clan Crusher
	UniqueZakarumsHand                         // Zakarum's Hand
	UniqueTheFetidSprinkler                    // The Fetid Sprinkler
	UniqueHandOfBlessedLight                   // Hand of Blessed Light
	UniqueFleshrender                          // Fleshrender
	UniqueSureshrillFrost                      // Sureshrill Frost
	UniqueMoonfall                             // Moonfall
	UniqueBaezilsVortex                        // Baezil's Vortex
	UniqueEarthshaker                          // Earthshaker
	UniqueBloodtreeStump                       // Bloodtree Stump
	UniqueTheGavelOfPain                       // The Gavel of Pain
	UniqueBloodletter                          // Bloodletter
	UniqueColdsteelEye                         // Coldsteel Eye
	UniqueHexfire                              // Hexfire
	UniqueBladeOfAliBaba                       // Blade of Ali Baba
	UniqueGinthersRift                         // Ginther's Rift
	UniqueHeadstriker                          // Headstriker
	UniquePlagueBearer                         // Plague Bearer
	UniqueTheAtlantian                         // The Atlantian
	UniqueCrainteVomir                         // Crainte Vomir
	UniqueBingSzWang                           // Bing Sz Wang
	UniqueTheVileHusk                          // The Vile Husk
	UniqueCloudcrack                           // Cloudcrack
	UniqueTodesfaelleFlamme                    // Todesfaelle Flamme
	UniqueSwordguard                           // Swordguard
	UniqueSpineripper                          // Spineripper
	UniqueHeartCarver                          // Heart Carver
	UniqueBlackbogsSharp                       // Blackbog's Sharp
	UniqueStormspike                           // Stormspike
	UniqueTheImpaler                           // The Impaler
	UniqueKelpieSnare                          // Kelpie Snare
	UniqueSoulfeastTine                        // Soulfeast Tine
	UniqueHoneSundan                           // Hone Sundan
	UniqueSpireOfHonor                         // Spire of Honor
	UniqueTheMeatScraper                       // The Meat Scraper
	UniqueBlackleachBlade                      // Blackleach Blade
	UniqueAthenasWrath                         // Athena's Wrath
	UniquePierreTombaleCouant                  // Pierre Tombale Couant
	UniqueHusoldalEvo                          // Husoldal Evo
	UniqueGrimsBurningDead                     // Grim's Burning Dead
	UniqueRazorswitch                          // Razorswitch
	UniqueRibcracker                           // Ribcracker
	UniqueChromaticIre                         // Chromatic Ire
	UniqueWarpspear                            // Warpspear
	UniqueSkullcollector                       // Skullcollector
	UniqueSkystrike                            // Skystrike
	UniqueRiphook                              // Riphook
	UniqueKukoShakaku                          // Kuko Shakaku
	UniqueEndlesshail                          // Endlesshail
	UniqueWhichwildString                      // Whichwild String
	UniqueCliffkiller                          // Cliffkiller
	UniqueMagewrath                            // Magewrath
	UniqueGodstrikeArch                        // Godstrike Arch
	UniqueLangerBriser                         // Langer Briser
	UniquePusSpiter                            // Pus Spiter
	UniqueBurizaDoKyanon                       // Buriza-Do Kyanon
	UniqueDemonMachine                         // Demon Machine
	UniqueUnknownArmor                         // Armor (Unknown)
	UniquePeasentCrown                         // Peasent Crown
	UniqueRockstopper                          // Rockstopper
	UniqueStealskull                           // Stealskull
	UniqueDarksightHelm                        // Darksight Helm
	UniqueValkyrieWing                         // Valkyrie Wing
	UniqueCrownOfThieves                       // Crown of Thieves
	UniqueBlckhornsFace                        // Blckhorn's Face
	UniqueVampireGaze                          // Vampire Gaze
	UniqueTheSpiritShroud                      // The Spirit Shroud
	UniqueSkinOftheVipermagi                   // Skin of the Vipermagi
	UniqueSkinOfTheFlayedOne                   // Skin of the Flayed One
	UniqueIronpelt                             // Ironpelt
	UniqueSpiritforge                          // Spiritforge
	UniqueCrowCaw                              // Crow Caw
	UniqueShaftstop                            // Shaftstop
	UniqueDurielsShell                         // Duriel's Shell
	UniqueSkulldersIre                         // Skullder's Ire
	UniqueGuardianAngel                        // Guardian Angel
	UniqueToothrow                             // Toothrow
	UniqueAtmasWail                            // Atma's Wail
	UniqueBlackHades                           // Black Hades
	UniqueCorpsemourn                          // Corpsemourn
	UniqueQueHegansWisdom                      // Que-Hegan's Wisdom
	UniqueVisceratuant                         // Visceratuant
	UniqueMosersBlessedCircle                  // Mosers Blessed Circle
	UniqueStormchaser                          // Stormchaser
	UniqueTiamatsRebuke                        // Tiamat's Rebuke
	UniqueGerkesSanctuary                      // Gerke's Sanctuary
	UniqueRadimantsSphere                      // Radimant's Sphere
	UniqueLidlessWall                          // Lidless Wall
	UniqueLanceGuard                           // Lance Guard
	UniqueVenomGrip                            // Venom Grip
	UniqueGravepalm                            // Gravepalm
	UniqueGhoulhide                            // Ghoulhide
	UniqueLavagout                             // Lavagout
	UniqueHellmouth                            // Hellmouth
	UniqueInfernostride                        // Infernostride
	UniqueWaterwalk                            // Waterwalk
	UniqueSilkweave                            // Silkweave
	UniqueWartraveler                          // Wartraveler
	UniqueGorerider                            // Gorerider
	UniqueStringOfEars                         // String of Ears
	UniqueRazortail                            // Razortail
	UniqueGloomstrap                           // Gloomstrap
	UniqueSnowclash                            // Snowclash
	UniqueThundergodsVigor                     // Thundergod's Vigor
	UniqueEliteunique                          // Elite unique
	UniqueHarlequinCrest                       // Harlequin Crest
	UniqueVeilOfSteel                          // Veil of Steel
	UniqueTheGladiatorsBane                    // The Gladiator's Bane
	UniqueArkainesValor                        // Arkaine's Valor
	UniqueBlackoakShield                       // Blackoak Shield
	UniqueStormshield                          // Stormshield
	UniqueHellslayer                           // Hellslayer
	UniqueMesserschmidtsReaver                 // Messerschmidt's Reaver
	UniqueBaranarsStar                         // Baranar's Star
	UniqueSchaefersHammer                      // Schaefer's Hammer
	UniqueTheCraniumBasher                     // The Cranium Basher
	UniqueLightsabre                           // Lightsabre
	UniqueDoombringer                          // Doombringer
	UniqueTheGrandfather                       // The Grandfather
	UniqueWizardspike                          // Wizardspike
	UniqueConstrictingRing                     // Constricting Ring
	UniqueStormspire                           // Stormspire
	UniqueEaglehorn                            // Eaglehorn
	UniqueWindforce                            // Windforce
	UniqueRing                                 // Ring
	UniqueBulKathosWeddingBand                 // Bul Katho's Wedding Band
	UniqueTheCatsEye                           // The Cat's Eye
	UniqueTheRisingSun                         // The Rising Sun
	UniqueCrescentMoon                         // Crescent Moon
	UniqueMarasKaleidoscope                    // Mara's Kaleidoscope
	UniqueAtmasScarab                          // Atma's Scarab
	UniqueDwarfStar                            // Dwarf Star
	UniqueRavenFrost                           // Raven Frost
	UniqueHighlordsWrath                       // Highlord's Wrath
	UniqueSaracensChance                       // Saracen's Chance
	UniqueClassspecific                        // Class specific
	UniqueArreatsFace                          // Arreat's Face
	UniqueHomunculus                           // Homunculus
	UniqueTitansRevenge                        // Titan's Revenge
	UniqueLycandersAim                         // Lycander's Aim
	UniqueLycandersFlank                       // Lycander's Flank
	UniqueTheOculus                            // The Oculus
	UniqueHeraldOfZakarum                      // Herald of Zakarum
	UniqueBartucsCutThroat                     // Bartuc's Cut-Throat
	UniqueJalalsMane                           // Jalal's Mane
	UniqueTheScalper                           // The Scalper
	UniqueBloodmoon                            // Bloodmoon
	UniqueDjinnslayer                          // Djinnslayer
	UniqueDeathbit                             // Deathbit
	UniqueWarshrike                            // Warshrike
	UniqueGutsiphon                            // Gutsiphon
	UniqueRazoredge                            // Razoredge
	UniqueGoreRipper                           // Gore Ripper
	UniqueDemonLimb                            // Demon Limb
	UniqueSteelShade                           // Steel Shade
	UniqueTombReaver                           // Tomb Reaver
	UniqueDeathsWeb                            // Death's Web
	UniqueNaturesPeace                         // Nature's Peace
	UniqueAzurewrath2                          // Azurewrath
	UniqueSeraphsHymn                          // Seraph's Hymn
	UniqueZakarumsSalvation                    // Zakarum's Salvation
	UniqueFleshripper                          // Fleshripper
	UniqueOdium                                // Odium
	UniqueHorizonsTornado                      // Horizon's Tornado
	UniqueStoneCrusher                         // Stone Crusher
	UniqueJadeTalon                            // Jade Talon
	UniqueShadowDancer                         // Shadow Dancer
	UniqueCerebusBite                          // Cerebus' Bite
	UniqueTyraelsMight                         // Tyrael's Might
	UniqueSoulDrainer                          // Soul Drainer
	UniqueRuneMaster                           // Rune Master
	UniqueDeathCleaver                         // Death Cleaver
	UniqueExecutionersJustice                  // Executioner's Justice
	UniqueStoneraven                           // Stoneraven
	UniqueLeviathan                            // Leviathan
	UniqueLarzuksChampion                      // Larzuk's Champion
	UniqueWispProjector                        // Wisp Projector
	UniqueGargoylesBite                        // Gargoyle's Bite
	UniqueLacerator                            // Lacerator
	UniqueMangSongsLesson                      // Mang Song's Lesson
	UniqueViperfork                            // Viperfork
	UniqueEtherealEdge                         // Ethereal Edge
	UniqueDemonhornsEdge                       // Demonhorn's Edge
	UniqueTheReapersToll                       // The Reaper's Toll
	UniqueSpiritkeeper                         // Spiritkeeper
	UniqueHellrack                             // Hellrack
	UniqueAlmaNegra                            // Alma Negra
	UniqueDarkforgeSpawn                       // Darkforge Spawn
	UniqueWidowmaker                           // Widowmaker
	UniqueBloodravensCharge                    // Bloodraven's Charge
	UniqueGhostflame                           // Ghostflame
	UniqueShadowkiller                         // Shadowkiller
	UniqueGimmershred                          // Gimmershred
	UniqueGriffonsEye                          // Griffon's Eye
	UniqueWindhammer                           // Windhammer
	UniqueThunderstroke                        // Thunderstroke
	UniqueGiantMaimer                          // Giant Maimer
	UniqueDemonsArch                           // Demon's Arch
	UniqueBoneflame                            // Boneflame
	UniqueSteelpillar                          // Steelpillar
	UniqueNightwingsVeil                       // Nightwing's Veil
	UniqueCrownOfAges                          // Crown of Ages
	UniqueAndarielsVisage                      // Andariel's Visage
	UniqueDarkfear                             // Darkfear
	UniqueDragonscale                          // Dragonscale
	UniqueSteelCarapice                        // Steel Carapice
	UniqueMedusasGaze                          // Medusa's Gaze
	UniqueRavenlore                            // Ravenlore
	UniqueBoneshade                            // Boneshade
	UniqueNethercrow                           // Nethercrow
	UniqueFlamebellow                          // Flamebellow
	UniqueFathom                               // Fathom
	UniqueWolfhowl                             // Wolfhowl
	UniqueSpiritWard                           // Spirit Ward
	UniqueKirasGuardian                        // Kira's Guardian
	UniqueOrmusRobes                           // Ormus Robes
	UniqueGheedsFortune                        // Gheed's Fortune
	UniqueStormlash                            // Stormlash
	UniqueHalaberdsReign                       // Halaberd's Reign
	UniqueWarrivsWarder                        // Warriv's Warder
	UniqueSpikeThorn                           // Spike Thorn
	UniqueDraculsGrasp                         // Dracul's Grasp
	UniqueFrostwind                            // Frostwind
	UniqueTemplarsMight                        // Templar's Might
	UniqueEschutasTemper                       // Eschuta's Temper
	UniqueFirelizardsTalons                    // Firelizard's Talons
	UniqueSandstormTrek                        // Sandstorm Trek
	UniqueMarrowwalk                           // Marrowwalk
	UniqueHeavensLight                         // Heaven's Light
	UniqueMermansSpeed                         // Merman's Speed
	UniqueArachnidMesh                         // Arachnid Mesh
	UniqueNosferatusCoil                       // Nosferatu's Coil
	UniqueMetalgrid                            // Metalgrid
	UniqueVerdugosHeartyCord                   // Verdugo's Hearty Cord
	UniqueSigurdsStaunch                       // Sigurd's Staunch
	UniqueCarrionWind                          // Carrion Wind
	UniqueGiantskull                           // Giantskull
	UniqueIronward                             // Ironward
	UniqueAnnihilus                            // Annihilus
	UniqueAriocsNeedle                         // Arioc's Needle
	UniqueCranebeak                            // Cranebeak
	UniqueNordsTenderizer                      // Nord's Tenderizer
	UniqueEarthshifter                         // Earthshifter
	UniqueWraithflight                         // Wraithflight
	UniqueBonehew                              // Bonehew
	UniqueOndalsWisdom                         // Ondal's Wisdom
	UniqueTheReedeemer                         // The Reedeemer
	UniqueHeadhuntersGlory                     // Headhunter's Glory
	UniqueSteelrend                            // Steelrend
	UniqueRainbowFacet1                        // Rainbow Facet
	UniqueRainbowFacet2                        // Rainbow Facet
	UniqueRainbowFacet3                        // Rainbow Facet
	UniqueRainbowFacet4                        // Rainbow Facet
	UniqueRainbowFacet5                        // Rainbow Facet
	UniqueRainbowFacet6                        // Rainbow Facet
	UniqueRainbowFacet7                        // Rainbow Facet
	UniqueRainbowFacet8                        // Rainbow Facet
	UniqueHellfireTorch                        // Hellfire Torch
)
