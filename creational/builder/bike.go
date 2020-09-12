package builder

// BikeProduct _
type BikeProduct struct {
	Wheels int
	Seats  int
	Engine int
	Doors  int
	Fuel   string
}

// BikeBuilder _
type BikeBuilder struct {
	bike BikeProduct
}

// Drive _
func (b *BikeProduct) Drive() string {
	return "Driving a bike"
}

// SetWheels _
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.bike.Wheels = 2

	return b
}

// SetSeats _
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.bike.Seats = 1

	return b
}

// SetEngine _
func (b *BikeBuilder) SetEngine() BuildProcess {
	b.bike.Engine = 500

	return b
}

// SetDoors _
func (b *BikeBuilder) SetDoors() BuildProcess {
	b.bike.Doors = 0

	return b
}

// SetFuel _
func (b *BikeBuilder) SetFuel() BuildProcess {
	b.bike.Fuel = "petrol"

	return b
}

// GetVehicle _
func (b *BikeBuilder) GetVehicle() Vehicle {
	return &b.bike
}
