package phi

import (
	"errors"
	"reflect"
)

func GetStructField[T any](value T, fieldName string) (bool, *reflect.StructField, any) {
	if Kind[T]() == reflect.Struct {
		valueOf := Value(value)
		for i := 0; i < valueOf.NumField(); i++ {
			field := valueOf.Type().Field(i)
			if field.Name == fieldName {
				if field.IsExported() {
					return true, &field, valueOf.FieldByName(fieldName).Interface()
				}
				return true, &field, errors.New("cannot access non exported field value")
			}
		}
	}
	return false, nil, nil
}

func HasStructField[T any](fieldName string) bool {
	present, _, _ := GetStructField(Empty[T](), fieldName)
	return present
}

func GetEmbeddedStructField[T any](value T, fieldName string) (bool, *reflect.StructField, any) {
	if present, field, fieldValue := GetStructField[T](value, fieldName); present && field.Anonymous {
		return present, field, fieldValue
	}
	return false, nil, nil
}

func HasEmbeddedStructField[T any](fieldName string) bool {
	present, _, _ := GetEmbeddedStructField(Empty[T](), fieldName)
	return present
}
