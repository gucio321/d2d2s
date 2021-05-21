package d2d2s

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/d2scorpse"
	"github.com/gucio321/d2d2s/d2sdifficulty"
	"github.com/gucio321/d2d2s/d2senums"
	"github.com/gucio321/d2d2s/d2sirongolem"
	"github.com/gucio321/d2d2s/d2sitems"
	"github.com/gucio321/d2d2s/d2smercenary"
	"github.com/gucio321/d2d2s/d2snpc"
	"github.com/gucio321/d2d2s/d2squests"
	"github.com/gucio321/d2d2s/d2sskills"
	"github.com/gucio321/d2d2s/d2sstats"
	"github.com/gucio321/d2d2s/d2sstatus"
	"github.com/gucio321/d2d2s/d2swaypoints"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	saveFileSignature  = 0xaa55aa55
	characterNameSize  = 16
	skillHotKeys       = 16
	unknown6BytesCount = 32
	unknown8BytesCount = 144
	int32Size          = 4
	byteLen            = 8
	fileSizePosition   = 8
	checksumPosition   = 12
)

type hotkeys map[byte]d2senums.SkillID

// D2S represents a Diablo II character save file structure
type D2S struct {
	Version     d2senums.Version
	unknown1    uint32
	Name        string
	Status      *d2sstatus.Status
	Progression byte
	unknown2    uint16
	Class       d2senums.CharacterClass
	unknown3    uint16
	Level       byte
	unknown4    uint32
	Time        uint32
	unknown5    uint32
	Hotkeys     hotkeys
	LeftSkill,
	RightSkill,
	LeftSkillSwitch,
	RightSkillSwitch d2senums.SkillID
	unknown6   [unknown6BytesCount]byte // probably character apperence in char select menu
	Difficulty *d2sdifficulty.Difficulty
	MapID      uint32
	unknown7   uint16
	Mercenary  *d2smercenary.Mercenary
	unknown8   [unknown8BytesCount]byte
	Quests     *d2squests.Quests
	Waypoints  *d2swaypoints.Waypoints
	NPC        *d2snpc.NPC
	Stats      *d2sstats.Stats
	Skills     *d2sskills.Skills
	Items      *d2sitems.Items
	Corpse     *d2scorpse.Corpse

	// necromancer only
	IronGolem *d2sirongolem.IronGolem
}

// New creates a new D2S structure
func New() *D2S {
	result := &D2S{
		Version:    d2senums.VersionLODLatest,
		Status:     d2sstatus.New(),
		Hotkeys:    make(hotkeys),
		Difficulty: d2sdifficulty.New(),
		Mercenary:  d2smercenary.New(),
		Quests:     d2squests.New(),
		Waypoints:  d2swaypoints.New(),
		NPC:        d2snpc.New(),
		Stats:      d2sstats.New(),
		Skills:     d2sskills.New(),
		Items:      &d2sitems.Items{},
		Corpse:     d2scorpse.New(),
		IronGolem:  d2sirongolem.New(),
	}

	return result
}

