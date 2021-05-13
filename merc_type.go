package d2d2s

import (
	"errors"

	"github.com/gucio321/d2d2s/datautils"
)

type mercenary struct {
	Died uint16 // ?
	ID   uint32
	Name uint16
	Type struct {
		Code uint16
	}
	Experience uint32
	Items      *Items
}

// todo; see: https://user.xmission.com/~trevin/DiabloIIv1.09_Mercenaries.html#code
func (m *mercenary) LoadType(data uint16) {
	m.Type.Code = data
}

func (m *mercenary) EncodeType() (result uint16) {
	return m.Type.Code
}

func (m *mercenary) LoadMercItems(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2)
	if string(id) != "jf" {
		return errors.New("unexpected merc header")
	}

	// we need to check if there is a merc at all :)
	if m.ID == 0 {
		return nil // just no merc
	}

	m.Items = &Items{}
	numItems, err := m.Items.LoadHeader(sr)
	if err != nil {
		return err
	}
	if err := m.Items.LoadList(sr, numItems); err != nil {
		return err
	}

	return nil
}
