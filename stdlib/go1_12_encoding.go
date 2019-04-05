// +build go1.12,!go1.13

package stdlib

// Code generated by 'goexports encoding'. DO NOT EDIT.

import (
	"encoding"
	"reflect"
)

func init() {
	Value["encoding"] = map[string]reflect.Value{
		// function, constant and variable definitions

		// type definitions
		"BinaryMarshaler":   reflect.ValueOf((*encoding.BinaryMarshaler)(nil)),
		"BinaryUnmarshaler": reflect.ValueOf((*encoding.BinaryUnmarshaler)(nil)),
		"TextMarshaler":     reflect.ValueOf((*encoding.TextMarshaler)(nil)),
		"TextUnmarshaler":   reflect.ValueOf((*encoding.TextUnmarshaler)(nil)),
	}
}