package main

import (
	"fmt"
	"reflect"
)

type User struct {
	FirstName string `json:"first_name" binding:"min=8,max=255"`
	LastName  string `json:"last-name" binding:"min=8,max=255"`
	Age       int    `json:"age" binding:"min=8,max=255"`
	Password  string `json:"password" binding:"min=8,max=255"`
}

func reflect1(f interface{}) {
	val := reflect.ValueOf(f).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Type, valueField.Interface(), tag.Get("json"))
		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Type, valueField.Interface(), tag.Get("binding"))
	}
}

func main() {
	f := &User{
		FirstName: "B man",
	}
	reflect1(f)
}
