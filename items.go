package d2d2s

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/gucio321/d2d2s/itemdata"
)

const (
	itemListID       = "JM"
	characterNameLen = 15
)

// Items represents items list
type Items []Item

func (i *Items) LoadHeader(sr *datautils.BitMuncher) (numItems uint16, err error) {
	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != itemListID {
		return 0, errors.New("unexpected item header")
	}

	numItems = sr.GetUInt16()

	return numItems, nil
}

// Load loads items list data into Items structure
// nolint:funlen // TODO: check, if it is possible to write encoder for d2s.Item
// If not, theis function must be changed
func (i *Items) LoadList(sr *datautils.BitMuncher, numItems uint16) error {
	*i = make([]Item, numItems)
	// note: if item has sockets, it is followed by item socketed in!
	for n := uint16(0); n < numItems; n++ {
		if err := (*i)[n].Load(sr); err != nil {
			return err
		}

		// if item is socketed into another item ( last on list) we need to append it
		if !(*i)[n].IsSimple {
			for s := byte(0); s < (*i)[n].NumberOfItemsInSockets; s++ {
				item := &Item{}
				if err := item.Load(sr); err != nil {
					return err
				}
				(*i)[n].SocketedItems = append((*i)[n].SocketedItems, *item)
			}
		}

	}

	return nil
}

type Item struct {
	unknown1  byte // 4 bits
	unknown2  byte // 6 bits
	unknown3  bool // 1 bit
	unknown4  byte // 2 bits
	unknown5  byte // 3 bits
	unknown6  bool // 1 bit
	unknown7  bool // 1 bit
	unknown8  byte // 5 bits
	unknown9  byte // 2 bits
	unknown10 bool // 1 bit
	unknown11 byte // 5 bits (tomes only)
	unknown12 bool // 1 bit

	// Part 1: simple item
	Identified bool
	JustPicked bool // This bit is set on items which you have picked up since the last time the game was saved.  Why?...
	IsEar      bool
	NewbieItem bool // it is set when th item is an item generated with a new character
	IsSimple   bool // only 111 bits of data
	Etheral    bool
	Version    byte // byte
	Location   struct {
		LocationID ItemLocationType  // 3 bits
		EquippedID ItemEquippedPlace // 4 bits
		X          byte              // 4 bits
		Y          byte              // 3 bits
		StorageID  StoragePlace      // 3 bits
	}
	Ear struct {
		Class CharacterClass // 3 bits
		Level byte           // 7 bits
		Name  string         // len(Name) * 7 bits
	}
	Type       string
	TypeID     itemdata.ItemTypeID
	TypeName   string
	BaseDamage struct {
		Min        int
		Max        int
		TwoHandMin int
		TwoHandMax int
	}
	NumberOfItemsInSockets byte

	// Part 2; extended
	ID              uint32      // 32 bits (just uint32)
	Level           byte        // 7 bits
	Quality         ItemQuality // 4 bits
	MultiplePicture struct {
		HasMultiplePicture bool // 1 bit
		ID                 byte // 7 bits
	}
	ClassSpecific struct {
		IsClassSpecific bool   // 1 bit
		Data            uint16 // 11 bits
	}
	QualityData struct {
		LowQualityID byte // 3 bits

		HighQualityData byte // 3 bits (TODO)

		// Magically enchanced, Rare, Crafted
		// in case of magically enchanced items size is 11,
		// in case of rare/crafted - 12 * len()
		MagicPrefix []MagicModifier
		MagicSuffix []MagicModifier
		// Rare, Crafted only
		RareNames []MagicModifier // len() * 8 bits

		SetID         uint16 // 12 bits
		SetName       string
		SetListID     byte // 5 bits
		SetListCount  uint64
		SetAttributes []MagicAttributes

		UniqueID   uint16 // 12 bits
		UniqueName string
	}
	RuneWord struct {
		HasRuneWord bool
		ID          uint16
		Name        string
		unknown     byte // 4 bits (brobably always value 5)
		Attributes  MagicAttributes
	}
	Personalization struct {
		IsPersonalized bool
		Name           string
	}
	Timestamp  bool
	ExtraStats struct {
		DefenseRating uint16 // 11 bits //  warning: = defRate - 10
		Durability    struct {
			Max     byte // 8 bits
			Current byte // 8 bits
		}
		Quantity uint16 // 9 bits
	}
	Socketed struct {
		IsInSocket           bool
		TotalNumberOfSockets byte // 4 bits
	}

	Attributes    MagicAttributes
	SocketedItems []Item
}

