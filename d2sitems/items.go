package d2sitems

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/d2sitems/d2smagicattributes"
	"github.com/gucio321/d2d2s/d2sitems/itemdata"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	byteLen   = 8
	uint32Len = 4 * byteLen

	itemListID = "JM"

	etheralItemsDamageModifier = 1.5
	defenseRatingModifier      = 10

	unknown1Len  = 4
	unknown2Len  = 6
	unknown4Len  = 2
	unknown5Len  = 3
	unknown8Len  = 5
	unknown9Len  = 2
	unknown11Len = 5

	locationLocationIDLen = 3
	locationEquippedIDLen = 4
	locationXLen          = 4
	locationYLen          = 3
	locationStorageIDLen  = 3
	// ear
	earClassLen          = 3
	earLevelLen          = levelLen
	numItemsInSocketsLen = 3

	typeLen              = 4
	levelLen             = 7
	qualityLen           = 4
	multiplePictureIDLen = 3
	classSpecificDataLen = 11
	lowQualityIDLen      = 3
	highQualityDataLen   = 3
	setIDLen             = 12
	setListIDLen         = 5
	uniqueIDLen          = 12
	runeWordIDLen        = 12
	runeWordUnknownLen   = 4
	defenseRatingLen     = 11
	quantityLen          = 9
	totalNSocketsLen     = 4

	characterLen       = 7
	magicModifierIDLen = 11
	rareNameIDLen      = 8
)

// Items represents items list
type Items []Item

// LoadHeader loads items header and returns items count
func (i *Items) LoadHeader(sr *datautils.BitMuncher) (numItems uint16, err error) {
	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != itemListID {
		return 0, errors.New("unexpected item header")
	}

	numItems = sr.GetUInt16()

	return numItems, nil
}

// LoadList loads items list data into Items structure
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

// Encode encodes items list back into a byte slice
func (i *Items) Encode() []byte {
	sw := datautils.CreateStreamWriter()

	// header
	sw.PushBytes([]byte(itemListID)...)
	sw.PushUint16(uint16(len(*i)))

	for n := range *i {
		// for n := 0; n < 42; n++ {
		item := (*i)[n]
		sw.PushBytes(item.Encode()...)

		for s := 0; s < len(item.SocketedItems); s++ {
			sw.PushBytes(item.SocketedItems[s].Encode()...)
		}
	}

	return sw.GetBytes()
}

// Item represents an item
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
		LocationID d2senums.ItemLocationType  // 3 bits
		EquippedID d2senums.ItemEquippedPlace // 4 bits
		X          byte                       // 4 bits
		Y          byte                       // 3 bits
		StorageID  d2senums.StoragePlace      // 3 bits
	}
	Ear struct {
		Class d2senums.CharacterClass // 3 bits
		Level byte                    // 7 bits
		Name  string                  // len(Name) * 7 bits
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

	NumberOfItemsInSockets byte // 3 bits

	// Part 2; extended
	ID              uint32               // 32 bits (just uint32)
	Level           byte                 // 7 bits
	Quality         d2senums.ItemQuality // 4 bits
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
		SetAttributes []d2smagicattributes.MagicAttributes

		UniqueID   uint16 // 12 bits
		UniqueName string
	}
	RuneWord struct {
		HasRuneWord bool
		ID          uint16 // 12 bits
		Name        string
		unknown     byte // 4 bits (brobably always value 5)
		Attributes  d2smagicattributes.MagicAttributes
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

	Attributes    d2smagicattributes.MagicAttributes
	SocketedItems []Item
}

// Load loads an item
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

