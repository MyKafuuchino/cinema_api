package helper

import (
	"reflect"
	"strings"
)

func MapUpdateField(existingProduct interface{}, updateData interface{}) {
	existingVal := reflect.ValueOf(existingProduct).Elem()
	updateVal := reflect.ValueOf(updateData).Elem()

	for i := 0; i < updateVal.NumField(); i++ {
		field := updateVal.Field(i)
		fieldName := updateVal.Type().Field(i).Name

		// Optional: Add json tag support
		jsonTag := updateVal.Type().Field(i).Tag.Get("json")
		if jsonTag != "" {
			fieldName = strings.Split(jsonTag, ",")[0]
		}

		if !field.IsNil() {
			existingField := existingVal.FieldByName(fieldName)
			if existingField.IsValid() && existingField.CanSet() {
				existingField.Set(field.Elem())
			}
		}
	}
}
