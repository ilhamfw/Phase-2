package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age      int    `validate:"min=18,max=60"`
	Password string `validate:"minLen=6,maxLen=12"`
}

func validateMin(field reflect.Value, tag string, fieldName string) error {
	minValue, err := getTagValue(tag)
	if err != nil {
		return err
	}

	if field.Kind() != reflect.Int {
		return fmt.Errorf("Field '%s' is not an integer and cannot be validated for minimum value", fieldName)
	}

	if int(field.Int()) < int(minValue) {
		return fmt.Errorf("Field '%s' must be greater than or equal to %d", fieldName, minValue)
	}

	return nil
}

func validateMax(field reflect.Value, tag string, fieldName string) error {
	maxValue, err := getTagValue(tag)
	if err != nil {
		return err
	}

	if field.Kind() != reflect.Int {
		return fmt.Errorf("Field '%s' is not an integer and cannot be validated for maximum value", fieldName)
	}

	if int(field.Int()) > int(maxValue) {
		return fmt.Errorf("Field '%s' must be less than or equal to %d", fieldName, maxValue)
	}

	return nil
}

func validateMinLen(field reflect.Value, tag string, fieldName string) error {
	minLen, err := getTagValue(tag)
	if err != nil {
		return err
	}

	if field.Kind() != reflect.String {
		return fmt.Errorf("Field '%s' is not a string and cannot be validated for minimum length", fieldName)
	}

	if len(field.String()) < int(minLen) {
		return fmt.Errorf("Field '%s' must have a minimum length of %d", fieldName, minLen)
	}

	return nil
}

func validateMaxLen(field reflect.Value, tag string, fieldName string) error {
	maxLen, err := getTagValue(tag)
	if err != nil {
		return err
	}

	if field.Kind() != reflect.String {
		return fmt.Errorf("Field '%s' is not a string and cannot be validated for maximum length", fieldName)
	}

	if len(field.String()) > int(maxLen) {
		return fmt.Errorf("Field '%s' must have a maximum length of %d", fieldName, maxLen)
	}

	return nil
}

func getTagValue(tag string) (int, error) {
	return 10, nil // Mengganti dengan nilai yang sesuai dengan tag
}

func main() {
	p1 := Person{
		Age:      25,
		Password: "password123",
	}

	errors1 := validateMin(reflect.ValueOf(p1.Age), "min=18", "Age")
	if errors1 != nil {
		fmt.Println(errors1)
	}

	errors2 := validateMax(reflect.ValueOf(p1.Age), "max=60", "Age")
	if errors2 != nil {
		fmt.Println(errors2)
	}

	errors3 := validateMinLen(reflect.ValueOf(p1.Password), "minLen=6", "Password")
	if errors3 != nil {
		fmt.Println(errors3)
	}

	errors4 := validateMaxLen(reflect.ValueOf(p1.Password), "maxLen=12", "Password")
	if errors4 != nil {
		fmt.Println(errors4)
	}
}