// Load loads d2s file into D2S structure
// nolint:funlen,gocyclo // can't reduce
func Load(data []byte) (*D2S, error) {
	result := New()

	sr := datautils.CreateBitMuncher(data, 0)

	if err := result.loadHeader(sr); err != nil {
		return nil, err
	}

	status := sr.GetByte()
	result.Status.Load(status)

	/*
		from: https://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml
		Character progression.  This number tells (sort of) how many acts you have completed from all
		difficulty levels.  It appears to be incremented when you kill the final demon in an act -- i.e.,
		Andarial, Duriel, Mephisto, and Diablo / Baal.  There's a catch to that last one: in an
		Expansion game, the value is not incremented after killing Diablo, but is incremented by 2 after killing Baal.
		(The reason is unknown.)  So it skips the values 4, 9, and 14.
	*/
	result.Progression = sr.GetByte()
	result.unknown2 = sr.GetUInt16()

	class := sr.GetByte()
	result.Class = d2senums.CharacterClass(class)

	result.unknown3 = sr.GetUInt16()
	result.Level = sr.GetByte()
	result.unknown4 = sr.GetUInt32()
	result.Time = sr.GetUInt32()
	result.unknown5 = sr.GetUInt32()
	result.loadHotkeys(sr)

	lsk := sr.GetUInt32()
	result.LeftSkill = d2senums.SkillID(lsk)

	rsk := sr.GetUInt32()
	result.RightSkill = d2senums.SkillID(rsk)

	if result.Status.Expansion {
		alsk := sr.GetUInt32()
		result.LeftSkillSwitch = d2senums.SkillID(alsk)

		arsk := sr.GetUInt32()
		result.RightSkillSwitch = d2senums.SkillID(arsk)
	}

	unknown6 := sr.GetBytes(unknown6BytesCount)

	copy(result.unknown6[:], unknown6[:unknown6BytesCount])

	result.Difficulty.Load(sr)

	result.MapID = sr.GetUInt32()
	result.unknown7 = sr.GetUInt16()

	result.Mercenary.LoadHeader(sr)

	unknown8 := sr.GetBytes(unknown8BytesCount)

	copy(result.unknown8[:], unknown8[:unknown8BytesCount])

	qd := sr.GetBytes(d2squests.NumQuestsBytes)

	var questsData [d2squests.NumQuestsBytes]byte

	copy(questsData[:], qd[:d2squests.NumQuestsBytes])

	err := result.Quests.Unmarshal(&questsData)
	if err != nil {
		return nil, fmt.Errorf("error loading quests: %w", err)
	}

	wd := sr.GetBytes(d2swaypoints.NumWaypointsBytes)

	var waypointsData [d2swaypoints.NumWaypointsBytes]byte

	copy(waypointsData[:], wd[:d2swaypoints.NumWaypointsBytes])

	if err := result.Waypoints.Load(&waypointsData); err != nil { // nolint:govet // it is ok
		return nil, fmt.Errorf("error loading waypoints data: %w", err)
	}

	nd := sr.GetBytes(d2snpc.NumNPCBytes)

	var npcData [d2snpc.NumNPCBytes]byte

	copy(npcData[:], nd[:d2snpc.NumNPCBytes])

	if err := result.NPC.Load(npcData); err != nil { // nolint:govet // it is og
		return nil, fmt.Errorf("error loading npcs data: %w", err)
	}

	if err := result.Stats.Load(sr); err != nil { // nolint:govet // it is ok
		return nil, fmt.Errorf("error loading character stats: %w", err)
	}

	if skillErr := result.Skills.Load(sr, result.Class); skillErr != nil {
		return nil, fmt.Errorf("error loading skills: %w", skillErr)
	}

	numItems, err := result.Items.LoadHeader(sr)
	if err != nil {
		return nil, fmt.Errorf("error loading items header: %w", err)
	}

	if err := result.Items.LoadList(sr, numItems); err != nil {
		return nil, fmt.Errorf("error loading items list: %w", err)
	}

	// thanks to @nokka <https://github.com/nokka/d2s> for figuring out these fields!
	if err := result.Corpse.Load(sr); err != nil {
		return nil, fmt.Errorf("error loading corpse: %w", err)
	}

	if result.Status.Expansion {
		if err := result.Mercenary.LoadMercItems(sr); err != nil {
			return nil, fmt.Errorf("error loading merc items: %w", err)
		}
	}

	// iron golem for necromancer
	if result.Class == d2senums.CharacterClassNecromancer && result.Status.Expansion {
		if err := result.IronGolem.Load(sr); err != nil {
			return nil, fmt.Errorf("error loading iron golem: %w", err)
		}
	}

	return result, nil
}

func (d *D2S) loadHeader(sr *datautils.BitMuncher) error {
	if signature := sr.GetUInt32(); signature != saveFileSignature {
		return errors.New("unexpected file signature")
	}

	v := sr.GetUInt32()
	version := d2senums.Version(v)

	if version != d2senums.VersionLODLatest {
		log.Printf("Warning! wrong version %s. It might be unsupported", version.String())
	}

	d.Version = version

	// file size in bytes ( len(data) )
	_ = sr.GetInt32()

	// checksum (32-bit checksum)
	_ = sr.GetUInt32()

	d.unknown1 = sr.GetUInt32()

	name := sr.GetBytes(characterNameSize)
	d.Name = strings.ReplaceAll(string(name), string(rune(0)), "")

	return nil
}

func (d *D2S) loadHotkeys(sr *datautils.BitMuncher) {
	for i := byte(0); i < skillHotKeys; i++ {
		id := sr.GetUInt32()
		d.Hotkeys[i] = d2senums.SkillID(id)
	}
}