// nolint:funlen,gocognit,gocyclo // will reduce later
func (i *Item) loadExtendedFields(sr *datautils.BitMuncher) (err error) {
	i.ID = sr.GetUInt32() // probably 4 * 8 chars
	i.Level = byte(sr.GetBits(levelLen))
	i.Quality = d2senums.ItemQuality(sr.GetBits(qualityLen))

	// multiple picture
	i.MultiplePicture.HasMultiplePicture = sr.GetBit() == 1
	if i.MultiplePicture.HasMultiplePicture {
		i.MultiplePicture.ID = byte(sr.GetBits(multiplePictureIDLen))
	}

	// class specific
	i.ClassSpecific.IsClassSpecific = sr.GetBit() == 1
	if i.ClassSpecific.IsClassSpecific {
		// probably class is somewhere here ... ?
		i.ClassSpecific.Data = uint16(sr.GetBits(classSpecificDataLen))
	}

	switch i.Quality {
	case d2senums.ItemQualityLow:
		i.QualityData.LowQualityID = byte(sr.GetBits(lowQualityIDLen))
	case d2senums.ItemQualityNormal:
		// noop
	case d2senums.ItemQualityHigh:
		i.QualityData.HighQualityData = byte(sr.GetBits(highQualityDataLen))
	case d2senums.ItemQualityEnchanced:
		i.QualityData.MagicPrefix = make([]MagicModifier, 1)
		id := uint16(sr.GetBits(magicModifierIDLen)) // helper variable (avoid noise with x.y.z.id ;-)
		i.QualityData.MagicPrefix[0].ID = id
		prefixName, ok := itemdata.MagicalPrefixes[id]

		if ok {
			i.QualityData.MagicPrefix[0].Name = prefixName
		}

		i.QualityData.MagicSuffix = make([]MagicModifier, 1)
		id = uint16(sr.GetBits(magicModifierIDLen)) // helper variable (avoid noise with x.y.z.id ;-)
		i.QualityData.MagicSuffix[0].ID = id
		suffixName, ok := itemdata.MagicalSuffixes[id]

		if ok {
			i.QualityData.MagicSuffix[0].Name = suffixName
		}
	case d2senums.ItemQualitySet:
		id := uint16(sr.GetBits(setIDLen))
		i.QualityData.SetID = id
		setName, ok := itemdata.SetNames[id]

		if ok {
			i.QualityData.SetName = setName
		}
	case d2senums.ItemQualityRare, d2senums.ItemQualityCrafted:
		i.QualityData.RareNames = make([]MagicModifier, 2)

		for n := 0; n < 2; n++ {
			id := uint16(sr.GetBits(rareNameIDLen))
			i.QualityData.RareNames[n].ID = id
			name, ok := itemdata.RareNames[id]

			if ok {
				i.QualityData.RareNames[n].Name = name
			}
		}

		i.QualityData.MagicPrefix = make([]MagicModifier, 0)
		i.QualityData.MagicSuffix = make([]MagicModifier, 0)

		for p := 0; p < 3; p++ {
			prefixExist := sr.GetBit() == 1
			if prefixExist {
				i.QualityData.MagicPrefix = append(i.QualityData.MagicPrefix, MagicModifier{})
				id := uint16(sr.GetBits(magicModifierIDLen))
				i.QualityData.MagicPrefix[p].ID = id
				name, ok := itemdata.MagicalPrefixes[id]

				if ok {
					i.QualityData.MagicPrefix[p].Name = name
				}
			}

			suffixExist := sr.GetBit() == 1
			if suffixExist {
				i.QualityData.MagicSuffix = append(i.QualityData.MagicSuffix, MagicModifier{})
				id := uint16(sr.GetBits(magicModifierIDLen))
				i.QualityData.MagicSuffix[p].ID = id
				name, ok := itemdata.MagicalSuffixes[id]

				if ok {
					i.QualityData.MagicSuffix[p].Name = name
				}
			}
		}
	case d2senums.ItemQualityUnique:
		id := uint16(sr.GetBits(uniqueIDLen))
		i.QualityData.UniqueID = id
		uniqueName, ok := itemdata.UniqueNames[id]

		if ok {
			i.QualityData.UniqueName = uniqueName
		}
	}

	if i.RuneWord.HasRuneWord {
		id := uint16(sr.GetBits(runeWordIDLen))
		i.RuneWord.ID = id
		name, ok := itemdata.RunewordNames[id]

		if ok {
			i.RuneWord.Name = name
		}

		i.RuneWord.unknown = byte(sr.GetBits(runeWordUnknownLen))
	}

	if i.Personalization.IsPersonalized {
		name := make([]byte, 0)

		for {
			char := byte(sr.GetBits(characterLen))
			if char == 0 {
				break
			}

			name = append(name, char)
		}

		i.Personalization.Name = string(name)
	}

	if itemdata.TomeMap[i.Type] {
		i.unknown11 = byte(sr.GetBits(unknown11Len))
	}

	i.Timestamp = sr.GetBit() == 1

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDShield {
		defRating := uint16(sr.GetBits(defenseRatingLen))
		i.ExtraStats.DefenseRating = defRating - defenseRatingModifier
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
		i.ExtraStats.Quantity = uint16(sr.GetBits(quantityLen))
	}

	if i.Socketed.IsInSocket {
		i.Socketed.TotalNumberOfSockets = byte(sr.GetBits(totalNSocketsLen))
	}

	var setListValue byte
	if i.Quality == d2senums.ItemQualitySet {
		setListValue = byte(sr.GetBits(setListIDLen))
		listCount, ok := itemdata.SetListMap[setListValue]
		i.QualityData.SetListID = setListValue

		if !ok {
			return fmt.Errorf("unknown set list value %d", setListValue)
		}

		i.QualityData.SetListCount = listCount
	}

	if err := i.Attributes.Load(sr); err != nil {
		return fmt.Errorf("error loading item attributes: %w", err)
	}

	// nolint:wsl // note at the end of block
	if c := i.QualityData.SetListCount; c > 0 {
		for j := 0; j < int(c); j++ {
			i.QualityData.SetAttributes = append(i.QualityData.SetAttributes, d2smagicattributes.MagicAttributes{})
			if err := i.QualityData.SetAttributes[len(i.QualityData.SetAttributes)-1].Load(sr); err != nil {
				return fmt.Errorf("error loading set attributes: %w", err)
			}
		}

		// here could be interpretation of above data
		// however it is optional because doesn't affect stream
	}

	if i.RuneWord.HasRuneWord {
		if err := i.RuneWord.Attributes.Load(sr); err != nil {
			return fmt.Errorf("error loading runeword attributes: %w", err)
		}
	}

	return nil
}

