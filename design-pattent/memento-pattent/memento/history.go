package memento

type History struct {
	states []EditorState
}

func (h *History) Push(state EditorState) {
	h.states = append(h.states, state)
}

func (h *History) Pop() EditorState {
	lastIndex := len(h.states) - 1
	lastState := h.states[lastIndex]
	h.states = h.states[:lastIndex]
	return lastState
}
