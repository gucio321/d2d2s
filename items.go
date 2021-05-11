package d2d2s

import (
	"errors"

	"github.com/nokka/d2s"

	"github.com/gucio321/d2d2s/datautils"
	"github.com/gucio321/d2d2s/itemdata"
)

const itemListID = "JM"

type Items struct {
	Items []d2s.Item
}

func (i *Items) Load(sr *datautils.BitMuncher) error {
	var err error

	id := sr.GetBytes(2)
	if string(id) != itemListID {
		return errors.New("unexpected item header")
	}

	numItems := sr.GetUInt16()

	items, err := d2s.ParseItemList(sr, int(numItems))
	if err != nil {
		return err
	}

	i.Items = items
	/*
		i.Items = make([]Item, numItems)

		for n := range i.Items {
			id := string(sr.GetBytes(2))
			if id != itemListID {
				return errors.New("Unexpected id")
			}

			sr.SkipBits(4) // unknown
			i.Items[n].Identified = sr.GetBit() == 1
			sr.SkipBits(6) // unknown
			i.Items[n].HasSockets = sr.GetBit() == 1
			sr.SkipBits(4) // bit 2 is set, on items which you have picked up since the last time the game was saved.  Why?...
			isEar := sr.GetBit() == 1
			_ = isEar // compile-fix will use later (TODO)

			i.Items[n].NewbieItem = sr.GetBit() == 1
			sr.SkipBits(3)
			isSimple := sr.GetBit() == 1

			i.Items[n].Etheral = sr.GetBit() == 1
			sr.SkipBits(1) // probably always 1
			isPersonalized := sr.GetBit() == 1
			sr.SkipBits(1)                            // unknown
			i.Items[n].HasRuneWord = sr.GetBit() == 1 // ?
			sr.SkipBits(15)                           // unknown (may be set)

			i.Items[n].Position.Location = ItemLocationType(sr.GetBits(3))
			i.Items[n].Position.EquippedPlace = ItemEquippedPlace(sr.GetBits(4))
			i.Items[n].Position.Column = byte(sr.GetBits(4))
			i.Items[n].Position.Row = byte(sr.GetBits(3))
			sr.SkipBits(1) // unknown
			i.Items[n].Position.StoragePlace = StoragePlace(sr.GetBits(3))
			// e.g. "amu ", "2hs "
			// it would be cool, to use d2enum.WeaponClass here (but this type is incomplete for now)

			// WARNING, if ear, need to use isEar here
			i.Items[n].Type = strings.ReplaceAll(string(sr.GetBytes(4)), " ", "")
			i.Items[n].TypeID = itemdata.GetTypeID(i.Items[n].Type)
			fmt.Println(i.Items[n].TypeID)
			i.Items[n].GemsNumber = byte(sr.GetBits(3))
			sr.SkipBits(1) // skip 1 bit in order to fix offset

			if isSimple {
				fmt.Println("simple")
				continue
			}

			i.Items[n].UID = sr.GetUInt32()
			i.Items[n].Level = byte(sr.GetBits(7))
			i.Items[n].Quality = ItemQuality(sr.GetBits(4))

			i.Items[n].HasMultiplePicture = sr.GetBit() == 1
			if i.Items[n].HasMultiplePicture {
				i.Items[n].MultiplePictureID = byte(sr.GetBits(3))
			}

			i.Items[n].IsClassSpecific = sr.GetBit() == 1
			if i.Items[n].IsClassSpecific {
				for j := 0; j < 11; j++ {
					i.Items[n].ClassSpecificData[j] = sr.GetBit() == 1
				}
			}

			switch i.Items[n].Quality {
			case ItemQualityLow:
				i.Items[n].LowQualityItemType = LowQualityItemType(sr.GetBits(3))
			case ItemQualityNormal:
				// noop
			case ItemQualityHigh:
				// just save data, TODO: figure it out
				for j := 0; j < 3; j++ {
					i.Items[n].HighQualityItemData[j] = sr.GetBit() == 1
				}
			case ItemQualityEnchanced:
				// TODO: make a list of these values
				i.Items[n].MagicPreffix = uint16(sr.GetBits(11))
				i.Items[n].MagicSuffix = uint16(sr.GetBits(11))
			case ItemQualitySet:
				i.Items[n].ItemsSetID = uint16(sr.GetBits(12))
			case ItemQualityRare, ItemQualityCrafted:
				i.Items[n].RareItemData.FirstWordID = sr.GetByte()
				i.Items[n].RareItemData.SecondWordID = sr.GetByte()
				if sr.GetBit() == 1 {
					i.Items[n].RareItemData.FirstMagicPreffix = uint16(sr.GetBits(11))
				}
				if sr.GetBit() == 1 {
					i.Items[n].RareItemData.FirstMagicSuffix = uint16(sr.GetBits(11))
				}
				if sr.GetBit() == 1 {
					i.Items[n].RareItemData.SecondMagicPreffix = uint16(sr.GetBits(11))
				}
				if sr.GetBit() == 1 {
					i.Items[n].RareItemData.SecondMagicSuffix = uint16(sr.GetBits(11))
				}
				if sr.GetBit() == 1 {
					i.Items[n].RareItemData.ThirdMagicPreffix = uint16(sr.GetBits(11))
				}
				if sr.GetBit() == 1 {
					i.Items[n].RareItemData.ThirdMagicSuffix = uint16(sr.GetBits(11))
				}
			case ItemQualityUnique:
				// TODO: make a list of these stuff
				i.Items[n].UniqueItemID = uint16(sr.GetBits(12))
			}

			if i.Items[n].HasRuneWord {
				i.Items[n].RuneWordData.ID = uint16(sr.GetBits(12))
				sr.SkipBits(4)
			}

			if isPersonalized {
				personalizedNameData := make([]byte, 0)
				for {
					if c := byte(sr.GetBits(7)); c != 0 {
						personalizedNameData = append(personalizedNameData, c)
					} else {
						i.Items[n].PersonalizedName = string(personalizedNameData)
						break
					}
				}
			}

			if itemdata.TomeMap[i.Items[n].Type] {
				sr.SkipBits(5)
			}

			sr.SkipBits(1) // unknown

			if id := i.Items[n].TypeID; id == itemdata.ItemTypeIDArmor || id == itemdata.ItemTypeIDShield {
				defRating := uint16(sr.GetBits(11))
				i.Items[n].SpecificItemData.DefenseRating = defRating - 10
			}

			if id := i.Items[n].TypeID; id == itemdata.ItemTypeIDArmor ||
				id == itemdata.ItemTypeIDWeapon || id == itemdata.ItemTypeIDShield {
				i.Items[n].SpecificItemData.MaxDurability = sr.GetByte()
				if i.Items[n].SpecificItemData.MaxDurability > 0 {
					i.Items[n].SpecificItemData.CurrentDurability = sr.GetByte()
					sr.SkipBits(1)
				}
			}

			if itemdata.QuantityMap[i.Items[n].Type] {
				i.Items[n].Quantity = uint16(sr.GetBits(9))
			}

			if i.Items[n].HasSockets {
				i.Items[n].NumberOfSockets = byte(sr.GetBits(4))
			}

			// part of https://github.com/nokka/d2s
			var setListValue byte = 0
			if i.Items[n].Quality == ItemQualitySet {
				setListValue = byte(sr.GetBits(5))

				listCount, ok := itemdata.SetListMap[uint64(setListValue)]
				if !ok {
					return errors.New("unexpected set list number")
				}

				i.Items[n].SetListCount = byte(listCount)
			}

			i.LoadMagicAttributes(sr, n)

			if i.Items[n].SetListCount > 0 {
			}

			if i.Items[n].HasRuneWord {
			}

			if i.Items[n].Position.Location == ItemLocationInSocked {
			} else {
			}

			fmt.Println(i.Items[n].NumberOfSockets)
			sr.SkipBits(8 - sr.Offset())
			sr.SkipBits(8)
		}
	*/

	return nil
}

