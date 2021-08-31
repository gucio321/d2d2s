package itemdata

// ItemTypeID represents an item type ID
type ItemTypeID byte

// Item types, used to decide what attribute to give an item socketed with
// gems or runes mostly.
const (
	ItemTypeIDArmor ItemTypeID = iota
	ItemTypeIDShield
	ItemTypeIDWeapon
	ItemTypeIDOther
)

// GetTypeID returns type id basing on item code given
func (c ItemCode) GetTypeID() ItemTypeID {
	if _, ok := ArmorCodes[c]; ok {
		return ItemTypeIDArmor
	}

	if _, ok := ShieldCodes[c]; ok {
		return ItemTypeIDShield
	}

	if _, ok := WeaponCodes[c]; ok {
		return ItemTypeIDWeapon
	}

	return ItemTypeIDOther
}
