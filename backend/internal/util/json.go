package util

import "encoding/json"

func MarshalJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func UnmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func MarshalJSONToString(v interface{}) (string, error) {
	data, err := MarshalJSON(v)
	return string(data), err
}
