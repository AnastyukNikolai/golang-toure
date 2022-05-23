package models

func (m *TodoItem) GetId() int {
	return m.Id
}

func (m *TodoItem) SetId(val int) {
	m.Id = val
}

func (m *TodoItem) GetTitle() string {
	return m.Title
}

func (m *TodoItem) SetTitle(val string) {
	m.Title = val
}

func (m *TodoItem) GetDescription() string {
	return m.Description
}

func (m *TodoItem) SetDescription(val string) {
	m.Description = val
}

func (m *TodoItem) GetStatus() string {
	return m.Status
}

func (m *TodoItem) SetStatus(val string) {
	m.Status = val
}

func (m *TodoItem) GetDone() bool {
	return m.Done
}

func (m *TodoItem) SetDone(val bool) {
	m.Done = val
}
