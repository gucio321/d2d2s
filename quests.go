package d2d2s

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
)

const (
	numQuestsBytes               = 298
	questHeaderUnknownBytesCount = 6
)

type Quests struct {
	unknown1 [questHeaderUnknownBytesCount]byte
	QuestSets
}

func (q *Quests) Unmarshal(data [numQuestsBytes]byte) error {
	q.QuestSets = make(QuestSets)

	var err error

	sr := d2datautils.CreateStreamReader(data[:])

	questHeaderID, err := sr.ReadBytes(4)
	if err != nil {
		return err
	}

	if string(questHeaderID) != expectedQuestHeaderID {
		return errors.New("unexpected quest header")
	}

	unknown1, err := sr.ReadBytes(questHeaderUnknownBytesCount)
	if err != nil {
		return err
	}

	copy(q.unknown1[:], unknown1[:questHeaderUnknownBytesCount])

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		q.QuestSets[i] = &QuestSet{}
		q.QuestSets[i].Unmarshal(sr)
	}

	return nil
}

type QuestSets map[d2enum.DifficultyType]*QuestSet

type QuestSet struct {
	Introduced bool
}

func (q *QuestSet) Unmarshal(sr *d2datautils.StreamReader) error {
	var err error

	introduced, err := sr.ReadUInt16()
	if err != nil {
		return err
	}

	q.Introduced = introduced == 1

	return nil
}
