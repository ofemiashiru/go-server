package user

import (
	"acme/model"
	"errors"
	"slices"
)

type InMemoryUserRepository struct{}

// mocking ID
var count int = 3

// declare users
var users []model.User

func NewInMemoryRepository() *InMemoryUserRepository {
	InitDB() // Initialize the in-memory database with sample data
	return &InMemoryUserRepository{}
}

func InitDB() {
	// initialise in memory db with data (Create hardcoded slice using struct)
	users = []model.User{
		{ID: 1, Name: "Name 1"},
		{ID: 2, Name: "Name 2"},
		{ID: 3, Name: "Name 3"},
	}
}

func (repo *InMemoryUserRepository) GetUsers() ([]model.User, error) {
	return users, nil
}

func (repo *InMemoryUserRepository) AddUser(user []model.User) (id int, err error) {
	count++
	user[0].ID = count

	users = append(users, user[0])

	return count, nil
}

func (repo *InMemoryUserRepository) GetUser(id int) ([]model.User, error) {
	var user []model.User

	for _, user := range users {
		if user.ID == id {
			return []model.User{user}, nil
		}
	}

	return user, errors.New("user id not found")

}

func (repo *InMemoryUserRepository) DeleteUser(id int) error {

	for index, user := range users {
		if user.ID == id {
			users = slices.Delete(users, index, index+1)
			return nil
		}
	}

	return errors.New("user id not found to delete")

}

func (repo *InMemoryUserRepository) UpdateUser(id int, updatedUser []model.User) ([]model.User, error) {

	for index, user := range users {
		if user.ID == id {
			users[index].Name = updatedUser[0].Name
			return []model.User{user}, nil
		}
	}

	return []model.User{}, errors.New("user id not found to update")

}

func (repo *InMemoryUserRepository) Close() {

}
