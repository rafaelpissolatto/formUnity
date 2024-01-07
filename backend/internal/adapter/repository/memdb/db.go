package repository

import "github.com/rafaelpissolatto/formUnity/backend/internal/core/domain"

// memdb is a in-memory database adapter for the repository.
type memDB struct {
	m map[string]*domain.Volunteer
}

// Delete implements domain.VolunteerRepository.
func (*memDB) Delete(volunteer *domain.Volunteer) error {
	panic("unimplemented")
}

// FindAll implements domain.VolunteerRepository.
func (*memDB) FindAll() ([]domain.Volunteer, error) {
	panic("unimplemented")
}

// FindByEmail implements domain.VolunteerRepository.
func (*memDB) FindByEmail(email string) (*domain.Volunteer, error) {
	panic("unimplemented")
}

// FindByID implements domain.VolunteerRepository.
func (*memDB) FindByID(id string) (*domain.Volunteer, error) {
	panic("unimplemented")
}

// FindByPhone implements domain.VolunteerRepository.
func (*memDB) FindByPhone(phone string) (*domain.Volunteer, error) {
	panic("unimplemented")
}

// Store implements domain.VolunteerRepository.
func (*memDB) Store(volunteer *domain.Volunteer) error {
	panic("unimplemented")
}

// Update implements domain.VolunteerRepository.
func (*memDB) Update(volunteer *domain.Volunteer) error {
	panic("unimplemented")
}

// NewMemDB creates a new in-memory database.
func NewMemDB() domain.VolunteerRepository {
	return &memDB{
		m: make(map[string]*domain.Volunteer),
	}
}
