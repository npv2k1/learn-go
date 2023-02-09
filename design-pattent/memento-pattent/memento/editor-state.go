package memento

type EditorState struct {
	content string
}

func (e *EditorState) GetContent() string {
	return e.content
}

func (e *EditorState) SetContent(content string) {
	e.content = content
}
