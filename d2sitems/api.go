package d2sitems

import (
	"log"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/d2sitems/itemdata"
)

// Ear sets ear attributes.
// BUG - doesn't want to work with game saves
func NewEar(
	class d2senums.CharacterClass,
	level byte,
	name string,
) *Item {
	result := &Item{}
	result.IsSimple = true
	result.IsEar = true
	result.Ear.Class = class
	result.Ear.Level = level
	result.Ear.Name = name

	return result
}

// NewItem creates a new item
// NOTE: calling SetLocation is expected!
func NewItem(
	itemType itemdata.ItemCode,
) *Item {
	result := &Item{
		Type:     itemType,
		IsSimple: true,
		IsEar:    false,
	}

	result.loadTypeInfo()

	return result
}

func (i *Item) SetIdentified(idt bool) *Item {
	i.Identified = idt

	return i
}

func (i *Item) SetJustPicked(jp bool) *Item {
	i.JustPicked = jp
	return i
}

// SetLocation sets location of the item. Should be called for each new item
func (i *Item) SetLocation(
	locationID d2senums.ItemLocationType,
	equippedID d2senums.ItemEquippedPlace,
	x byte,
	y byte,
	storageID d2senums.StoragePlace,
) *Item {
	i.Location.LocationID = locationID
	i.Location.EquippedID = equippedID
	i.Location.X = x
	i.Location.Y = y
	i.Location.StorageID = storageID

	return i
}

// PART 2: Extended fields

// NOTE: one of the following method must be used, however using more than one of them
// will cause overwritting data and only the latest used method will be processed in fact

func (i *Item) SetMultiplePicture(pictureID byte) *Item {
	i.ensureExtended()
	i.MultiplePicture.HasMultiplePicture = true
	i.MultiplePicture.ID = pictureID

	return i
}

func (i *Item) SetLowQuality(id byte) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityLow
	i.QualityData.LowQualityID = id

	return i
}

func (i *Item) SetNormalQuality() *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityNormal

	return i
}

func (i *Item) SetHighQuality(data byte) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityHigh
	i.QualityData.HighQualityData = data

	return i
}

func (i *Item) SetMagicalyEnchanced(p itemdata.MagicalPrefix, s itemdata.MagicalSuffix) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityEnchanced
	i.QualityData.MagicPrefix = []itemdata.MagicalPrefix{p}
	i.QualityData.MagicSuffix = []itemdata.MagicalSuffix{s}

	return i
}

func (i *Item) setRareCrafted(r []itemdata.RareName, p []itemdata.MagicalPrefix, s []itemdata.MagicalSuffix) {
	i.QualityData.MagicPrefix = p
	i.QualityData.MagicSuffix = s
	i.QualityData.RareNames = r
}

func (i *Item) SetQualityRare(r []itemdata.RareName, p []itemdata.MagicalPrefix, s []itemdata.MagicalSuffix) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityRare
	i.setRareCrafted(r, p, s)

	return i
}

func (i *Item) SetQualityCrafted(r []itemdata.RareName, p []itemdata.MagicalPrefix, s []itemdata.MagicalSuffix) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityCrafted
	i.setRareCrafted(r, p, s)

	return i
}

func (i *Item) SetQualityPartOfSet(id itemdata.SetID, listID byte) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualitySet
	i.QualityData.SetID = id
	i.QualityData.SetListID = listID

	return i
}

func (i *Item) SetQualityUnique(id itemdata.UniqueID) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityUnique
	i.QualityData.UniqueID = id

	return i
}

// runeword methods

func (i *Item) SetRuneword(id itemdata.RunewordID) *Item {
	i.ensureExtended()
	i.RuneWord.HasRuneWord = true
	i.RuneWord.ID = id

	// TODO: need to load attributes + name here

	return i
}

func (i *Item) SetPersonalization(name string) *Item {
	i.Personalization.IsPersonalized = true
	i.Personalization.Name = name

	return i
}

func (i *Item) AddItemToSocket(items ...*Item) *Item {
	for _, item := range items {
		if item.Location.LocationID != d2senums.LocationInSocket {
			log.Panicf("Item: AddItemToSocket: input item must be located in socket (LocationID = LocationInSocket)")
		}

		i.NumberOfItemsInSockets++
		i.SocketedItems = append(i.SocketedItems, item)
	}

	return i
}

func (i *Item) ensureExtended() {
	i.IsSimple = true
}
