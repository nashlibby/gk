package gk

import "encoding/json"

func TransformMap(value interface{}) (output map[string]interface{}, err error) {
	jValue, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = json.Unmarshal(jValue, &output)
	if err != nil {
		return
	}
	return
}
