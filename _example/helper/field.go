package helper

import (
	"reflect"
	"strings"
)

func contains(ss map[string]interface{}, s string) bool {
	_, ok := ss[s]

	return ok
}

func merge(m1, m2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for k, v := range m1 {
		result[k] = v
	}

	for k, v := range m2 {
		result[k] = v
	}

	return result
}

func ParseFields(fields string) map[string]interface{} {
	result := make(map[string]interface{})

	if fields == "*" {
		result["*"] = nil
		return result
	}

	for _, field := range strings.Split(fields, ",") {
		parts := strings.SplitN(field, ".", 2)

		if len(parts) == 2 {
			if result[parts[0]] == nil {
				result[parts[0]] = ParseFields(parts[1])
			} else {
				result[parts[0]] = merge(result[parts[0]].(map[string]interface{}), ParseFields(parts[1]))
			}
		} else {
			result[parts[0]] = nil
		}
	}

	return result
}

func FieldToMap(model interface{}, fields map[string]interface{}) map[string]interface{} {
	u := make(map[string]interface{})
	ts, vs := reflect.TypeOf(model), reflect.ValueOf(model)

	for i := 0; i < ts.NumField(); i++ {
		var jsonKey string
		field := ts.Field(i)
		jsonTag := field.Tag.Get("json")

		if jsonTag == "" {
			jsonKey = field.Name
		} else {
			jsonKey = strings.Split(jsonTag, ",")[0]
		}

		if contains(fields, "*") {
			u[jsonKey] = vs.Field(i).Interface()
			continue
		}

		if contains(fields, jsonKey) {
			v := fields[jsonKey]

			if vs.Field(i).Kind() == reflect.Ptr {
				if !vs.Field(i).IsNil() {
					if v == nil {
						u[jsonKey] = vs.Field(i).Elem().Interface()
					} else {
						u[jsonKey] = FieldToMap(vs.Field(i).Elem().Interface(), v.(map[string]interface{}))
					}
				} else {
					u[jsonKey] = nil
				}
			} else if vs.Field(i).Kind() == reflect.Slice {
				var fieldMap []interface{}
				s := reflect.ValueOf(vs.Field(i).Interface())

				for i := 0; i < s.Len(); i++ {
					if v == nil {
						fieldMap = append(fieldMap, s.Index(i).Interface())
					} else {
						fieldMap = append(fieldMap, FieldToMap(s.Index(i).Interface(), v.(map[string]interface{})))
					}
				}

				u[jsonKey] = fieldMap
			} else {
				if v == nil {
					u[jsonKey] = vs.Field(i).Interface()
				} else {
					u[jsonKey] = FieldToMap(vs.Field(i).Interface(), v.(map[string]interface{}))
				}
			}
		}
	}

	return u
}
