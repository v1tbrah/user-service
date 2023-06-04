package storage

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type user struct {
	create *sql.Stmt
	get    *sql.Stmt
}

func (su *user) prepare(db *sql.DB) (err error) {
	const createUser = `
		INSERT INTO table_user (name, surname, city_id)
		VALUES ($1, $2, $3)
		RETURNING id;
`

	if su.create, err = db.Prepare(createUser); err != nil {
		return errors.Wrapf(err, "prepare 'create' stmt")
	}

	const getUser = `
		SELECT 
		    id,
		    name,
		    surname,
		    city_id
		FROM table_user
		WHERE id = $1
`

	if su.get, err = db.Prepare(getUser); err != nil {
		return errors.Wrapf(err, "prepare 'get' stmt")
	}

	return nil
}

func (su *user) close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = su.create.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'create'")
			closeEnded <- struct{}{}
			return
		}

		if err = su.get.Close(); err != nil {
			err = errors.Wrapf(err, "close stmt 'get'")
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
