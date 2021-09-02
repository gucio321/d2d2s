// Package d2swidget contains a simple d2s editor written with giu.
package d2swidget

import (
	"fmt"
	"math"
	"strconv"

	"github.com/AllenDang/giu"

	"github.com/gucio321/d2d2s/pkg/d2s"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sstats"
)

var _ giu.Widget = &D2SWidget{}

// D2SWidget represents d2s widget
type D2SWidget struct {
	d2s *d2s.D2S
	id  string
}

// D2S creates d2s widget
func D2S(d2sData *d2s.D2S) *D2SWidget {
	result := &D2SWidget{
		d2s: d2sData,
		id:  giu.GenAutoID("D2SWidget"),
	}

	return result
}

// Build implements giu.Widget
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
		giu.TreeNode("Skills").Layout(w.skills()),
		giu.TreeNode("Difficulty").Layout(w.difficulties()),
		giu.TreeNode("Map").Layout(
			giu.Label(fmt.Sprintf("ID: %v", w.d2s.MapID)),
		),
		giu.TreeNode("Mercenary").Layout(w.mercenary()),
		giu.TreeNode("Quests").Layout(w.quests()),
		giu.TreeNode("Waypoints").Layout(w.waypoints()),
		giu.TreeNode("NPC").Layout(w.npc()),
		giu.TreeNode("Stats").Layout(w.stats()),
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
	return giu.Layout{giu.Label("TODO")}
}

func (w *D2SWidget) skills() giu.Layout {
	return giu.Layout{
		giu.Row(
			giu.Label("Left Skill: "),
			skillCombo(&w.d2s.LeftSkill),
		),
		giu.Row(
			giu.Label("Right Skill: "),
			skillCombo(&w.d2s.RightSkill),
		),
		giu.Row(
			giu.Label("Left Skill (switch): "),
			skillCombo(&w.d2s.LeftSkillSwitch),
		),
		giu.Row(
			giu.Label("Right Skill (switch): "),
			skillCombo(&w.d2s.RightSkillSwitch),
		),
	}
}

func (w *D2SWidget) difficulties() giu.Layout {
	state := w.getState()
	diffStatus := (*w.d2s.Difficulty)[state.difficultyDifficulty]
	act := int32(diffStatus.Act)

	return giu.Layout{
		difficultySlider(&state.difficultyDifficulty),
		giu.Checkbox("Active: ", &diffStatus.Active),
		giu.Row(
			giu.Label("Act: "),
			giu.InputInt(&act).OnChange(func() {
				diffStatus.SetAct(int(act))
			}),
		),
	}
}

func (w *D2SWidget) mercenary() giu.Layout {
	died := int32(w.d2s.Mercenary.Died)

	return giu.Layout{
		giu.Row(
			giu.Label("Died: "),
			giu.InputInt(&died).OnChange(func() {
				if died < 0 || died > math.MaxUint16 {
					return
				}

				w.d2s.Mercenary.Died = uint16(died)
			}),
		),
		giu.Label(fmt.Sprintf("ID: %v", w.d2s.Mercenary.ID)),                 // TODO: add editor for this
		giu.Label(fmt.Sprintf("Name ID: %v", w.d2s.Mercenary.Name)),          // TODO: add editor for this
		giu.Label(fmt.Sprintf("Type: %s", w.d2s.Mercenary.Type)),             // TODO: add editor for this
		giu.Label(fmt.Sprintf("Experience: %v", w.d2s.Mercenary.Experience)), // TODO: add editor for this
		giu.TreeNode("Items").Layout(
			items(w.d2s.Mercenary.Items),
		),
	}
}

func (w *D2SWidget) quests() giu.Layout {
	return giu.Layout{giu.Label("TODO")}
}

