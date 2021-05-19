package itemdata

import "log"

// GetSetAttributesLen : each set item has 5 bits of data containing the number of set lists follow
// the magical attributes list, this map tells us how many lists to read
// depending on the value given from the 5 bits. A number of 0-5 set lists.
// nolint:gomnd // data method
func GetSetAttributesLen(id byte) byte {
	lookup := map[byte]byte{
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		4:  1,
		6:  2,
		7:  3,
		10: 2,
		12: 2,
		15: 4,
		31: 5,
	}

	v, ok := lookup[id]
	if !ok {
		log.Panicf("unknown set attributes id: %v", id)
	}

	return v
}

// SetReqID - Certain set items (only Civerb's Ward in unmodded D2) have bonuses
// that require certain other set items in order to be activated
// (instead of the normal requirements of just 'wearing > x of any
// items in the set'); determined by add_func=1 in SetItems.txt
func SetReqID(id byte) []uint64 {
	SetReqIDsMap := map[byte][]uint64{
		// Civerb's Ward: [Civerb's Icon, Civerb's Cudgel]
		0: {1, 2},
	}

	r, ok := SetReqIDsMap[id]
	if !ok {
		return nil
	}

	return r
}
