package main

import (
	"fmt"

	"github.com/MihaiBlebea/go-patterns/creational/builder"
)

func main() {
	dir := builder.Director{}

	dir.SetBuilder(&builder.BikeBuilder{})

	bike := dir.Construct()

	fmt.Println(bike.Drive())
}
