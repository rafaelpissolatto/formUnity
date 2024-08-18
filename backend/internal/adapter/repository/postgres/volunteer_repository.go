package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/rafaelpissolatto/formUnity/backend/internal/domain/volunteer"
)

// PostgresVolunteerRepository is a repository for volunteers backed by PostgreSQL.
type PostgresVolunteerRepository struct {
	DB *sql.DB
}

// NewPostgresVolunteerRepository creates a new PostgresVolunteerRepository.
func NewPostgresVolunteerRepository(db *sql.DB) *PostgresVolunteerRepository {
	return &PostgresVolunteerRepository{
		DB: db,
	}
}

// AddVolunteer adds a new volunteer to the database.
func (r *PostgresVolunteerRepository) AddVolunteer(v volunteer.Volunteer) error {
	query := `INSERT INTO volunteers (id, name, email, phone, address) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.DB.Exec(query, uuid.New().String(), v.Name, v.Email, v.Phone, v.Address)
	return err
}

// GetVolunteer retrieves a volunteer by its ID.
func (r *PostgresVolunteerRepository) GetVolunteer(id volunteer.ID) (volunteer.Volunteer, error) {
	query := `SELECT id, name, email, phone, address, created_at FROM volunteers WHERE id = $1`
	row := r.DB.QueryRow(query, id)
	v := &volunteer.Volunteer{}
	err := row.Scan(&v.ID, &v.Name, &v.Email, &v.Phone, &v.Address, &v.CreatedAt)
	return *v, err
}

// GetAllVolunteers retrieves all volunteers from the database.
func (r *PostgresVolunteerRepository) GetAllVolunteers() ([]volunteer.Volunteer, error) {
	query := `SELECT id, name, email, phone, address, created_at FROM volunteers`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	volunteers := []volunteer.Volunteer{}
	for rows.Next() {
		v := volunteer.Volunteer{}
		err := rows.Scan(&v.ID, &v.Name, &v.Email, &v.Phone, &v.Address, &v.CreatedAt)
		if err != nil {
			return nil, err
		}
		volunteers = append(volunteers, v)
	}
	return volunteers, nil
}
