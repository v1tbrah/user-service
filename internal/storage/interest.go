package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gitlab.com/pet-pr-social-network/user-service/internal/model"
)

func (s *Storage) CreateInterest(ctx context.Context, interest model.Interest) (id int64, err error) {
	row := s.stmtInterest.stmtCreateInterest.QueryRowContext(ctx, interest.Name)
	if err = row.Scan(&id); err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok && pgError.Code == pgerrcode.UniqueViolation {
			return -1, ErrInterestAlreadyExists
		}
		return -1, fmt.Errorf("scan created interest id: %w", err)
	}
	if row.Err() != nil {
		return -1, fmt.Errorf("check scan err: %w", row.Err())
	}

	return id, nil
}

func (s *Storage) GetInterest(ctx context.Context, id int64) (interest model.Interest, err error) {
	row := s.stmtInterest.stmtGetInterest.QueryRowContext(ctx, id)
	if err = row.Scan(&interest.ID, &interest.Name); err != nil {
		return interest, fmt.Errorf("scan get interest by id: %w", err)
	}
	if row.Err() != nil {
		return interest, fmt.Errorf("check scan err: %w", row.Err())
	}

	return interest, nil
}

func (s *Storage) GetAllInterests(ctx context.Context) (interests []model.Interest, err error) {
	rows, err := s.stmtInterest.stmtGetAllInterests.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all interests: %w", err)
	}
	defer rows.Close()
	var (
		tempID   int64
		tempName string
	)
	for rows.Next() {
		if err = rows.Scan(&tempID, &tempName); err != nil {
			return nil, fmt.Errorf("scan interest: %w", err)
		}
		interests = append(interests, model.Interest{ID: tempID, Name: tempName})
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("check scan err: %w", rows.Err())
	}

	return interests, nil
}
