package repository

import (
	"context"
	"fmt"

	"github.com/yakob-abada/go-api/go/app/entity"
)

type SessionRepository struct {
	dBConnection DatabaseConnection
	ctx          context.Context
}

func NewSessionRepository(dbConnection DatabaseConnection, ctx context.Context) *SessionRepository {
	return &SessionRepository{
		dBConnection: dbConnection,
		ctx:          ctx,
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

func (sr *SessionRepository) FindAll() ([]entity.Session, error) {
	sessions := []entity.Session{}

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
		if err := rows.Scan(&session.Id, &session.Name, &session.Time, &session.Duration, &session.IsFull); err != nil {
			return nil, fmt.Errorf("session: %v", err)
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (sr *SessionRepository) FindActive() ([]entity.Session, error) {
	sessions := []entity.Session{}

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

	return sessions, nil
}

func (sr *SessionRepository) Join(sessionId int8, userId int8) error {

	db, err := sr.dBConnection.Connect()

	if err != nil {
		return err
	}

	tx, err := db.BeginTx(sr.ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(sr.ctx, "INSERT INTO session_user (session_id, user_id) values (?, ?)", sessionId, userId)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(sr.ctx, `
		UPDATE session SET is_full = 1 WHERE id =
		(
			select id from (
				SELECT count(session_id) AS count, session.id, class.max_participant AS max_participant 
				FROM session 
				INNER JOIN class ON class.id = session.class_id
				LEFT JOIN session_user ON session.id = session_id
				WHERE session.id = ?
				GROUP BY session.id
				HAVING count >= max_participant
			) x
		);
	`, sessionId)

	if err != nil {
		return err
	}

	return tx.Commit()
}
