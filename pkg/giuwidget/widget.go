// Package giuwidget contains a simple d2s editor written with giu.
package giuwidget

import (
	"fmt"

	"github.com/AllenDang/giu"

	"github.com/gucio321/d2d2s/pkg/d2s"
)

var _ giu.Widget = &D2SWidget{}

type D2SWidget struct {
	d2s *d2s.D2S
}

func D2S(d2s *d2s.D2S) *D2SWidget {
	result := &D2SWidget{
		d2s: d2s,
	}

	return result
}

func (w *D2SWidget) Build() {
	giu.Layout{
		giu.Label(fmt.Sprintf("Version: %s (%d)", w.d2s.Version, w.d2s.Version)),
		giu.Row(
			giu.Label("Name: "),
			giu.InputText(&w.d2s.Name),
		),
	}.Build()
}
