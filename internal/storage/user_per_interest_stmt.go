package storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type userPerInterest struct {
	addInterestToUser     *sql.Stmt
	getInterestListByUser *sql.Stmt
}

func (supi *userPerInterest) prepare(db *sql.DB) (err error) {
	const addInterestToUser = `
		INSERT INTO table_user_per_interest (user_id, interest_id)
		VALUES ($1, $2);
`

	if supi.addInterestToUser, err = db.Prepare(addInterestToUser); err != nil {
		return errors.Wrapf(err, "prepare 'add interest to user' stmt")
	}

	const getInterestsByUser = `
		SELECT
			interest_id
		FROM table_user_per_interest
		WHERE user_id=$1;
`

	if supi.getInterestListByUser, err = db.Prepare(getInterestsByUser); err != nil {
		return errors.Wrapf(err, "prepare 'get interest list by user' stmt")
	}

	return nil
}

func (supi *userPerInterest) close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = supi.addInterestToUser.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'add interest to user'")
			closeEnded <- struct{}{}
			return
		}

		if err = supi.getInterestListByUser.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'get interests by user'")
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
