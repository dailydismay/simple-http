package viewmodels

import "simple-http/internal/domain"

type UserAgent struct {
	Agent string `json:"useragent"`
}
func (m *UserAgent) Model(d *domain.UserAgent) {
	m.Agent = d.Agent
}

