package storage

import (
	"context"
	"database/sql"
	"fmt"
)

const createTableUserTmpl = `
CREATE TABLE IF NOT EXISTS %s
(
	id      serial  PRIMARY KEY,
	name    varchar NOT NULL,
	surname varchar NOT NULL,
	city_id bigint,
	CONSTRAINT fk_city
      FOREIGN KEY(city_id) 
	  REFERENCES %s(id)
);
`

type StmtUser struct {
	stmtCreateUser *sql.Stmt
	stmtGetUser    *sql.Stmt
}

func (su *StmtUser) prepare(dbConn *sql.DB, userTableName string) (err error) {
	const createUser = `
		INSERT INTO %s (name, surname, city_id)
		VALUES ($1, $2, $3)
		RETURNING id;
`

	if su.stmtCreateUser, err = dbConn.Prepare(fmt.Sprintf(createUser, userTableName)); err != nil {
		return fmt.Errorf("prepare 'create user' stmt: %w", err)
	}

	const getUser = `
		SELECT users.id, users.name, users.surname, users.city_id
		FROM %s AS users
		WHERE users.id = $1
`

	if su.stmtGetUser, err = dbConn.Prepare(fmt.Sprintf(getUser, userTableName)); err != nil {
		return fmt.Errorf("prepare 'get user' stmt: %w", err)
	}

	return nil
}

func (su *StmtUser) Close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = su.stmtCreateUser.Close(); err != nil {
			err = fmt.Errorf("close stmt 'create user': %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = su.stmtGetUser.Close(); err != nil {
			err = fmt.Errorf("close stmt 'get user': %w", err)
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
