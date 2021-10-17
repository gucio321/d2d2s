package d2smercenary

import (
	"fmt"

	"github.com/gucio321/d2d2s/internal/datareader"
	"github.com/gucio321/d2d2s/internal/datautils"
	"github.com/gucio321/d2d2s/pkg/common"
	"github.com/gucio321/d2d2s/pkg/d2s/d2sitems"
	"github.com/gucio321/d2d2s/pkg/d2s/d2smercenary/d2smercenarytype"
)

const headerLen = 2

// Mercenary stores mercenary data
type Mercenary struct {
	Died       uint16 // ?
	ID         uint32
	Name       uint16
	Type       *d2smercenarytype.MercenaryType
	Experience uint32
	Items      *d2sitems.Items
}

// New creates a new mercenary structure
func New() *Mercenary {
	result := &Mercenary{
		Type:  &d2smercenarytype.MercenaryType{},
		Items: &d2sitems.Items{},
	}

	return result
}

// LoadHeader loads merc header
func (m *Mercenary) LoadHeader(sr *datareader.Reader) {
	m.Died = sr.GetUint16()
	m.ID = sr.GetUint32()
	m.Name = sr.GetUint16()

	mercType := sr.GetUint16()
	m.Type = d2smercenarytype.Load(mercType)

	m.Experience = sr.GetUint32()
}

// LoadMercItems loads merc items
func (m *Mercenary) LoadMercItems(sr *datareader.Reader) error {
	if id := sr.GetBytes(headerLen); string(id) != "jf" {
		return fmt.Errorf("merc header: %w", common.ErrUnexpectedHeader)
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

// EncodeHeader encodes merc header into a byte slice
func (m *Mercenary) EncodeHeader(sw *datautils.StreamWriter) {
	sw.PushUint16(m.Died)
	sw.PushUint32(m.ID)
	sw.PushUint16(m.Name)
	sw.PushUint16(m.Type.Encode())
	sw.PushUint32(m.Experience)
}

// EncodeItems encodes merc items data back into byte slice
func (m *Mercenary) EncodeItems(sw *datautils.StreamWriter) {
	sw.PushBytes([]byte("jf")...)

	if m.ID == 0 {
		return
	}

	sw.PushBytes(m.Items.Encode()...)
}
