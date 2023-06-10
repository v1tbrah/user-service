package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/v1tbrah/user-service/internal/model"
)

func (s *Storage) CreateCity(ctx context.Context, city model.City) (id int64, err error) {
	row := s.city.create.QueryRowContext(ctx, city.Name)
	if err = row.Scan(&id); err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok && pgError.Code == pgerrcode.UniqueViolation {
			return -1, ErrCityAlreadyExists
		}
		return -1, fmt.Errorf("scan created city id: %w", err)
	}
	if row.Err() != nil {
		return -1, fmt.Errorf("check scan err: %w", row.Err())
	}

	return id, nil
}

func (s *Storage) GetCity(ctx context.Context, id int64) (city model.City, err error) {
	row := s.city.get.QueryRowContext(ctx, id)
	if err = row.Scan(&city.ID, &city.Name); err != nil {
		return city, fmt.Errorf("scan get city by id: %w", err)
	}
	if row.Err() != nil {
		return city, fmt.Errorf("check scan err: %w", row.Err())
	}

	return city, nil
}

func (s *Storage) GetAllCities(ctx context.Context) (cities []model.City, err error) {
	rows, err := s.city.getAll.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all cities: %v", err)
	}
	defer rows.Close()
	var (
		tempID   int64
		tempName string
	)
	for rows.Next() {
		if err = rows.Scan(&tempID, &tempName); err != nil {
			return nil, fmt.Errorf("scan city: %w", err)
		}
		cities = append(cities, model.City{ID: tempID, Name: tempName})
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("check scan err: %w", rows.Err())
	}

	return cities, nil
}
