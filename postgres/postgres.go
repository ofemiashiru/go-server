package postgres

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

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var DB *sqlx.DB

func InitDB(connectionString string) error {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}

	DB = db

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name VARCHAR(50) NOT NULL);`)

	if err != nil {
		return err
	}

	// attempt to ping database to check for a successful connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error pinging the database: %w", err)
	}

	fmt.Println("Successfully connected to database")
	return nil

}

func AddUser(user model.User) (id int, err error) {
	err = DB.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id;", user.Name).Scan(&id)
	if err != nil {
		fmt.Println("Error inserting user into the database:", err)
		return 0, errors.New("could not insert user")
	}

	return id, nil
}

func GetUser(id int) ([]model.User, error) {

	user := []model.User{}

	err := sqlx.Select(DB, &user, "SELECT * FROM users WHERE id=($1);", id)

	if err != nil {
		fmt.Println("Error querying the database:", err)
	}

	return user, nil
}

func DeleteUser(id int) error {
	user := []model.User{}

	err := sqlx.Select(DB, &user, "DELETE FROM users WHERE id=($1);", id)

	if err != nil {
		fmt.Println("Error deleting from the database", err)
		return errors.New("database could not be queried")
	}

	return nil
}

func GetUsers() ([]model.User, error) {

	users := []model.User{}

	err := sqlx.Select(DB, &users, "SELECT * FROM users;")
	if err != nil {
		fmt.Println("Error querying the database:", err)
		return []model.User{}, errors.New("database could not be queried")
	}

	return users, nil
}

func UpdateUser(id int, user []model.User) error {

	err := sqlx.Select(DB, &user, "UPDATE users SET name =($1) WHERE id=($2)", user[0].Name, id)
	if err != nil {
		fmt.Println("Error updating table:", err)
		return errors.New("database could not be updated")
	}
	return nil
}
