package storage

import (
	"context"
	"database/sql"
	"fmt"
)

const createTableInterestTmpl = `
CREATE TABLE IF NOT EXISTS %s
	(
		id   serial  PRIMARY KEY,
		name varchar UNIQUE NOT NULL
	);
`

type StmtInterest struct {
	stmtCreateInterest  *sql.Stmt
	stmtGetInterest     *sql.Stmt
	stmtGetAllInterests *sql.Stmt
}

func (si *StmtInterest) prepare(dbConn *sql.DB, interestTableName string) (err error) {
	const createInterest = `
		INSERT INTO %s (name)
		VALUES ($1)
		RETURNING id;
`

	if si.stmtCreateInterest, err = dbConn.Prepare(fmt.Sprintf(createInterest, interestTableName)); err != nil {
		return fmt.Errorf("prepare 'create interest' stmt: %w", err)
	}

	const getInterest = `
		SELECT interests.id, interests.name
		FROM %s AS interests
		WHERE interests.id = $1
`

	if si.stmtGetInterest, err = dbConn.Prepare(fmt.Sprintf(getInterest, interestTableName)); err != nil {
		return fmt.Errorf("prepare 'get interest' stmt: %w", err)
	}

	const getAllInterests = `
		SELECT interests.id, interests.name
		FROM %s AS interests
`

	if si.stmtGetAllInterests, err = dbConn.Prepare(fmt.Sprintf(getAllInterests, interestTableName)); err != nil {
		return fmt.Errorf("prepare 'get all interests' stmt: %w", err)
	}

	return nil
}

func (si *StmtInterest) Close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = si.stmtCreateInterest.Close(); err != nil {
			err = fmt.Errorf("close stmt 'create interest': %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = si.stmtGetInterest.Close(); err != nil {
			err = fmt.Errorf("close stmt 'get interest': %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = si.stmtGetAllInterests.Close(); err != nil {
			err = fmt.Errorf("close stmt 'get all interests': %w", err)
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
