package itemdata

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
