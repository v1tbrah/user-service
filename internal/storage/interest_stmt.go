package storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type interest struct {
	create *sql.Stmt
	get    *sql.Stmt
	getAll *sql.Stmt
}

func (si *interest) prepare(db *sql.DB) (err error) {
	const createInterest = `
		INSERT INTO table_interest (name)
		VALUES ($1)
		RETURNING id;
`

	if si.create, err = db.Prepare(createInterest); err != nil {
		return errors.Wrapf(err, "prepare 'create' stmt")
	}

	const getInterest = `
		SELECT 
		    id, 
		    name
		FROM table_interest
		WHERE id = $1;
`

	if si.get, err = db.Prepare(getInterest); err != nil {
		return errors.Wrapf(err, "prepare 'get' stmt")
	}

	const getAllInterests = `
		SELECT 
		    id, 
		    name
		FROM table_interest;
`

	if si.getAll, err = db.Prepare(getAllInterests); err != nil {
		return errors.Wrapf(err, "prepare 'get all' stmt")
	}

	return nil
}

func (si *interest) close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = si.create.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'create'")
			closeEnded <- struct{}{}
			return
		}

		if err = si.get.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'get'")
			closeEnded <- struct{}{}
			return
		}

		if err = si.getAll.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'get all'")
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
