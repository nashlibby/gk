package gk

import "encoding/json"

func Scan(origin, dest any) error {
	j, _ := json.Marshal(origin)
	return json.Unmarshal(j, dest)
}
