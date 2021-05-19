package itemdata

//go:generate stringer -linecomment -type MagicalPrefix -output magical_prefix_string.go

// MagicalPrefix represents a magical prefix
type MagicalPrefix uint16

// magical prefixes TODO: names
const (
	MagicalPrefix2 MagicalPrefix = iota + 2 //   Sturdy
	MagicalPrefix3                          //   Strong
	MagicalPrefix4                          //   Glorious
	MagicalPrefix5                          //   Blessed
	MagicalPrefix6                          //   Saintly
	MagicalPrefix7                          //   Holy
	MagicalPrefix8                          //   Devious
	MagicalPrefix9                          //   Fortified
	_
	_
	_
	MagicalPrefix13 //  Jagged
	MagicalPrefix14 //  Deadly
	MagicalPrefix15 //  Vicious
	MagicalPrefix16 //  Brutal
	MagicalPrefix17 //  Massive
	MagicalPrefix18 //  Savage
	MagicalPrefix19 //  Merciless
	MagicalPrefix20 //  Vulpine
	_
	_
	_
	_
	MagicalPrefix25 //  Tireless
	MagicalPrefix26 //  Rugged
	MagicalPrefix27 //  Bronze
	MagicalPrefix28 //  Iron
	MagicalPrefix29 //  Steel
	MagicalPrefix30 //  Silver
	_
	MagicalPrefix32 //  Gold
	MagicalPrefix33 //  Platinum
	MagicalPrefix34 //  Meteoric
	MagicalPrefix35 //  Sharp
	MagicalPrefix36 //  Fine
	MagicalPrefix37 //  Warrior's
	MagicalPrefix38 //  Soldier's
	MagicalPrefix39 //  Knight's
	MagicalPrefix40 //  Lord's
	MagicalPrefix41 //  King's
	MagicalPrefix42 //  Howling
	MagicalPrefix43 //  Fortuitous
	_
	_
	_
	_
	_
	MagicalPrefix49 //  Glimmering
	MagicalPrefix50 //  Glowing
	_
	_
	MagicalPrefix53 //  Lizard's
	_
	MagicalPrefix55 //  Snake's
	MagicalPrefix56 //  Serpent's
	MagicalPrefix57 //  Serpent's
	MagicalPrefix58 //  Drake's
	MagicalPrefix59 //  Dragon's
	MagicalPrefix60 //  Dragon's
	MagicalPrefix61 //  Wyrm's
	_
	_
	MagicalPrefix64 //  Prismatic
	MagicalPrefix65 //  Prismatic
	MagicalPrefix66 //  Azure
	MagicalPrefix67 //  Lapis
	MagicalPrefix68 //  Lapis
	MagicalPrefix69 //  Cobalt
	MagicalPrefix70 //  Cobalt
	_
	MagicalPrefix72 //  Sapphire
	_
	_
	MagicalPrefix75 //  Crimson
	MagicalPrefix76 //  Burgundy
	MagicalPrefix77 //  Burgundy
	MagicalPrefix78 //  Garnet
	MagicalPrefix79 //  Garnet
	_
	MagicalPrefix81 //  Ruby
	_
	_
	MagicalPrefix84 //  Ocher
	MagicalPrefix85 //  Tangerine
	MagicalPrefix86 //  Tangerine
	MagicalPrefix87 //  Coral
	MagicalPrefix88 //  Coral
	_
	MagicalPrefix90 //  Amber
	_
	_
	MagicalPrefix93 //  Beryl
	MagicalPrefix94 //  Jade
	MagicalPrefix95 //  Jade
	MagicalPrefix96 //  Viridian
	MagicalPrefix97 //  Viridian
	_
	MagicalPrefix99 //  Emerald
	_
	MagicalPrefix101 // Fletcher's
	MagicalPrefix102 // Archer's
	MagicalPrefix103 // Archer's
	MagicalPrefix104 // Monk's
	MagicalPrefix105 // Priest's
	MagicalPrefix106 // Priest's
	MagicalPrefix107 // Summoner's
	MagicalPrefix108 // Necromancer's
	MagicalPrefix109 // Necromancer's
	MagicalPrefix110 // Angel's
	MagicalPrefix111 // Arch-Angel's
	MagicalPrefix112 // Arch-Angel's
	MagicalPrefix113 // Slayer's
	MagicalPrefix114 // Berserker's
	MagicalPrefix115 // Berserker's
	_
	_
	MagicalPrefix118 // Triumphant
	MagicalPrefix119 // Stout
	MagicalPrefix120 // Stout
	MagicalPrefix121 // Stout
	MagicalPrefix122 // Burly
	MagicalPrefix123 // Burly
	MagicalPrefix124 // Burly
	MagicalPrefix125 // Stalwart
	MagicalPrefix126 // Stalwart
	MagicalPrefix127 // Stalwart
	MagicalPrefix128 // Stout
	MagicalPrefix129 // Stout
	MagicalPrefix130 // Stout
	MagicalPrefix131 // Burly
	MagicalPrefix132 // Burly
	MagicalPrefix133 // Stalwart
	MagicalPrefix134 // Stalwart
	MagicalPrefix135 // Stout
	MagicalPrefix136 // Stout
	MagicalPrefix137 // Burly
	MagicalPrefix138 // Stalwart
	MagicalPrefix139 // Blanched
	MagicalPrefix140 // Eburin
	MagicalPrefix141 // Bone
	MagicalPrefix142 // Ivory
	MagicalPrefix143 // Sturdy
	MagicalPrefix144 // Sturdy
	MagicalPrefix145 // Strong
	MagicalPrefix146 // Glorious
	MagicalPrefix147 // Blessed
	MagicalPrefix148 // Saintly
	MagicalPrefix149 // Holy
	MagicalPrefix150 // Godly
	MagicalPrefix151 // Devious
	MagicalPrefix152 // Blank
	MagicalPrefix153 // Null
	MagicalPrefix154 // Antimagic
	MagicalPrefix155 // Red
	MagicalPrefix156 // Red
	MagicalPrefix157 // Sanguinary
	MagicalPrefix158 // Sanguinary
	MagicalPrefix159 // Bloody
	MagicalPrefix160 // Red
	MagicalPrefix161 // Sanguinary
	MagicalPrefix162 // Bloody
	MagicalPrefix163 // Red
	MagicalPrefix164 // Sanguinary
	MagicalPrefix165 // Bloody
	MagicalPrefix166 // Scarlet
	MagicalPrefix167 // Crimson
	MagicalPrefix168 // Jagged
	MagicalPrefix169 // Jagged
	MagicalPrefix170 // Jagged
	MagicalPrefix171 // Forked
	MagicalPrefix172 // Forked
	MagicalPrefix173 // Serrated
	MagicalPrefix174 // Serrated
	MagicalPrefix175 // Jagged
	MagicalPrefix176 // Jagged
	MagicalPrefix177 // Forked
	MagicalPrefix178 // Forked
	MagicalPrefix179 // Serrated
	MagicalPrefix180 // Jagged
	MagicalPrefix181 // Forked
	MagicalPrefix182 // Serrated
	MagicalPrefix183 // Carbuncle
	MagicalPrefix184 // Carmine
	MagicalPrefix185 // Vermillion
	MagicalPrefix186 // Jagged
	MagicalPrefix187 // Deadly
	MagicalPrefix188 // Vicious
	MagicalPrefix189 // Brutal
	MagicalPrefix190 // Massive
	MagicalPrefix191 // Savage
	MagicalPrefix192 // Merciless
	MagicalPrefix193 // Ferocious
	MagicalPrefix194 // Cruel
	MagicalPrefix195 // Cinnabar
	MagicalPrefix196 // Rusty
	MagicalPrefix197 // Realgar
	MagicalPrefix198 // Ruby
	MagicalPrefix199 // Vulpine
	MagicalPrefix200 // Dun
	MagicalPrefix201 // Tireless
	MagicalPrefix202 // Tireless
	MagicalPrefix203 // Brown
	MagicalPrefix204 // Rugged
	MagicalPrefix205 // Rugged
	MagicalPrefix206 // Rugged
	MagicalPrefix207 // Rugged
	MagicalPrefix208 // Rugged
	MagicalPrefix209 // Rugged
	MagicalPrefix210 // Rugged
	MagicalPrefix211 // Rugged
	MagicalPrefix212 // Rugged
	MagicalPrefix213 // Rugged
	MagicalPrefix214 // Rugged
	MagicalPrefix215 // Vigorous
	MagicalPrefix216 // Chestnut
	MagicalPrefix217 // Maroon
	MagicalPrefix218 // Bronze
	MagicalPrefix219 // Bronze
	MagicalPrefix220 // Bronze
	MagicalPrefix221 // Iron
	MagicalPrefix222 // Iron
	MagicalPrefix223 // Iron
	MagicalPrefix224 // Steel
	MagicalPrefix225 // Steel
	MagicalPrefix226 // Steel
	MagicalPrefix227 // Bronze
	MagicalPrefix228 // Bronze
	MagicalPrefix229 // Bronze
	MagicalPrefix230 // Iron
	MagicalPrefix231 // Iron
	MagicalPrefix232 // Steel
	MagicalPrefix233 // Steel
	MagicalPrefix234 // Bronze
	MagicalPrefix235 // Bronze
	MagicalPrefix236 // Iron
	MagicalPrefix237 // Steel
	MagicalPrefix238 // Bronze
	MagicalPrefix239 // Iron
	MagicalPrefix240 // Steel
	MagicalPrefix241 // Silver
	MagicalPrefix242 // Gold
	MagicalPrefix243 // Platinum
	MagicalPrefix244 // Meteoric
	MagicalPrefix245 // Strange
	MagicalPrefix246 // Weird
	MagicalPrefix247 // Nickel
	MagicalPrefix248 // Tin
	MagicalPrefix249 // Silver
	MagicalPrefix250 // Argent
	MagicalPrefix251 // Fine
	MagicalPrefix252 // Fine
	MagicalPrefix253 // Sharp
	MagicalPrefix254 // Fine
	MagicalPrefix255 // Sharp
	MagicalPrefix256 // Fine
	MagicalPrefix257 // Sharp
	MagicalPrefix258 // Fine
	MagicalPrefix259 // Warrior's
	MagicalPrefix260 // Soldier's
	MagicalPrefix261 // Knight's
	MagicalPrefix262 // Lord's
	MagicalPrefix263 // King's
	MagicalPrefix264 // Master's
	MagicalPrefix265 // Grandmaster's
	MagicalPrefix266 // Glimmering
	MagicalPrefix267 // Glowing
	MagicalPrefix268 // Bright
	MagicalPrefix269 // Screaming
	MagicalPrefix270 // Howling
	MagicalPrefix271 // Wailing
	MagicalPrefix272 // Screaming
	MagicalPrefix273 // Howling
	MagicalPrefix274 // Wailing
	MagicalPrefix275 // Lucky
	MagicalPrefix276 // Lucky
	MagicalPrefix277 // Lucky
	MagicalPrefix278 // Lucky
	MagicalPrefix279 // Lucky
	MagicalPrefix280 // Lucky
	MagicalPrefix281 // Felicitous
	MagicalPrefix282 // Fortuitous
	MagicalPrefix283 // Emerald
	MagicalPrefix284 // Lizard's
	MagicalPrefix285 // Lizard's
	MagicalPrefix286 // Lizard's
	MagicalPrefix287 // Snake's
	MagicalPrefix288 // Snake's
	MagicalPrefix289 // Snake's
	MagicalPrefix290 // Serpent's
	MagicalPrefix291 // Serpent's
	MagicalPrefix292 // Serpent's
	MagicalPrefix293 // Lizard's
	MagicalPrefix294 // Lizard's
	MagicalPrefix295 // Lizard's
	MagicalPrefix296 // Snake's
	MagicalPrefix297 // Snake's
	MagicalPrefix298 // Serpent's
	MagicalPrefix299 // Serpent's
	MagicalPrefix300 // Lizard's
	MagicalPrefix301 // Lizard's
	MagicalPrefix302 // Snake's
	MagicalPrefix303 // Serpent's
	MagicalPrefix304 // Lizard's
	MagicalPrefix305 // Snake's
	MagicalPrefix306 // Serpent's
	MagicalPrefix307 // Serpent's
	MagicalPrefix308 // Drake's
	MagicalPrefix309 // Dragon's
	MagicalPrefix310 // Dragon's
	MagicalPrefix311 // Wyrm's
	MagicalPrefix312 // Great Wyrm's
	MagicalPrefix313 // Bahamut's
	MagicalPrefix314 // Zircon
	MagicalPrefix315 // Jacinth
	MagicalPrefix316 // Turquoise
	MagicalPrefix317 // Shimmering
	MagicalPrefix318 // Shimmering
	MagicalPrefix319 // Shimmering
	MagicalPrefix320 // Shimmering
	MagicalPrefix321 // Shimmering
	MagicalPrefix322 // Shimmering
	MagicalPrefix323 // Shimmering
	MagicalPrefix324 // Rainbow
	MagicalPrefix325 // Scintillating
	MagicalPrefix326 // Prismatic
	MagicalPrefix327 // Chromatic
	MagicalPrefix328 // Shimmering
	MagicalPrefix329 // Rainbow
	MagicalPrefix330 // Scintillating
	MagicalPrefix331 // Prismatic
	MagicalPrefix332 // Chromatic
	MagicalPrefix333 // Shimmering
	MagicalPrefix334 // Rainbow
	MagicalPrefix335 // Scintillating
	MagicalPrefix336 // Shimmering
	MagicalPrefix337 // Scintillating
	MagicalPrefix338 // Azure
	MagicalPrefix339 // Lapis
	MagicalPrefix340 // Cobalt
	MagicalPrefix341 // Sapphire
	MagicalPrefix342 // Azure
	MagicalPrefix343 // Lapis
	MagicalPrefix344 // Cobalt
	MagicalPrefix345 // Sapphire
	MagicalPrefix346 // Azure
	MagicalPrefix347 // Lapis
	MagicalPrefix348 // Cobalt
	MagicalPrefix349 // Sapphire
	MagicalPrefix350 // Azure
	MagicalPrefix351 // Lapis
	MagicalPrefix352 // Lapis
	MagicalPrefix353 // Cobalt
	MagicalPrefix354 // Cobalt
	MagicalPrefix355 // Sapphire
	MagicalPrefix356 // Lapis Lazuli
	MagicalPrefix357 // Sapphire
	MagicalPrefix358 // Crimson
	MagicalPrefix359 // Russet
	MagicalPrefix360 // Garnet
	MagicalPrefix361 // Ruby
	MagicalPrefix362 // Crimson
	MagicalPrefix363 // Russet
	MagicalPrefix364 // Garnet
	MagicalPrefix365 // Ruby
	MagicalPrefix366 // Crimson
	MagicalPrefix367 // Russet
	MagicalPrefix368 // Garnet
	MagicalPrefix369 // Ruby
	MagicalPrefix370 // Russet
	MagicalPrefix371 // Russet
	MagicalPrefix372 // Garnet
	MagicalPrefix373 // Garnet
	MagicalPrefix374 // Ruby
	MagicalPrefix375 // Garnet
	MagicalPrefix376 // Ruby
	MagicalPrefix377 // Tangerine
	MagicalPrefix378 // Ocher
	MagicalPrefix379 // Coral
	MagicalPrefix380 // Amber
	MagicalPrefix381 // Tangerine
	MagicalPrefix382 // Ocher
	MagicalPrefix383 // Coral
	MagicalPrefix384 // Amber
	MagicalPrefix385 // Tangerine
	MagicalPrefix386 // Ocher
	MagicalPrefix387 // Coral
	MagicalPrefix388 // Amber
	MagicalPrefix389 // Tangerine
	MagicalPrefix390 // Ocher
	MagicalPrefix391 // Ocher
	MagicalPrefix392 // Coral
	MagicalPrefix393 // Coral
	MagicalPrefix394 // Amber
	MagicalPrefix395 // Camphor
	MagicalPrefix396 //  Ambergris
	MagicalPrefix397 // Beryl
	MagicalPrefix398 // Viridian
	MagicalPrefix399 // Jade
	MagicalPrefix400 // Emerald
	MagicalPrefix401 // Beryl
	MagicalPrefix402 // Viridian
	MagicalPrefix403 // Jade
	MagicalPrefix404 // Emerald
	MagicalPrefix405 // Beryl
	MagicalPrefix406 // Viridian
	MagicalPrefix407 // Jade
	MagicalPrefix408 // Emerald
	MagicalPrefix409 // Beryl
	MagicalPrefix410 // Viridian
	MagicalPrefix411 // Viridian
	MagicalPrefix412 // Jade
	MagicalPrefix413 // Jade
	MagicalPrefix414 // Emerald
	MagicalPrefix415 // Beryl
	MagicalPrefix416 // Jade
	MagicalPrefix417 // Triumphant
	MagicalPrefix418 // Victorious
	MagicalPrefix419 // Aureolin
	MagicalPrefix420 // Mechanist's
	MagicalPrefix421 // Artificer's
	MagicalPrefix422 // Jeweler's
	MagicalPrefix423 // Assamic
	MagicalPrefix424 // Arcadian
	MagicalPrefix425 // Unearthly
	MagicalPrefix426 // Astral
	MagicalPrefix427 // Elysian
	MagicalPrefix428 // Celestial
	MagicalPrefix429 // Diamond
	MagicalPrefix430 // Fletcher's
	MagicalPrefix431 // Acrobat's
	MagicalPrefix432 // Harpoonist's
	MagicalPrefix433 // Fletcher's
	MagicalPrefix434 // Bowyer's
	MagicalPrefix435 // Archer's
	MagicalPrefix436 // Acrobat's
	MagicalPrefix437 // Gymnast's
	MagicalPrefix438 // Athlete's
	MagicalPrefix439 // Harpoonist's
	MagicalPrefix440 // Spearmaiden's
	MagicalPrefix441 // Lancer's
	MagicalPrefix442 // Burning
	MagicalPrefix443 // Sparking
	MagicalPrefix444 // Chilling
	MagicalPrefix445 // Burning
	MagicalPrefix446 // Blazing
	MagicalPrefix447 // Volcanic
	MagicalPrefix448 // Sparking
	MagicalPrefix449 // Charged
	MagicalPrefix450 // Powered
	MagicalPrefix451 // Chilling
	MagicalPrefix452 // Freezing
	MagicalPrefix453 // Glacial
	MagicalPrefix454 // Hexing
	MagicalPrefix455 // Fungal
	MagicalPrefix456 // Graverobber's
	MagicalPrefix457 // Hexing
	MagicalPrefix458 // Blighting
	MagicalPrefix459 // Accursed
	MagicalPrefix460 // Fungal
	MagicalPrefix461 // Noxious
	MagicalPrefix462 // Venomous
	MagicalPrefix463 // Graverobber's
	MagicalPrefix464 // Vodoun
	MagicalPrefix465 // Golemlord's
	MagicalPrefix466 // Lion Branded
	MagicalPrefix467 // Captain's
	MagicalPrefix468 // Preserver's
	MagicalPrefix469 // Lion Branded
	MagicalPrefix470 // Hawk Branded
	MagicalPrefix471 // Rose Branded
	MagicalPrefix472 // Captain's
	MagicalPrefix473 // Commander's
	MagicalPrefix474 // Marshal's
	MagicalPrefix475 // Preserver's
	MagicalPrefix476 // Warder's
	MagicalPrefix477 // Guardian's
	MagicalPrefix478 // Expert's
	MagicalPrefix479 // Fanatic
	MagicalPrefix480 // Sounding
	MagicalPrefix481 // Expert's
	MagicalPrefix482 // Veteran's
	MagicalPrefix483 // Master's
	MagicalPrefix484 // Fanatic
	MagicalPrefix485 // Raging
	MagicalPrefix486 // Furious
	MagicalPrefix487 // Sounding
	MagicalPrefix488 // Resonant
	MagicalPrefix489 // Echoing
	MagicalPrefix490 // Trainer's
	MagicalPrefix491 // Spiritual
	MagicalPrefix492 // Nature's
	MagicalPrefix493 // Trainer's
	MagicalPrefix494 // Caretaker's
	MagicalPrefix495 // Keeper's
	MagicalPrefix496 // Spiritual
	MagicalPrefix497 // Feral
	MagicalPrefix498 // Communal
	MagicalPrefix499 // Nature's
	MagicalPrefix500 // Terra's
	MagicalPrefix501 // Gaea's
	MagicalPrefix502 // Entrapping
	MagicalPrefix503 // Mentalist's
	MagicalPrefix504 // Shogukusha's
	MagicalPrefix505 // Entrapping
	MagicalPrefix506 // Trickster's
	MagicalPrefix507 // Cunning
	MagicalPrefix508 // Mentalist's
	MagicalPrefix509 // Psychic
	MagicalPrefix510 // Shadow
	MagicalPrefix511 // Shogukusha's
	MagicalPrefix512 // Sensei's
	MagicalPrefix513 // Kenshi's
	MagicalPrefix514 // Miocene
	MagicalPrefix515 // Miocene
	MagicalPrefix516 // Oligocene
	MagicalPrefix517 // Oligocene
	MagicalPrefix518 // Eocene
	MagicalPrefix519 // Eocene
	MagicalPrefix520 // Paleocene
	MagicalPrefix521 // Paleocene
	MagicalPrefix522 // Knave's
	MagicalPrefix523 // Jack's
	MagicalPrefix524 // Jester's
	MagicalPrefix525 // Joker's
	MagicalPrefix526 // Trump
	MagicalPrefix527 // Loud
	MagicalPrefix528 // Calling
	MagicalPrefix529 // Yelling
	MagicalPrefix530 // Shouting
	MagicalPrefix531 // Screaming
	MagicalPrefix532 // Paradox
	MagicalPrefix533 // Paradox
	MagicalPrefix534 // Robineye
	MagicalPrefix535 // Sparroweye
	MagicalPrefix536 // Falconeye
	MagicalPrefix537 // Hawkeye
	MagicalPrefix538 // Eagleeye
	MagicalPrefix539 // Visionary
	MagicalPrefix540 // Mnemonic
	MagicalPrefix541 // Snowflake
	MagicalPrefix542 // Shivering
	MagicalPrefix543 // Boreal
	MagicalPrefix544 // Hibernal
	MagicalPrefix545 // Ember
	MagicalPrefix546 // Smoldering
	MagicalPrefix547 // Smoking
	MagicalPrefix548 // Flaming
	MagicalPrefix549 // Scorching
	MagicalPrefix550 // Static
	MagicalPrefix551 // Glowing
	MagicalPrefix552 // Buzzing
	MagicalPrefix553 // Arcing
	MagicalPrefix554 // Shocking
	MagicalPrefix555 // Septic
	MagicalPrefix556 // Envenomed
	MagicalPrefix557 // Corosive
	MagicalPrefix558 // Toxic
	MagicalPrefix559 // Pestilent
	MagicalPrefix560 // Maiden's
	MagicalPrefix561 // Valkyrie's
	MagicalPrefix562 // Maiden's
	MagicalPrefix563 // Valkyrie's
	MagicalPrefix564 // Monk's
	MagicalPrefix565 // Priest's
	MagicalPrefix566 // Monk's
	MagicalPrefix567 // Priest's
	MagicalPrefix568 // Monk's
	MagicalPrefix569 // Priest's
	MagicalPrefix570 // Summoner's
	MagicalPrefix571 // Necromancer's
	MagicalPrefix572 // Summoner's
	MagicalPrefix573 // Necromancer's
	MagicalPrefix574 // Angel's
	MagicalPrefix575 // Arch-Angel's
	MagicalPrefix576 // Angel's
	MagicalPrefix577 // Arch-Angel's
	MagicalPrefix578 // Slayer's
	MagicalPrefix579 // Berserker's
	MagicalPrefix580 // Slayer's
	MagicalPrefix581 // Berserker's
	MagicalPrefix582 // Slayer's
	MagicalPrefix583 // Berserker's
	MagicalPrefix584 // Shaman's
	MagicalPrefix585 // Hierophant's
	MagicalPrefix586 // Shaman's
	MagicalPrefix587 // Hierophant's
	MagicalPrefix588 // Magekiller's
	MagicalPrefix589 // Witch-hunter's
	MagicalPrefix590 // Magekiller's
	MagicalPrefix591 // Witch-hunter's
	MagicalPrefix592 // Compact
	MagicalPrefix593 // Thin
	MagicalPrefix594 // Dense
	MagicalPrefix595 // Consecrated
	MagicalPrefix596 // Pure
	MagicalPrefix597 // Sacred
	MagicalPrefix598 // Hallowed
	MagicalPrefix599 // Divine
	MagicalPrefix600 // Pearl
	MagicalPrefix601 // Crimson
	MagicalPrefix602 // Red
	MagicalPrefix603 // Sanguinary
	MagicalPrefix604 // Bloody
	MagicalPrefix605 // Red
	MagicalPrefix606 // Sanguinary
	MagicalPrefix607 // Red
	MagicalPrefix608 // Jagged
	MagicalPrefix609 // Forked
	MagicalPrefix610 // Serrated
	MagicalPrefix611 // Jagged
	MagicalPrefix612 // Forked
	MagicalPrefix613 // Jagged
	MagicalPrefix614 // Snowflake
	MagicalPrefix615 // Shivering
	MagicalPrefix616 // Boreal
	MagicalPrefix617 // Hibernal
	MagicalPrefix618 // Snowflake
	MagicalPrefix619 // Shivering
	MagicalPrefix620 // Boreal
	MagicalPrefix621 // Hibernal
	MagicalPrefix622 // Snowflake
	MagicalPrefix623 // Shivering
	MagicalPrefix624 // Boreal
	MagicalPrefix625 // Hibernal
	MagicalPrefix626 // Ember
	MagicalPrefix627 // Smoldering
	MagicalPrefix628 // Smoking
	MagicalPrefix629 // Flaming
	MagicalPrefix630 // Ember
	MagicalPrefix631 // Smoldering
	MagicalPrefix632 // Smoking
	MagicalPrefix633 // Flaming
	MagicalPrefix634 // Ember
	MagicalPrefix635 // Smoldering
	MagicalPrefix636 // Smoking
	MagicalPrefix637 // Flaming
	MagicalPrefix638 // Static
	MagicalPrefix639 // Glowing
	MagicalPrefix640 // Arcing
	MagicalPrefix641 // Shocking
	MagicalPrefix642 // Static
	MagicalPrefix643 // Glowing
	MagicalPrefix644 // Arcing
	MagicalPrefix645 // Shocking
	MagicalPrefix646 // Static
	MagicalPrefix647 // Glowing
	MagicalPrefix648 // Arcing
	MagicalPrefix649 // Shocking
	MagicalPrefix650 // Septic
	MagicalPrefix651 // Envenomed
	MagicalPrefix652 // Toxic
	MagicalPrefix653 // Pestilent
	MagicalPrefix654 // Septic
	MagicalPrefix655 // Envenomed
	MagicalPrefix656 // Toxic
	MagicalPrefix657 // Pestilent
	MagicalPrefix658 // Septic
	MagicalPrefix659 // Envenomed
	MagicalPrefix660 // Toxic
	MagicalPrefix661 // Pestilent
	MagicalPrefix662 // Tireless
	MagicalPrefix663 // Lizard's
	MagicalPrefix664 // Azure
	MagicalPrefix665 // Crimson
	MagicalPrefix666 // Tangerine
	MagicalPrefix667 // Beryl
	MagicalPrefix668 // Godly
	MagicalPrefix669 // Cruel
)

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
