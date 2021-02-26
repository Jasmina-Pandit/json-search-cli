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

func CheckTrimmedValueInArrayString(value string, lookupValue string) bool {
	value = strings.TrimPrefix(value, "[")
	value = strings.TrimSuffix(value, "]")
	splitStrs := strings.Fields(value)
	for _, a := range splitStrs {
		if a == lookupValue {
			return true
		}
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