func (i *Item) loadSimpleFields(sr *datautils.BitMuncher) (err error) {
	id := sr.GetBytes(2) // nolint:gomnd // header
	if string(id) != itemListID {
		return errors.New("unexpected item signature")
	}

	i.unknown1 = byte(sr.GetBits(unknown1Len))
	i.Identified = sr.GetBit() == 1
	i.unknown2 = byte(sr.GetBits(unknown2Len))
	i.Socketed.IsInSocket = sr.GetBit() == 1
	i.unknown3 = sr.GetBit() == 1
	i.JustPicked = sr.GetBit() == 1
	i.unknown4 = byte(sr.GetBits(unknown4Len))
	i.IsEar = sr.GetBit() == 1
	i.NewbieItem = sr.GetBit() == 1
	i.unknown5 = byte(sr.GetBits(unknown5Len))
	i.IsSimple = sr.GetBit() == 1
	i.Etheral = sr.GetBit() == 1
	i.unknown6 = sr.GetBit() == 1
	i.Personalization.IsPersonalized = sr.GetBit() == 1
	i.unknown7 = sr.GetBit() == 1
	i.RuneWord.HasRuneWord = sr.GetBit() == 1
	i.unknown8 = byte(sr.GetBits(unknown8Len))
	i.Version = sr.GetByte()
	i.unknown9 = byte(sr.GetBits(unknown9Len))
	i.Location.LocationID = d2senums.ItemLocationType(sr.GetBits(locationLocationIDLen))
	i.Location.EquippedID = d2senums.ItemEquippedPlace(sr.GetBits(locationEquippedIDLen))
	i.Location.X = byte(sr.GetBits(locationXLen))
	i.Location.Y = byte(sr.GetBits(locationYLen))
	i.unknown10 = sr.GetBit() == 1
	i.Location.StorageID = d2senums.StoragePlace(sr.GetBits(locationStorageIDLen))

	if i.IsEar {
		i.Ear.Class = d2senums.CharacterClass(sr.GetBits(earClassLen))
		i.Ear.Level = byte(sr.GetBits(earLevelLen))

		var name []byte

		for {
			char := byte(sr.GetBits(characterLen))
			if char == 0 {
				break
			}

			name = append(name, char)
		}

		i.Ear.Name = string(name)

		sr.AlignToBytes()
	} else {
		t := sr.GetBytes(typeLen)
		i.Type = strings.Trim(string(t), " ")
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
					i.BaseDamage.Min = int((float64(baseDamage.Min) * etheralItemsDamageModifier))
					i.BaseDamage.Max = int((float64(baseDamage.Max) * etheralItemsDamageModifier))
					i.BaseDamage.TwoHandMin = int((float64(baseDamage.TwoMin) * etheralItemsDamageModifier))
					i.BaseDamage.TwoHandMax = int((float64(baseDamage.TwoMax) * etheralItemsDamageModifier))
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

		i.NumberOfItemsInSockets = byte(sr.GetBits(numItemsInSocketsLen))
	}

	return nil
}

