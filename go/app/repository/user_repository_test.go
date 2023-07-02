package repository

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yakob-abada/go-api/go/app/entity"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestUserRepositoryFindByUsername(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "username", "password", "is_active"}).
			AddRow(1, "Yakob", "Abada", "yakob.abada", "test", true)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, first_name, last_name, username, password, is_active FROM user`)).WithArgs("username").WillReturnRows(rows)

		user := entity.User{
			Id:        1,
			FirstName: "Yakob",
			LastName:  "Abada",
			Username:  "yakob.abada",
			Password:  "test",
			IsActive:  true,
		}

		sut := NewUserRepository(mockMysqlConnection)

		result, err := sut.FindByUsername("username")

		assert.Exactly(t, &user, result)
		assert.Nil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("failed", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, first_name, last_name, username, password, is_active FROM user`)).WithArgs("username").WillReturnError(fmt.Errorf("No Rows found"))

		sut := NewUserRepository(mockMysqlConnection)

		_, err = sut.FindByUsername("username")

		assert.EqualError(t, err, "user with username: 'username' has Error: No Rows found")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
