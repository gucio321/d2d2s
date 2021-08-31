package itemdata

import "log"

// GetSetAttributesLen : each set item has 5 bits of data containing the number of set lists follow
// the magical attributes list, this map tells us how many lists to read
// depending on the value given from the 5 bits. A number of 0-5 set lists.
// nolint:gomnd // data method
func GetSetAttributesLen(id byte) byte {
	lookup := map[byte]byte{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  1,
		6:  2,
		7:  3,
		10: 2,
		12: 2,
		15: 4,
		31: 5,
	}

	v, ok := lookup[id]
	if !ok {
		log.Panicf("unknown set attributes id: %v", id)
	}

	return v
}

// SetReqID - Certain set items (only Civerb's Ward in unmodded D2) have bonuses
// that require certain other set items in order to be activated
// (instead of the normal requirements of just 'wearing > x of any
// items in the set'); determined by add_func=1 in SetItems.txt
func SetReqID(id byte) []uint64 {
	SetReqIDsMap := map[byte][]uint64{
		// Civerb's Ward: [Civerb's Icon, Civerb's Cudgel]
		0: {1, 2},
	}

	r, ok := SetReqIDsMap[id]
	if !ok {
		return nil
	}

	return r
}

//go:generate stringer -linecomment -type SetID -output set_ids_string.go

// SetID represents a set ID
type SetID uint16

