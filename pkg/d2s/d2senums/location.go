package d2senums

//go:generate stringer -linecomment -type ItemLocationType -output location_type_string.go
//go:generate stringer -linecomment -type StoragePlace -output storage_place_string.go
//go:generate stringer -linecomment -type ItemEquippedPlace -output equipped_place_string.go

// ItemLocationType represents an item location
type ItemLocationType byte

// item locations
const (
	LocationStored   ItemLocationType = iota // Stored
	LocationEquipped                         // Equipped
	LocationBelt                             // in belt
	_
	// i.e., has been picked up by the mouse
	LocationMoved // moving
	_
	LocationInSocket // in socket
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
	EquippedAnywhere    ItemEquippedPlace = iota //
	EquippedHead                                 // head
	EquippedNeck                                 // neck
	EquippedTorse                                // torse
	EquippedRightHand                            // right hand
	EquippedLeftHand                             // left hand
	EquippedRightFinger                          // right finger
	EquippedLeftFinger                           // left winger
	EquippedWaist                                // waist
	EquippedFeet                                 // feet
	EquippedHands                                // hands
	EquippedARightHand                           // alternative right hand
	EquippedALeftHand                            // alternative left hand
)
