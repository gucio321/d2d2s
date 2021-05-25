package d2d2s

import (
	"log"
	"strings"

	"github.com/gucio321/d2d2s/d2senums"
)

const (
	maxCharNameLen = 15
	NameFilter     = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz_-"

	maxAllowedLevel = 99
)

// NewCharacter creates a new character (for now, wrapps New)
func NewCharacter() *D2S {
	result := New()

	return result
}

// SetName sets character name.
// NOTE: maximal allowed name length is 15, the name can contain only lower and upper cases (a-Z)
// and 1 "-" or 1 "_"  but it cannot be a first or the last character. the name should
func (d *D2S) SetName(name string) *D2S {
	switch x := name[0]; x {
	case '_', '-':
		log.Printf("D2S: SetName: Exception name[0] == %v: disallowed", x)

		name = name[1:]
	}

	if len(name) > maxCharNameLen {
		log.Printf("D2S: SetName: Exception len(%s) > %d: name is too long, cannot set it!\n%s", name, maxCharNameLen,
			"the length of name was reduced to characters")

		name = name[:maxCharNameLen+1]
	}

	switch x := name[len(name)-1]; x {
	case '_', '-':
		log.Printf("D2S: SetName: Exception name[len(name)-1] == %v: disallowed", x)
		name = name[:]
	}

	if x := strings.Count(name, "-") + strings.Count(name, "_"); x > 1 {
		log.Printf("D2s: SetName: Exception - name contains more than 1 \"-\" and \"_\" - disallowed")
	}

	n := name
	for _, char := range NameFilter {
		n = strings.ReplaceAll(n, string(char), "")
	}

	if n != "" {
		for _, illegalChar := range n {
			log.Printf("D2S: SetName: Exception - name contains illegal character \"%c\". Will be removed.", illegalChar)
			name = strings.ReplaceAll(name, string(illegalChar), "")
		}
	}

	d.Name = name

	return d
}

// SetLevel sets character's level
func (d *D2S) SetLevel(level byte) *D2S {
	if level > maxAllowedLevel {
		log.Printf("D2S: SetLevel: exception - level is too large (%d); it must be in range 0-%d", level, maxAllowedLevel)
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
