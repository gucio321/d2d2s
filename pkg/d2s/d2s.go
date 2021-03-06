package d2s

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2scorpse"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sdifficulty"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
	"github.com/gucio321/d2d2s/pkg/d2s/d2shotkeys"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sirongolem"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
	"github.com/gucio321/d2d2s/pkg/d2s/d2smercenary"
	"github.com/gucio321/d2d2s/pkg/d2s/d2snpc"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sprogression"
	"github.com/gucio321/d2d2s/pkg/d2s/d2squests"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sskills"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sstats"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sstatus"
	"github.com/gucio321/d2d2s/pkg/d2s/d2swaypoints"
)

const (
	saveFileSignature  = 0xaa55aa55
	characterNameSize  = 16
	unknown6BytesCount = 32
	unknown8BytesCount = 144
	int32Size          = 4
	byteLen            = 8
	fileSizePosition   = 8
	checksumPosition   = 12
)

// D2S represents a Diablo II character save file structure
type D2S struct {
	Version     d2senums.Version
	unknown1    uint32
	Name        string
	Status      *d2sstatus.Status
	Progression *d2sprogression.Progression
	unknown2    uint16
	Class       d2senums.CharacterClass
	unknown3    uint16
	Level       byte
	unknown4    uint32
	Time        time.Time
	unknown5    uint32
	Hotkeys     *d2shotkeys.Hotkeys
	LeftSkill,
	RightSkill,
	LeftSkillSwitch,
	RightSkillSwitch d2senums.SkillID
	unknown6   [unknown6BytesCount]byte // probably character appearance in char select menu
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
		Version:     d2senums.VersionLODLatest,
		Status:      d2sstatus.New(),
		Progression: d2sprogression.New(),
		Hotkeys:     d2shotkeys.New(),
		Difficulty:  d2sdifficulty.New(),
		Mercenary:   d2smercenary.New(),
		Quests:      d2squests.New(),
		Waypoints:   d2swaypoints.New(),
		NPC:         d2snpc.New(),
		Stats:       d2sstats.New(),
		Skills:      d2sskills.New(),
		Items:       d2sitems.New(),
		Corpse:      d2scorpse.New(),
		IronGolem:   d2sirongolem.New(),
	}

	return result
}

// Load creates a new d2s structure and loads d2s file's data into it.
func Load(data []byte) (*D2S, error) {
	result := New()
	err := result.Load(data)

	return result, err
}

// Load loads d2s file into D2S structure
func (d2s *D2S) Load(data []byte) error {
	sr := datareader.NewReader(data)

	if err := d2s.loadHeader(sr); err != nil {
		return fmt.Errorf("loading header: %w", err)
	}

	d2s.loadCharDetails(sr)

	d2s.loadInterfaceState(sr)

	unknown6 := sr.GetBytes(unknown6BytesCount)

	copy(d2s.unknown6[:], unknown6[:unknown6BytesCount])

	dd := sr.GetBytes(d2sdifficulty.NumDifficultyBytes)

	var difficultyData [d2sdifficulty.NumDifficultyBytes]byte

	copy(difficultyData[:], dd[:d2sdifficulty.NumDifficultyBytes])

	d2s.Difficulty.Load(difficultyData)

	if err := d2s.loadMapDetails(sr); err != nil {
		return fmt.Errorf("loading map data: %w", err)
	}

	if err := d2s.loadCharacterDetails(sr); err != nil {
		return fmt.Errorf("loading character details: %w", err)
	}

	return nil
}

func (d2s *D2S) loadHeader(sr *datareader.Reader) error {
	if signature := sr.GetUint32(); signature != saveFileSignature {
		return fmt.Errorf("unexpected file signature: %w", common.ErrUnexpectedHeader)
	}

	v := sr.GetUint32()
	version := d2senums.Version(v)

	if version != d2senums.VersionLODLatest {
		log.Printf("Warning! wrong version %s. It might be unsupported", version.String())
	}

	d2s.Version = version

	// file size in bytes ( len(data) )
	_ = sr.GetInt32()

	// checksum (32-bit checksum)
	_ = sr.GetUint32()

	d2s.unknown1 = sr.GetUint32()

	name := sr.GetBytes(characterNameSize)
	d2s.Name = strings.ReplaceAll(string(name), string(rune(0)), "")

	return nil
}

