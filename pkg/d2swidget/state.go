package d2swidget

import (
	"fmt"

	"github.com/AllenDang/giu"

	"github.com/gucio321/d2d2s/pkg/d2s/d2senums"
)

type widgetState struct {
	difficultyDifficulty d2senums.DifficultyType
	waypointDifficulty   d2senums.DifficultyType
	waypointAct          int32
}

// DIspose implements giu.Disposable
func (s *widgetState) Dispose() {
	// noop
}

func (w *D2SWidget) getState() *widgetState {
	stateID := giu.ID(fmt.Sprintf("%s_state", w.id))

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
