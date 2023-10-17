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
	Password string `validate:"minLen=6"`
}

func validateStruct(obj interface{}) []string {
	v := reflect.ValueOf(obj)
	t := v.Type()
	var errors []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("validate")
		fieldName := t.Field(i).Name

		switch tag {
		case "required":
			if isZero(field) {
				errors = append(errors, fmt.Sprintf("Field '%s' is required", fieldName))
			}
		case "min":
			if err := validateMin(field, tag, fieldName); err != nil {
				errors = append(errors, err.Error())
			}
		case "max":
			if err := validateMax(field, tag, fieldName); err != nil {
				errors = append(errors, err.Error())
			}
		case "email":
			if err := validateEmail(field, fieldName); err != nil {
				errors = append(errors, err.Error())
			}
		case "minLen":
			if err := validateMinLen(field, tag, fieldName); err != nil {
				errors = append(errors, err.Error())
			}
		case "maxLen":
			if err := validateMaxLen(field, tag, fieldName); err != nil {
				errors = append(errors, err.Error())
			}
		}
	}

	return errors
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int:
		return v.Int() == 0
	default:
		return false
	}
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

func validateEmail(field reflect.Value, fieldName string) error {
	if field.Kind() != reflect.String {
		return fmt.Errorf("Field '%s' is not a string and cannot be validated for email format", fieldName)
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	valid := regexp.MustCompile(emailPattern).MatchString(field.String())
	if !valid {
		return fmt.Errorf("Field '%s' has an invalid email format", fieldName)
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
	tagValue := reflect.StructTag(tag).Get("validate")
	if tagValue == "" {
		return 0, errors.New("Tag value not provided")
	}

	value, err := strconv.Atoi(tagValue)
	if err != nil {
		return 0, fmt.Errorf("Invalid tag value: %s", tagValue)
	}

	return value, nil
}

func main() {
	p1 := Person{
		Name:     "Tanos",
		Age:      1,
		Email:    "asal",
		Password: "qw",
	}

	errors1 := validateStruct(p1)
	if len(errors1) == 0 {
		fmt.Println("Validation successful for p1.")
	} else {
		for _, err := range errors1 {
			fmt.Println(err)
		}
	}

	// errors1 = validateMin(reflect.ValueOf(p1.Age), "min=18", "Age")
	// if errors1 != nil {
	// 	fmt.Println(errors1)
	// }

	// errors2 := validateMax(reflect.ValueOf(p1.Age), "max=60", "Age")
	// if errors2 != nil {
	// 	fmt.Println(errors2)
	// }

	// errors3 := validateMinLen(reflect.ValueOf(p1.Password), "minLen=6", "Password")
	// if errors3 != nil {
	// 	fmt.Println(errors3)
	// }

	// errors4 := validateMaxLen(reflect.ValueOf(p1.Password), "maxLen=12", "Password")
	// if errors4 != nil {
	// 	fmt.Println(errors4)
	// }

	p2 := Person{
		Name:     "",
		Age:      1,
		Email:    "invalidemail@mail.com",
		Password: "short",
	}

	errors2 := validateStruct(p2)
	if len(errors2) == 0 {
		fmt.Println("Validation successful for p2.")
	} else {
		for _, err := range errors2 {
			fmt.Println(err)
		}
	}
}
