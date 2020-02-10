package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, code int, payload interface{}) error {
	resp, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("cannot marshal json: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(resp); err != nil {
		return fmt.Errorf("cannot write response: %w", err)
	}

	return nil
}

func writeJSONError(w http.ResponseWriter, code int, err error) error {
	resp, err := json.Marshal(map[string]interface{}{
		"code":    code,
		"message": err.Error(),
	})
	if err != nil {
		return fmt.Errorf("cannot marshal json: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(resp); err != nil {
		return fmt.Errorf("cannot write response: %w", err)
	}

	return nil
}

