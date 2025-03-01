package volunteer

type VolunteerService struct {
	Repository VolunteerRepository
	Notifier   VolunteerNotifier
}

func NewVolunteerService(r VolunteerRepository, n VolunteerNotifier) *VolunteerService {
	return &VolunteerService{
		Repository: r,
		Notifier:   n,
	}
}

func (s *VolunteerService) AddVolunteer(v Volunteer) error {
	return s.Repository.AddVolunteer(v)
}

func (s *VolunteerService) GetVolunteer(id ID) (Volunteer, error) {
	v, err := s.Repository.GetVolunteer(id)
	if err != nil {
		return Volunteer{}, err
	}
	return v, nil
}

func (s *VolunteerService) GetAllVolunteers() ([]Volunteer, error) {
	volunteers, err := s.Repository.GetAllVolunteers()
	if err != nil {
		return []Volunteer{}, err
	}
	return volunteers, nil
}

func (s *VolunteerService) Notify(v Volunteer, message string) error {
	return s.Notifier.Notify(v, message)
}
