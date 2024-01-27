package models

import "interview-test/car-rental/pkg/utilities/standard"

type Car struct {
	Id       string
	Name     string
	Price    int
	Discount int
}

type CarResponse struct {
	standard.StandardReponse
	Data interface{}
}
