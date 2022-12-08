package repository

import (
	"github.com/daniilmikhaylov2005/town-shop-rest-api/models"
)

func InsertUser(user models.User) (int, error) {
	db := getConnection()
	defer db.Close()

	query := `INSERT INTO users (name, username, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	row := db.QueryRow(query, user.Name, user.Username, user.Password, user.Role)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func FindUserByUsername(username string) (models.User, error) {
	db := getConnection()
	defer db.Close()

	query := `SELECT * FROM users WHERE username=$1`
	row := db.QueryRow(query, username)

	var user models.User

	if err := row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Role); err != nil {
		return models.User{}, err
	}

	return user, nil
}
