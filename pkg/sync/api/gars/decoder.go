package gars

import (
	"encoding/json"
	"fmt"
)

func decodeResponse(data []byte, target interface{}) error {
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}
