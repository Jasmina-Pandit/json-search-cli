package helper

import (
	"reflect"
	"strings"
)

func CaseAndUnderscoreInsenstiveFieldByName(v reflect.Value, name string) reflect.Value {
	name = strings.ToLower(name)
	name = strings.Replace(name, "_", "", -1)
	return v.FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == name })
}

func CheckTrimmedValueInArrayString(key string, keysOfTypeArr []string, value reflect.Value, lookupValue string) bool {
	if IsCaseAndUnderscoreInsenKeyInArray(keysOfTypeArr, key) {
		arr := value.Interface().([]string)
		for _, a := range arr {
			if a == lookupValue {
				return true
			}
		}
		return false
	}
	return false
}

func IsCaseAndUnderscoreInsenKeyInArray(keys []string, key string) bool {
	key = strings.ToLower(key)
	key = strings.Replace(key, "_", "", -1)
	for _, k := range keys {
		k = strings.ToLower(k)
		k = strings.Replace(k, "_", "", -1)
		if k == key {
			return true
		}
	}
	return false
}
