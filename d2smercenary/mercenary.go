package d2smercenary

import (
	"errors"
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2datautils"

	"github.com/gucio321/d2d2s/d2sitems"
	"github.com/gucio321/d2d2s/datautils"
)

// Mercenary stores mercenary data
type Mercenary struct {
	Died uint16 // ?
	ID   uint32
	Name uint16
	Type struct {
		Code uint16
	}
	Experience uint32
	Items      *d2sitems.Items
}

// New creates a new mercenary structure
func New() *Mercenary {
	result := &Mercenary{
		Items: &d2sitems.Items{},
	}

	return result
}

// LoadHeader loads merc header
func (m *Mercenary) LoadHeader(sr *datautils.BitMuncher) {
	m.Died = sr.GetUInt16()
	m.ID = sr.GetUInt32()
	m.Name = sr.GetUInt16()

	mercType := sr.GetUInt16()
	m.LoadType(mercType)

	m.Experience = sr.GetUInt32()
}

// LoadType load merc type: TODO; see: https://user.xmission.com/~trevin/DiabloIIv1.09_Mercenaries.html#code
func (m *Mercenary) LoadType(data uint16) {
	m.Type.Code = data
}

// EncodeType encodes merc type (TODO)
func (m *Mercenary) EncodeType() (result uint16) {
	return m.Type.Code
}

// LoadMercItems loads merc items
func (m *Mercenary) LoadMercItems(sr *datautils.BitMuncher) error {
	id := sr.GetBytes(2) // nolint:gomnd // header size
	if string(id) != "jf" {
		return errors.New("unexpected merc header")
	}

	// we need to check if there is a merc at all :)
	if m.ID == 0 {
		return nil // just no merc
	}

	numItems, err := m.Items.LoadHeader(sr)
	if err != nil {
		return fmt.Errorf("error loading mercenary items: %w", err)
	}

	if err := m.Items.LoadList(sr, numItems); err != nil {
		return fmt.Errorf("error reading mercenary items list: %w", err)
	}

	return nil
}

// EncodeItems encodes merc items data back into byte slice
func (m *Mercenary) EncodeItems(sw *d2datautils.StreamWriter) {
	sw.PushBytes([]byte("jf")...)

	if m.ID == 0 {
		return
	}

	sw.PushBytes(m.Items.Encode()...)
}
