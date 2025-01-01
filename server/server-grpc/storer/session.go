package storer

import (
	"context"
	"portfolio/models"
)

func (s *PSQLStorer) CreateSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	_, err := s.db.NamedExecContext(ctx, "INSERT INTO sessions (id, user_id, refresh_token, email, is_revoked, expires_at) VALUES (:id, :user_id, :refresh_token, :email, :is_revoked, :expires_at)", session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *PSQLStorer) RevokeSession(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, "UPDATE sessions SET is_revoked = true WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PSQLStorer) UpdateSession(ctx context.Context, session *models.Session) (*models.Session, error) {
	_, err := s.db.NamedExecContext(ctx, "UPDATE sessions SET refresh_token=:refresh_token, expires_at=:expires_at WHERE id=:id", session)
	if err != nil {
		return nil, err
	}

	return session, nil
}