// Encode encodes character save back into a byte slice (WIP)
func (d *D2S) Encode() ([]byte, error) {
	sw := d2datautils.CreateStreamWriter()

	if err := d.encodeHeader(sw); err != nil {
		return nil, err
	}

	sw.PushBytes(d.Status.Encode())
	sw.PushBytes(d.Progression)
	sw.PushUint16(d.unknown2)
	sw.PushBytes(byte(d.Class))
	sw.PushUint16(d.unknown3)
	sw.PushBytes(d.Level)
	sw.PushUint32(d.unknown4)
	sw.PushUint32(d.Time)
	sw.PushUint32(d.unknown5)

	d.encodeHotkeys(sw)

	sw.PushUint32(uint32(d.LeftSkill))
	sw.PushUint32(uint32(d.RightSkill))

	if d.Status.Expansion {
		sw.PushUint32(uint32(d.LeftSkillSwitch))
		sw.PushUint32(uint32(d.RightSkillSwitch))
	}

	sw.PushBytes(d.unknown6[:]...)

	d.Difficulty.Encode(sw)

	sw.PushUint32(d.MapID)
	sw.PushUint16(d.unknown7)

	d.Mercenary.EncodeHeader(sw)

	sw.PushBytes(d.unknown8[:]...)

	qd := d.Quests.Encode()
	sw.PushBytes(qd[:]...)

	wd := d.Waypoints.Encode()
	sw.PushBytes(wd[:]...)

	nd := d.NPC.Encode()
	sw.PushBytes(nd[:]...)

	sd, err := d.Stats.Encode()
	if err != nil {
		return nil, fmt.Errorf("error encoding stats: %w", err)
	}

	sw.PushBytes(sd...)

	d.Skills.Encode(sw, d.Class)

	sw.PushBytes(d.Items.Encode()...)

	if err := d.Corpse.Encode(sw); err != nil {
		return nil, fmt.Errorf("error encoding corpse: %w", err)
	}

	d.Mercenary.EncodeItems(sw)

	if d.Class == d2senums.CharacterClassNecromancer && d.Status.Expansion {
		d.IronGolem.Encode(sw)
	}

	// we need to write file size and checksum here:
	data := sw.GetBytes()
	CalculateFileSize(&data)
	CalculateChecksum(&data)

	return data, nil
}

func (d *D2S) encodeHeader(sw *d2datautils.StreamWriter) error {
	sw.PushUint32(saveFileSignature)
	sw.PushUint32(uint32(d.Version))
	// file size, 0 for now
	sw.PushUint32(0)
	// checksum - 0 for now
	sw.PushUint32(0)

	name := []byte(d.Name)
	if len(name) > characterNameSize {
		return errors.New("wrong character name! (len(name) > 16)")
	}

	sw.PushUint32(d.unknown1)

	sw.PushBytes(name...)

	for i := 0; i < characterNameSize-len(name); i++ {
		sw.PushBytes(0)
	}

	return nil
}

func (d *D2S) encodeHotkeys(sw *d2datautils.StreamWriter) {
	for i := byte(0); i < skillHotKeys; i++ {
		sw.PushUint32(uint32(d.Hotkeys[i]))
	}
}

// CalculateChecksum calculates a checksum and saves in a byte slice
func CalculateChecksum(data *[]byte) {
	var sum uint32
	for i := range *data {
		// thanks goes to <https://github.com/pairofdocs>
		// this shift expresion was translated from here: https://github.com/pairofdocs/d2s_edit_recalc
		// origin:
		// ```python
		// # from stackoverflow ref
		// def leftshift(int_in, shift_n, tot_bits):
		//    # returns an int.  take bin() to get  '0b10101'
		//    # bin(leftshift(int('10000000000000000000000000000000',2), 1, 32) )  ---> 0000...1
		//    return ((int_in << shift_n) % (1 << tot_bits)) | (int_in >> (tot_bits - shift_n))
		// ```
		sum = ((sum << 1) % math.MaxUint32) | (sum >> (int32Size*byteLen - 1))

		sum += uint32((*data)[i])
	}

	sumBytes := make([]byte, int32Size)
	binary.LittleEndian.PutUint32(sumBytes, sum)

	for i := 0; i < int32Size; i++ {
		(*data)[checksumPosition+i] = sumBytes[i]
	}
}

// CalculateFileSize calculates a file size in bytes and writes it in byte slice
func CalculateFileSize(data *[]byte) {
	fileSize := uint32(len(*data))

	fileSizeBytes := make([]byte, int32Size)
	binary.LittleEndian.PutUint32(fileSizeBytes, fileSize)

	for i := 0; i < int32Size; i++ {
		(*data)[fileSizePosition+i] = fileSizeBytes[i]
	}
}
