package storer

import (
	"context"
	"fmt"
	"portfolio/models"
	"portfolio/server-grpc/pb"

	"github.com/jmoiron/sqlx"
)

func insertEducation(ctx context.Context, tx *sqlx.Tx, e models.Education) (string, error) {
	var id string
	query := `
		INSERT INTO educations (school, course, started, ended, description) 
		VALUES (:school, :course, :started, :ended, :description) 
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

func updateEducation(ctx context.Context, tx *sqlx.Tx, e *models.Education) error {
	_, err := tx.NamedExecContext(ctx, `
		UPDATE educations 
		SET school=:school, course=:course, started=:started, ended=:ended, description=:description, updated_at=now() 
		WHERE id=:id`, e)
	return err
}

func batchInsertEduSkills(ctx context.Context, tx *sqlx.Tx, skills []models.EduSkill, educationID string) error {
	for _, es := range skills {
		es.EducationID = educationID
		if err := insertEduSkill(ctx, tx, es); err != nil {
			return err
		}
	}
	return nil
}

func insertEduSkill(ctx context.Context, tx *sqlx.Tx, es models.EduSkill) error {
	_, err := tx.NamedExecContext(ctx, `
		INSERT INTO edu_skill (education_id, name, percentage) 
		VALUES (:education_id, :name, :percentage)`, es)
	return err
}

func (s *PSQLStorer) CreateEducation(ctx context.Context, e *models.Education) error {
	err := s.execTx(ctx, func(tx *sqlx.Tx) error {
		id, err := insertEducation(ctx, tx, *e)
		if err != nil {
			return err
		}

		if err := batchInsertEduSkills(ctx, tx, e.Skills, id); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error creating education: %w", err)
	}

	return nil
}

func (s *PSQLStorer) UpdateEducation(ctx context.Context, e *models.Education) (*pb.Empty, error) {
	err := s.execTx(ctx, func(tx *sqlx.Tx) error {
		if _, err := tx.ExecContext(ctx, "DELETE FROM edu_skill WHERE education_id=$1", e.ID); err != nil {
			return err
		}

		if err := batchInsertEduSkills(ctx, tx, e.Skills, e.ID); err != nil {
			return err
		}

		if err := updateEducation(ctx, tx, e); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error updating education: %w", err)
	}

	return &pb.Empty{}, nil
}
