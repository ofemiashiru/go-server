package service

import (
	"acme/model"
	"acme/postgres"
	"errors"
	"fmt"
)

func CreateUser(user model.User) (id int, err error) {
	id, err = postgres.AddUser(user) // << was db.AddUser(user)

	if err != nil {
		fmt.Println("Error creating user in DB:", err)
		return 0, errors.New("could not create user")
	}

	return id, nil
}

func GetUsers() ([]model.User, error) {
	users, err := postgres.GetUsers()

	if err != nil {
		fmt.Println("Error getting users from DB:", err)
		return nil, errors.New("there was an error getting the users from the database")
	}

	return users, nil

}

func GetUser(id int) ([]model.User, error) {
	user, err := postgres.GetUser(id)

	if err != nil {
		fmt.Println("Error getting user from DB:", err)
		return nil, errors.New("there was an error getting the user from database")
	}

	return user, nil
}

func DeleteUser(id int) error {
	err := postgres.DeleteUser(id)

	if err != nil {
		fmt.Println("Error deleting user from DB:", err)
		return errors.New("could not delete user")
	}

	return nil
}

func UpdateUser(id int, user []model.User) error {
	err := postgres.UpdateUser(id, user)

	if err != nil {
		fmt.Println("Error updating user in DB:", err)
		return errors.New("could not update user")
	}

	return nil

}