func (i *Item) Load(sr *datautils.BitMuncher) (err error) {
	if err := i.loadSimpleFields(sr); err != nil {
		return err
	}

	// if is not a simple item, continue work
	if !i.IsSimple {
		if err := i.loadExtendedFields(sr); err != nil {
			return err
		}
	}

	sr.AlignToBytes()

	return nil
}

func (i *Item) loadExtendedFields(sr *datautils.BitMuncher) (err error) {
	i.ID = sr.GetUInt32() // probably 4 * 8 chars
	i.Level = byte(sr.GetBits(7))
	i.Quality = ItemQuality(sr.GetBits(4))

	// multiple picture
	i.MultiplePicture.HasMultiplePicture = sr.GetBit() == 1
	if i.MultiplePicture.HasMultiplePicture {
		i.MultiplePicture.ID = byte(sr.GetBits(3))
	}

	// class specific
	i.ClassSpecific.IsClassSpecific = sr.GetBit() == 1
	if i.ClassSpecific.IsClassSpecific {
		// probably class is somewhere here ... ?
		i.ClassSpecific.Data = uint16(sr.GetBits(11))
	}

	switch i.Quality {
	case ItemQualityLow:
		i.QualityData.LowQualityID = byte(sr.GetBits(3))
	case ItemQualityNormal:
		// noop
	case ItemQualityHigh:
		i.QualityData.HighQualityData = byte(sr.GetBits(3))
	case ItemQualityEnchanced:
		i.QualityData.MagicPrefix = make([]MagicModifier, 1)
		id := uint16(sr.GetBits(11)) // helper variable (avoid noise with x.y.z.id ;-)
		i.QualityData.MagicPrefix[0].ID = id
		prefixName, ok := itemdata.MagicalPrefixes[id]
		if ok {
			i.QualityData.MagicPrefix[0].Name = prefixName
		}

		i.QualityData.MagicSuffix = make([]MagicModifier, 1)
		id = uint16(sr.GetBits(11)) // helper variable (avoid noise with x.y.z.id ;-)
		i.QualityData.MagicSuffix[0].ID = id
		suffixName, ok := itemdata.MagicalSuffixes[id]
		if ok {
			i.QualityData.MagicSuffix[0].Name = suffixName
		}
	case ItemQualitySet:
		id := uint16(sr.GetBits(12))
		i.QualityData.SetID = id
		setName, ok := itemdata.SetNames[id]
		if ok {
			i.QualityData.SetName = setName
		}
	case ItemQualityRare, ItemQualityCrafted:
		i.QualityData.RareNames = make([]MagicModifier, 2)
		for n := 0; n < 2; n++ {
			id := uint16(sr.GetBits(8))
			i.QualityData.RareNames[n].ID = id
			name, ok := itemdata.RareNames[id]
			if ok {
				i.QualityData.RareNames[n].Name = name
			}
		}

		i.QualityData.MagicPrefix = make([]MagicModifier, 0)
		i.QualityData.MagicSuffix = make([]MagicModifier, 0)
		for p := 0; p < 3; p++ { // nolint:gomnd // 3 sets of modifiers allowed
			prefixExist := sr.GetBit() == 1
			if prefixExist {
				i.QualityData.MagicPrefix = append(i.QualityData.MagicPrefix, MagicModifier{})
				id := uint16(sr.GetBits(11))
				i.QualityData.MagicPrefix[p].ID = id
				name, ok := itemdata.MagicalPrefixes[id]
				if ok {
					i.QualityData.MagicPrefix[p].Name = name
				}
			}

			suffixExist := sr.GetBit() == 1
			if suffixExist {
				i.QualityData.MagicSuffix = append(i.QualityData.MagicPrefix, MagicModifier{})
				id := uint16(sr.GetBits(11))
				i.QualityData.MagicSuffix[p].ID = id
				name, ok := itemdata.MagicalSuffixes[id]
				if ok {
					i.QualityData.MagicSuffix[p].Name = name
				}
			}
		}
	case ItemQualityUnique:
		id := uint16(sr.GetBits(12))
		i.QualityData.UniqueID = id
		uniqueName, ok := itemdata.UniqueNames[id]
		if ok {
			i.QualityData.UniqueName = uniqueName
		}
	}

	if i.RuneWord.HasRuneWord {
		id := uint16(sr.GetBits(12))
		i.RuneWord.ID = id
		name, ok := itemdata.RunewordNames[id]
		if ok {
			i.RuneWord.Name = name
		}

		i.RuneWord.unknown = byte(sr.GetBits(4))
	}

	if i.Personalization.IsPersonalized {
		name := make([]byte, 0)
		for {
			char := byte(sr.GetBits(7))
			if char == 0 {
				break
			}

			name = append(name, char)
		}

		i.Personalization.Name = string(name)
	}

	if itemdata.TomeMap[i.Type] {
		i.unknown11 = byte(sr.GetBits(5))
	}

	i.Timestamp = sr.GetBit() == 1

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDShield {
		defRating := uint16(sr.GetBits(11))
		i.ExtraStats.DefenseRating = defRating - 10
	}

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDWeapon ||
		i.TypeID == itemdata.ItemTypeIDShield {
		i.ExtraStats.Durability.Max = sr.GetByte()
		if i.ExtraStats.Durability.Max > 0 {
			i.ExtraStats.Durability.Current = sr.GetByte()
			i.unknown12 = sr.GetBit() == 1
		}
	}

	if itemdata.QuantityMap[i.Type] {
		i.ExtraStats.Quantity = uint16(sr.GetBits(9))
	}

	if i.Socketed.IsInSocket {
		i.Socketed.TotalNumberOfSockets = byte(sr.GetBits(4))
	}

	var setListValue byte = 0
	if i.Quality == ItemQualitySet {
		setListValue = byte(sr.GetBits(5))
		listCount, ok := itemdata.SetListMap[setListValue]
		i.QualityData.SetListID = setListValue
		if !ok {
			return fmt.Errorf("unknown set list value %d", setListValue)
		}
		i.QualityData.SetListCount = listCount
	}

	if err := i.Attributes.Load(sr); err != nil {
		return err
	}

	if c := i.QualityData.SetListCount; c > 0 {
		for j := 0; j < int(c); j++ {
			i.QualityData.SetAttributes = append(i.QualityData.SetAttributes, MagicAttributes{})
			i.QualityData.SetAttributes[len(i.QualityData.SetAttributes)-1].Load(sr)
		}
		// here could be interpretation of above data
		// however it is optional because doesn't affect stream
	}

	if i.RuneWord.HasRuneWord {
		i.RuneWord.Attributes.Load(sr)
	}

	return nil
}

