package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// User represents a user in the database.
type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

// CreateUser creates a new user in the database.
func CreateUser(c sqlx.Conn, user User) (err error) {
	stmt, err := sqlx.Preparex(DB, "INSERT INTO users (name) VALUES (?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// GetUserById returns a user from the database by ID.
func GetUserById(c sqlx.Conn, id int) (user User, err error) {
	stmt, err := sqlx.Preparex(DB, "SELECT id, name FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Get(&user, id)
	return
}

// GetUserByName returns a user from the database by name.
func UpdateUser(c sqlx.Conn, user User) (err error) {
	stmt, err := sqlx.Preparex(DB, "UPDATE users SET name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.ID)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// DeleteUser deletes a user from the database by ID.
func DeleteUser(c sqlx.Conn, id int) (err error) {
	stmt, err := sqlx.Preparex(DB, "DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	return
}
