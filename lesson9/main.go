package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

func valueToByteSlice(value reflect.Value) []byte {
	switch value.Interface().(type) {
	case string:
		return []byte(value.Interface().(string))
	case int:
		return []byte(strconv.Itoa(value.Interface().(int)))
	case bool:
		return []byte(strconv.FormatBool(value.Interface().(bool)))
	default:
		return []byte{}
	}
}

func parseTag(tag string) (name string, omitempty bool) {
	if len(tag) <= 0 { // tag is empty
		return "", false
	}

	splitted := strings.Split(tag, ",")

	if len(splitted) == 1 { // only first element (name) is present
		return splitted[0], false
	}

	if len(splitted) == 2 && splitted[1] == "omitempty" { // second element is 'omitempty'
		return splitted[0], true
	}

	return splitted[0], false // second element is not 'omitempty'
}

func Serialize(person Person) string {
	v := reflect.ValueOf(person)
	t := v.Type()

	sb := strings.Builder{}

	for idx := 0; idx < v.NumField(); idx++ {
		name, omit := parseTag(t.Field(idx).Tag.Get("properties"))
		valueAsByteSlice := valueToByteSlice(v.Field(idx))

		if omit == true && len(valueAsByteSlice) <= 0 {
			continue
		}

		sb.Write([]byte(name))
		sb.Write([]byte("="))
		sb.Write([]byte(valueToByteSlice(v.Field(idx))))

		if idx < v.NumField()-1 {
			sb.Write([]byte("\n"))
		}
	}

	return sb.String()
}

func main() {
	person := Person{
		Name:    "Dima",
		Address: "Saint-Petersburg",
		Age:     99,
		Married: true,
	}

	fmt.Println(person)
	fmt.Println(Serialize(person))
}
