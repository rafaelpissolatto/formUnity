package port

import "github.com/rafaelpissolatto/formUnity/backend/internal/core/domain"

type VolunteerRepository interface {
	FindAll() ([]domain.Volunteer, error)
	FindByID(id string) (*domain.Volunteer, error)
	FindByEmail(email string) (*domain.Volunteer, error)
	FindByPhone(phone string) (*domain.Volunteer, error)
	Store(volunteer *domain.Volunteer) error
	Update(volunteer *domain.Volunteer) error
	Delete(volunteer *domain.Volunteer) error

	// implement more methods here
}

type VolunteerService interface {
	FindAll() ([]domain.Volunteer, error)
	FindByID(id string) (*domain.Volunteer, error)
	FindByEmail(email string) (*domain.Volunteer, error)
	FindByPhone(phone string) (*domain.Volunteer, error)
	Store(volunteer *domain.Volunteer) error
	Update(volunteer *domain.Volunteer) error
	Delete(volunteer *domain.Volunteer) error

	// implement more methods here
}
