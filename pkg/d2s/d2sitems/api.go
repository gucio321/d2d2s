package d2sitems

/*
 * this file contains a methods, which can be used to create a new
 * items in a character save file
 *
 * because of each method returns *Item type, they could
 * be used in a chain of methods or separately
 * for example: i := NewItem(...).SetLocation(...).SetIdentified(...).SetQualityNormal()
 *          or: i := NewItem(...)
 *              i.SetLocation(...)
 *              i.SetIdentified(...)
 *              i.SetQualityNormal()
 */

import (
	"log"

	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems/itemdata"
)

// NewEar creates a new ear
// NOTE: any other method than SetLocation() shouldn't be called after NewEar(...)
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

// SetIdentified sets if item is identified
func (i *Item) SetIdentified(idt bool) *Item {
	i.Identified = idt

	return i
}

// SetJustPicked sets if item was picket in the last game
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

// SetMultiplePicture sets multiple picture and multiple picture ID
func (i *Item) SetMultiplePicture(pictureID byte) *Item {
	i.ensureExtended()
	i.MultiplePicture.HasMultiplePicture = true
	i.MultiplePicture.ID = pictureID

	return i
}

// NOTE: one of the following method must be used, however using more than one of them
// will cause overwritting data and only the latest used method will be processed in fact

// SetQualityLow sets quality to low and sets LowQualityID
func (i *Item) SetQualityLow(id byte) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityLow
	i.QualityData.LowQualityID = id

	return i
}

// SetQualityNormal sets normal quality
func (i *Item) SetQualityNormal() *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityNormal

	return i
}

// SetQualityHigh sets high quality and its data
func (i *Item) SetQualityHigh(data byte) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityHigh
	i.QualityData.HighQualityData = data

	return i
}

// SetMagicalyEnchanced sets quality to magicaly enchanced and sets prefix/suffix
func (i *Item) SetMagicalyEnchanced(p itemdata.MagicalPrefix, s itemdata.MagicalSuffix) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityEnchanced
	i.QualityData.MagicPrefix = []itemdata.MagicalPrefix{p}
	i.QualityData.MagicSuffix = []itemdata.MagicalSuffix{s}

	return i
}

// setRareCrafted sets magic modifiers and rare names
func (i *Item) setRareCrafted(r []itemdata.RareName, p []itemdata.MagicalPrefix, s []itemdata.MagicalSuffix) {
	i.QualityData.MagicPrefix = p
	i.QualityData.MagicSuffix = s
	i.QualityData.RareNames = r
}

// SetQualityRare sets item's quality to rare
// and sets rare names and modifiers
func (i *Item) SetQualityRare(r []itemdata.RareName, p []itemdata.MagicalPrefix, s []itemdata.MagicalSuffix) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityRare
	i.setRareCrafted(r, p, s)

	return i
}

// SetQualityCrafted sets item's quality to crafted
// and sets rare names and modifiers
func (i *Item) SetQualityCrafted(r []itemdata.RareName, p []itemdata.MagicalPrefix, s []itemdata.MagicalSuffix) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityCrafted
	i.setRareCrafted(r, p, s)

	return i
}

// SetQualityPartOfSet sets items quality as a part of set and sets its ID
func (i *Item) SetQualityPartOfSet(id itemdata.SetID, listID byte) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualitySet
	i.QualityData.SetID = id
	i.QualityData.SetListID = listID

	return i
}

// SetQualityUnique sets quality to unique and unique ID
func (i *Item) SetQualityUnique(id itemdata.UniqueID) *Item {
	i.ensureExtended()
	i.Quality = d2senums.ItemQualityUnique
	i.QualityData.UniqueID = id

	return i
}

// runeword methods

// SetRuneword sets a rneword
func (i *Item) SetRuneword(id itemdata.RunewordID) *Item {
	i.ensureExtended()
	i.RuneWord.HasRuneWord = true
	i.RuneWord.ID = id

	// TODO: need to load attributes + name here

	return i
}

// SetPersonalization sets item's personalization name
func (i *Item) SetPersonalization(name string) *Item {
	i.Personalization.IsPersonalized = true
	i.Personalization.Name = name

	return i
}

// AddItemToSocket adds items to the sockets
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
