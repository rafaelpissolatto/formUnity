package service

import (
	"context"

	"github.com/rafaelpissolatto/formUnity/backend/internal/core/domain"
	"github.com/rafaelpissolatto/formUnity/backend/internal/core/port"
)

/**
* VolunteerService implements port.VolunteerService interface
* and provides an acces to the volunteer repository
* and volunteer service
 */

type VolunteerService struct {
	repo      port.VolunteerRepository
	volunteer port.VolunteerService
}

// NewVolunteerService creates a new volunteer service
func NewVolunteerService(repo port.VolunteerRepository, volunteer port.VolunteerService) *VolunteerService {
	return &VolunteerService{
		repo:      repo,
		volunteer: volunteer,
	}
}

// FindAll returns all volunteers
func (vs *VolunteerService) FindAll(ctx context.Context) ([]domain.Volunteer, error) {
	return vs.repo.FindAll()
}

// FindByID returns a volunteer by id
func (vs *VolunteerService) FindByID(ctx context.Context, id string) (*domain.Volunteer, error) {
	return vs.repo.FindByID(id)
}

// FindByEmail returns a volunteer by email
func (vs *VolunteerService) FindByEmail(ctx context.Context, email string) (*domain.Volunteer, error) {
	return vs.repo.FindByEmail(email)
}

// FindByPhone returns a volunteer by phone
func (vs *VolunteerService) FindByPhone(ctx context.Context, phone string) (*domain.Volunteer, error) {
	return vs.repo.FindByPhone(phone)
}

// Store creates a new volunteer
func (vs *VolunteerService) Store(ctx context.Context, volunteer *domain.Volunteer) error {
	return vs.repo.Store(volunteer)
}

// Update updates a volunteer
func (vs *VolunteerService) Update(ctx context.Context, volunteer *domain.Volunteer) error {
	return vs.repo.Update(volunteer)
}

// Delete deletes a volunteer
func (vs *VolunteerService) Delete(ctx context.Context, volunteer *domain.Volunteer) error {
	return vs.repo.Delete(volunteer)
}
