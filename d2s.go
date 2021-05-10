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
}

func Unmarshal(data []byte) (*D2S, error) {
	var err error

	fmt.Println("exec")
	result := &D2S{
		Status:     &Status{},
		Hotkeys:    make(hotkeys),
		Difficulty: make(Difficulty),
		Quests:     &Quests{},
		Waypoints:  make(Waypoints),
		NPC:        &NPC{},
		Stats:      &Stats{},
	}
	sr := datautils.CreateStreamReader(data)

	signature, err := sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	if signature != saveFileSignature {
		return nil, errors.New("Unexpected file signature")
	}

	v, err := sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.version = version(v)

	// file size in bytes ( len(data) )
	_, err = sr.ReadInt32()
	if err != nil {
		return nil, err
	}

	// checksum (32-bit checksum)
	_, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.unknown1, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	name, err := sr.ReadBytes(characterNameSize)
	if err != nil {
		return nil, err
	}

	result.Name = string(name)

	status, err := sr.ReadByte()
	if err != nil {
		return nil, err
	}

	result.Status.Unmarshal(status)

	/*
		from: https://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml
		Character progression.  This number tells (sort of) how many acts you have completed from all
		difficulty levels.  It appears to be incremented when you kill the final demon in an act -- i.e.,
		Andarial, Duriel, Mephisto, and Diablo / Baal.  There's a catch to that last one: in an
		Expansion game, the value is not incremented after killing Diablo, but is incremented by 2 after killing Baal.
		(The reason is unknown.)  So it skips the values 4, 9, and 14.
	*/
	_, err = sr.ReadByte()
	if err != nil {
		return nil, err
	}

	result.unknown2, err = sr.ReadUInt16()
	if err != nil {
		return nil, err
	}

	class, err := sr.ReadByte()
	if err != nil {
		return nil, err
	}

	result.Class = CharacterClass(class)

	result.unknown3, err = sr.ReadUInt16()
	if err != nil {
		return nil, err
	}

	result.Level, err = sr.ReadByte()
	if err != nil {
		return nil, err
	}

	result.unknown4, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.Time, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.unknown5, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	for i := byte(0); i < skillHotKeys; i++ {
		id, err := sr.ReadUInt32()
		if err != nil {
			return nil, err
		}
		result.Hotkeys[i] = SkillID(id)
	}

	lsk, err := sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.LeftSkill = SkillID(lsk)

	rsk, err := sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.RightSkill = SkillID(rsk)

	if result.Status.Expansion {
		alsk, err := sr.ReadUInt32()
		if err != nil {
			return nil, err
		}

		result.LeftSkillSwitch = SkillID(alsk)

		arsk, err := sr.ReadUInt32()
		if err != nil {
			return nil, err
		}

		result.RightSkillSwitch = SkillID(arsk)
	}

	unknown6, err := sr.ReadBytes(unknown6BytesCount)
	if err != nil {
		return nil, err
	}

	copy(result.unknown6[:], unknown6[:unknown6BytesCount])

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		d, err := sr.ReadByte()
		if err != nil {
			return nil, err
		}

		if d == 0 {
			continue
		}

		result.Difficulty[i] = &DifficultyLevelStatus{}
		result.Difficulty[i].Unmarshal(d)
	}

	result.MapID, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.unknown7, err = sr.ReadUInt16()
	if err != nil {
		return nil, err
	}

	result.Mercenary.Died, err = sr.ReadUInt16()
	if err != nil {
		return nil, err
	}

	result.Mercenary.ID, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	result.Mercenary.Name, err = sr.ReadUInt16()
	if err != nil {
		return nil, err
	}

	mercType, err := sr.ReadUInt16()
	if err != nil {
		return nil, err
	}

	result.Mercenary.LoadType(mercType)

	result.Mercenary.Experience, err = sr.ReadUInt32()
	if err != nil {
		return nil, err
	}

	unknown8, err := sr.ReadBytes(unknown8BytesCount)
	if err != nil {
		return nil, err
	}

	copy(result.unknown8[:], unknown8[:unknown8BytesCount])

	qd, err := sr.ReadBytes(numQuestsBytes)

	var questsData [numQuestsBytes]byte
	copy(questsData[:], qd[:numQuestsBytes])

	err = result.Quests.Unmarshal(questsData)
	if err != nil {
		return nil, fmt.Errorf("error loading quests: %w", err)
	}

	wd, err := sr.ReadBytes(numWaypointsBytes)
	var waypointsData [numWaypointsBytes]byte
	copy(waypointsData[:], wd[:numWaypointsBytes])

	if err := result.Waypoints.Load(waypointsData); err != nil {
		return nil, fmt.Errorf("error loading waypoints data: %w", err)
	}

	nd, err := sr.ReadBytes(numNPCBytes)
	var npcData [numNPCBytes]byte
	copy(npcData[:], nd[:numNPCBytes])

	if err := result.NPC.Load(npcData); err != nil {
		return nil, fmt.Errorf("error loading npcs data: %w", err)
	}

	if err := result.Stats.Load(sr); err != nil {
		return nil, fmt.Errorf("error loading character stats: %w", err)
	}

	return result, nil
}
