package d2d2s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetName(t *testing.T) {
	tested := "_some!toolongnam_ewith@#!illegalcharacters"
	expected := "sometoolongnam"

	d := NewCharacter()
	d.SetName(tested)

	assert.Equal(t, expected, d.Name, "unexpected name set")
}

func Test_SetLevel(t *testing.T) {
	a := assert.New(t)

	d := NewCharacter()

	level := byte(89)
	d.SetLevel(level)

	a.Equal(level, d.Level, "unexpected level set")
	a.Equal(level, byte(d.Stats.Level), "unexpected level set")

	d.SetLevel(100)

	a.Equal(byte(maxAllowedLevel), d.Level, "unexpected level set")
	a.Equal(byte(maxAllowedLevel), byte(d.Stats.Level), "unexpected level set")
}
