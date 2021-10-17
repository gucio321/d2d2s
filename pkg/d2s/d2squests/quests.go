package d2squests

import (
	"fmt"

	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

const (
	// NumQuestsBytes is a total number of quests data bytes
	NumQuestsBytes = 298

	expectedQuestHeaderID        = "Woo!"
	questHeaderIDLen             = 4
	questHeaderUnknownBytesCount = 6
	defaultQuestsCount           = 6
	act4QuestsCount              = 3
	act4                         = 4
	act5                         = 5
	act4UnknownBytesCount        = 3
	act5Unknown1BytesCount       = 2
	act5Unknown2BytesCount       = 7
)

// Quests represents quests status structure
type Quests map[d2senums.DifficultyType]*[d2senums.NumActs]*QuestsSet

// New creates a new quests structure
func New() *Quests {
	result := &Quests{}
	*result = make(Quests)

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		(*result)[i] = &[d2senums.NumActs]*QuestsSet{}

		for act := 1; act <= d2senums.NumActs; act++ {
			var l int
			if act == act4 {
				l = act4QuestsCount
			} else {
				l = defaultQuestsCount
			}

			(*result)[i][act-1] = &QuestsSet{
				Act:    act,
				Quests: make([]*Quest, l),
			}

			for qst := range (*result)[i][act-1].Quests {
				(*result)[i][act-1].Quests[qst] = &Quest{}
			}

			switch act {
			case act4:
				(*result)[i][act-1].unknown1 = make([]byte, act4UnknownBytesCount)
			case act5:
				(*result)[i][act-1].unknown1 = make([]byte, act5Unknown1BytesCount)
				(*result)[i][act-1].unknown2 = make([]byte, act5Unknown2BytesCount)
			}
		}
	}

	return result
}

func unknownQuestsHeaderBytes() [questHeaderUnknownBytesCount]byte {
	return [questHeaderUnknownBytesCount]byte{6, 0, 0, 0, 42, 1}
}

// Unmarshal unmarshals quests status data
func (q *Quests) Unmarshal(data *[NumQuestsBytes]byte) error {
	sr := datareader.NewReader((*data)[:])

	questHeaderID := sr.GetBytes(questHeaderIDLen)

	if string(questHeaderID) != expectedQuestHeaderID {
		return fmt.Errorf("quest header: %w", common.ErrUnexpectedHeader)
	}

	_ = sr.GetBytes(questHeaderUnknownBytesCount)

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		for act := 1; act <= d2senums.NumActs; act++ {
			err := (*q)[i][act-1].Unmarshal(sr, act)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Encode encodes quests back into byte array
func (q *Quests) Encode() (result [NumQuestsBytes]byte) {
	sw := datautils.CreateStreamWriter()

	sw.PushBytes([]byte(expectedQuestHeaderID)...)

	unknown := unknownQuestsHeaderBytes()
	sw.PushBytes(unknown[:]...)

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		for act := 1; act <= d2senums.NumActs; act++ {
			data := (*q)[i][act-1].Encode()
			sw.PushBytes(data...)
		}
	}

	data := sw.GetBytes()
	copy(result[:], data[:NumQuestsBytes])

	return result
}

// QuestsSet represents a set of 6 (in case of act 4 - 3) quests
type QuestsSet struct {
	Introduced bool
	Quests     []*Quest
	ActEnd     bool // uncertain
	Act        int
	unknown1   []byte
	unknown2   []byte
}

// Unmarshal unmarshals quests set
func (q *QuestsSet) Unmarshal(sr *datareader.Reader, act int) (err error) {
	q.Act = act

	loadQuests := func(sr *datareader.Reader) {
		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			data := sr.GetBytes(2) // nolint:gomnd // quest bytes. TODO: https://github.com/gucio321/d2d2s/issues/11
			copy(q.Quests[qst].Data[:], data[:2])
			q.Quests[qst].Load()
		}
	}

	switch act {
	case act4:
		q.Introduced = sr.GetUint16() == 1

		loadQuests(sr)

		q.ActEnd = sr.GetUint16() == 1
		q.unknown1 = sr.GetBytes(act4UnknownBytesCount)
	case act5:
		q.Introduced = sr.GetUint16() == 1
		q.unknown1 = sr.GetBytes(act5Unknown1BytesCount)

		loadQuests(sr)

		q.ActEnd = sr.GetUint16() == 1

		q.unknown2 = sr.GetBytes(act5Unknown2BytesCount)
	default:
		q.Introduced = sr.GetUint16() == 1

		loadQuests(sr)

		q.ActEnd = sr.GetUint16() == 1
	}

	return nil
}

// Encode encodes quests set into a byte slice
func (q *QuestsSet) Encode() []byte {
	sw := datautils.CreateStreamWriter()

	switch q.Act {
	case act4:
		if q.Introduced {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		for qst := range q.Quests {
			// TODO: Quest.Encode()
			sw.PushBytes(q.Quests[qst].Data[:]...)
		}

		if q.ActEnd {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		sw.PushBytes(q.unknown1...)
	case act5:
		if q.Introduced {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		sw.PushBytes(q.unknown1...)

		for qst := range q.Quests {
			// TODO: Quest.Encode()
			sw.PushBytes(q.Quests[qst].Data[:]...)
		}

		if q.ActEnd {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		sw.PushBytes(q.unknown2...)
	default:
		if q.Introduced {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		for qst := range q.Quests {
			// TODO: Quest.Encode()
			sw.PushBytes(q.Quests[qst].Data[:]...)
		}

		if q.ActEnd {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}
	}

	data := sw.GetBytes()

	return data
}

// Quest represents a single quest's state
type Quest struct {
	Data          [2]byte
	Completed     bool
	Done          bool // all requirements has been completed (need to get reward)
	Started       bool // when NPC gives you quest
	Body          [9]bool
	Closed        bool //  you have seen the swirling fire animation that closes a quest icon.
	JustCompleted bool // gets set, when completed in current game
}

// Load loads quest into Quest structure
func (q *Quest) Load() {
	bm := datareader.NewReader(q.Data[:])

	q.Completed = bm.GetBit()
	q.Done = bm.GetBit()
	q.Started = bm.GetBit()

	for i := 0; i < 9; i++ {
		q.Body[i] = bm.GetBit()
	}

	q.Closed = bm.GetBit()
	q.JustCompleted = bm.GetBit()
}