func (i *Item) loadSimpleFields(sr *datautils.BitMuncher) (err error) {
	id := sr.GetBytes(2)
	if string(id) != itemListID {
		return errors.New("unexpected item signature")
	}

	i.unknown1 = byte(sr.GetBits(4))
	i.Identified = sr.GetBit() == 1
	i.unknown2 = byte(sr.GetBits(6))
	i.Socketed.IsInSocket = sr.GetBit() == 1
	i.unknown3 = sr.GetBit() == 1
	i.JustPicked = sr.GetBit() == 1
	i.unknown4 = byte(sr.GetBits(2))
	i.IsEar = sr.GetBit() == 1
	i.NewbieItem = sr.GetBit() == 1
	i.unknown5 = byte(sr.GetBits(3))
	i.IsSimple = sr.GetBit() == 1
	i.Etheral = sr.GetBit() == 1
	i.unknown6 = sr.GetBit() == 1
	i.Personalization.IsPersonalized = sr.GetBit() == 1
	i.unknown7 = sr.GetBit() == 1
	i.RuneWord.HasRuneWord = sr.GetBit() == 1
	i.unknown8 = byte(sr.GetBits(5))
	i.Version = sr.GetByte()
	i.unknown9 = byte(sr.GetBits(2))
	i.Location.LocationID = ItemLocationType(sr.GetBits(3))
	i.Location.EquippedID = ItemEquippedPlace(sr.GetBits(4))
	i.Location.X = byte(sr.GetBits(4))
	i.Location.Y = byte(sr.GetBits(3))
	i.unknown10 = sr.GetBit() == 1
	i.Location.StorageID = StoragePlace(sr.GetBits(3))

	if i.IsEar {
		i.Ear.Class = CharacterClass(sr.GetBits(3))
		i.Ear.Level = byte(sr.GetBits(7))
		var name []byte
		for {
			char := byte(sr.GetBits(7))
			if char == 0 {
				break
			}

			name = append(name, char)
		}

		i.Ear.Name = string(name)

		sr.AlignToBytes()
	} else {
		i.Type = strings.Trim(string(sr.GetBytes(4)), " ")
		i.TypeID = itemdata.GetTypeID(i.Type)
		switch i.TypeID {
		case itemdata.ItemTypeIDArmor:
			typeName, ok := itemdata.ArmorCodes[i.Type]
			if ok {
				i.TypeName = typeName
			}
		case itemdata.ItemTypeIDShield:
			typeName, ok := itemdata.ShieldCodes[i.Type]
			if ok {
				i.TypeName = typeName
			}
		case itemdata.ItemTypeIDWeapon:
			typeName, ok := itemdata.WeaponCodes[i.Type]
			if ok {
				i.TypeName = typeName
			}

			baseDamage, ok := itemdata.WeaponDamageMap[i.Type]
			if ok {
				// If the item is ethereal we need to add 50% enhanced
				// damage to the base damage.
				if i.Etheral {
					i.BaseDamage.Min = int((float64(baseDamage.Min) * 1.5))
					i.BaseDamage.Max = int((float64(baseDamage.Max) * 1.5))
					i.BaseDamage.TwoHandMin = int((float64(baseDamage.TwoMin) * 1.5))
					i.BaseDamage.TwoHandMax = int((float64(baseDamage.TwoMax) * 1.5))
				} else {
					i.BaseDamage.Min = baseDamage.Min
					i.BaseDamage.Max = baseDamage.Max
					i.BaseDamage.TwoHandMin = baseDamage.TwoMin
					i.BaseDamage.TwoHandMax = baseDamage.TwoMax
				}
			}
		case itemdata.ItemTypeIDOther:
			typeName, ok := itemdata.MiscCodes[i.Type]
			if ok {
				i.TypeName = typeName
			}
		}

		i.NumberOfItemsInSockets = byte(sr.GetBits(3))
	}

	return nil
}

