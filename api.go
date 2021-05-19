package d2d2s

// SetLevel sets character's level
func (d *D2S) SetLevel(level byte) {
	d.Level = level
	d.Stats.Level = uint32(level)
}
