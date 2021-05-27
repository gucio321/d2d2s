package d2d2s

import (
	"log"
	"strings"

	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/d2sitems"
)

const (
	maxCharNameLen = 15
	nameFilter     = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz_-"

	maxAllowedLevel = 99
)

// NewCharacter creates a new character (for now, wrapps New)
func NewCharacter(
	name string,
	class d2senums.CharacterClass,
) *D2S {
	result := New().SetName(name).SetClass(class)

	return result
}

// SetName sets character name.
// NOTE: maximal allowed name length is 15, the name can contain only lower and upper cases (a-Z)
// and 1 "-" or 1 "_"  but it cannot be a first or the last character. the name should
func (d *D2S) SetName(name string) *D2S {
	// first remove illegal characters (like numbers or special chars)
	// copy name, remove all allowed characters from the copy
	n := name
	for _, char := range nameFilter {
		n = strings.ReplaceAll(n, string(char), "")
	}

	// check if there are some illegal chars
	if n != "" {
		// if so, print warning and replace them
		for _, illegalChar := range n {
			log.Printf("D2S: SetName: string name contains illegal character \"%c\".", illegalChar)
			name = strings.ReplaceAll(name, string(illegalChar), "")
		}
	}

	// remove `_`s and `-`s from the first index
cullingCharsAtTheStartOfName:
	for {
		switch x := name[0]; x {
		case '_', '-':
			log.Printf("D2S: SetName: Exception name[0] == %v: disallowed", x)

			name = name[1:]
		default:
			break cullingCharsAtTheStartOfName
		}
	}

	// if the name has more than one `_` and `-`, it should be cleanded
	if x := strings.Count(name, "-") + strings.Count(name, "_"); x > 1 {
		log.Printf("D2s: SetName: Exception - name contains more than 1 \"-\" and \"_\" - disallowed")

		n1 := strings.SplitAfter(name, "-")

		n := make([]string, 0)
		for i := range n1 {
			n = append(n, strings.SplitAfter(n1[i], "_")...)
		}

		for i := 1; i < len(n); i++ {
			s := strings.ReplaceAll(n[i], "-", "")
			s = strings.ReplaceAll(s, "_", "")
			n[i] = s
		}

		name = strings.Join(n, "")
	}

	// cut name
	if len(name) > maxCharNameLen {
		log.Printf("D2S: SetName: Assertion: len(%s) > %d: name is too long, cannot set it!\n%s", name, maxCharNameLen,
			"the length of name was reduced to characters")

		name = name[:maxCharNameLen]
	}

	// replace al `_`s and `-`s from the last position
cullingCharsAtTheEndOfName:
	for {
		switch x := name[len(name)-1]; x {
		case '_', '-':
			log.Printf("D2S: SetName: Exception name[len(name)-1] == %v: disallowed", x)
			name = name[:len(name)-1]
		default:
			break cullingCharsAtTheEndOfName
		}
	}

	d.Name = name

	return d
}

// SetLevel sets character's level
func (d *D2S) SetLevel(level byte) *D2S {
	if level > maxAllowedLevel {
		log.Printf("D2S: SetLevel: assertion: %d > %d:  level is too large; it must be in range 0-%d", level, maxAllowedLevel, maxAllowedLevel)
		level = maxAllowedLevel
	}

	d.Level = level
	d.Stats.Level = uint32(level)

	return d
}

// SetIsHardcore sets if character is hardcore character
func (d *D2S) SetIsHardcore(b bool) *D2S {
	d.Status.Hardcore = b
	return d
}

// SetClass sets character class (default is 0 - amazon)
func (d *D2S) SetClass(c d2senums.CharacterClass) *D2S {
	d.Class = c
	return d
}

// PushItems adds items to item list
func (d *D2S) PushItems(items ...*d2sitems.Item) *D2S {
	d.Items.Add(items...)
	return d
}
