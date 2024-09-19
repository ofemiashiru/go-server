package user

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	/*
		ABOVE
		importing with underscore means we only want the
		package for it's side effects not any of its exported identifiers
	*/
	"acme/model"
)

type PostgresUserRepository struct {
	DB *sqlx.DB
}

// New constructor
func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (repo *PostgresUserRepository) AddUser(user []model.User) (id int, err error) {
	err = sqlx.Select(repo.DB, &user, "INSERT INTO users (name) VALUES ($1) RETURNING id;", user[0].Name)
	if err != nil {
		fmt.Println("Error inserting user into the database:", err)
		return 0, errors.New("could not insert user")
	}

	return user[0].ID, nil
}

func (repo *PostgresUserRepository) GetUser(id int) ([]model.User, error) {

	user := []model.User{}

	err := sqlx.Select(repo.DB, &user, "SELECT * FROM users WHERE id=($1);", id)

	if err != nil {
		fmt.Println("Error querying the database:", err)
	}

	return user, nil
}

func (repo *PostgresUserRepository) DeleteUser(id int) error {
	user := []model.User{}

	err := sqlx.Select(repo.DB, &user, "DELETE FROM users WHERE id=($1);", id)

	if err != nil {
		fmt.Println("Error deleting from the database", err)
		return errors.New("database could not be queried")
	}

	return nil
}

func (repo *PostgresUserRepository) GetUsers() ([]model.User, error) {

	users := []model.User{}

	err := sqlx.Select(repo.DB, &users, "SELECT * FROM users;")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.User{}, errors.New("database could not be queried")
	}

	return users, nil
}

func (repo *PostgresUserRepository) UpdateUser(id int, user []model.User) ([]model.User, error) {

	err := sqlx.Select(repo.DB, &user, "UPDATE users SET name =($1) WHERE id=($2)", user[0].Name, id)
	if err != nil {
		fmt.Println("Error updating table:", err)
		return []model.User{}, errors.New("database could not be updated")
	}

	return user, nil
}

func (repo *PostgresUserRepository) Close() {

}