// Encode encodes an item into a byte slice
func (i *Item) Encode() []byte {
	sw := datautils.CreateStreamWriter()

	i.encodeSimpleFields(sw)

	if !i.IsSimple {
		if err := i.encodeExtendedFields(sw); err != nil {
			log.Fatal(err)
		}
	}

	sw.Align()

	return sw.GetBytes()
}

// nolint:funlen,gocognit,gocyclo // will reduce later
func (i *Item) encodeExtendedFields(sw *datautils.StreamWriter) (err error) {
	sw.PushBits32(i.ID, uint32Len)
	sw.PushBits(i.Level, levelLen)
	sw.PushBits(byte(i.Quality), qualityLen)
	sw.PushBit(i.MultiplePicture.HasMultiplePicture)

	if i.MultiplePicture.HasMultiplePicture {
		sw.PushBits(i.MultiplePicture.ID, 3)
	}

	sw.PushBit(i.ClassSpecific.IsClassSpecific)

	if i.ClassSpecific.IsClassSpecific {
		sw.PushBits16(i.ClassSpecific.Data, classSpecificDataLen)
	}

	switch i.Quality {
	case d2senums.ItemQualityLow:
		sw.PushBits(i.QualityData.LowQualityID, lowQualityIDLen)
	case d2senums.ItemQualityNormal:
		// noop
	case d2senums.ItemQualityHigh:
		sw.PushBits(i.QualityData.HighQualityData, highQualityDataLen)
	case d2senums.ItemQualityEnchanced:
		sw.PushBits16(i.QualityData.MagicPrefix[0].ID, 11)
		sw.PushBits16(i.QualityData.MagicSuffix[0].ID, 11)
	case d2senums.ItemQualitySet:
		sw.PushBits16(i.QualityData.SetID, setIDLen)
	case d2senums.ItemQualityRare, d2senums.ItemQualityCrafted:
		for _, name := range i.QualityData.RareNames {
			sw.PushBits(byte(name.ID), 8)
		}

		for n := 0; n < 3; n++ {
			l := len(i.QualityData.MagicPrefix)
			hasPrefix := n < l
			sw.PushBit(hasPrefix)

			if hasPrefix {
				sw.PushBits16(i.QualityData.MagicPrefix[n].ID, 11)
			}

			l = len(i.QualityData.MagicSuffix)
			hasSuffix := n < l
			sw.PushBit(hasSuffix)

			if hasSuffix {
				sw.PushBits16(i.QualityData.MagicSuffix[n].ID, 11)
			}
		}
	case d2senums.ItemQualityUnique:
		sw.PushBits16(i.QualityData.UniqueID, uniqueIDLen)
	}

	if i.RuneWord.HasRuneWord {
		sw.PushBits16(i.RuneWord.ID, runeWordIDLen)
		sw.PushBits(i.RuneWord.unknown, runeWordUnknownLen)
	}

	if i.Personalization.IsPersonalized {
		name := []byte(i.Personalization.Name)
		for _, c := range name {
			sw.PushBits(c, characterLen)
		}

		sw.PushBits(byte(' '), characterLen)
	}

	if itemdata.TomeMap[i.Type] {
		sw.PushBits(i.unknown11, unknown11Len)
	}

	sw.PushBit(i.Timestamp)

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDShield {
		sw.PushBits16(i.ExtraStats.DefenseRating+defenseRatingModifier, defenseRatingLen)
	}

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDWeapon ||
		i.TypeID == itemdata.ItemTypeIDShield {
		sw.PushBits(i.ExtraStats.Durability.Max, byteLen)

		if i.ExtraStats.Durability.Max > 0 {
			sw.PushBits(i.ExtraStats.Durability.Current, byteLen)
			sw.PushBit(i.unknown12)
		}
	}

	if itemdata.QuantityMap[i.Type] {
		sw.PushBits16(i.ExtraStats.Quantity, quantityLen)
	}

	if i.Socketed.IsInSocket {
		sw.PushBits(i.Socketed.TotalNumberOfSockets, totalNSocketsLen)
	}

	if i.Quality == d2senums.ItemQualitySet {
		sw.PushBits(i.QualityData.SetListID, setListIDLen)
	}

	if err := i.Attributes.Encode(sw); err != nil {
		return fmt.Errorf("error encoding item attributes: %w", err)
	}

	// OFC length of set attributes is > 0 fo sets
	for _, a := range i.QualityData.SetAttributes {
		if err := a.Encode(sw); err != nil {
			return fmt.Errorf("error encoding set attributes: %w", err)
		}
	}

	if i.RuneWord.HasRuneWord {
		if err := i.RuneWord.Attributes.Encode(sw); err != nil {
			return fmt.Errorf("error encoding runeword's attributes: %w", err)
		}
	}

	return nil
}

