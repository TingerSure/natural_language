package nl_interface

import (
	"reflect"
)

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	return reflect.ValueOf(i).IsNil()
}
