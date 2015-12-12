package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func main() {
	file, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var obj interface{}
	if err := json.Unmarshal(file, &obj); err != nil {
		log.Fatal(err)
	}

	sum := recursiveSum(obj)
	fmt.Printf("Sum: %d\n", sum)
}

func recursiveSum(obj interface{}) int {
	original := reflect.ValueOf(obj)

	switch original.Kind() {
	case reflect.Slice:
		sum := 0
		for i := 0; i < original.Len(); i++ {
			sum += recursiveSum(original.Index(i).Interface())
		}
		return sum

	case reflect.Map:
		sum := 0
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			reflected := originalValue.Elem()
			fmt.Printf("? %#v ?= %#v %#v\n", reflected.Kind(), reflect.String, reflected)
			if reflected.Kind() == reflect.String && reflected.String() == "red" {
				return 0
			}
			sum += recursiveSum(originalValue.Interface())
		}
		return sum
	case reflect.String:
		return 0
	case reflect.Int:
		return int(original.Int())
	case reflect.Float64:
		return int(original.Float())
	}
	panic(fmt.Sprintf("unexpected: %v\n", original.Kind()))
}
