package d2d2s

import (
	"errors"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/gucio321/d2d2s/datautils"
)

const (
	numWaypointsBytes              = 81
	waypointHeaderID               = "WS"
	numUnknownWaypointsHeaderBytes = 6
	numActs                        = 5
)

func unknownWaypointsHeaderBytes() [numUnknownWaypointsHeaderBytes]byte {
	return [numUnknownWaypointsHeaderBytes]byte{0x01, 0x00, 0x00, 0x00, 0x50, 0x00}
}

// Waypoints contains state (active = true / inactive = false) of any waypoint in game (difficulty level / act)
type Waypoints map[d2enum.DifficultyType]*[numActs][]bool

func (w *Waypoints) Load(data [numWaypointsBytes]byte) error {
	var err error

	sr := datautils.CreateStreamReader(data[:])

	id, err := sr.ReadBytes(2)
	if err != nil {
		return err
	}

	if string(id) != waypointHeaderID {
		return errors.New("unexpected header identifier")
	}

	// unknownwaypointsheaderbytes
	sr.SkipBytes(numUnknownWaypointsHeaderBytes)

	for i := d2enum.DifficultyNormal; i <= d2enum.DifficultyHell; i++ {
		sr.SkipBytes(2)

		d, err := sr.ReadBytes(5)
		if err != nil {
			return err
		}

		sr.SkipBytes(17)

		bm := datautils.CreateBitMuncher(d, 0)
		_ = bm

		(*w)[i] = &[numActs][]bool{}

		for act := 1; act <= numActs; act++ {
			var l byte
			if act == 4 { // for act 4
				l = 3
			} else {
				l = 9
			}

			(*w)[i][act-1] = make([]bool, l)
			for wp := range (*w)[i][act-1] {
				(*w)[i][act-1][wp] = (bm.GetBit() == 1)
			}
		}
	}

	return nil
}
