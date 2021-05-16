package enums

//go:generate stringer -linecomment -type ItemLocationType -output location_type_string.go
//go:generate stringer -linecomment -type StoragePlace -output storage_place_string.go
//go:generate stringer -linecomment -type ItemEquippedPlace -output equipped_place_string.go

// ItemLocationType represents an item location
type ItemLocationType byte

// item locations
const (
	ItemLocationStored   ItemLocationType = iota // Stored
	ItemLocationEquipped                         // Equipped
	ItemLocationBelt                             // in belt
	_
	// i.e., has been picked up by the mouse
	ItemLocationMoved // moving
	_
	ItemLocationInSocket // in socket
)

// StoragePlace represents a plece, where the item is stored
type StoragePlace byte

// storage places
const (
	StorageNone      StoragePlace = iota //
	StorageInventory                     // inventory
	_
	_
	StorageCube  // Horadric Cube
	StorageStash // stash
)

// ItemEquippedPlace represents a place, where the item is equipped
type ItemEquippedPlace byte

// item equipped places
const (
	ItemEquippedAnywhere    ItemEquippedPlace = iota //
	ItemEquippedHead                                 // head
	ItemEquippedNeck                                 // neck
	ItemEquippedTorse                                // torse
	ItemEquippedRightHand                            // right hand
	ItemEquippedLeftHand                             // left hand
	ItemEquippedRightFinger                          // right finger
	ItemEquippedLeftFinger                           // left winger
	ItemEquippedWaist                                // waist
	ItemEquippedFeet                                 // feet
	ItemEquippedHands                                // hands
	ItemEquippedARightHand                           // alternative right hand
	ItemEquippedALeftHand                            // alternative left hand
)
