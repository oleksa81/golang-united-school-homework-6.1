package main

import (
	"fmt"
	"golang_united_school_homework"
)

func main() {
	box := golang_united_school_homework.NewBox(3)
	circle := new(golang_united_school_homework.Circle)
	circle.Radius = 4
	box.AddShape(circle)
	box.AddShape(circle)
	box.AddShape(circle)

	shape, err := box.GetByIndex(3)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	fmt.Println(shape)
}
