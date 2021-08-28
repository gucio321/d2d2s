// Package testdata contains embeded by go:embed testing data
package testdata

import "embed"

//go:embed *
var Testdata embed.FS
