// Package pkg for all porject
package pkg

import (
	"reflect"
)

// URL: https://github.com/attapon-th/go-pkgs
// github.com/attapon-th/go-pkgs/task
// github.com/attapon-th/go-pkgs/zlog
// github.com/attapon-th/go-pkgs/zlog/log
// github.com/attapon-th/go-pkgs/validstruct

// IsEmptyValue check variable is empty == true
func IsEmptyValue(i interface{}) bool {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
