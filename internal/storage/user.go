package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
	"gitlab.com/pet-pr-social-network/user-service/internal/model"
)

func (s *Storage) CreateUser(ctx context.Context, user model.User) (id int64, err error) {
	tx, err := s.dbConn.Begin()
	if err != nil {
		return -1, err
	}
	defer func() {
		if errRollback := tx.Rollback(); errRollback != nil && errRollback != sql.ErrTxDone {
			log.Error().Err(errRollback).Msg("storage.CreateUser tx.Rollback")
		}
	}()

	row := tx.Stmt(s.stmtUser.stmtCreateUser).QueryRowContext(ctx, user.Name, user.Surname, user.CityID)
	if err = row.Scan(&id); err != nil {
		return -1, fmt.Errorf("scan created user id: %w", err)
	}
	if row.Err() != nil {
		return -1, fmt.Errorf("check scan err: %w", row.Err())
	}

	for _, interestID := range user.InterestsID {
		if _, err = tx.Stmt(s.stmtUserPerInterest.stmtAddInterestToUser).ExecContext(ctx, id, interestID); err != nil {
			return -1, fmt.Errorf("add user interest: %w", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return id, nil
}

func (s *Storage) GetUser(ctx context.Context, id int64) (user model.User, err error) {
	row := s.stmtUser.stmtGetUser.QueryRowContext(ctx, id)
	if err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.CityID); err != nil {
		return user, fmt.Errorf("scan get user by id: %w", err)
	}
	if row.Err() != nil {
		return user, fmt.Errorf("check scan err: %w", row.Err())
	}

	rows, err := s.stmtUserPerInterest.stmtGetInterestsByUser.QueryContext(ctx, id)
	if err != nil {
		return user, fmt.Errorf("get interests by user: %w", err)
	}
	defer rows.Close()

	var tempInterestID int64
	for rows.Next() {
		if err = rows.Scan(&tempInterestID); err != nil {
			return user, fmt.Errorf("scan interest id: %w", err)
		}
		user.InterestsID = append(user.InterestsID, tempInterestID)
	}
	if rows.Err() != nil {
		return user, fmt.Errorf("check scan err: %w", rows.Err())
	}

	return user, nil
}
