// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Users struct {
// 	Name      string
// 	Username  string
// 	Password  string
// 	Level     string
// 	CreatedAt string
// 	UpdateAt  string
// }

// func ValidateStruct(s interface{}) error {
// 	t := reflect.TypeOf(s)
// 	for i := 0; i<t.NumField(); i++ {
// 		field := t.Field(i)
// 		if field.Tag.Get("Required") == "true" {
// 			value := reflect.ValueOf(s).Field(i).Interface()
// 			if value == "" {
// 				return fmt.Errorf("%s is required", field.Name)
// 			}
// 		} else {
// 			return fmt.Errorf("%s is required field", field.Name)
// 		}
// 	}
// 	return nil
// }
// func main() {
// 	// var number float64 = 23.42
// 	// var reflectValue = reflect.ValueOf(number)

// 	// fmt.Println("Tipe Variabel :", reflectValue.Type())

// 	// if reflectValue.Kind() == reflect.Float64 {
// 	// 	fmt.Println("NIlai Variabel :", reflectValue.Float())
// 	// }

// 	// if reflectValue.Kind() == reflect.Int {
// 	// 	fmt.Println("Nilai Variabel :", reflectValue.Int())
// 	// }

// 	newUser := Users {
// 		Name: "Rizky",
// 		Username : "rizky01",
// 		Password : "123456",
// 		Level : "admin",
// 		CreatedAt : "2021-01-01",
// 		UpdateAt : "2021-01-01",
// 	}

// 	err := ValidateStruct



// // 	userValue := reflect.ValueOf(newUser)
// // 	fmt.Println(userValue)

// // 	userType := reflect.TypeOf(newUser)
// // 	fmt.Println(userValue)

// // 	userField := userType.NumField()
// // 	fmt.Println(userField)

// // 	fmt.Println("--------------")
// // 	fmt.Println(userType.Field(0))
// // 	fmt.Println(userType.Field(2).Name)
// // }
