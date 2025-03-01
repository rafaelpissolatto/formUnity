package postgres

import (
	"github.com/rafaelpissolatto/formUnity/backend/internal/domain/volunteer"
	"gorm.io/gorm"
)

// PostgresVolunteerRepository is a repository for volunteers backed by PostgreSQL.
type PostgresVolunteerRepository struct {
	DB *gorm.DB
}

// NewPostgresVolunteerRepository creates a new PostgresVolunteerRepository.
func NewPostgresVolunteerRepository(db *gorm.DB) *PostgresVolunteerRepository {
	return &PostgresVolunteerRepository{
		DB: db,
	}
}

// AddVolunteer adds a new volunteer to the database.
func (r *PostgresVolunteerRepository) AddVolunteer(v volunteer.Volunteer) error {
	return r.DB.Create(&v).Error
}

// GetVolunteer retrieves a volunteer by its ID.
func (r *PostgresVolunteerRepository) GetVolunteer(id volunteer.ID) (volunteer.Volunteer, error) {
	var v volunteer.Volunteer
	err := r.DB.First(&v, "id = ?", id).Error
	return v, err
}

// GetAllVolunteers retrieves all volunteers from the database.
func (r *PostgresVolunteerRepository) GetAllVolunteers() ([]volunteer.Volunteer, error) {
	var volunteers []volunteer.Volunteer
	err := r.DB.Find(&volunteers).Error
	return volunteers, err
}

func (r *PostgresVolunteerRepository) VolunteerMigration() {
	r.DB.AutoMigrate(&volunteer.Volunteer{})
}
