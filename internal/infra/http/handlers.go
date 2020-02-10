package http

import (
	"net/http"
	"simple-http/internal/infra/http/viewmodels"
)

func (a *adapter) getUserAgent(w http.ResponseWriter, r *http.Request) {
	agent, err := a.service.GetUserAgent(r)
	if err != nil {
		writeJSONError(w, 400, err)
	}

	var vm viewmodels.UserAgent
	vm.Model(agent)
	writeJSON(w, 200, vm)
}