func (w *D2SWidget) waypoints() giu.Layout {
	state := w.getState()

	return giu.Layout{
		difficultySlider(&state.waypointDifficulty),
		giu.SliderInt(giu.GenAutoID("##waypointsAct"), &state.waypointAct, 1, d2senums.NumActs).Format("Act: %d"),
		giu.Custom(func() {
			for i := range (*w.d2s.Waypoints)[state.waypointDifficulty][state.waypointAct] {
				giu.Checkbox(fmt.Sprintf("waypoint %d", i), &(*w.d2s.Waypoints)[state.waypointDifficulty][state.waypointAct][i]).Build()
			}
		}),
	}
}

func (w *D2SWidget) npc() giu.Layout {
	return giu.Layout{giu.Label("TODO")}
}

func (w *D2SWidget) stats() giu.Layout {
	const decimal = 10

	strength := strconv.Itoa(int(w.d2s.Stats.Strength))
	strengthLen, _ := d2sstats.Strength.GetStatLen()
	energy := strconv.Itoa(int(w.d2s.Stats.Energy))
	energyLen, _ := d2sstats.Energy.GetStatLen()
	dexterity := strconv.Itoa(int(w.d2s.Stats.Dexterity))
	dexterityLen, _ := d2sstats.Dexterity.GetStatLen()

	return giu.Layout{
		giu.Row(
			giu.Label("Strength: "),
			giu.InputText(&strength).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				s, _ := strconv.ParseUint(strength, decimal, strengthLen)
				w.d2s.Stats.Strength = uint32(s)
			}),
		),
		giu.Row(
			giu.Label("Energy: "),
			giu.InputText(&energy).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				e, _ := strconv.ParseUint(energy, decimal, energyLen)
				w.d2s.Stats.Energy = uint32(e)
			}),
		),
		giu.Row(
			giu.Label("Dexterity: "),
			giu.InputText(&dexterity).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				d, _ := strconv.ParseUint(dexterity, decimal, dexterityLen)
				w.d2s.Stats.Dexterity = uint32(d)
			}),
		),
	}
}

// nolint:unparam // TODO: write items method
func items(items *d2sitems.Items) giu.Layout {
	return giu.Layout{giu.Label("TODO")}
}

func difficultyCombo(value *d2senums.DifficultyType) giu.Widget {
	list := make([]string, 0)
	for d := d2senums.DifficultyNormal; d <= d2senums.DifficultyHell; d++ {
		list = append(list, d.String())
	}

	v := int32(*value)

	return giu.Combo(giu.GenAutoID("##difficultyCombo"), list[v], list, &v).OnChange(func() {
		*value = d2senums.DifficultyType(v)
	})
}

func charClassCombo(value *d2senums.CharacterClass) giu.Widget {
	list := make([]string, 0)
	for d := d2senums.CharacterAmazon; d <= d2senums.CharacterAssassin; d++ {
		list = append(list, d.String())
	}

	v := int32(*value)

	return giu.Combo(giu.GenAutoID("##charClassCombo"), list[v], list, &v).OnChange(func() {
		*value = d2senums.CharacterClass(v)
	})
}

func skillCombo(value *d2senums.SkillID) giu.Widget {
	list := make([]string, 0)
	for d := d2senums.SkillAttack; d <= d2senums.SkillBarbarianBattleCommands; d++ {
		list = append(list, d.String())
	}

	for d := d2senums.SkillScrollIdentify; d <= d2senums.SkillAssasinRoyalStrike; d++ {
		list = append(list, d.String())
	}

	v := int32(*value)

	return giu.Combo(giu.GenAutoID("##skillCombo"), list[v], list, &v).OnChange(func() {
		*value = d2senums.SkillID(v)
	})
}

func difficultySlider(value *d2senums.DifficultyType) giu.Widget {
	v := int32(*value)

	return giu.SliderInt("##difficultySlider", &v, int32(d2senums.DifficultyNormal), int32(d2senums.DifficultyHell)).
		Format(value.String()).OnChange(func() {
		*value = d2senums.DifficultyType(v)
	})
}