func (d2s *D2S) loadCharDetails(sr *datareader.Reader) {
	status := sr.GetByte()
	d2s.Status.Load(status)

	d2s.Progression.Load(sr.GetByte())
	d2s.unknown2 = sr.GetUint16()

	class := sr.GetByte()
	d2s.Class = d2senums.CharacterClass(class)

	d2s.unknown3 = sr.GetUint16()
	d2s.Level = sr.GetByte()
	d2s.unknown4 = sr.GetUint32()
	d2s.Time = time.Unix(int64(sr.GetUint32()), 0)
	d2s.unknown5 = sr.GetUint32()
}

func (d2s *D2S) loadInterfaceState(sr *datareader.Reader) {
	hd := sr.GetBytes(d2shotkeys.NumHotkeysBytes)

	var hotkeysData [d2shotkeys.NumHotkeysBytes]byte

	copy(hotkeysData[:], hd[:d2shotkeys.NumHotkeysBytes])

	d2s.Hotkeys.Load(hotkeysData)

	lsk := sr.GetUint32()
	d2s.LeftSkill = d2senums.SkillID(lsk)

	rsk := sr.GetUint32()
	d2s.RightSkill = d2senums.SkillID(rsk)

	if d2s.Status.Expansion {
		alternativeLeftSkill := sr.GetUint32()
		d2s.LeftSkillSwitch = d2senums.SkillID(alternativeLeftSkill)

		alternativeRightSkill := sr.GetUint32()
		d2s.RightSkillSwitch = d2senums.SkillID(alternativeRightSkill)
	}
}

func (d2s *D2S) loadMapDetails(sr *datareader.Reader) error {
	d2s.MapID = sr.GetUint32()
	d2s.unknown7 = sr.GetUint16()

	d2s.Mercenary.LoadHeader(sr)

	unknown8 := sr.GetBytes(unknown8BytesCount)

	copy(d2s.unknown8[:], unknown8[:unknown8BytesCount])

	qd := sr.GetBytes(d2squests.NumQuestsBytes)

	var questsData [d2squests.NumQuestsBytes]byte

	copy(questsData[:], qd[:d2squests.NumQuestsBytes])

	if err := d2s.Quests.Unmarshal(&questsData); err != nil {
		return fmt.Errorf("loading quests: %w", err)
	}

	wd := sr.GetBytes(d2swaypoints.NumWaypointsBytes)

	var waypointsData [d2swaypoints.NumWaypointsBytes]byte

	copy(waypointsData[:], wd[:d2swaypoints.NumWaypointsBytes])

	if err := d2s.Waypoints.Load(&waypointsData); err != nil {
		return fmt.Errorf("loading waypoints data: %w", err)
	}

	nd := sr.GetBytes(d2snpc.NumNPCBytes)

	var npcData [d2snpc.NumNPCBytes]byte

	copy(npcData[:], nd[:d2snpc.NumNPCBytes])

	if err := d2s.NPC.Load(npcData); err != nil {
		return fmt.Errorf("loading npcs data: %w", err)
	}

	return nil
}

func (d2s *D2S) loadCharacterDetails(sr *datareader.Reader) error {
	if err := d2s.Stats.Load(sr); err != nil {
		return fmt.Errorf("loading character stats: %w", err)
	}

	if skillErr := d2s.Skills.Load(sr, d2s.Class); skillErr != nil {
		return fmt.Errorf("loading skills: %w", skillErr)
	}

	numItems, err := d2s.Items.LoadHeader(sr)
	if err != nil {
		return fmt.Errorf("loading items header: %w", err)
	}

	if err := d2s.Items.LoadList(sr, numItems); err != nil {
		return fmt.Errorf("loading items list: %w", err)
	}

	// thanks to @nokka <https://github.com/nokka/d2s> for figuring out that fields!
	if err := d2s.Corpse.Load(sr); err != nil {
		return fmt.Errorf("loading corpse: %w", err)
	}

	if d2s.Status.Expansion {
		if err := d2s.Mercenary.LoadMercItems(sr); err != nil {
			return fmt.Errorf("loading merc items: %w", err)
		}
	}

	// iron golem for necromancer
	if d2s.Class == d2senums.CharacterNecromancer && d2s.Status.Expansion {
		if err := d2s.IronGolem.Load(sr); err != nil {
			return fmt.Errorf("loading iron golem: %w", err)
		}
	}

	return nil
}

