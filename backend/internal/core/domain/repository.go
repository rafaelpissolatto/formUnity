package domain

// VolunteerRepository represents a volunteer repository
type VolunteerRepository interface {
	FindAll() ([]Volunteer, error)
	FindByID(id string) (*Volunteer, error)
	FindByEmail(email string) (*Volunteer, error)
	FindByPhone(phone string) (*Volunteer, error)
	Store(volunteer *Volunteer) error
	Update(volunteer *Volunteer) error
	Delete(volunteer *Volunteer) error
}
