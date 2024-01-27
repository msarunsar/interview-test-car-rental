package models

import "interview-test/car-rental/pkg/utilities/standard"

type Car struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Discount int    `json:"discount"`
}

type CarResponse struct {
	standard.StandardReponse
	Data interface{} `json:"data"`
}
