package storage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gitlab.com/pet-pr-social-network/user-service/internal/config"
)

type Storage struct {
	dbConn *sql.DB

	stmtCity            StmtCity
	stmtInterest        StmtInterest
	stmtUser            StmtUser
	stmtUserPerInterest StmtUserPerInterest

	cfg config.StorageConfig
}

func Init(cfg config.StorageConfig) (*Storage, error) {
	newStorage := &Storage{cfg: cfg}

	dbConn, err := sql.Open("pgx", connString(cfg))
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	newStorage.dbConn = dbConn

	dbConn.SetMaxOpenConns(maxOpenConns)
	dbConn.SetMaxIdleConns(maxIdleConns)
	dbConn.SetConnMaxIdleTime(maxIdleTime)
	dbConn.SetConnMaxLifetime(maxLifeTime)

	if err = dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("dbConn.Ping: %w", err)
	}

	tx, err := dbConn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	if _, err = tx.Exec(fmt.Sprintf(createTableCityTmpl, cfg.CityTableName)); err != nil {
		return nil, fmt.Errorf("CreateTableCity: %w", err)
	}

	if _, err = tx.Exec(fmt.Sprintf(createTableInterestTmpl, cfg.InterestTableName)); err != nil {
		return nil, fmt.Errorf("CreateTableInterest: %w", err)
	}

	if _, err = tx.Exec(fmt.Sprintf(createTableUserTmpl, cfg.UserTableName, cfg.CityTableName)); err != nil {
		return nil, fmt.Errorf("CreateTableUser: %w", err)
	}

	if _, err = tx.Exec(fmt.Sprintf(createTableUserPerInterestTmpl, cfg.UserPerInterestTableName)); err != nil {
		return nil, fmt.Errorf("CreateTableUserPerInterest: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	if err = newStorage.stmtCity.prepare(dbConn, cfg.CityTableName); err != nil {
		return nil, fmt.Errorf("prepare 'city' stmts: %w", err)
	}

	if err = newStorage.stmtInterest.prepare(dbConn, cfg.InterestTableName); err != nil {
		return nil, fmt.Errorf("prepare 'interest' stmts: %w", err)
	}

	if err = newStorage.stmtUser.prepare(dbConn, cfg.UserTableName); err != nil {
		return nil, fmt.Errorf("prepare 'user' stmts: %w", err)
	}

	if err = newStorage.stmtUserPerInterest.prepare(dbConn, cfg.UserPerInterestTableName); err != nil {
		return nil, fmt.Errorf("prepare 'user per interest' stmts: %w", err)
	}

	return newStorage, nil
}

func (s *Storage) Close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = s.stmtCity.Close(ctx); err != nil {
			err = fmt.Errorf("close stmt city: %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = s.stmtInterest.Close(ctx); err != nil {
			err = fmt.Errorf("close stmt interest: %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = s.stmtUser.Close(ctx); err != nil {
			err = fmt.Errorf("close stmt user: %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = s.stmtUserPerInterest.Close(ctx); err != nil {
			err = fmt.Errorf("close stmt user per interest: %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = s.dbConn.Close(); err != nil {
			err = fmt.Errorf("close db conn: %w", err)
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

func connString(cfg config.StorageConfig) string {
	// example: "postgres://username:password@localhost:5432/database_name"
	return "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.UserDBName
}
