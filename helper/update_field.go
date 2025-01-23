package helper

import (
	"reflect"
)

func UpdateFields(target interface{}, source interface{}) {
	targetValue := reflect.ValueOf(target).Elem()
	sourceValue := reflect.ValueOf(source).Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		sourceField := sourceValue.Field(i)
		targetField := targetValue.Field(i)

		if sourceField.Kind() == reflect.Ptr && !sourceField.IsNil() {
			if targetField.CanSet() {
				targetField.Set(sourceField.Elem())
			}
		}
	}
}
