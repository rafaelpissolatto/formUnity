package domain

import (
	"time"

	"github.com/google/uuid"
)

// Volunteer represents a volunteer
type Volunteer struct {
	ID               uuid.UUID        `json:"id"`
	FirstName        string           `json:"firstName"`
	LastName         string           `json:"lastName"`
	Email            string           `json:"email"`
	Phone            string           `json:"phone"`
	Nationality      string           `json:"nationality"`
	Occupation       string           `json:"occupation"`
	BirthDate        time.Time        `json:"birthDate"`
	VolunteerDetails VolunteerDetails `json:"volunteerDetails"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
}

// VolunteerDetails represents a volunteer details
type VolunteerDetails struct {
	Availability      string            `json:"availability"`
	AvailabilityTime  string            `json:"availabilityTime"`
	AvailabilityDay   string            `json:"availabilityDay"`
	AvailabilityDate  string            `json:"availabilityDate"`
	AvailabilityNote  string            `json:"availabilityNote"`
	Language          string            `json:"language"`
	LanguageNote      string            `json:"languageNote"`
	Interest          string            `json:"interest"`
	InterestNote      string            `json:"interestNote"`
	Comment           string            `json:"comment"`
	VolunteerDocument VolunteerDocument `json:"volunteerDocument"`
}

// VolunteerDocument represents a volunteer document
type VolunteerDocument struct {
	ID        uuid.UUID `json:"id"`
	Document  Document  `json:"document"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Document represents a document
type Document struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Extension string    `json:"extension"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
