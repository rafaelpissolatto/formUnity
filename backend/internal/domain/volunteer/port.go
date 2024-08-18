package volunteer

// VolunteerRepository provides access to the volunteer storage.
type VolunteerRepository interface {
	// AddVolunteer saves a given volunteer to the repository.
	AddVolunteer(v Volunteer) error
	// GetVolunteer retrieves a volunteer by its ID.
	GetVolunteer(id ID) (Volunteer, error)
	// GetAllVolunteers returns all volunteers stored in the repository.
	GetAllVolunteers() ([]Volunteer, error)
}

// VolunteerNotifier sends notifications about volunteers.
type VolunteerNotifier interface {
	// Notify sends a notification about a volunteer.
	Notify(v Volunteer, message string) error
}
