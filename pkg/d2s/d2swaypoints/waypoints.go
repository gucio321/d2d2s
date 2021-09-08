package d2swaypoints

import (
	"errors"

	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

const (
	// NumWaypointsBytes is a total number of waypoints data bytes
	NumWaypointsBytes = 81

	waypointHeaderID               = "WS"
	numUnknownWaypointsHeaderBytes = 6
	defaultWPCound                 = 9
	act4WPCount                    = 3
	unknownWaypointsBytesCount     = 17
	waypointDataBytesCount         = 5
)

func unknownWaypointsHeaderBytes() [numUnknownWaypointsHeaderBytes]byte {
	return [numUnknownWaypointsHeaderBytes]byte{0x01, 0x00, 0x00, 0x00, 0x50, 0x00}
}

// Waypoints contains state (active = true / inactive = false) of any waypoint in game (difficulty level / act)
type Waypoints map[d2senums.DifficultyType]*[d2senums.NumActs][]bool

// New creates a new waypoints structure
func New() *Waypoints {
	result := &Waypoints{}
	*result = make(Waypoints)

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		(*result)[i] = &[d2senums.NumActs][]bool{}

		for act := 1; act <= d2senums.NumActs; act++ {
			var l byte

			const act4 = 4

			if act == act4 {
				l = act4WPCount
			} else {
				l = defaultWPCound
			}

			(*result)[i][act-1] = make([]bool, l)
		}
	}

	return result
}

// Load loads waypoints data
func (w *Waypoints) Load(data *[NumWaypointsBytes]byte) error {
	sr := datautils.CreateBitMuncher((*data)[:], 0)

	id := sr.GetBytes(len(waypointHeaderID))
	if string(id) != waypointHeaderID {
		return errors.New("unexpected header identifier")
	}

	// unknownwaypointsheaderbytes
	sr.SkipBytes(numUnknownWaypointsHeaderBytes)

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		const unknownWPBytesCount = 2

		sr.SkipBytes(unknownWPBytesCount)

		d := sr.GetBytes(waypointDataBytesCount)

		sr.SkipBytes(unknownWaypointsBytesCount)

		bm := datautils.CreateBitMuncher(d, 0)

		for act := 1; act <= d2senums.NumActs; act++ {
			for wp := range (*w)[i][act-1] {
				(*w)[i][act-1][wp] = (bm.GetBit() == 1)
			}
		}
	}

	return nil
}

// Encode encodes waypoints data back into byte array
func (w *Waypoints) Encode() (result [NumWaypointsBytes]byte) {
	sw := datautils.CreateStreamWriter()

	sw.PushBytes([]byte(waypointHeaderID)...)

	unknown := unknownWaypointsHeaderBytes()
	sw.PushBytes(unknown[:]...)

	for i := d2senums.DifficultyNormal; i <= d2senums.DifficultyHell; i++ {
		// https://user.xmission.com/~trevin/DiabloIIv1.09_File_Format.shtml
		//	 	unknown; I always see the values { 2, 1 } here.
		sw.PushBytes([]byte{2, 1}...)

		data := datautils.CreateStreamWriter()
		// here 5 bytes
		for act := 0; act < d2senums.NumActs; act++ {
			for _, wp := range (*w)[i][act] {
				data.PushBit(wp)
			}
		}

		data.PushBit(false)

		d := data.GetBytes()
		sw.PushBytes(d...)

		unknownData := [unknownWaypointsBytesCount]byte{}
		sw.PushBytes(unknownData[:]...)
	}

	sw.PushBytes(1) // uncertain
	data := sw.GetBytes()
	copy(result[:], data[:NumWaypointsBytes])

	return result
}
