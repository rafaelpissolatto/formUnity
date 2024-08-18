package httpadapter

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelpissolatto/formUnity/backend/internal/domain/volunteer"
)

// VolunteerHandler is a http handler for volunteer
type VolunteerHandler struct {
	Service *volunteer.VolunteerService
}

// NewVolunteerHandler creates a new VolunteerHandler
func NewVolunteerHandler(s *volunteer.VolunteerService) *VolunteerHandler {
	return &VolunteerHandler{
		Service: s,
	}
}

// NewVolunteerHandler creates a new VolunteerHandler
func (h *VolunteerHandler) AddVolunteer(w http.ResponseWriter, r *http.Request) {
	var v volunteer.Volunteer
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.AddVolunteer(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetVolunteer returns a volunteer by its ID
func (h *VolunteerHandler) GetVolunteer(w http.ResponseWriter, r *http.Request) {
	id := volunteer.ID(r.URL.Query().Get("id"))
	v, err := h.Service.GetVolunteer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetAllVolunteers returns all volunteers
func (h *VolunteerHandler) GetAllVolunteers(w http.ResponseWriter, r *http.Request) {
	volunteers, err := h.Service.GetAllVolunteers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(volunteers); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
