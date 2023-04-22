package storage

import (
	"context"
	"database/sql"
	"fmt"
)

const createTableUserPerInterestTmpl = `
CREATE TABLE IF NOT EXISTS %s
(
	user_id     bigint  NOT NULL,
	interest_id bigint NOT NULL,
	PRIMARY KEY (user_id, interest_id)
);
`

type StmtUserPerInterest struct {
	stmtAddInterestToUser  *sql.Stmt
	stmtGetInterestsByUser *sql.Stmt
}

func (supi *StmtUserPerInterest) prepare(dbConn *sql.DB, userPerInterestTableName string) (err error) {
	const addInterestToUser = `
		INSERT INTO %s (user_id, interest_id)
		VALUES ($1, $2)
`

	if supi.stmtAddInterestToUser, err = dbConn.Prepare(fmt.Sprintf(addInterestToUser, userPerInterestTableName)); err != nil {
		return fmt.Errorf("prepare 'add interest to user' stmt: %w", err)
	}

	const getInterestsByUser = `
		SELECT interest_id FROM %s 
		WHERE user_id=$1
`

	if supi.stmtGetInterestsByUser, err = dbConn.Prepare(fmt.Sprintf(getInterestsByUser, userPerInterestTableName)); err != nil {
		return fmt.Errorf("prepare 'get interests by user' stmt: %w", err)
	}

	return nil
}

func (supi *StmtUserPerInterest) Close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = supi.stmtAddInterestToUser.Close(); err != nil {
			err = fmt.Errorf("close stmt 'add interest to user': %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = supi.stmtGetInterestsByUser.Close(); err != nil {
			err = fmt.Errorf("close stmt 'get interests by user': %w", err)
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
