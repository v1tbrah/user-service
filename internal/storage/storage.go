package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"

	"gitlab.com/pet-pr-social-network/user-service/config"
)

type Storage struct {
	db *sql.DB

	city            city
	interest        interest
	user            user
	userPerInterest userPerInterest

	cfg config.Storage
}

func Init(cfg config.Storage) (*Storage, error) {
	newStorage := &Storage{cfg: cfg}

	db, err := sql.Open("pgx", connString(cfg))
	if err != nil {
		return nil, errors.Wrapf(err, "sql.Open")
	}
	newStorage.db = db

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(maxIdleTime)
	db.SetConnMaxLifetime(maxLifeTime)

	if err = db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "db.Ping")
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, errors.Wrapf(err, "db.Begin")
	}
	defer tx.Rollback()

	if err = tx.Commit(); err != nil {
		return nil, errors.Wrapf(err, "tx.Commit")
	}

	if err = newStorage.city.prepare(db); err != nil {
		return nil, errors.Wrapf(err, "prepare 'city' stmts")
	}

	if err = newStorage.interest.prepare(db); err != nil {
		return nil, errors.Wrapf(err, "prepare 'interest' stmts")
	}

	if err = newStorage.user.prepare(db); err != nil {
		return nil, errors.Wrapf(err, "prepare 'user' stmts")
	}

	if err = newStorage.userPerInterest.prepare(db); err != nil {
		return nil, fmt.Errorf("prepare 'user per interest' stmts: %w", err)
	}

	return newStorage, nil
}

func (s *Storage) Close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = s.city.close(ctx); err != nil {
			err = errors.Wrapf(err, "close stmt 'city'")
			closeEnded <- struct{}{}
			return
		}

		if err = s.interest.close(ctx); err != nil {
			err = errors.Wrapf(err, "close stmt 'interest'")
			closeEnded <- struct{}{}
			return
		}

		if err = s.user.close(ctx); err != nil {
			err = errors.Wrapf(err, "close stmt 'user'")
			closeEnded <- struct{}{}
			return
		}

		if err = s.userPerInterest.close(ctx); err != nil {
			err = errors.Wrapf(err, "close stmt 'user per interest'")
			closeEnded <- struct{}{}
			return
		}

		if err = s.db.Close(); err != nil {
			err = errors.Wrapf(err, "close db conn")
			closeEnded <- struct{}{}
			return
		}

		closeEnded <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-closeEnded:
		return err
	}
}

func connString(cfg config.Storage) string {
	// example: "postgres://username:password@localhost:5432/database_name"
	return "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.UserDBName
}
