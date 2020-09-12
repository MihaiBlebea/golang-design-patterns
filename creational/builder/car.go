package builder

// CarProduct _
type CarProduct struct {
	Wheels int
	Seats  int
	Engine int
	Doors  int
	Fuel   string
}

// CarBuilder _
type CarBuilder struct {
	car CarProduct
}

// Drive _
func (cp *CarProduct) Drive() string {
	return "Driving a car"
}

// SetWheels _
func (cb *CarBuilder) SetWheels() BuildProcess {
	cb.car.Wheels = 4

	return cb
}

// SetSeats _
func (cb *CarBuilder) SetSeats() BuildProcess {
	cb.car.Seats = 5

	return cb
}

// SetEngine _
func (cb *CarBuilder) SetEngine() BuildProcess {
	cb.car.Engine = 1500

	return cb
}

// SetDoors _
func (cb *CarBuilder) SetDoors() BuildProcess {
	cb.car.Doors = 4

	return cb
}

// SetFuel _
func (cb *CarBuilder) SetFuel() BuildProcess {
	cb.car.Fuel = "diesel"

	return cb
}

// GetVehicle _
func (cb *CarBuilder) GetVehicle() Vehicle {
	return &cb.car
}
