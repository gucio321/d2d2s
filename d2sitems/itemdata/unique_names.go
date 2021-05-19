package itemdata

// UniqueID represents an uinque item's id
type UniqueID uint16

// unique IDs TODO: name these constants
const (
	Unique0   UniqueID = iota // The Gnasher
	Unique1                   // Deathspade
	Unique2                   // Bladebone
	Unique3                   // Skull splitter
	Unique4                   // Rakescar
	Unique5                   // Axe of Fechmar
	Unique6                   // Goreshovel
	Unique7                   // The Chiefthan
	Unique8                   // Brainhew
	Unique9                   // Humongous
	Unique10                  // Torch of Iros
	Unique11                  // Maelstorm
	Unique12                  // Gravenspine
	Unique13                  // Umes Lament
	Unique14                  // Felloak
	Unique15                  // Knell Striker
	Unique16                  // Rusthandle
	Unique17                  // Stormeye
	Unique18                  // Stoutnail
	Unique19                  // Crushflange
	Unique20                  // Bloodrise
	Unique21                  // The Generals Tan Do Li Ga
	Unique22                  // Ironstone
	Unique23                  // Bonesnap
	Unique24                  // Steeldriver
	Unique25                  // Rixot's Keen
	Unique26                  // Blood Crescent
	Unique27                  // Skewer of Krintiz
	Unique28                  // Gleamscythe
	Unique29                  // Azurewrath
	Unique30                  // Griswold's Edge
	Unique31                  // Hellplague
	Unique32                  // Culwens Point
	Unique33                  // Shadowfang
	Unique34                  // Soulflay
	Unique35                  // Kinemils Awl
	Unique36                  // Blacktongue
	Unique37                  // Ripsaw
	Unique38                  // The Patriarch
	Unique39                  // Gull
	Unique40                  // The Diggler
	Unique41                  // The Jade Tan Do
	Unique42                  // Spectral Shard
	Unique43                  // The Dragon Chang
	Unique44                  // Razortine
	Unique45                  // Bloodthief
	Unique46                  // Lance of Yaggai
	Unique47                  // The Tannr Gorerod
	Unique48                  // Dimoaks Hew
	Unique49                  // Steelgoad
	Unique50                  // Soul Harvest
	Unique51                  // The Battlebranch
	Unique52                  // Woestave
	Unique53                  // The Grim Reaper
	Unique54                  // Bane Ash
	Unique55                  // Serpent Lord
	Unique56                  // Spire of Lazarus
	Unique57                  // The Salamander
	Unique58                  // The Iron Jang Bong
	Unique59                  // Pluckeye
	Unique60                  // Witherstring
	Unique61                  // Raven Claw
	Unique62                  // Rogue's Bow
	Unique63                  // Stormstrike
	Unique64                  // Wizendraw
	Unique65                  // Hellclap
	Unique66                  // Blastbark
	Unique67                  // Leadcrow
	Unique68                  // Ichorsting
	Unique69                  // Hellcast
	Unique70                  // Doomslinger
	Unique71                  // Biggin's Bonnet
	Unique72                  // Tarnhelm
	Unique73                  // Coif of Glory
	Unique74                  // Duskdeep
	Unique75                  // Wormskull
	Unique76                  // Howltusk
	Unique77                  // Undead Crown
	Unique78                  // The Face of Horror
	Unique79                  // Greyform
	Unique80                  // Blinkbat's Form
	Unique81                  // The Centurion
	Unique82                  // Twitchthroe
	Unique83                  // Darkglow
	Unique84                  // Hawkmail
	Unique85                  // Sparking Mail
	Unique86                  // Venom Ward
	Unique87                  // Iceblink
	Unique88                  // Boneflesh
	Unique89                  // Rockfleece
	Unique90                  // Rattlecage
	Unique91                  // Goldskin
	Unique92                  // Victors Silk
	Unique93                  // Heavenly Garb
	Unique94                  // Pelta Lunata
	Unique95                  // Umbral Disk
	Unique96                  // Stormguild
	Unique97                  // Wall of the Eyeless
	Unique98                  // Swordback Hold
	Unique99                  // Steelclash
	Unique100                 // Bverrit Keep
	Unique101                 // The Ward
	Unique102                 // The Hand of Broc
	Unique103                 // Bloodfist
	Unique104                 // Chance Guards
	Unique105                 // Magefist
	Unique106                 // Frostburn
	Unique107                 // Hotspur
	Unique108                 // Gorefoot
	Unique109                 // Treads of Cthon
	Unique110                 // Goblin Toe
	Unique111                 // Tearhaunch
	Unique112                 // Lenymo
	Unique113                 // Snakecord
	Unique114                 // Nightsmoke
	Unique115                 // Goldwrap
	Unique116                 // Bladebuckle
	Unique117                 // Nokozan Relic
	Unique118                 // The Eye of Etlich
	Unique119                 // The Mahim-Oak Curio
	Unique120                 // Nagelring
	Unique121                 // Manald Heal
	Unique122                 // The Stone of Jordan
	Unique123                 // Amulet of the Viper
	Unique124                 // Staff of Kings
	Unique125                 // Horadric Staff
	Unique126                 // Hell Forge Hammer
	Unique127                 // Khalim's Flail
	Unique128                 // Super Khalim's Flail
	Unique129                 // Coldkill
	Unique130                 // Butcher's Pupil
	Unique131                 // Islestrike
	Unique132                 // Pompe's Wrath
	Unique133                 // Guardian Naga
	Unique134                 // Warlord's Trust
	Unique135                 // Spellsteel
	Unique136                 // Stormrider
	Unique137                 // Boneslayer Blade
	Unique138                 // The Minataur
	Unique139                 // Suicide Branch
	Unique140                 // Carin Shard
	Unique141                 // Arm of King Leoric
	Unique142                 // Blackhand Key
	Unique143                 // Dark Clan Crusher
	Unique144                 // Zakarum's Hand
	Unique145                 // The Fetid Sprinkler
	Unique146                 // Hand of Blessed Light
	Unique147                 // Fleshrender
	Unique148                 // Sureshrill Frost
	Unique149                 // Moonfall
	Unique150                 // Baezil's Vortex
	Unique151                 // Earthshaker
	Unique152                 // Bloodtree Stump
	Unique153                 // The Gavel of Pain
	Unique154                 // Bloodletter
	Unique155                 // Coldsteel Eye
	Unique156                 // Hexfire
	Unique157                 // Blade of Ali Baba
	Unique158                 // Ginther's Rift
	Unique159                 // Headstriker
	Unique160                 // Plague Bearer
	Unique161                 // The Atlantian
	Unique162                 // Crainte Vomir
	Unique163                 // Bing Sz Wang
	Unique164                 // The Vile Husk
	Unique165                 // Cloudcrack
	Unique166                 // Todesfaelle Flamme
	Unique167                 // Swordguard
	Unique168                 // Spineripper
	Unique169                 // Heart Carver
	Unique170                 // Blackbog's Sharp
	Unique171                 // Stormspike
	Unique172                 // The Impaler
	Unique173                 // Kelpie Snare
	Unique174                 // Soulfeast Tine
	Unique175                 // Hone Sundan
	Unique176                 // Spire of Honor
	Unique177                 // The Meat Scraper
	Unique178                 // Blackleach Blade
	Unique179                 // Athena's Wrath
	Unique180                 // Pierre Tombale Couant
	Unique181                 // Husoldal Evo
	Unique182                 // Grim's Burning Dead
	Unique183                 // Razorswitch
	Unique184                 // Ribcracker
	Unique185                 // Chromatic Ire
	Unique186                 // Warpspear
	Unique187                 // Skullcollector
	Unique188                 // Skystrike
	Unique189                 // Riphook
	Unique190                 // Kuko Shakaku
	Unique191                 // Endlesshail
	Unique192                 // Whichwild String
	Unique193                 // Cliffkiller
	Unique194                 // Magewrath
	Unique195                 // Godstrike Arch
	Unique196                 // Langer Briser
	Unique197                 // Pus Spiter
	Unique198                 // Buriza-Do Kyanon
	Unique199                 // Demon Machine
	Unique200                 // Armor (Unknown)
	Unique201                 // Peasent Crown
	Unique202                 // Rockstopper
	Unique203                 // Stealskull
	Unique204                 // Darksight Helm
	Unique205                 // Valkyrie Wing
	Unique206                 // Crown of Thieves
	Unique207                 // Blckhorn's Face
	Unique208                 // Vampire Gaze
	Unique209                 // The Spirit Shroud
	Unique210                 // Skin of the Vipermagi
	Unique211                 // Skin of the Flayed One
	Unique212                 // Ironpelt
	Unique213                 // Spiritforge
	Unique214                 // Crow Caw
	Unique215                 // Shaftstop
	Unique216                 // Duriel's Shell
	Unique217                 // Skullder's Ire
	Unique218                 // Guardian Angel
	Unique219                 // Toothrow
	Unique220                 // Atma's Wail
	Unique221                 // Black Hades
	Unique222                 // Corpsemourn
	Unique223                 // Que-Hegan's Wisdom
	Unique224                 // Visceratuant
	Unique225                 // Mosers Blessed Circle
	Unique226                 // Stormchaser
	Unique227                 // Tiamat's Rebuke
	Unique228                 // Gerke's Sanctuary
	Unique229                 // Radimant's Sphere
	Unique230                 // Lidless Wall
	Unique231                 // Lance Guard
	Unique232                 // Venom Grip
	Unique233                 // Gravepalm
	Unique234                 // Ghoulhide
	Unique235                 // Lavagout
	Unique236                 // Hellmouth
	Unique237                 // Infernostride
	Unique238                 // Waterwalk
	Unique239                 // Silkweave
	Unique240                 // Wartraveler
	Unique241                 // Gorerider
	Unique242                 // String of Ears
	Unique243                 // Razortail
	Unique244                 // Gloomstrap
	Unique245                 // Snowclash
	Unique246                 // Thundergod's Vigor
	Unique247                 // Elite unique
	Unique248                 // Harlequin Crest
	Unique249                 // Veil of Steel
	Unique250                 // The Gladiator's Bane
	Unique251                 // Arkaine's Valor
	Unique252                 // Blackoak Shield
	Unique253                 // Stormshield
	Unique254                 // Hellslayer
	Unique255                 // Messerschmidt's Reaver
	Unique256                 // Baranar's Star
	Unique257                 // Schaefer's Hammer
	Unique258                 // The Cranium Basher
	Unique259                 // Lightsabre
	Unique260                 // Doombringer
	Unique261                 // The Grandfather
	Unique262                 // Wizardspike
	Unique263                 // Constricting Ring
	Unique264                 // Stormspire
	Unique265                 // Eaglehorn
	Unique266                 // Windforce
	Unique267                 // Ring
	Unique268                 // Bul Katho's Wedding Band
	Unique269                 // The Cat's Eye
	Unique270                 // The Rising Sun
	Unique271                 // Crescent Moon
	Unique272                 // Mara's Kaleidoscope
	Unique273                 // Atma's Scarab
	Unique274                 // Dwarf Star
	Unique275                 // Raven Frost
	Unique276                 // Highlord's Wrath
	Unique277                 // Saracen's Chance
	Unique278                 // Class specific
	Unique279                 // Arreat's Face
	Unique280                 // Homunculus
	Unique281                 // Titan's Revenge
	Unique282                 // Lycander's Aim
	Unique283                 // Lycander's Flank
	Unique284                 // The Oculus
	Unique285                 // Herald of Zakarum
	Unique286                 // Bartuc's Cut-Throat
	Unique287                 // Jalal's Mane
	Unique288                 // The Scalper
	Unique289                 // Bloodmoon
	Unique290                 // Djinnslayer
	Unique291                 // Deathbit
	Unique292                 // Warshrike
	Unique293                 // Gutsiphon
	Unique294                 // Razoredge
	Unique295                 // Gore Ripper
	Unique296                 // Demon Limb
	Unique297                 // Steel Shade
	Unique298                 // Tomb Reaver
	Unique299                 // Death's Web
	Unique300                 // Nature's Peace
	Unique301                 // Azurewrath
	Unique302                 // Seraph's Hymn
	Unique303                 // Zakarum's Salvation
	Unique304                 // Fleshripper
	Unique305                 // Odium
	Unique306                 // Horizon's Tornado
	Unique307                 // Stone Crusher
	Unique308                 // Jade Talon
	Unique309                 // Shadow Dancer
	Unique310                 // Cerebus' Bite
	Unique311                 // Tyrael's Might
	Unique312                 // Soul Drainer
	Unique313                 // Rune Master
	Unique314                 // Death Cleaver
	Unique315                 // Executioner's Justice
	Unique316                 // Stoneraven
	Unique317                 // Leviathan
	Unique318                 // Larzuk's Champion
	Unique319                 // Wisp Projector
	Unique320                 // Gargoyle's Bite
	Unique321                 // Lacerator
	Unique322                 // Mang Song's Lesson
	Unique323                 // Viperfork
	Unique324                 // Ethereal Edge
	Unique325                 // Demonhorn's Edge
	Unique326                 // The Reaper's Toll
	Unique327                 // Spiritkeeper
	Unique328                 // Hellrack
	Unique329                 // Alma Negra
	Unique330                 // Darkforge Spawn
	Unique331                 // Widowmaker
	Unique332                 // Bloodraven's Charge
	Unique333                 // Ghostflame
	Unique334                 // Shadowkiller
	Unique335                 // Gimmershred
	Unique336                 // Griffon's Eye
	Unique337                 // Windhammer
	Unique338                 // Thunderstroke
	Unique339                 // Giant Maimer
	Unique340                 // Demon's Arch
	Unique341                 // Boneflame
	Unique342                 // Steelpillar
	Unique343                 // Nightwing's Veil
	Unique344                 // Crown of Ages
	Unique345                 // Andariel's Visage
	Unique346                 // Darkfear
	Unique347                 // Dragonscale
	Unique348                 // Steel Carapice
	Unique349                 // Medusa's Gaze
	Unique350                 // Ravenlore
	Unique351                 // Boneshade
	Unique352                 // Nethercrow
	Unique353                 // Flamebellow
	Unique354                 // Fathom
	Unique355                 // Wolfhowl
	Unique356                 // Spirit Ward
	Unique357                 // Kira's Guardian
	Unique358                 // Ormus Robes
	Unique359                 // Gheed's Fortune
	Unique360                 // Stormlash
	Unique361                 // Halaberd's Reign
	Unique362                 // Warriv's Warder
	Unique363                 // Spike Thorn
	Unique364                 // Dracul's Grasp
	Unique365                 // Frostwind
	Unique366                 // Templar's Might
	Unique367                 // Eschuta's Temper
	Unique368                 // Firelizard's Talons
	Unique369                 // Sandstorm Trek
	Unique370                 // Marrowwalk
	Unique371                 // Heaven's Light
	Unique372                 // Merman's Speed
	Unique373                 // Arachnid Mesh
	Unique374                 // Nosferatu's Coil
	Unique375                 // Metalgrid
	Unique376                 // Verdugo's Hearty Cord
	Unique377                 // Sigurd's Staunch
	Unique378                 // Carrion Wind
	Unique379                 // Giantskull
	Unique380                 // Ironward
	Unique381                 // Annihilus
	Unique382                 // Arioc's Needle
	Unique383                 // Cranebeak
	Unique384                 // Nord's Tenderizer
	Unique385                 // Earthshifter
	Unique386                 // Wraithflight
	Unique387                 // Bonehew
	Unique388                 // Ondal's Wisdom
	Unique389                 // The Reedeemer
	Unique390                 // Headhunter's Glory
	Unique391                 // Steelrend
	Unique392                 // Rainbow Facet
	Unique393                 // Rainbow Facet
	Unique394                 // Rainbow Facet
	Unique395                 // Rainbow Facet
	Unique396                 // Rainbow Facet
	Unique397                 // Rainbow Facet
	Unique398                 // Rainbow Facet
	Unique399                 // Rainbow Facet
	Unique400                 // Hellfire Torch
)