func (i *Item) encodeSimpleFields(sw *datautils.StreamWriter) {
	sw.PushBytes([]byte(itemListID)...)
	sw.PushBits(i.unknown1, unknown1Len)
	sw.PushBit(i.Identified)
	sw.PushBits(i.unknown2, unknown2Len)
	sw.PushBit(i.Socketed.IsInSocket)
	sw.PushBit(i.unknown3)
	sw.PushBit(i.JustPicked)
	sw.PushBits(i.unknown4, unknown4Len)
	sw.PushBit(i.IsEar)
	sw.PushBit(i.NewbieItem)
	sw.PushBits(i.unknown5, unknown5Len)
	sw.PushBit(i.IsSimple)
	sw.PushBit(i.Etheral)
	sw.PushBit(i.unknown6)
	sw.PushBit(len(i.Personalization.Name) > 0)
	sw.PushBit(i.unknown7)
	sw.PushBit(i.RuneWord.HasRuneWord)
	sw.PushBits(i.unknown8, unknown8Len)
	sw.PushBytes(i.Version)
	sw.PushBits(i.unknown9, unknown9Len)
	sw.PushBits(byte(i.Location.LocationID), locationLocationIDLen)
	sw.PushBits(byte(i.Location.EquippedID), locationEquippedIDLen)
	sw.PushBits(i.Location.X, locationXLen)
	sw.PushBits(i.Location.Y, locationYLen)
	sw.PushBit(i.unknown10)
	sw.PushBits(byte(i.Location.StorageID), locationStorageIDLen)

	if i.IsEar {
		sw.PushBits(byte(i.Ear.Class), earClassLen)
		sw.PushBits(i.Ear.Level, earLevelLen)

		for _, c := range i.Ear.Name {
			sw.PushBits(byte(c), characterLen)
		}

		sw.PushBits(0, characterLen)
		sw.Align()
	} else {
		name := []byte(i.Type)
		for _, c := range name {
			sw.PushBits(c, 8)
		}
		sw.PushBits(byte(' '), 8)
		sw.PushBits(i.NumberOfItemsInSockets, numItemsInSocketsLen)
	}
}

// MagicModifier represents a magic modifier
type MagicModifier struct {
	ID   uint16
	Name string
}
