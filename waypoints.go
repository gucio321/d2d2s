package d2d2s

import (
	"errors"
	"fmt"

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

type Waypoints map[d2enum.DifficultyType][39]Waypoint

func (w *Waypoints) Load(data [numWaypointsBytes]byte) error {
	var err error

	sr := datautils.CreateStreamReader(data[:])

	id, err := sr.ReadBytes(2)
	fmt.Println(string(id))
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
		for range (*w)[i] {
			// fmt.Println(bm.GetBit())
			// w[i][idx].Active = (bm.GetBit() == 1)
		}
	}

	return nil
}

type Waypoint struct {
	Active bool
}