// ItemLocationType represents an item location
type ItemLocationType byte

// item locations
const (
	ItemLocationStored ItemLocationType = iota
	ItemLocationEquipped
	ItemLocationBelt
	_
	ItemLocationMoved // i.e., has been picked up by the mouse
	_
	ItemLocationInSocket // ?
)

// ItemEquippedPlace represents a place, where the item is equipped
type ItemEquippedPlace byte

// item equipped places
const (
	ItemEquippedAnywhere ItemEquippedPlace = iota
	ItemEquippedHead
	ItemEquippedNeck
	ItemEquippedTorse
	ItemEquippedRightHand
	ItemEquippedLeftHand
	ItemEquippedRightFinger
	ItemEquippedLeftFinger
	ItemEquippedWaist
	ItemEquippedFeet
	ItemEquippedHands
	ItemEquippedARightHand
	ItemEquippedALeftHand
)

// StoragePlace represents a plece, where the item is stored
type StoragePlace byte

// storage places
const (
	StorageNone StoragePlace = iota
	StorageInventory
	_
	_
	StorageCube
	StorageStash
)

// ItemQuality represents a quality of an item
type ItemQuality byte

// item qualities
const (
	ItemQualityLow ItemQuality = iota + 1
	ItemQualityNormal
	ItemQualityHigh
	ItemQualityEnchanced
	ItemQualitySet
	ItemQualityRare
	ItemQualityUnique
	ItemQualityCrafted
)

// LowQualityItemType represents a type of low quality item
type LowQualityItemType byte

// low quality item types
const (
	LowQualityItemCrude LowQualityItemType = iota
	LowQualityItemCracked
	LowQualityItemDamaged
	LowQualityItemLowQuality
)

type MagicModifier struct {
	ID   uint16
	Name string
}
