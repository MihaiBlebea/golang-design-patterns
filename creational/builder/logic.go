package builder

// Director _
type Director struct {
	builder BuildProcess
}

// BuildProcess _
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetEngine() BuildProcess
	SetDoors() BuildProcess
	SetFuel() BuildProcess
	GetVehicle() Vehicle
}

// Vehicle _
type Vehicle interface {
	Drive() string
}

// Construct _
func (d *Director) Construct() Vehicle {
	return d.builder.
		SetEngine().
		SetFuel().
		SetSeats().
		SetDoors().
		SetWheels().
		GetVehicle()
}

// SetBuilder _
func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}
