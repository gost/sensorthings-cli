package cmd

import (
	"encoding/json"
	"net/http"
)

// getJson return json from server
func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
