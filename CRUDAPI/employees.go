package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmpName   string  `json:"empname"`
	EmpSalary float64 `json:"salary"`
	EmpEmail  string  `json:"email"`
}
