package repository

import (
	"fmt"

	"github.com/yakob-abada/go-api/go/app/entity"
)

type UserRepository struct {
	dBConnection DatabaseConnection
}

func NewUserRepository(dbConnection DatabaseConnection) *UserRepository {
	return &UserRepository{
		dBConnection: dbConnection,
	}
}

func (sr *UserRepository) FindByUsernameAndPass(username string, password string) (*entity.User, error) {
	var user entity.User

	db, err := sr.dBConnection.Connect()

	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`
		SELECT id, first_name, last_name, username, is_active FROM application.user
		WHERE username = ? AND password = ?
	`, username, password)

	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.IsActive); err != nil {
		return nil, fmt.Errorf("user with username %s: %v", username, err)
	}

	return &user, nil
}
