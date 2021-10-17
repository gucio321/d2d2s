package d2sitems

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems/d2smagicattributes"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems/itemdata"
)

const (
	byteLen   = 8
	uint32Len = 4 * byteLen

	itemListID     = "JM"
	numHeaderBytes = 2

	defenseRatingModifier = 10

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

	rareNamesCount = 2
)

// New creates a new Items list
func New() *Items {
	result := &Items{}
	*result = make([]*Item, 0)

	return result
}

// Items represents items list
type Items []*Item

// LoadHeader loads items header and returns items count
func (i *Items) LoadHeader(sr *datareader.Reader) (numItems uint16, err error) {
	id := sr.GetBytes(numHeaderBytes)
	if string(id) != itemListID {
		return 0, fmt.Errorf("item header list: %w", common.ErrUnexpectedHeader)
	}

	numItems = sr.GetUint16()

	return numItems, nil
}

// LoadList loads items list data into Items structure
// If not, theis function must be changed
func (i *Items) LoadList(sr *datareader.Reader, numItems uint16) error {
	// note: if item has sockets, it is followed by item socketed in!
	for n := uint16(0); n < numItems; n++ {
		item := &Item{}
		if err := item.Load(sr); err != nil {
			return err
		}

		i.Add(item)

		// if item is socketed into another item ( last on list) we need to append it
		if !(*i)[n].IsSimple {
			for s := byte(0); s < (*i)[n].NumberOfItemsInSockets; s++ {
				item := &Item{}
				if err := item.Load(sr); err != nil {
					return err
				}

				(*i)[n].SocketedItems = append((*i)[n].SocketedItems, item)
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

// Add adds an item / items given
func (i *Items) Add(items ...*Item) {
	*i = append(*i, items...)
}

// DeleteAt deletes an item at given index
func (i *Items) DeleteAt(idx int) {
	if idx > len(*i) {
		return
	}

	*i = append((*i)[:idx], (*i)[idx+1:]...)
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
	Type       itemdata.ItemCode
	TypeID     itemdata.ItemTypeID
	TypeName   string
	BaseDamage *itemdata.WeaponDamage

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
		MagicPrefix []itemdata.MagicalPrefix
		MagicSuffix []itemdata.MagicalSuffix
		// Rare, Crafted only
		RareNames []itemdata.RareName // len() * 8 bits

		SetID         itemdata.SetID // 12 bits
		SetName       string
		SetListID     byte // 5 bits
		SetListCount  byte
		SetAttributes []d2smagicattributes.MagicAttributes

		UniqueID   itemdata.UniqueID // 12 bits
		UniqueName string
	}
	RuneWord struct {
		HasRuneWord bool
		ID          itemdata.RunewordID // 12 bits
		unknown     byte                // 4 bits (brobably always value 5)
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
	SocketedItems []*Item
}

// Load loads an item
func (i *Item) Load(sr *datareader.Reader) (err error) {
	if err := i.loadSimpleFields(sr); err != nil {
		return err
	}

	// if is not a simple item, continue work
	if !i.IsSimple {
		if err := i.loadExtendedFields(sr); err != nil {
			return err
		}
	}

	sr.Align()

	return nil
}

func (i *Item) loadExtendedFields(sr *datareader.Reader) (err error) {
	i.ID = sr.GetUint32() // probably 4 * 8 chars
	i.Level = sr.GetBits(levelLen)
	i.Quality = d2senums.ItemQuality(sr.GetBits(qualityLen))

	i.loadMultiplePicture(sr)

	i.loadClassSpecific(sr)

	if err := i.loadQualityData(sr); err != nil {
		return fmt.Errorf("loading item quality data: %w", err)
	}

	i.loadRuneword(sr)

	i.loadPersonalizationData(sr)

	if i.Type.IsTome() {
		i.unknown11 = sr.GetBits(unknown11Len)
	}

	i.Timestamp = sr.GetBit()

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDShield {
		defRating := sr.GetBits16(defenseRatingLen)
		i.ExtraStats.DefenseRating = defRating - defenseRatingModifier
	}

	i.loadDurability(sr)

	if i.Type.HasQuantity() {
		i.ExtraStats.Quantity = sr.GetBits16(quantityLen)
	}

	if i.Socketed.IsInSocket {
		i.Socketed.TotalNumberOfSockets = sr.GetBits(totalNSocketsLen)
	}

	i.loadSetAttrList(sr)

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

// nolint:gocyclo // will fix later
func (i *Item) loadQualityData(sr *datareader.Reader) error {
	switch i.Quality {
	case d2senums.ItemQualityLow:
		i.QualityData.LowQualityID = sr.GetBits(lowQualityIDLen)
	case d2senums.ItemQualityNormal:
		// noop
	case d2senums.ItemQualityHigh:
		i.QualityData.HighQualityData = sr.GetBits(highQualityDataLen)
	case d2senums.ItemQualityEnchanced:
		i.QualityData.MagicPrefix = make([]itemdata.MagicalPrefix, 1)
		id := sr.GetBits16(magicModifierIDLen) // helper variable (avoid noise with x.y.z.id ;-)

		var found bool

		if i.QualityData.MagicPrefix[0], found = itemdata.GetMagicalPrefix(id); !found {
			return errors.New("loading item prefix: unknown id")
		}

		i.QualityData.MagicSuffix = make([]itemdata.MagicalSuffix, 1)
		id = sr.GetBits16(magicModifierIDLen) // helper variable (avoid noise with x.y.z.id ;-)

		if i.QualityData.MagicSuffix[0], found = itemdata.GetMagicalSuffix(id); !found {
			return errors.New("loading item suffix: unknown id")
		}
	case d2senums.ItemQualitySet:
		id := sr.GetBits16(setIDLen)
		i.QualityData.SetID = itemdata.SetID(id)
	case d2senums.ItemQualityRare, d2senums.ItemQualityCrafted:
		i.QualityData.RareNames = make([]itemdata.RareName, rareNamesCount)

		for n := 0; n < rareNamesCount; n++ {
			id := sr.GetBits16(rareNameIDLen)
			i.QualityData.RareNames[n] = itemdata.RareName(id)
		}

		i.QualityData.MagicPrefix = make([]itemdata.MagicalPrefix, 0)
		i.QualityData.MagicSuffix = make([]itemdata.MagicalSuffix, 0)

		for p := 0; p < 3; p++ {
			prefixExist := sr.GetBit()
			if prefixExist {
				id := sr.GetBits16(magicModifierIDLen)

				prefix, found := itemdata.GetMagicalPrefix(id)
				if !found {
					return errors.New("loading item prefix: unknown id")
				}

				i.QualityData.MagicPrefix = append(i.QualityData.MagicPrefix, prefix)
			}

			suffixExist := sr.GetBit()
			if suffixExist {
				id, found := itemdata.GetMagicalSuffix(sr.GetBits16(magicModifierIDLen))
				if !found {
					return errors.New("loading item suffix: unknown id")
				}

				i.QualityData.MagicSuffix = append(i.QualityData.MagicSuffix, id)
			}
		}
	case d2senums.ItemQualityUnique:
		id := sr.GetBits16(uniqueIDLen)
		i.QualityData.UniqueID = itemdata.UniqueID(id)
	}

	return nil
}

func (i *Item) loadRuneword(sr *datareader.Reader) {
	if i.RuneWord.HasRuneWord {
		id := sr.GetBits16(runeWordIDLen)
		i.RuneWord.ID = itemdata.RunewordID(id)

		i.RuneWord.unknown = sr.GetBits(runeWordUnknownLen)
	}
}

func (i *Item) loadPersonalizationData(sr *datareader.Reader) {
	if i.Personalization.IsPersonalized {
		name := make([]byte, 0)

		for {
			char := sr.GetBits(characterLen)
			if char == 0 {
				break
			}

			name = append(name, char)
		}

		i.Personalization.Name = string(name)
	}
}

func (i *Item) loadDurability(sr *datareader.Reader) {
	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDWeapon ||
		i.TypeID == itemdata.ItemTypeIDShield {
		i.ExtraStats.Durability.Max = sr.GetByte()
		if i.ExtraStats.Durability.Max > 0 {
			i.ExtraStats.Durability.Current = sr.GetByte()
			i.unknown12 = sr.GetBit()
		}
	}
}

func (i *Item) loadSetAttrList(sr *datareader.Reader) {
	var setListValue byte
	if i.Quality == d2senums.ItemQualitySet {
		setListValue = sr.GetBits(setListIDLen)
		listCount := itemdata.GetSetAttributesLen(setListValue)
		i.QualityData.SetListID = setListValue

		i.QualityData.SetListCount = listCount
	}
}

func (i *Item) loadMultiplePicture(sr *datareader.Reader) {
	i.MultiplePicture.HasMultiplePicture = sr.GetBit()
	if i.MultiplePicture.HasMultiplePicture {
		i.MultiplePicture.ID = sr.GetBits(multiplePictureIDLen)
	}
}

func (i *Item) loadClassSpecific(sr *datareader.Reader) {
	i.ClassSpecific.IsClassSpecific = sr.GetBit()
	if i.ClassSpecific.IsClassSpecific {
		// probably class is somewhere here ... ?
		i.ClassSpecific.Data = sr.GetBits16(classSpecificDataLen)
	}
}

func (i *Item) loadSimpleFields(sr *datareader.Reader) (err error) {
	id := sr.GetBytes(numHeaderBytes)
	if string(id) != itemListID {
		return errors.New("unexpected item signature")
	}

	i.unknown1 = sr.GetBits(unknown1Len)
	i.Identified = sr.GetBit()
	i.unknown2 = sr.GetBits(unknown2Len)
	i.Socketed.IsInSocket = sr.GetBit()
	i.unknown3 = sr.GetBit()
	i.JustPicked = sr.GetBit()
	i.unknown4 = sr.GetBits(unknown4Len)
	i.IsEar = sr.GetBit()
	i.NewbieItem = sr.GetBit()
	i.unknown5 = sr.GetBits(unknown5Len)
	i.IsSimple = sr.GetBit()
	i.Etheral = sr.GetBit()
	i.unknown6 = sr.GetBit()
	i.Personalization.IsPersonalized = sr.GetBit()
	i.unknown7 = sr.GetBit()
	i.RuneWord.HasRuneWord = sr.GetBit()
	i.unknown8 = sr.GetBits(unknown8Len)
	i.Version = sr.GetByte()
	i.unknown9 = sr.GetBits(unknown9Len)
	i.Location.LocationID = d2senums.ItemLocationType(sr.GetBits(locationLocationIDLen))
	i.Location.EquippedID = d2senums.ItemEquippedPlace(sr.GetBits(locationEquippedIDLen))
	i.Location.X = sr.GetBits(locationXLen)
	i.Location.Y = sr.GetBits(locationYLen)
	i.unknown10 = sr.GetBit()
	i.Location.StorageID = d2senums.StoragePlace(sr.GetBits(locationStorageIDLen))

	if i.IsEar {
		i.Ear.Class = d2senums.CharacterClass(sr.GetBits(earClassLen))
		i.Ear.Level = sr.GetBits(earLevelLen)

		var name []byte

		for {
			char := sr.GetBits(characterLen)
			if char == 0 {
				break
			}

			name = append(name, char)
		}

		i.Ear.Name = string(name)

		sr.Align()
	} else {
		t := sr.GetBytes(typeLen)
		i.Type = itemdata.ItemCodeFromString(strings.Trim(string(t), " "))
		i.loadTypeInfo()
		i.NumberOfItemsInSockets = sr.GetBits(numItemsInSocketsLen)
	}

	return nil
}

func (i *Item) loadTypeInfo() {
	i.TypeID = i.Type.GetTypeID()
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

		baseDamage := i.Type.WeaponDamage()
		if baseDamage != nil {
			// If the item is ethereal we need to add 50% enhanced
			// damage to the base damage.
			if i.Etheral {
				baseDamage.Etheral()
			}

			i.BaseDamage = baseDamage
		}
	case itemdata.ItemTypeIDOther:
		typeName, ok := itemdata.MiscCodes[i.Type]
		if ok {
			i.TypeName = typeName
		}
	}
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

func (i *Item) encodeExtendedFields(sw *datautils.StreamWriter) (err error) {
	sw.PushBits32(i.ID, uint32Len)
	sw.PushBits(i.Level, levelLen)
	sw.PushBits(byte(i.Quality), qualityLen)
	sw.PushBit(i.MultiplePicture.HasMultiplePicture)

	i.encodeMultiplePicture(sw)

	sw.PushBit(i.ClassSpecific.IsClassSpecific)

	i.encodeClassSpecific(sw)

	i.encodeQuality(sw)

	i.encodeRuneword(sw)

	i.encodePersonalization(sw)

	if i.Type.IsTome() {
		sw.PushBits(i.unknown11, unknown11Len)
	}

	sw.PushBit(i.Timestamp)

	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDShield {
		sw.PushBits16(i.ExtraStats.DefenseRating+defenseRatingModifier, defenseRatingLen)
	}

	i.encodeDurability(sw)

	if i.Type.HasQuantity() {
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

	// OFC length of set attributes is > 0 for sets only
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

func (i *Item) encodeMultiplePicture(sw *datautils.StreamWriter) {
	if i.MultiplePicture.HasMultiplePicture {
		sw.PushBits(i.MultiplePicture.ID, multiplePictureIDLen)
	}
}

func (i *Item) encodeClassSpecific(sw *datautils.StreamWriter) {
	if i.ClassSpecific.IsClassSpecific {
		sw.PushBits16(i.ClassSpecific.Data, classSpecificDataLen)
	}
}

func (i *Item) encodeRuneword(sw *datautils.StreamWriter) {
	if i.RuneWord.HasRuneWord {
		sw.PushBits16(uint16(i.RuneWord.ID), runeWordIDLen)
		sw.PushBits(i.RuneWord.unknown, runeWordUnknownLen)
	}
}

func (i *Item) encodePersonalization(sw *datautils.StreamWriter) {
	if i.Personalization.IsPersonalized {
		name := []byte(i.Personalization.Name)
		for _, c := range name {
			sw.PushBits(c, characterLen)
		}

		sw.PushBits(byte(' '), characterLen)
	}
}

func (i *Item) encodeDurability(sw *datautils.StreamWriter) {
	if i.TypeID == itemdata.ItemTypeIDArmor ||
		i.TypeID == itemdata.ItemTypeIDWeapon ||
		i.TypeID == itemdata.ItemTypeIDShield {
		sw.PushBits(i.ExtraStats.Durability.Max, byteLen)

		if i.ExtraStats.Durability.Max > 0 {
			sw.PushBits(i.ExtraStats.Durability.Current, byteLen)
			sw.PushBit(i.unknown12)
		}
	}
}

func (i *Item) encodeQuality(sw *datautils.StreamWriter) {
	switch i.Quality {
	case d2senums.ItemQualityLow:
		sw.PushBits(i.QualityData.LowQualityID, lowQualityIDLen)
	case d2senums.ItemQualityNormal:
		// noop
	case d2senums.ItemQualityHigh:
		sw.PushBits(i.QualityData.HighQualityData, highQualityDataLen)
	case d2senums.ItemQualityEnchanced:
		sw.PushBits16(i.QualityData.MagicPrefix[0].GetID(), magicModifierIDLen)
		sw.PushBits16(i.QualityData.MagicSuffix[0].GetID(), magicModifierIDLen)
	case d2senums.ItemQualitySet:
		sw.PushBits16(uint16(i.QualityData.SetID), setIDLen)
	case d2senums.ItemQualityRare, d2senums.ItemQualityCrafted:
		for _, name := range i.QualityData.RareNames {
			sw.PushBits(byte(name), byteLen)
		}

		for n := 0; n < 3; n++ {
			l := len(i.QualityData.MagicPrefix)
			hasPrefix := n < l
			sw.PushBit(hasPrefix)

			if hasPrefix {
				sw.PushBits16(i.QualityData.MagicPrefix[n].GetID(), magicModifierIDLen)
			}

			l = len(i.QualityData.MagicSuffix)
			hasSuffix := n < l
			sw.PushBit(hasSuffix)

			if hasSuffix {
				sw.PushBits16(i.QualityData.MagicSuffix[n].GetID(), magicModifierIDLen)
			}
		}
	case d2senums.ItemQualityUnique:
		sw.PushBits16(uint16(i.QualityData.UniqueID), uniqueIDLen)
	}
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
		name := []byte(i.Type.String())
		for _, c := range name {
			sw.PushBits(c, byteLen)
		}
		sw.PushBits(byte(' '), byteLen)
		sw.PushBits(i.NumberOfItemsInSockets, numItemsInSocketsLen)
	}
}

var _ fmt.Stringer = &Item{}

func (i *Item) String() string {
	return i.TypeName
}

// MagicModifier represents a magic modifier
type MagicModifier struct {
	ID   uint16
	Name string
}
