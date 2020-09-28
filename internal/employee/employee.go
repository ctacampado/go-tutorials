package employee

// Employee struct
type Employee struct {
	id            string
	Name, Address string
	Position      string
	Salary        int
	ManagerID     string
}

// GetID returns the employee's ID
func (e *Employee) GetID() string { return e.id }

// SetID sets the employee's ID
func (e *Employee) SetID(id string) { e.id = id }
