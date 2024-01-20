package protocol

import (
	"encoding/json"
	"reflect"

	"znci.dev/lamp-v2/utils"
)

func SerializeJSONToBytes(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		utils.Error("Failed to serialize JSON to bytes")
	}

	return bytes
}

func unpackArray(array interface{}) []interface{} {
	value := reflect.ValueOf(array)
	length := value.Len()

	unpacked := make([]interface{}, length)

	for i := 0; i < length; i++ {
		unpacked[i] = value.Index(i).Interface()
	}

	return unpacked
}
