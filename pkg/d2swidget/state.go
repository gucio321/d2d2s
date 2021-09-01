package d2swidget

import (
	"fmt"

	"github.com/AllenDang/giu"
)

type widgetState struct{}

// DIspose implements giu.Disposable
func (s *widgetState) Dispose() {
	// noop
}

func (w *D2SWidget) getState() *widgetState {
	stateID := fmt.Sprintf("%s_state", w.id)

	s := giu.Context.GetState(stateID)
	if s == nil {
		giu.Context.SetState(stateID, w.initState())
		return w.getState()
	}

	return s.(*widgetState)
}

func (w *D2SWidget) initState() *widgetState {
	return &widgetState{}
}
