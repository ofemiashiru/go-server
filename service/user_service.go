package service

import (
	"acme/model"
	"acme/repository/user"
	"errors"
	"fmt"
)

type UserService struct {
	repository user.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) CreateUser(user []model.User) (id int, err error) {

	id, err = s.repository.AddUser(user)

	if err != nil {
		fmt.Println("Error creating user in DB:", err)
		return 0, errors.New("could not create user")
	}

	return id, nil
}

func (s *UserService) GetUsers() ([]model.User, error) {
	users, err := s.repository.GetUsers()

	if err != nil {
		fmt.Println("Error getting users from DB:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return users, nil

}

func (s *UserService) GetUser(id int) ([]model.User, error) {
	user, err := s.repository.GetUser(id)

	if err != nil {
		fmt.Println("Error getting user from DB:", err)
		return []model.User{}, errors.New("there was an error getting the user from database")
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)

	if err != nil {
		fmt.Println("Error deleting user from DB:", err)
		return errors.New("could not delete user")
	}

	return nil
}

func (s *UserService) UpdateUser(id int, user []model.User) error {
	_, err := s.repository.UpdateUser(id, user)

	if err != nil {
		fmt.Println("Error updating user in DB:", err)
		return errors.New("could not update user")
	}

	return nil

}
