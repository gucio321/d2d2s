package d2sitems

import (
	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/d2sitems/itemdata"
)

// Ear sets ear attributes.
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
