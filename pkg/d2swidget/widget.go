// Package d2swidget contains a simple d2s editor written with giu.
package d2swidget

import (
	"fmt"
	"math"
	"strconv"

	"github.com/AllenDang/giu"
	"golang.org/x/image/colornames"

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

// nolint:funlen // will fix later
func (w *D2SWidget) stats() giu.Layout {
	const (
		decimal         = 10
		floatSize       = 32
		progressBarSize = 50
		inputIntW       = 60
	)

	strength := strconv.Itoa(int(w.d2s.Stats.Strength))
	strengthLen, _ := d2sstats.Strength.GetStatLen(nil)
	energy := strconv.Itoa(int(w.d2s.Stats.Energy))
	energyLen, _ := d2sstats.Energy.GetStatLen(nil)
	dexterity := strconv.Itoa(int(w.d2s.Stats.Dexterity))
	dexterityLen, _ := d2sstats.Dexterity.GetStatLen(nil)
	vitality := strconv.Itoa(int(w.d2s.Stats.Vitality))
	vitalityLen, _ := d2sstats.Vitality.GetStatLen(nil)
	ustats := strconv.Itoa(int(w.d2s.Stats.UnusedStats))
	ustatsLen, _ := d2sstats.UnusedStats.GetStatLen(nil)
	uskills := strconv.Itoa(int(w.d2s.Stats.UnusedSkillPoints))
	uskillsLen, _ := d2sstats.UnusedSkills.GetStatLen(nil)

	currentHP := strconv.Itoa(int(w.d2s.Stats.CurrentHP))
	maxHP := strconv.Itoa(int(w.d2s.Stats.MaxHP))
	currentMana := strconv.Itoa(int(w.d2s.Stats.CurrentMana))
	maxMana := strconv.Itoa(int(w.d2s.Stats.MaxMana))
	currentStamina := strconv.Itoa(int(w.d2s.Stats.CurrentStamina))
	maxStamina := strconv.Itoa(int(w.d2s.Stats.MaxStamina))

	exp := strconv.Itoa(int(w.d2s.Stats.Experience))
	expLen, _ := d2sstats.Experience.GetStatLen(nil)
	gold := strconv.Itoa(int(w.d2s.Stats.Gold))
	goldLen, _ := d2sstats.Gold.GetStatLen(nil)
	stashGold := strconv.Itoa(int(w.d2s.Stats.StashedGold))
	stashGoldLen, _ := d2sstats.StashedGold.GetStatLen(nil)

	return giu.Layout{
		giu.Row(
			giu.Label("Strength: "),
			giu.InputText(&strength).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				s, _ := strconv.ParseUint(strength, decimal, strengthLen)
				w.d2s.Stats.Strength = s
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Energy: "),
			giu.InputText(&energy).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				e, _ := strconv.ParseUint(energy, decimal, energyLen)
				w.d2s.Stats.Energy = e
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Dexterity: "),
			giu.InputText(&dexterity).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				d, _ := strconv.ParseUint(dexterity, decimal, dexterityLen)
				w.d2s.Stats.Dexterity = d
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Vitality: "),
			giu.InputText(&vitality).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				v, _ := strconv.ParseUint(vitality, decimal, vitalityLen)
				w.d2s.Stats.Vitality = v
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Unused stat points: "),
			giu.InputText(&ustats).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				u, _ := strconv.ParseUint(ustats, decimal, ustatsLen)
				w.d2s.Stats.UnusedStats = u
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Unused skill points: "),
			giu.InputText(&uskills).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				u, _ := strconv.ParseUint(uskills, decimal, uskillsLen)
				w.d2s.Stats.UnusedSkillPoints = u
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("HP: "),
			giu.InputText(&currentHP).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				c, _ := strconv.ParseFloat(currentHP, floatSize)
				w.d2s.Stats.CurrentHP = c
			}).Size(inputIntW),
			giu.Label("/"),
			giu.InputText(&maxHP).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				c, _ := strconv.ParseFloat(maxHP, floatSize)
				w.d2s.Stats.MaxHP = c
			}).Size(inputIntW),
		),
		giu.Style().SetColor(giu.StyleColorProgressBarActive, colornames.Red).To(
			giu.ProgressBar(float32(w.d2s.Stats.CurrentHP/w.d2s.Stats.MaxHP)).Size(progressBarSize, progressBarSize),
		),
		giu.Row(
			giu.Label("Mana: "),
			giu.InputText(&currentMana).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				c, _ := strconv.ParseFloat(currentMana, floatSize)
				w.d2s.Stats.CurrentMana = c
			}).Size(inputIntW),
			giu.Label("/"),
			giu.InputText(&maxMana).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				c, _ := strconv.ParseFloat(maxMana, floatSize)
				w.d2s.Stats.MaxMana = c
			}).Size(inputIntW),
		),
		giu.Style().SetColor(giu.StyleColorProgressBarActive, colornames.Blue).To(
			giu.ProgressBar(float32(w.d2s.Stats.CurrentMana/w.d2s.Stats.MaxMana)).Size(progressBarSize, progressBarSize),
		),
		giu.Row(
			giu.Label("Stamina: "),
			giu.InputText(&currentStamina).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				c, _ := strconv.ParseFloat(currentStamina, floatSize)
				w.d2s.Stats.CurrentStamina = c
			}).Size(inputIntW),
			giu.Label("/"),
			giu.InputText(&maxStamina).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				c, _ := strconv.ParseFloat(maxStamina, floatSize)
				w.d2s.Stats.MaxStamina = c
			}).Size(inputIntW),
		),
		giu.ProgressBar(float32(w.d2s.Stats.CurrentStamina / w.d2s.Stats.MaxStamina)),
		giu.Label(fmt.Sprintf("Level: %v", w.d2s.Stats.Level)),
		giu.Row(
			giu.Label("Experience: "),
			giu.InputText(&strength).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				e, _ := strconv.ParseUint(exp, decimal, expLen)
				w.d2s.Stats.Experience = e
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Gold: "),
			giu.InputText(&gold).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				g, _ := strconv.ParseUint(gold, decimal, goldLen)
				w.d2s.Stats.Gold = g
			}).Size(inputIntW),
		),
		giu.Row(
			giu.Label("Gold in stash: "),
			giu.InputText(&stashGold).
				Flags(giu.InputTextFlagsCharsDecimal).OnChange(func() {
				s, _ := strconv.ParseUint(stashGold, decimal, stashGoldLen)
				w.d2s.Stats.StashedGold = s
			}).Size(inputIntW),
		),
	}
}

func items(items *d2sitems.Items) giu.Layout {
	_ = items
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
