package storage

import (
	"context"
	"database/sql"
	"fmt"
)

const createTableCityTmpl = `
CREATE TABLE IF NOT EXISTS %s
	(
		id   serial  PRIMARY KEY,
		name varchar UNIQUE NOT NULL
	);
`

type StmtCity struct {
	stmtCreateCity   *sql.Stmt
	stmtGetCity      *sql.Stmt
	stmtGetAllCities *sql.Stmt
}

func (sc *StmtCity) prepare(dbConn *sql.DB, cityTableName string) (err error) {
	const createCity = `
		INSERT INTO %s (name)
		VALUES ($1)
		RETURNING id;
`

	if sc.stmtCreateCity, err = dbConn.Prepare(fmt.Sprintf(createCity, cityTableName)); err != nil {
		return fmt.Errorf("prepare 'create city' stmt: %w", err)
	}

	const getCity = `
		SELECT
			cities.id,
			cities.name
		FROM %s AS cities
		WHERE cities.id = $1
`

	if sc.stmtGetCity, err = dbConn.Prepare(fmt.Sprintf(getCity, cityTableName)); err != nil {
		return fmt.Errorf("prepare 'get city' stmt: %w", err)
	}

	const getAllCities = `
		SELECT cities.id, cities.name
		FROM %s AS cities
`

	if sc.stmtGetAllCities, err = dbConn.Prepare(fmt.Sprintf(getAllCities, cityTableName)); err != nil {
		return fmt.Errorf("prepare 'get all cities' stmt: %w", err)
	}

	return nil
}

func (sc *StmtCity) Close(ctx context.Context) (err error) {
	closeEnded := make(chan struct{})

	go func() {
		if err = sc.stmtCreateCity.Close(); err != nil {
			err = fmt.Errorf("close stmt 'create city': %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = sc.stmtGetCity.Close(); err != nil {
			err = fmt.Errorf("close stmt 'get city': %w", err)
			closeEnded <- struct{}{}
			return
		}

		if err = sc.stmtGetAllCities.Close(); err != nil {
			err = fmt.Errorf("close stmt 'get all cities': %w", err)
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
