package d2senums

//go:generate stringer -linecomment -type ItemQuality -output item_quality_string.go
//go:generate stringer -linecomment -type LowQualityItemType -output low_quality_string.go

// ItemQuality represents a quality of an item
type ItemQuality byte

// item qualities
const (
	ItemQualityLow       ItemQuality = iota + 1 // low
	ItemQualityNormal                           // normal
	ItemQualityHigh                             // high
	ItemQualityEnchanced                        // magically enchanced
	ItemQualitySet                              // part of set
	ItemQualityRare                             // rare
	ItemQualityUnique                           // unique
	ItemQualityCrafted                          // crafted
)

// LowQualityItemType represents a type of low quality item
type LowQualityItemType byte

// low quality item types
const (
	LowQualityItemCrude      LowQualityItemType = iota // crude
	LowQualityItemCracked                              // cracked
	LowQualityItemDamaged                              // damaged
	LowQualityItemLowQuality                           // low quality
)
