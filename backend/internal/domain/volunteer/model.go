package volunteer

import "time"

// ID is a unique identifier for a volunteer.
type ID string

// Volunteer represents a volunteer in the system.
type Volunteer struct {
	ID        ID
	Name      string
	Email     string
	Phone     string
	Address   string
	CreatedAt time.Time
}
