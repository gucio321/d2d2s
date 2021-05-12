package d2d2s

type mercenary struct {
	Died uint16 // ?
	ID   uint32
	Name uint16
	Type struct {
		Code uint16
	}
	Experience uint32
}

// todo; see: https://user.xmission.com/~trevin/DiabloIIv1.09_Mercenaries.html#code
func (m *mercenary) LoadType(data uint16) {
	m.Type.Code = data
}

func (m *mercenary) EncodeType() (result uint16) {
	return m.Type.Code
}
