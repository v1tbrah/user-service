package storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type city struct {
	create *sql.Stmt
	get    *sql.Stmt
	getAll *sql.Stmt
}

func (sc *city) prepare(db *sql.DB) (err error) {
	const createCity = `
		INSERT INTO table_city (name)
		VALUES ($1)
		RETURNING id;
`

	if sc.create, err = db.Prepare(createCity); err != nil {
		return errors.Wrapf(err, "prepare 'create' stmt")
	}

	const getCity = `
		SELECT
			id,
			name
		FROM table_city
		WHERE id = $1;
`

	if sc.get, err = db.Prepare(getCity); err != nil {
		return errors.Wrapf(err, "prepare 'get' stmt")
	}

	const getAllCities = `
		SELECT
			id,
			name
		FROM table_city;
`

	if sc.getAll, err = db.Prepare(getAllCities); err != nil {
		return errors.Wrapf(err, "prepare 'get all' stmt")
	}

	return nil
}

func (sc *city) close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = sc.create.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'create'")
			closeEnded <- struct{}{}
			return
		}

		if err = sc.get.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'get'")
			closeEnded <- struct{}{}
			return
		}

		if err = sc.getAll.Close(); err != nil {
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
