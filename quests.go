package d2d2s

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	numQuestsBytes               = 298
	questHeaderUnknownBytesCount = 6
)

/*
type Quests struct {
	unknown1 [questHeaderUnknownBytesCount]byte
	QuestSets
}*/
type Quests map[d2enum.DifficultyType]*[numActs]*QuestsSet

func unknownQuestsHeaderBytes() [questHeaderUnknownBytesCount]byte {
	return [questHeaderUnknownBytesCount]byte{6, 0, 0, 0, 42, 1}
}

func (q *Quests) Unmarshal(data [numQuestsBytes]byte) error {
	sr := datautils.CreateBitMuncher(data[:], 0)

	questHeaderID := sr.GetBytes(4)

	if string(questHeaderID) != expectedQuestHeaderID {
		return errors.New("unexpected quest header")
	}

	_ = sr.GetBytes(questHeaderUnknownBytesCount)

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		(*q)[i] = &[numActs]*QuestsSet{}
		for act := 1; act <= numActs; act++ {
			var l int
			if act == 4 {
				l = 3
			} else {
				l = 6
			}

			(*q)[i][act-1] = &QuestsSet{
				Quests: make([]*Quest, l),
			}

			(*q)[i][act-1].Unmarshal(sr, act)
		}
	}

	return nil
}

type QuestsSet struct {
	Introduced bool
	Quests     []*Quest
	ActEnd     bool // uncertain
}

func (q *QuestsSet) Unmarshal(sr *datautils.BitMuncher, act int) (err error) {
	switch act {
	case 4:
		q.Introduced = sr.GetUInt16() == 1

		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			q.Quests[qst].Data = sr.GetBytes(2)
			q.Quests[qst].Load()
		}

		q.ActEnd = sr.GetUInt16() == 1
		sr.SkipBits(3 * 8)
	case 5:
		q.Introduced = sr.GetUInt16() == 1
		sr.SkipBits(2 * 8)

		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			q.Quests[qst].Data = sr.GetBytes(2)
			q.Quests[qst].Load()
		}

		q.ActEnd = sr.GetUInt16() == 1

		sr.SkipBits(7 * 8)

	default:
		q.Introduced = sr.GetUInt16() == 1

		for qst := range q.Quests {
			q.Quests[qst] = &Quest{}
			q.Quests[qst].Data = sr.GetBytes(2)
			q.Quests[qst].Load()
		}

		q.ActEnd = sr.GetUInt16() == 1
	}

	return nil
}

type Quest struct {
	Data          []byte
	Completed     bool
	Done          bool // all requirements has been completed (need to get reward)
	Started       bool // when NPC gives you quest
	Body          []bool
	Closed        bool //  you have seen the swirling fire animation that closes a quest icon.
	JustCompleted bool // gets set, when completed in current game
}

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
	fmt.Println(bm.Offset())
}
