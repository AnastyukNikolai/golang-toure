package models

func (m *TodoItem) GetIdV2() int {
	return m.Id
}

func (m *TodoItem) SetIdV2(val int) {
	m.Id = val
}

func (m *TodoItem) GetTitleV2() string {
	return m.Title
}

func (m *TodoItem) SetTitleV2(val string) {
	m.Title = val
}

func (m *TodoItem) GetDescriptionV2() string {
	return m.Description
}

func (m *TodoItem) SetDescriptionV2(val string) {
	m.Description = val
}

func (m *TodoItem) GetStatusV2() string {
	return m.Status
}

func (m *TodoItem) SetStatusV2(val string) {
	m.Status = val
}

func (m *TodoItem) GetDoneV2() bool {
	return m.Done
}

func (m *TodoItem) SetDoneV2(val bool) {
	m.Done = val
}
