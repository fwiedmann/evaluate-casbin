package user

type Repository interface {
	Create(m Model) error
	FindById(id string) (Model, error)
}

type Service struct {
	Repo Repository
}

func (s *Service) FindById(id string) (Model, error) {
	return s.Repo.FindById(id)
}
func (s *Service) Create(m Model) error {
	return s.Repo.Create(m)
}
