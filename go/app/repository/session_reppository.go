package repository

import (
	"fmt"

	"github.com/yakob-abada/go-api/go/app/entity"
)

type ISessionRepository interface {
	FindById(id string) (*entity.Session, error)
	FindAll() (*[]entity.Session, error)
	FindActive() (*[]entity.Session, error)
	Join(sessionId int8, userId int8) error
	SetSessionIsFullSatistfaction(sessionId int8) error
}

type SessionRepository struct {
	dBConnection DatabaseConnection
}

func NewSessionRepository(dbConnection DatabaseConnection) *SessionRepository {
	return &SessionRepository{
		dBConnection: dbConnection,
	}
}

func (sr *SessionRepository) FindById(id string) (*entity.Session, error) {
	var session entity.Session

	db, err := sr.dBConnection.Connect()

	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`
		SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id 
		WHERE session.id = ? 
	`, id)

	if err := row.Scan(&session.Id, &session.Name, &session.Time, &session.Duration, &session.IsFull); err != nil {
		return nil, fmt.Errorf("session with id %s: %v", id, err)
	}

	return &session, nil
}

func (sr *SessionRepository) FindAll() (*[]entity.Session, error) {
	var sessions []entity.Session = []entity.Session{}

	db, err := sr.dBConnection.Connect()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var session entity.Session
		if err := rows.Scan(&session.Id); err != nil {
			return nil, fmt.Errorf("session: %v", err)
		}

		sessions = append(sessions, session)
	}

	return &sessions, nil
}

func (sr *SessionRepository) FindActive() (*[]entity.Session, error) {
	var sessions []entity.Session = []entity.Session{}

	db, err := sr.dBConnection.Connect()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`
		SELECT session.id, class.name, session.time, class.duration, is_full 
		FROM session 
		INNER JOIN class ON session.class_id = class.id 
		WHERE is_full = 0 AND session.time > now()
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var session entity.Session
		if err := rows.Scan(&session.Id, &session.Name, &session.Time, &session.Duration, &session.IsFull); err != nil {
			return nil, fmt.Errorf("session: %v", err)
		}

		sessions = append(sessions, session)
	}

	return &sessions, nil
}

func (sr *SessionRepository) Join(sessionId int8, userId int8) error {
	db, err := sr.dBConnection.Connect()

	if err != nil {
		return err
	}

	_, err = db.Query("INSERT INTO session_user (session_id, user_id) values (?, ?)", sessionId, userId)

	return err
}

func (sr *SessionRepository) SetSessionIsFullSatistfaction(sessionId int8) error {
	db, err := sr.dBConnection.Connect()

	if err != nil {
		return err
	}

	_, err = db.Query(`
		UPDATE session SET is_full = 1 WHERE id =
		(
			select id from (
				SELECT count(session_id) AS count, session.id, class.max_participant AS max_participant 
				FROM session 
				INNER JOIN class ON class.id = session.class_id
				LEFT JOIN session_user ON session.id = session_id
				WHERE session.id = ?
				HAVING count >= max_participant
			) x
		);
	`, sessionId)

	return err
}
