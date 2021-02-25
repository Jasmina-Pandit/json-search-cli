package helper

import (
	"reflect"
	"strings"
)

func  CaseAndUnderscoreInsenstiveFieldByName(v reflect.Value, name string) reflect.Value {
	name = strings.ToLower(name)
	name= strings.Replace(name,"_","",-1)
	return v.FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == name })
}

