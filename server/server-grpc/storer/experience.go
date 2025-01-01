package storer

import (
	"context"
	"fmt"
	"portfolio/models"
	"portfolio/server-grpc/pb"

	"github.com/jmoiron/sqlx"
)

func batchInsertExpSkills(ctx context.Context, tx *sqlx.Tx, skills []models.ExpSkill, experienceID string) error {
	for _, es := range skills {
		es.ExperienceID = experienceID
		if err := insertExpSkill(ctx, tx, &es); err != nil {
			return err
		}
	}
	return nil
}

func insertExpSkill(ctx context.Context, tx *sqlx.Tx, es *models.ExpSkill) error {
	_, err := tx.NamedExecContext(ctx, `
		INSERT INTO exp_skill (name, percentage, experience_id) 
		VALUES (:name, :percentage, :experience_id)`, es)
	return err
}

func insertExperience(ctx context.Context, tx *sqlx.Tx, e models.Experiences) (string, error) {
	var id string
	query := `
		INSERT INTO experiences (company, title, location, started, ended, descriptions) 
		VALUES (:company, :title, :location, :started, :ended, :descriptions)
		RETURNING id`

	nstmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return "", fmt.Errorf("error preparing statement: %w", err)
	}
	defer nstmt.Close()

	err = nstmt.GetContext(ctx, &id, e)
	if err != nil {
		return "", fmt.Errorf("error inserting education: %w", err)
	}

	return id, nil
}

func updateExperience(ctx context.Context, tx *sqlx.Tx, e *models.Experiences) error {
	_, err := tx.NamedExecContext(ctx, `
		UPDATE experiences 
		SET company=:company, title=:title, location=:location, started=:started, ended=:ended, descriptions=:descriptions, updated_at=now()
		WHERE id=:id`, e)
	return err
}

func (s *PSQLStorer) CreateExperience(ctx context.Context, e *models.Experiences) error {
	err := s.execTx(ctx, func(tx *sqlx.Tx) error {
		id, err := insertExperience(ctx, tx, *e)
		if err != nil {
			return err
		}

		if err := batchInsertExpSkills(ctx, tx, e.Skills, id); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error creating experience: %w", err)
	}

	return nil
}

func (s *PSQLStorer) UpdateExperience(ctx context.Context, e *models.Experiences) (*pb.Empty, error) {
	err := s.execTx(ctx, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, "DELETE FROM exp_skill WHERE experience_id=$1", e.ID); err != nil {
			return err
		}

		if err := batchInsertExpSkills(ctx, tx, e.Skills, e.ID); err != nil {
			return err
		}

		if err := updateExperience(ctx, tx, e); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error updating experience: %w", err)
	}

	return &pb.Empty{}, nil
}
