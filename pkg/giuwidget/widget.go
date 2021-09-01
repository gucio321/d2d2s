// Package giuwidget contains a simple d2s editor written with giu.
package giuwidget

import (
	"fmt"

	"github.com/AllenDang/giu"

	"github.com/gucio321/d2d2s/pkg/d2s"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

var _ giu.Widget = &D2SWidget{}

type D2SWidget struct {
	d2s *d2s.D2S
	id  string
}

func D2S(d2s *d2s.D2S) *D2SWidget {
	result := &D2SWidget{
		d2s: d2s,
		id:  giu.GenAutoID("D2SWidget"),
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
		giu.TreeNode("Status").Layout(w.status()),
		giu.TreeNode("Progression").Layout(w.progression()),
	}.Build()
}

func (w *D2SWidget) status() giu.Layout {
	return giu.Layout{
		giu.Checkbox("Hardcore", &w.d2s.Status.Hardcore),
		giu.Checkbox("has ever died", &w.d2s.Status.Died),
		giu.Tooltip("it is checked, if you hav died in some point in past"),
		giu.Checkbox("Character from Expansion set", &w.d2s.Status.Expansion),
		giu.Checkbox("Ladder", &w.d2s.Status.Ladder),
	}
}

func (w *D2SWidget) progression() giu.Layout {
	act := int32(w.d2s.Progression.Act)
	return giu.Layout{
		giu.Row(
			giu.Label("Difficulty: "),
			difficultyCombo(&w.d2s.Progression.DifficultyLevel),
		),
		giu.Row(
			giu.Label("Act"),
			giu.InputInt(&act).OnChange(func() {
				if act > 0 && act <= d2senums.NumActs {
					w.d2s.Progression.Act = int(act)
				}
			}),
		),
	}
}

func difficultyCombo(value *d2senums.DifficultyType) giu.Widget {
	list := make([]string, 0)
	for d := d2senums.DifficultyNormal; d <= d2senums.DifficultyHell; d++ {
		list = append(list, d.String())
	}

	v := int32(*value)

	return giu.Combo(giu.GenAutoID("difficultyCombo"), list[v], list, &v).OnChange(func() {
		*value = d2senums.DifficultyType(v)
	})
}
