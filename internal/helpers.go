package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetBodyData(dataStructure interface{}, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, dataStructure); err != nil {
		return err
	}
	return nil
}
