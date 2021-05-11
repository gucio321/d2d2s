package d2d2s

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	saveFileSignature     = 0xaa55aa55
	characterNameSize     = 16
	skillHotKeys          = 16
	unknown6BytesCount    = 32
	numDifficoultyLevels  = 3
	unknown8BytesCount    = 144
	expectedQuestHeaderID = "Woo!"
	skillsHeaderID        = "if"
	numSkills             = 30
)

type hotkeys map[byte]SkillID

type D2S struct {
	version  version
	unknown1 uint32
	Name     string
	Status   *Status
	unknown2 uint16
	Class    CharacterClass
	unknown3 uint16
	Level    byte
	unknown4 uint32
	Time     uint32
	unknown5 uint32
	Hotkeys  hotkeys
	LeftSkill,
	RightSkill,
	LeftSkillSwitch,
	RightSkillSwitch SkillID
	unknown6   [unknown6BytesCount]byte // probably character apperence in char select menu
	Difficulty Difficulty
	MapID      uint32
	unknown7   uint16
	Mercenary  mercenary
	unknown8   [unknown8BytesCount]byte
	Quests     *Quests
	Waypoints  Waypoints
	NPC        *NPC
	Stats      *Stats
	Skills     [numSkills]SkillID
	Items      *Items
}

func Unmarshal(data []byte) (*D2S, error) {
	var err error

	result := &D2S{
		Status:     &Status{},
		Hotkeys:    make(hotkeys),
		Difficulty: make(Difficulty),
		Quests:     &Quests{},
		Waypoints:  make(Waypoints),
		NPC:        &NPC{},
		Stats:      &Stats{},
		Items:      &Items{},
	}
	// sr := datautils.CreateStreamGeter(data)
	sr := datautils.CreateBitMuncher(data, 0)

	if signature := sr.GetUInt32(); signature != saveFileSignature {
		return nil, errors.New("unexpected file signature")
	}

	v := sr.GetUInt32()
	result.version = version(v)

	// file size in bytes ( len(data) )
	_ = sr.GetInt32()

	// checksum (32-bit checksum)
	_ = sr.GetUInt32()

	result.unknown1 = sr.GetUInt32()

	name := sr.GetBytes(characterNameSize)
	result.Name = string(name)

	status := sr.GetByte()
	result.Status.Unmarshal(status)

	/*
		from: https://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml
		Character progression.  This number tells (sort of) how many acts you have completed from all
		difficulty levels.  It appears to be incremented when you kill the final demon in an act -- i.e.,
		Andarial, Duriel, Mephisto, and Diablo / Baal.  There's a catch to that last one: in an
		Expansion game, the value is not incremented after killing Diablo, but is incremented by 2 after killing Baal.
		(The reason is unknown.)  So it skips the values 4, 9, and 14.
	*/
	_ = sr.GetByte()

	result.unknown2 = sr.GetUInt16()

	class := sr.GetByte()
	result.Class = CharacterClass(class)

	result.unknown3 = sr.GetUInt16()

	result.Level = sr.GetByte()

	result.unknown4 = sr.GetUInt32()

	result.Time = sr.GetUInt32()

	result.unknown5 = sr.GetUInt32()

	for i := byte(0); i < skillHotKeys; i++ {
		id := sr.GetUInt32()
		result.Hotkeys[i] = SkillID(id)
	}

	lsk := sr.GetUInt32()
	result.LeftSkill = SkillID(lsk)

	rsk := sr.GetUInt32()
	result.RightSkill = SkillID(rsk)

	if result.Status.Expansion {
		alsk := sr.GetUInt32()
		result.LeftSkillSwitch = SkillID(alsk)

		arsk := sr.GetUInt32()
		result.RightSkillSwitch = SkillID(arsk)
	}

	unknown6 := sr.GetBytes(unknown6BytesCount)

	copy(result.unknown6[:], unknown6[:unknown6BytesCount])

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		d := sr.GetByte()
		if d == 0 {
			continue
		}

		result.Difficulty[i] = &DifficultyLevelStatus{}
		result.Difficulty[i].Unmarshal(d)
	}

	result.MapID = sr.GetUInt32()

	result.unknown7 = sr.GetUInt16()

	result.Mercenary.Died = sr.GetUInt16()

	result.Mercenary.ID = sr.GetUInt32()

	result.Mercenary.Name = sr.GetUInt16()

	mercType := sr.GetUInt16()

	result.Mercenary.LoadType(mercType)

	result.Mercenary.Experience = sr.GetUInt32()

	unknown8 := sr.GetBytes(unknown8BytesCount)

	copy(result.unknown8[:], unknown8[:unknown8BytesCount])

	qd := sr.GetBytes(numQuestsBytes)

	var questsData [numQuestsBytes]byte

	copy(questsData[:], qd[:numQuestsBytes])

	err = result.Quests.Unmarshal(questsData)
	if err != nil {
		return nil, fmt.Errorf("error loading quests: %w", err)
	}

	wd := sr.GetBytes(numWaypointsBytes)

	var waypointsData [numWaypointsBytes]byte

	copy(waypointsData[:], wd[:numWaypointsBytes])

	if err := result.Waypoints.Load(waypointsData); err != nil {
		return nil, fmt.Errorf("error loading waypoints data: %w", err)
	}

	nd := sr.GetBytes(numNPCBytes)

	var npcData [numNPCBytes]byte

	copy(npcData[:], nd[:numNPCBytes])

	if err := result.NPC.Load(npcData); err != nil {
		return nil, fmt.Errorf("error loading npcs data: %w", err)
	}

	if err := result.Stats.Load(sr); err != nil {
		return nil, fmt.Errorf("error loading character stats: %w", err)
	}

	skillsID := sr.GetBytes(2)

	if string(skillsID) != skillsHeaderID {
		return nil, errors.New("unexpected skills section header")
	}

	for i := 0; i < numSkills; i++ {
		id := sr.GetByte()
		result.Skills[i] = SkillID(id)
	}

	if err := result.Items.Load(sr); err != nil {
		return nil, err
	}

	return result, nil
}
