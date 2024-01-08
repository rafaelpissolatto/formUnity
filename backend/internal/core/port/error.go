package port

import "errors"

var (
	// ErrNotFound is used when a resource is not found
	ErrNotFound = error.Error(errors.New("not found"))

	// ErrInvalidID is used when an ID is not valid
	ErrInvalidID = error.Error(errors.New("invalid ID"))

	// ErrInvalidEmail is used when an email is not valid
	ErrInvalidEmail = error.Error(errors.New("invalid email"))

	// ErrInvalidPhone is used when a phone is not valid
	ErrInvalidPhone = error.Error(errors.New("invalid phone"))

	// ErrVolunteerNotFound is used when a volunteer is not found
	ErrVolunteerNotFound = error.Error(errors.New("volunteer not found"))

	// ErrVolunteerAlreadyExists is used when a volunteer already exists
	ErrVolunteerAlreadyExists = error.Error(errors.New("volunteer already exists"))

	// ErrInvalidVolunteer is used when a volunteer is not valid
	ErrInvalidVolunteer = error.Error(errors.New("invalid volunteer"))
)
