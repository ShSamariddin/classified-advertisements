package utils

import (
	"reflect"
)

func GetStructFields(data interface{}, exclude map[string]bool) ([]interface{}, []string) {
	v := reflect.ValueOf(data)
	fieldNum := v.NumField()
	values := make([]interface{}, 0)
	names := make([]string, 0)
	for i := 0; i < fieldNum; i++ {
		name := v.Type().Field(i).Tag.Get("db")
		value := v.Field(i).Interface()
		if exclude[v.Type().Field(i).Name] {
			continue
		}
		values = append(values, value)
		names = append(names, name)
	}
	return values, names
}
