// Code generated by "stringer -linecomment -type Version"; DO NOT EDIT.

package d2senums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VersionOldD2-71]
	_ = x[VersionLOD08-87]
	_ = x[VersionD208-89]
	_ = x[Version09-92]
	_ = x[VersionLODLatest-96]
}

const (
	_Version_name_0 = "Diablo II v1.00 - v1.06"
	_Version_name_1 = "Diablo II v1.07 or Diablo II: Lord of Destruction v1.08"
	_Version_name_2 = "Diablo II v1.08"
	_Version_name_3 = "Diablo II or Diablo II: Lord of Destruction v.09"
	_Version_name_4 = "Diablo II Expansion v.10 or later"
)

func (i Version) String() string {
	switch {
	case i == 71:
		return _Version_name_0
	case i == 87:
		return _Version_name_1
	case i == 89:
		return _Version_name_2
	case i == 92:
		return _Version_name_3
	case i == 96:
		return _Version_name_4
	default:
		return "Version(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}