// Encode encodes character save back into a byte slice (WIP)
func (d2s *D2S) Encode() ([]byte, error) {
	sw := datautils.CreateStreamWriter()

	if err := d2s.encodeHeader(sw); err != nil {
		return nil, err
	}

	d2s.encodeCharDetails(sw)
	d2s.encodeInterfaceState(sw)

	sw.PushBytes(d2s.unknown6[:]...)

	dd := d2s.Difficulty.Encode()
	sw.PushBytes(dd[:]...)

	d2s.encodeMapDetails(sw)

	if err := d2s.encodeCharacterDetails(sw); err != nil {
		return nil, fmt.Errorf("encoding character details: %w", err)
	}

	// we need to write file size and checksum here:
	data := sw.GetBytes()
	CalculateFileSize(&data)
	CalculateChecksum(&data)

	return data, nil
}

func (d2s *D2S) encodeHeader(sw *datautils.StreamWriter) error {
	sw.PushUint32(saveFileSignature)
	sw.PushUint32(uint32(d2s.Version))
	// file size, 0 for now
	sw.PushUint32(0)
	// checksum - 0 for now
	sw.PushUint32(0)

	name := []byte(d2s.Name)
	if len(name) > characterNameSize {
		return errors.New("wrong character name! (len(name) > 16)")
	}

	sw.PushUint32(d2s.unknown1)

	sw.PushBytes(name...)

	for i := 0; i < characterNameSize-len(name); i++ {
		sw.PushBytes(0)
	}

	return nil
}

func (d2s *D2S) encodeCharDetails(sw *datautils.StreamWriter) {
	sw.PushBytes(d2s.Status.Encode())
	sw.PushBytes(d2s.Progression.Encode())
	sw.PushUint16(d2s.unknown2)
	sw.PushBytes(byte(d2s.Class))
	sw.PushUint16(d2s.unknown3)
	sw.PushBytes(d2s.Level)
	sw.PushUint32(d2s.unknown4)
	sw.PushUint32(uint32(d2s.Time.Unix()))
	sw.PushUint32(d2s.unknown5)
}

func (d2s *D2S) encodeInterfaceState(sw *datautils.StreamWriter) {
	hd := d2s.Hotkeys.Encode()
	sw.PushBytes(hd[:]...)

	sw.PushUint32(uint32(d2s.LeftSkill))
	sw.PushUint32(uint32(d2s.RightSkill))

	if d2s.Status.Expansion {
		sw.PushUint32(uint32(d2s.LeftSkillSwitch))
		sw.PushUint32(uint32(d2s.RightSkillSwitch))
	}
}

func (d2s *D2S) encodeMapDetails(sw *datautils.StreamWriter) {
	sw.PushUint32(d2s.MapID)
	sw.PushUint16(d2s.unknown7)

	d2s.Mercenary.EncodeHeader(sw)

	sw.PushBytes(d2s.unknown8[:]...)

	qd := d2s.Quests.Encode()
	sw.PushBytes(qd[:]...)

	wd := d2s.Waypoints.Encode()
	sw.PushBytes(wd[:]...)

	nd := d2s.NPC.Encode()
	sw.PushBytes(nd[:]...)
}

func (d2s *D2S) encodeCharacterDetails(sw *datautils.StreamWriter) error {
	sd, err := d2s.Stats.Encode()
	if err != nil {
		return fmt.Errorf("encoding stats: %w", err)
	}

	sw.PushBytes(sd...)

	d2s.Skills.Encode(sw, d2s.Class)

	sw.PushBytes(d2s.Items.Encode()...)

	if err := d2s.Corpse.Encode(sw); err != nil {
		return fmt.Errorf("encoding corpse: %w", err)
	}

	d2s.Mercenary.EncodeItems(sw)

	if d2s.Class == d2senums.CharacterNecromancer && d2s.Status.Expansion {
		d2s.IronGolem.Encode(sw)
	}

	return nil
}

// CalculateChecksum calculates a checksum and saves in a byte slice
func CalculateChecksum(data *[]byte) {
	var sum uint32
	for i := range *data {
		// thanks goes to <https://github.com/pairofdocs>
		// this shift expression was translated from here: https://github.com/pairofdocs/d2s_edit_recalc
		// origin:
		// ```python
		// # from stackoverflow ref
		// def leftShift(int_in, shift_n, tot_bits):
		//    # returns an int.  take bin() to get  '0b10101'
		//    # bin(leftShift(int('10000000000000000000000000000000',2), 1, 32) )  ---> 0000...1
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
