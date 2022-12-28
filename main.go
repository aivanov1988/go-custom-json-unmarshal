package gocustomjsonunmarshal

import (
	"encoding/json"
	"reflect"
	"strings"
)

// get this function from https://github.com/niski84/interfaceKeySearch.go and update it
// find key in interface (recursively) and return value as interface
func find(obj interface{}, keys []string) (interface{}, bool) {
	key := keys[0]

	//if the argument is not a map, ignore it
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}

	for k, v := range mobj {
		// key match, return value
		if k == key {
			if len(keys) == 1 {
				return v, true
			} else {
				return find(v, keys[1:])
			}
		}
	}
	return nil, false
}

func UnmarshalJSON(jsondata []byte, result interface{}) (interface{}, error) {
	var regulagUnmarshal interface{}
	parseError := json.Unmarshal(jsondata, &regulagUnmarshal)
	if parseError != nil {
		return nil, parseError
	}

	resultReflect := reflect.TypeOf(result)
	tempMap := make(map[string]interface{})
	for i := 0; i < resultReflect.NumField(); i++ {
		fieldEntity := resultReflect.Field(i)
		name := fieldEntity.Name
		tag := fieldEntity.Tag.Get("json")
		jsonPath := strings.Split(tag, ".")
		value, isExist := find(regulagUnmarshal, jsonPath)
		if isExist {
			tempMap[name] = value
		}
	}
	return tempMap, nil
}
