package d2senums

//go:generate stringer -linecomment -type Version

// Version represents a file version
type Version int32

// known versions
const (
	VersionOldD2     Version = 71 // Diablo II v1.00 - v1.06
	VersionLOD08     Version = 87 // Diablo II v1.07 or Diablo II: Lord of Destruction v1.08
	VersionD208      Version = 89 // Diablo II v1.08
	Version09        Version = 92 // Diablo II or Diablo II: Lord of Destruction v.09
	VersionLODLatest Version = 96 // Diablo II Expansion v.10 or later
)
