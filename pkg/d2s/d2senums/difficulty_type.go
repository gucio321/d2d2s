package d2senums

//go:generate stringer -linecomment -type DifficultyType -trimprefix Difficulty

// DifficultyType represents a difficulty level
type DifficultyType byte

// difficulty types
const (
	DifficultyNormal DifficultyType = iota
	DifficultyNightmare
	DifficultyHell
)
