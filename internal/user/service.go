package user

type service struct {
	repository Repository
}

func NewService(repo Repository) service {
	return service{repository: repo}
}

func (s *service) CreateUser(user User) {
	s.repository.Create(user)
}

func (s *service) GetUsers() []User {
	return s.repository.GetAll()
}

func (s *service) GetUserByName(name string) (user User, err error) {
	return s.repository.GetByName(name)
}

func (s *service) EditUser(id string, updatedUser User) error {
	err := s.repository.Update(id, updatedUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
