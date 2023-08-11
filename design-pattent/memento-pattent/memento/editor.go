package memento

type Editor struct {
	Content string
}

func (e *Editor) GetContent() string {
	return e.Content
}

func (e *Editor) SetContent(content string) {
	e.Content = content
}

func (e *Editor) CreateState() EditorState {
	return EditorState{content: e.Content}
}

func (e *Editor) Restore(state EditorState) {
	e.Content = state.GetContent()
}
