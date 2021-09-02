// Package d2swidget contains a simple d2s editor written with giu.
package d2swidget

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
	lvl := int32(w.d2s.Level)
	name := w.d2s.Name

	giu.Layout{
		giu.Label(fmt.Sprintf("Version: %s (%d)", w.d2s.Version, w.d2s.Version)),
		giu.Row(
			giu.Label("Name: "),
			giu.InputText(&name).OnChange(func() {
				w.d2s.SetName(name)
			}),
		),
		giu.Row(
			giu.Label("Class: "),
			charClassCombo(&w.d2s.Class),
		),
		giu.Row(
			giu.Label("Level: "),
			giu.InputInt(&lvl).OnChange(func() { w.d2s.SetLevel(byte(lvl)) }),
		),
		giu.TreeNode("Status").Layout(w.status()),
		giu.TreeNode("Progression").Layout(w.progression()),
		giu.TreeNode("Hotkeys").Layout(w.hotkeys()),
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

func (w *D2SWidget) hotkeys() giu.Layout {
	return giu.Layout{}
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

func charClassCombo(value *d2senums.CharacterClass) giu.Widget {
	list := make([]string, 0)
	for d := d2senums.CharacterAmazon; d <= d2senums.CharacterAssassin; d++ {
		list = append(list, d.String())
	}

	v := int32(*value)

	return giu.Combo(giu.GenAutoID("charClassCombo"), list[v], list, &v).OnChange(func() {
		*value = d2senums.CharacterClass(v)
	})
}
