package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func Fill(src interface{}, dst interface{}) error {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)

	if srcType.Kind() != reflect.Struct {
		return errors.New("src必须是struct")
	}
	if dstValue.Kind() != reflect.Ptr {
		return errors.New("dst必须是point")
	}

	for i := 0; i < srcType.NumField(); i++ {
		dstField :=  dstValue.Elem().FieldByName(srcType.Field(i).Name)
		if dstField.CanSet() {
			dstField.Set(srcValue.Field(i))
		}
	}

	return nil
}