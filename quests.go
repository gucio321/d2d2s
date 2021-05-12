package d2d2s

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/gucio321/d2d2s/datautils"
)

const (
	numQuestsBytes               = 298
	questHeaderUnknownBytesCount = 6
	defaultQuestsCount           = 6
	act4QuestsCount              = 3
)

// Quests represents quests status structure
type Quests map[d2enum.DifficultyType]*[numActs]*QuestsSet

func unknownQuestsHeaderBytes() [questHeaderUnknownBytesCount]byte {
	return [questHeaderUnknownBytesCount]byte{6, 0, 0, 0, 42, 1}
}

// Unmarshal unmarshals quests status data
func (q *Quests) Unmarshal(data *[numQuestsBytes]byte) error {
	sr := datautils.CreateBitMuncher((*data)[:], 0)

	questHeaderID := sr.GetBytes(4) // nolint:gomnd // header

	if string(questHeaderID) != expectedQuestHeaderID {
		return errors.New("unexpected quest header")
	}

	_ = sr.GetBytes(questHeaderUnknownBytesCount)

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		(*q)[i] = &[numActs]*QuestsSet{}

		for act := 1; act <= numActs; act++ {
			var l int
			if act == 4 { // nolint:gomnd // act 4
				l = act4QuestsCount
			} else {
				l = defaultQuestsCount
			}

			(*q)[i][act-1] = &QuestsSet{
				Quests: make([]*Quest, l),
			}

			err := (*q)[i][act-1].Unmarshal(sr, act)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Encode encodes quests back into byte array
func (q *Quests) Encode() (result [numQuestsBytes]byte) {
	sw := d2datautils.CreateStreamWriter()

	sw.PushBytes([]byte(expectedQuestHeaderID)...)

	unknown := unknownQuestsHeaderBytes()
	sw.PushBytes(unknown[:]...)

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		for act := 1; act <= numActs; act++ {
			data := (*q)[i][act-1].Encode()
			sw.PushBytes(data...)
		}
	}

	data := sw.GetBytes()
	copy(result[:], data[:numQuestsBytes])

	return result
}

// QuestsSet represents a set of 6 (in case of act 4 - 3) quests
type QuestsSet struct {
	Introduced bool
	Quests     []*Quest
	ActEnd     bool // uncertain
	Act        int
}

// Unmarshal unmarshals quests set
func (q *QuestsSet) Unmarshal(sr *datautils.BitMuncher, act int) (err error) {
	q.Act = act

	switch act {
	case 4: // nolint:gomnd // act 4
		q.Introduced = sr.GetUInt16() == 1

		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			q.Quests[qst].Data = sr.GetBytes(2) // nolint:gomnd // quest data size
			q.Quests[qst].Load()
		}

		q.ActEnd = sr.GetUInt16() == 1
		sr.SkipBits(3 * 8) // nolint:gomnd // 3 unknown bytes
	case 5: // nolint:gomnd // act 5
		q.Introduced = sr.GetUInt16() == 1
		sr.SkipBits(2 * 8) // nolint:gomnd // 2 unknown bytes

		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			q.Quests[qst].Data = sr.GetBytes(2) // nolint:gomnd // quest data size
			q.Quests[qst].Load()
		}

		q.ActEnd = sr.GetUInt16() == 1

		sr.SkipBits(7 * 8) // nolint:gomnd // 7 unknown bytes

	default:
		q.Introduced = sr.GetUInt16() == 1

		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			q.Quests[qst].Data = sr.GetBytes(2) // nolint:gomnd // quest data size
			q.Quests[qst].Load()
		}

		q.ActEnd = sr.GetUInt16() == 1
	}

	return nil
}

// Encode encodes quests set into a byte slice
func (q *QuestsSet) Encode() []byte {
	sw := d2datautils.CreateStreamWriter()

	switch q.Act {
	case 4: // nolint:gomnd // act 4
		if q.Introduced {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		for qst := range q.Quests {
			// TODO: Quest.Encode()
			sw.PushBytes(q.Quests[qst].Data...)
		}

		if q.ActEnd {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		sw.PushBytes(0, 0, 0)
	case 5: // nolint:gomnd // act 5
		if q.Introduced {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		sw.PushUint16(0)

		for qst := range q.Quests {
			// TODO: Quest.Encode()
			sw.PushBytes(q.Quests[qst].Data...)
		}

		if q.ActEnd {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		sw.PushBytes(0, 0, 0, 0, 0, 0, 0)
	default:
		if q.Introduced {
			sw.PushUint16(1)
		} else {
			sw.PushUint16(0)
		}

		for qst := range q.Quests {
			// TODO: Quest.Encode()
			sw.PushBytes(q.Quests[qst].Data...)
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
	Data          []byte
	Completed     bool
	Done          bool // all requirements has been completed (need to get reward)
	Started       bool // when NPC gives you quest
	Body          []bool
	Closed        bool //  you have seen the swirling fire animation that closes a quest icon.
	JustCompleted bool // gets set, when completed in current game
}

// Load loads quest into Quest structure
func (q *Quest) Load() {
	bm := datautils.CreateBitMuncher(q.Data, 0)

	q.Completed = bm.GetBit() == 1
	q.Done = bm.GetBit() == 1
	q.Started = bm.GetBit() == 1

	for i := 0; i < 9; i++ {
		q.Body = append(q.Body, bm.GetBit() == 1)
	}

	q.Closed = bm.GetBit() == 1
	q.JustCompleted = bm.GetBit() == 1
}
