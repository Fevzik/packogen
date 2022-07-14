package generator

import "strings"

func (m *Model) formatName() {
	m.Name = strings.Title(m.Name)
}