// set IDs
const (
	SetCiverbsWard               SetID = iota // Civerb's Ward
	SetCiverbsIcon                            // Civerbs Icon
	SetCiverbsCudgel                          // Civerbs Cudgel
	SetHsarusIronHeel                         // Hsarus Iron Heel
	SetHsarusIronFist                         // Hsarus Iron Fist
	SetHsarusIronStay                         // Hsarus Iron Stay
	SetCleglawsTooth                          // Cleglaws Tooth
	SetCleglawsClaw                           // Cleglaws Claw
	SetCleglawsPincers                        // Cleglaws Pincers
	SetIrathasCollar                          // Irathas Collar
	SetSetIrathasCuff                         // Irathas Cuff
	SetSetIrathasCoil                         // Irathas Coil
	SetSetIrathasCord                         // Irathas Cord
	SetSetIsenhartsLightbrand                 // Isenharts Lightbrand
	SetSetIsenhartsParry                      // Isenharts Parry
	SetSetIsenhartsCase                       // Isenharts Case
	SetSetIsenhartsHorns                      // Isenharts Horns
	SetSetVidalasBarb                         // Vidalas Barb
	SetSetVidalasFetlock                      // Vidalas Fetlock
	SetSetVidalasAmbush                       // Vidalas Ambush
	SetSetVidalasSnare                        // Vidalas Snare
	SetSetMilabregasOrb                       // Milabregas Orb
	SetSetMilabregasRod                       // Milabregas Rod
	SetSetMilabregasDiadem                    // Milabregas Diadem
	SetSetMialbregasRobe                      // Mialbregas Robe
	SetCathansRule                            // Cathans Rule
	SetCathansMesh                            // Cathans Mesh
	SetCathansVisage                          // Cathans Visage
	SetCathansSigil                           // Cathans Sigil
	SetCathansSeal                            // Cathans Seal
	SetTancredsCrowbill                       // Tancreds Crowbill
	SetTancredsSpine                          // Tancreds Spine
	SetTancredsHobnails                       // Tancreds Hobnails
	SetTancredsWeird                          // Tancreds Weird
	SetTancredsSkull                          // Tancreds Skull
	SetSigonsGage                             // Sigons Gage
	SetSigonsVisor                            // Sigons Visor
	SetSigonsShelter                          // Sigons Shelter
	SetSigonsSabot                            // Sigons Sabot
	SetSigonsWrap                             // Sigons Wrap
	SetSigonsGuard                            // Sigons Guard
	SetInfernalCranium                        // Infernal Cranium
	SetInfernalTorch                          // Infernal Torch
	SetInfernalSign                           // Infernal Sign
	SetBerserkersHeadgear                     // Berserkers Headgear
	SetBerserkersHauberk                      // Berserkers Hauberk
	SetBerserkersHatchet                      // Berserkers Hatchet
	SetDeathsHand                             // Deaths Hand
	SetDeathsGuard                            // Deaths Guard
	SetDeathsTouch                            // Deaths Touch
	SetAngelicSickle                          // Angelic Sickle
	SetAngelicMantle                          // Angelic Mantle
	SetAngelicHalo                            // Angelic Halo
	SetAngelicWings                           // Angelic Wings
	SetArcticHorn                             // Arctic Horn
	SetArcticFurs                             // Arctic Furs
	SetArcticBinding                          // Arctic Binding
	SetArcticMitts                            // Arctic Mitts
	SetArcannasSign                           // Arcannas Sign
	SetArcannasDeathwand                      // Arcannas Deathwand
	SetArcannasHead                           // Arcannas Head
	SetArcannasFlesh                          // Arcannas Flesh
	SetNatalyasTotem                          // Natalyas Totem
	SetNatalyasMark                           // Natalyas Mark
	SetNatalyasShadow                         // Natalyas Shadow
	SetNatalyasSoul                           // Natalyas Soul
	SetAldursStonyGaze                        // Aldurs Stony Gaze
	SetAldursDeception                        // Aldurs Deception
	SetAldursRhythm                           // Aldurs Rhythm
	SetAldursAdvance                          // Aldurs Advance
	SetImmortalKingsWill                      // Immortal Kings Will
	SetImmortalKingsSoulCage                  // Immortal Kings Soul Cage
	SetImmortalKingsDetail                    // Immortal Kings Detail
	SetImmortalKingsForge                     // Immortal Kings Forge
	SetImmortalKingsPillar                    // Immortal Kings Pillar
	SetImmortalKingsStoneCrusher              // Immortal Kings Stone Crusher
	SetTalRashasFineSpunCloth                 // Tal Rashas Fine-Spun Cloth
	SetTalRashasAdjudication                  // Tal Rashas Adjudication
	SetTalRashasLidlessEye                    // Tal Rashas Lidless Eye
	SetTalRashasGuardianship                  // Tal Rashas Guardianship
	SetTalRashasHoradricCrest                 // Tal Rashas Horadric Crest
	SetGriswoldsValor                         // Griswolds Valor
	SetGriswoldsHeart                         // Griswolds Heart
	SetGriswoldsRedemption                    // Griswolds Redemption
	SetGriswoldsHonor                         // Griswolds Honor
	SetTrangOulsGuise                         // Trang-Ouls Guise
	SetTrangOulsScales                        // Trang-Ouls Scales
	SetTrangOulsWing                          // Trang-Ouls Wing
	SetTrangOulsClaws                         // Trang-Ouls Claws
	SetTrangOulsGirth                         // Trang-Ouls Girth
	SetMavinasTrueSight                       // Mavinas True Sight
	SetMavinasEmbrace                         // Mavinas Embrace
	SetMavinasIcyClutch                       // Mavinas Icy Clutch
	SetMavinasTenet                           // Mavinas Tenet
	SetMavinasCaster                          // Mavinas Caster
	SetTellingofBeads                         // Telling of Beads
	SetLayingofHands                          // Laying of Hands
	SetRiteofPassage                          // Rite of Passage
	SetDarkAdherent                           // Dark Adherent
	SetCredendum                              // Credendum
	SetDangoonsTeaching                       // Dangoons Teaching
	SetTaebaeksGlory                          // Taebaeks Glory
	SetHaemosusAdament                        // Haemosus Adament
	SetOndalsAlmighty                         // Ondals Almighty
	SetGuillaumesFace                         // Guillaumes Face
	SetWilhelmsPride                          // Wilhelms Pride
	SetMagnusSkin                             // Magnus Skin
	SetWihtstansGuard                         // Wihtstans Guard
	SetHwaninsSplendor                        // Hwanins Splendor
	SetHwaninsRefuge                          // Hwanins Refuge
	SetHwaninsBlessing                        // Hwanins Blessing
	SetHwaninsJustice                         // Hwanins Justice
	SetSazabisCobaltRedeemer                  // Sazabis Cobalt Redeemer
	SetSazabisGhostLiberator                  // Sazabis Ghost Liberator
	SetSazabisMentalSheath                    // Sazabis Mental Sheath
	SetBulKathosSacredCharge                  // Bul-Kathos Sacred Charge
	SetBulKathosTribalGuardian                // Bul-Kathos Tribal Guardian
	SetCowKingsHorns                          // Cow Kings Horns
	SetCowKingsHide                           // Cow Kings Hide
	SetCowKingsHooves                         // Cow Kings Hooves
	SetNajsPuzzler                            // Najs Puzzler
	SetNajsLightPlate                         // Najs Light Plate
	SetNajsCirclet                            // Najs Circlet
	SetSandersParagon                         // Sanders Paragon
	SetSandersRiprap                          // Sanders Riprap
	SetSandersTaboo                           // Sanders Taboo
	SetSandersSuperstition                    // Sanders Superstition
)
