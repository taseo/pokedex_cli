package api

import (
	"encoding/json"
)

func ParseData[T any](data []byte, obj *T) (T, error) {
	err := json.Unmarshal(data, obj)

	if err != nil {
		return *obj, err
	}

	return *obj, nil
}
