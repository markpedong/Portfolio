package storer

import (
	"context"
	"fmt"
	"portfolio/models"
	"portfolio/server-grpc/pb"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type PSQLStorer struct {
	db *sqlx.DB
}

func NewPSQLStorer(db *sqlx.DB) *PSQLStorer {
	return &PSQLStorer{
		db: db,
	}
}

func (s *PSQLStorer) ToggleOrDelete(ctx context.Context, body *pb.IdModel) error {
	var query string
	var args []interface{}

	switch body.Type {
	case "toggle":
		query = fmt.Sprintf("UPDATE %s SET status = CASE WHEN status = 1 THEN 0 ELSE 1 END, updated_at = now() WHERE id = $1", body.Model)
		args = append(args, body.Id)

	case "delete":
		query = fmt.Sprintf("UPDATE %s SET deleted_at = now() WHERE id = $1", body.Model)
		args = append(args, body.Id)

	default:
		return fmt.Errorf("invalid operation type: %s", body.Type)
	}

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing %s operation on %s: %w", body.Type, body.Model, err)
	}

	return nil
}

func (s *PSQLStorer) GetByModel(ctx context.Context, resp interface{}, id, model string, child ...string) error {
	column := "id"
	if len(child) > 0 && child[0] != "" {
		column = child[0]
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=$1", model, column)
	if reflect.TypeOf(resp).Elem().Kind() == reflect.Slice {
		if err := s.db.SelectContext(ctx, resp, query, id); err != nil {
			return err
		}
	} else {
		if err := s.db.GetContext(ctx, resp, query, id); err != nil {
			return err
		}
	}

	return nil
}

// first index is no delete column and second index is for the status
func (s *PSQLStorer) GetAllByModel(ctx context.Context, r interface{}, model string, noDelColumn ...bool) error {
	query := fmt.Sprintf("SELECT * FROM %s", model)

	if len(noDelColumn) > 0 && noDelColumn[0] {
		query += " WHERE status = 1 AND deleted_at IS NULL"
	}

	query += " ORDER BY created_at DESC"

	return s.db.SelectContext(ctx, r, query)
}

func (s *PSQLStorer) UpdateRowByModel(ctx context.Context, i interface{}, query string) error {
	_, err := s.db.NamedExecContext(ctx, fmt.Sprintf("%s, updated_at=now() WHERE id=:id", query), i)
	if err != nil {
		return err
	}

	return nil
}

func (s *PSQLStorer) execTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("error rolling back transaction: %w", err)
		}
		return fmt.Errorf("error executing transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (s *PSQLStorer) CreateFile(ctx context.Context, f *models.Files) error {
	err := s.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, "INSERT INTO files (name, file) VALUES (:name, :file)", f)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	return nil
}

func (s *PSQLStorer) CreateRowByModel(ctx context.Context, i interface{}, q string) error {
	err := s.execTx(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, q, i)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
