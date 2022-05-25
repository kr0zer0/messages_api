package user

type Repository interface {
	Create(user User)
	GetByID(id string) (User, error)
	GetByName(name string) (User, error)
	GetAll() []User
	Update(id string, updatedUser User) error
	Delete(id string) error
}
