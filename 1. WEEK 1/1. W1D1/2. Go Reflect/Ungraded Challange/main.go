package main

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type Person struct {
	Name     string `validate:"required"`
	Age      int    `validate:"min=18,max=60"`
	Email    string `validate:"email"`
	Password string `validate:"minLen=6,maxLen=20"`
}

func validateStruct(obj interface{}) error {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag.Get("validate")

		switch tag {
		case "required":
			if isZero(field) {
				return fmt.Errorf("Field '%s' is required", v.Type().Field(i).Name)
			}
		case "min":
			if err := validateMin(field, tag, v.Type().Field(i).Name); err != nil {
				return err
			}
		case "max":
			if err := validateMax(field, tag, v.Type().Field(i).Name); err != nil {
				return err
			}
		case "minLen":
			if err := validateMinLen(field, tag, v.Type().Field(i).Name); err != nil {
				return err
			}
		case "maxLen":
			if err := validateMaxLen(field, tag, v.Type().Field(i).Name); err != nil {
				return err
			}
		case "email":
			if err := validateEmail(field, v.Type().Field(i).Name); err != nil {
				return err
			}
		}
	}
	return nil
}

func isZero(field reflect.Value) bool {
	return field.Interface() == reflect.Zero(field.Type()).Interface()
}

func validateMin(field reflect.Value, tag string, fieldName string) error {
	minValue, err := getTagValue(tag)
	if err != nil {
		return err
	}

	if field.Kind() != reflect.Int {
		return errors.New("Min validation is applicable only to integer fields")
	}

	if field.Int() < minValue {
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
		return errors.New("Max validation is applicable only to integer fields")
	}

	if field.Int() > maxValue {
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
		return errors.New("MinLen validation is applicable only to string fields")
	}

	if int64(len(field.String())) < minLen {
		return fmt.Errorf("Field '%s' length must be at least %d characters", fieldName, minLen)
	}

	return nil
}

func validateMaxLen(field reflect.Value, tag string, fieldName string) error {
	maxLen, err := getTagValue(tag)
	if err != nil {
		return err
	}

	if field.Kind() != reflect.String {
		return errors.New("MaxLen validation is applicable only to string fields")
	}

	if int64(len(field.String())) > maxLen {
		return fmt.Errorf("Field '%s' length must not exceed %d characters", fieldName, maxLen)
	}

	return nil
}

func validateEmail(field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.String {
		return errors.New("Email validation is applicable only to string fields")
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
	match, _ := regexp.MatchString(emailPattern, field.String())

	if !match {
		return fmt.Errorf("Field '%s' has an invalid email format", fieldName)
	}

	return nil
}


func getTagValue(tag string) (int64, error) {
	tagValue := reflect.StructTag(tag).Get("validate")
	if tagValue == "" {
		return 0, errors.New("Tag value not provided")
	}

	value, err := strconv.ParseInt(tagValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Invalid tag value: %s", tagValue)
	}

	return value, nil
}



func main() {
	p1 := Person{
		Name:     "John Doe",
		Age:      25,
		Email:    "john.doe@example.com",
		Password: "secret",
	}

	err1 := validateStruct(p1)
	if err1 != nil {
		fmt.Println("Validation error for p1:", err1)
	} else {
		fmt.Println("Validation successful for p1.")
	}

	p2 := Person{
		Name:     "miaw",             // Required field with an empty value
		Age:      16,             // Age is less than the minimum allowed value (18)
		Email:    "invalidemail", // Invalid email format
		Password: "short",        // Password length is less than the minimum allowed (6)
	}

	err2 := validateStruct(p2)
	if err2 != nil {
		fmt.Println("Validation error for p2:", err2)
	} else {
		fmt.Println("Validation successful for p2.")
	}

}
