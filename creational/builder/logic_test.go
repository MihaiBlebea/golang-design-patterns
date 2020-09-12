package builder_test

import (
	"testing"

	"github.com/MihaiBlebea/go-patterns/creational/builder"
)

func TestBuildCar(t *testing.T) {
	dir := builder.Director{}
	dir.SetBuilder(&builder.CarBuilder{})

	var car *builder.CarProduct = dir.Construct().(*builder.CarProduct)

	if car.Wheels != 4 {
		t.Errorf("Expected %d wheels, got %d", 4, car.Wheels)
	}

	if car.Doors != 4 {
		t.Errorf("Expected %d doors, got %d", 4, car.Doors)
	}

	if car.Engine != 1500 {
		t.Errorf("Expected %d cc engine, got %d", 1500, car.Engine)
	}

	if car.Fuel != "diesel" {
		t.Errorf("Expected %s fuel, got %s", "diesel", car.Fuel)
	}

	if car.Seats != 5 {
		t.Errorf("Expected %d seats, got %d", 5, car.Seats)
	}
}

func TestBuildBike(t *testing.T) {
	dir := builder.Director{}
	dir.SetBuilder(&builder.BikeBuilder{})

	var bike *builder.BikeProduct = dir.Construct().(*builder.BikeProduct)

	if bike.Wheels != 2 {
		t.Errorf("Expected %d wheels, got %d", 2, bike.Wheels)
	}

	if bike.Doors != 0 {
		t.Errorf("Expected %d doors, got %d", 0, bike.Doors)
	}

	if bike.Engine != 500 {
		t.Errorf("Expected %d cc engine, got %d", 500, bike.Engine)
	}

	if bike.Fuel != "petrol" {
		t.Errorf("Expected %s fuel, got %s", "petrol", bike.Fuel)
	}

	if bike.Seats != 1 {
		t.Errorf("Expected %d seats, got %d", 1, bike.Seats)
	}
}
