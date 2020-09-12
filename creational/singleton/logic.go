package singleton

// Counter struct
type Counter struct {
	value int
}

var counter *Counter

// Increment adds 1 to the counter value
func (p *Counter) Increment() int {
	p.value++

	return p.value
}

// Value returns the counter value
func (p *Counter) Value() int {
	return p.value
}

// GetInstance returns a pointer to the counter instance
func GetInstance() *Counter {
	if counter == nil {
		counter = new(Counter)
	}

	return counter
}
