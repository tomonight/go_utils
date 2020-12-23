package utils

import "encoding/json"

func JsonInterface2Struct(data interface{}, dest interface{}) (interface{}, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &dest)
	return dest, err
}
