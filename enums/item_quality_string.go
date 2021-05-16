// Code generated by "stringer -linecomment -type ItemQuality -output item_quality_string.go"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ItemQualityLow-1]
	_ = x[ItemQualityNormal-2]
	_ = x[ItemQualityHigh-3]
	_ = x[ItemQualityEnchanced-4]
	_ = x[ItemQualitySet-5]
	_ = x[ItemQualityRare-6]
	_ = x[ItemQualityUnique-7]
	_ = x[ItemQualityCrafted-8]
}

const _ItemQuality_name = "lownormalhighmagically enchancedpart of setrareuniquecrafted"

var _ItemQuality_index = [...]uint8{0, 3, 9, 13, 32, 43, 47, 53, 60}

func (i ItemQuality) String() string {
	i -= 1
	if i >= ItemQuality(len(_ItemQuality_index)-1) {
		return "ItemQuality(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ItemQuality_name[_ItemQuality_index[i]:_ItemQuality_index[i+1]]
}