type Item struct {
	Identified bool
	HasSockets bool
	// This bit is set on the weapon and shield your character is given when you start the game.  Apparently,
	// this gives the item the property of having a repair cost of 1gp, as well as a sell value of 1gp.
	NewbieItem  bool
	Etheral     bool
	HasRuneWord bool // need to check
	Position    struct {
		Location      ItemLocationType
		EquippedPlace ItemEquippedPlace
		Column        byte
		Row           byte
		StoragePlace  StoragePlace
	}
	Type       string
	TypeID     itemdata.ItemTypeID
	GemsNumber byte

	// extended item
	UID     uint32
	Level   byte
	Quality ItemQuality
	// If this is true, it means the item has more than one picture associated
	// with it.
	HasMultiplePicture bool
	MultiplePictureID  byte

	IsClassSpecific bool
	// TODO: see https://user.xmission.com/~trevin/DiabloIIv1.09_Item_Format.shtml#37
	ClassSpecificData [11]bool

	// low quality items onlly
	LowQualityItemType LowQualityItemType

	// high quality items only
	HighQualityItemData [3]bool

	// magically enhanced items only
	MagicPreffix uint16
	MagicSuffix  uint16

	// set items only
	ItemsSetID uint16

	// only for rare and crafted items
	RareItemData struct {
		FirstWordID        byte
		SecondWordID       byte
		FirstMagicPreffix  uint16
		FirstMagicSuffix   uint16
		SecondMagicPreffix uint16
		SecondMagicSuffix  uint16
		ThirdMagicPreffix  uint16
		ThirdMagicSuffix   uint16
	}

	// unique items only
	UniqueItemID uint16

	RuneWordData struct {
		ID uint16
	}

	PersonalizedName string

	SpecificItemData struct {
		DefenseRating     uint16
		MaxDurability     byte
		CurrentDurability byte
	}

	Quantity        uint16
	NumberOfSockets byte

	// only set items
	SetListCount    byte
	MagicAttributes []MagicAttribute
}

type ItemLocationType byte

const (
	ItemLocationStored ItemLocationType = iota
	ItemLocationEquipped
	ItemLocationBelt
	_
	ItemLocationMoved // i.e., has been picked up by the mouse
	_
	ItemLocationInSocked // ?
)

type ItemEquippedPlace byte

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

type StoragePlace byte

const (
	StorageNone StoragePlace = iota
	StorageInventory
	_
	_
	StorageCube
	StorageStash
)

type ItemQuality byte

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

type LowQualityItemType byte

const (
	LowQualityItemCrude LowQualityItemType = iota
	LowQualityItemCracked
	LowQualityItemDamaged
	LowQualityItemLowQuality
)
