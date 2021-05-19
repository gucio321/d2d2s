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

//go:generate stringer -linecomment -type SetID -output set_is_string.go

// SetID represents a set ID
type SetID uint16

// set IDs - TODO: give these constants some descriptive name
const (
	Set0   SetID = iota //  Civerb's Ward
	Set1                //  Civerb's Icon
	Set2                //  Civerb's Cudgel
	Set3                //  Hsaru's Iron Heel
	Set4                //  Hsaru's Iron Fist
	Set5                //  Hsaru's Iron Stay
	Set6                //  Cleglaw's Tooth
	Set7                //  Cleglaw's Claw
	Set8                //  Cleglaw's Pincers
	Set9                //  Iratha's Collar
	Set10               // Iratha's Cuff
	Set11               // Iratha's Coil
	Set12               // Iratha's Cord
	Set13               // Isenhart's Lightbrand
	Set14               // Isenhart's Parry
	Set15               // Isenhart's Case
	Set16               // Isenhart's Horns
	Set17               // Vidala's Barb
	Set18               // Vidala's Fetlock
	Set19               // Vidala's Ambush
	Set20               // Vidala's Snare
	Set21               // Milabrega's Orb
	Set22               // Milabrega's Rod
	Set23               // Milabrega's Diadem
	Set24               // Mialbrega's Robe
	Set25               // Cathan's Rule
	Set26               // Cathan's Mesh
	Set27               // Cathan's Visage
	Set28               // Cathan's Sigil
	Set29               // Cathan's Seal
	Set30               // Tancred's Crowbill
	Set31               // Tancred's Spine
	Set32               // Tancred's Hobnails
	Set33               // Tancred's Weird
	Set34               // Tancred's Skull
	Set35               // Sigon's Gage
	Set36               // Sigon's Visor
	Set37               // Sigon's Shelter
	Set38               // Sigon's Sabot
	Set39               // Sigon's Wrap
	Set40               // Sigon's Guard
	Set41               // Infernal Cranium
	Set42               // Infernal Torch
	Set43               // Infernal Sign
	Set44               // Berserker's Headgear
	Set45               // Berserker's Hauberk
	Set46               // Berserker's Hatchet
	Set47               // Death's Hand
	Set48               // Death's Guard
	Set49               // Death's Touch
	Set50               // Angelic Sickle
	Set51               // Angelic Mantle
	Set52               // Angelic Halo
	Set53               // Angelic Wings
	Set54               // Arctic Horn
	Set55               // Arctic Furs
	Set56               // Arctic Binding
	Set57               // Arctic Mitts
	Set58               // Arcanna's Sign
	Set59               // Arcanna's Deathwand
	Set60               // Arcanna's Head
	Set61               // Arcanna's Flesh
	Set62               // Natalya's Totem
	Set63               // Natalya's Mark
	Set64               // Natalya's Shadow
	Set65               // Natalya's Soul
	Set66               // Aldur's Stony Gaze
	Set67               // Aldur's Deception
	Set68               // Aldur's Rhythm
	Set69               // Aldur's Advance
	Set70               // Immortal King's Will
	Set71               // Immortal King's Soul Cage
	Set72               // Immortal King's Detail
	Set73               // Immortal King's Forge
	Set74               // Immortal King's Pillar
	Set75               // Immortal King's Stone Crusher
	Set76               // Tal Rasha's Fine-Spun Cloth
	Set77               // Tal Rasha's Adjudication
	Set78               // Tal Rasha's Lidless Eye
	Set79               // Tal Rasha's Guardianship
	Set80               // Tal Rasha's Horadric Crest
	Set81               // Griswold's Valor
	Set82               // Griswold's Heart
	Set83               // Griswold's Redemption
	Set84               // Griswold's Honor
	Set85               // Trang-Oul's Guise
	Set86               // Trang-Oul's Scales
	Set87               // Trang-Oul's Wing
	Set88               // Trang-Oul's Claws
	Set89               // Trang-Oul's Girth
	Set90               // M'avina's True Sight
	Set91               // M'avina's Embrace
	Set92               // M'avina's Icy Clutch
	Set93               // M'avina's Tenet
	Set94               // M'avina's Caster
	Set95               // Telling of Beads
	Set96               // Laying of Hands
	Set97               // Rite of Passage
	Set98               // Dark Adherent
	Set99               // Credendum
	Set100              // Dangoon's Teaching
	Set101              // Taebaek's Glory
	Set102              // Haemosu's Adament
	Set103              // Ondal's Almighty
	Set104              // Guillaume's Face
	Set105              // Wilhelm's Pride
	Set106              // Magnus' Skin
	Set107              // Wihtstan's Guard
	Set108              // Hwanin's Splendor
	Set109              // Hwanin's Refuge
	Set110              // Hwanin's Blessing
	Set111              // Hwanin's Justice
	Set112              // Sazabi's Cobalt Redeemer
	Set113              // Sazabi's Ghost Liberator
	Set114              // Sazabi's Mental Sheath
	Set115              // Bul-Katho's Sacred Charge
	Set116              // Bul-Katho's Tribal Guardian
	Set117              // Cow King's Horns
	Set118              // Cow King's Hide
	Set119              // Cow King's Hooves
	Set120              // Naj's Puzzler
	Set121              // Naj's Light Plate
	Set122              // Naj's Circlet
	Set123              // Sander's Paragon
	Set124              // Sander's Riprap
	Set125              // Sander's Taboo
	Set126              // Sander's Superstition
)
