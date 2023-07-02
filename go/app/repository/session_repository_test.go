package repository

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yakob-abada/go-api/go/app/entity"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFindById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		var time, _ = time.Parse("2006-01-02 15:04:00", "2023-06-26 07:00:00")

		rows := sqlmock.NewRows([]string{"session.id", "class.name", "session.time", "class.duration", "is_full"}).
			AddRow(1, "Spin", time, "30", false)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session`)).WithArgs("1").WillReturnRows(rows)

		var session = entity.Session{
			Id:       1,
			Time:     time,
			Name:     "Spin",
			Duration: 30,
			IsFull:   false,
		}

		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		result, err := sut.FindById("1")

		assert.Exactly(t, &session, result)
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

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session`)).WithArgs("1").WillReturnError(fmt.Errorf("No Rows found"))

		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		_, err = sut.FindById("1")

		assert.EqualError(t, err, "session with id 1: No Rows found")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestFindByAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		var time, _ = time.Parse("2006-01-02 15:04:00", "2023-06-26 07:00:00")

		rows := sqlmock.NewRows([]string{"session.id", "class.name", "session.time", "class.duration", "is_full"}).
			AddRow(1, "Spin", time, "30", false)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id`)).WillReturnRows(rows)

		var sessions = []entity.Session{
			{
				Id:       1,
				Time:     time,
				Name:     "Spin",
				Duration: 30,
				IsFull:   false,
			},
		}

		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		result, err := sut.FindAll()

		assert.Exactly(t, sessions, result)
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

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id`)).WillReturnError(fmt.Errorf("No Rows found"))

		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		_, err = sut.FindAll()

		assert.EqualError(t, err, "No Rows found")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestFindActive(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		var time, _ = time.Parse("2006-01-02 15:04:00", "2023-06-26 07:00:00")

		rows := sqlmock.NewRows([]string{"session.id", "class.name", "session.time", "class.duration", "is_full"}).
			AddRow(1, "Spin", time, "30", false)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id 
		WHERE is_full = 0 AND session.time > now()`)).WillReturnRows(rows)

		var sessions = []entity.Session{
			{
				Id:       1,
				Time:     time,
				Name:     "Spin",
				Duration: 30,
				IsFull:   false,
			},
		}

		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		result, err := sut.FindActive()

		assert.Exactly(t, sessions, result)
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

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id 
		WHERE is_full = 0 AND session.time > now()`)).WillReturnError(fmt.Errorf("No Rows found"))

		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		_, err = sut.FindActive()

		assert.EqualError(t, err, "No Rows found")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestJoin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO session_user").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE session SET is_full = 1 WHERE id ").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// now we execute our method
		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		err = sut.Join(1, 1)

		assert.Nil(t, err)

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("failed to insert", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO session_user").WithArgs(1, 1).WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()

		// now we execute our method
		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		err = sut.Join(1, 1)

		assert.EqualError(t, err, "some error")

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("failed to updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		mockMysqlConnection := &MockMysqlConnection{}
		mockMysqlConnection.On("Connect").Return(db, nil).Once()

		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO session_user").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("UPDATE session SET is_full = 1 WHERE id ").WillReturnError(fmt.Errorf("some error"))
		mock.ExpectRollback()

		// now we execute our method
		sut := NewSessionRepository(mockMysqlConnection, context.Background())

		err = sut.Join(1, 1)

		assert.EqualError(t, err, "some error")

		// we make sure that all expectations were met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
