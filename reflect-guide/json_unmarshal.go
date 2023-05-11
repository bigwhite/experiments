package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name      string
	Age       int
	IsStudent bool
}

func main() {
	jsonStr := `{
        "name": "John Doe",
        "age": 30,
        "isStudent": false
    }`

	person := Person{}
	parseJSONToStruct(jsonStr, &person)
	fmt.Printf("%+v\n", person)
}

func parseJSONToStruct(jsonStr string, v interface{}) {
	jsonLines := strings.Split(jsonStr, "\n")
	rv := reflect.ValueOf(v).Elem()

	for _, line := range jsonLines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "{") || strings.HasPrefix(line, "}") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		key := strings.TrimSpace(strings.Trim(parts[0], `"`))
		value := strings.TrimSpace(strings.Trim(parts[1], ","))

		// Find the corresponding field in the struct
		field := rv.FieldByNameFunc(func(fieldName string) bool {
			return strings.EqualFold(fieldName, key)
		})

		if field.IsValid() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(strings.Trim(value, `"`))
			case reflect.Int:
				intValue, _ := strconv.Atoi(value)
				field.SetInt(int64(intValue))
			case reflect.Bool:
				boolValue := strings.ToLower(value) == "true"
				field.SetBool(boolValue)
			}
		}
	}
}
