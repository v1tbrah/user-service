//go:build with_db

package storage

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/v1tbrah/user-service/internal/model"
)

func TestStorage_CreateCity(t *testing.T) {
	s := tHelperInitEmptyDB(t)

	testCity := model.City{Name: "testCityName"}
	idNewCity, err := s.CreateCity(context.Background(), testCity)
	if err != nil {
		t.Fatalf("s.CreateCity: %v", err)
	}

	row := s.db.QueryRow(fmt.Sprintf("SELECT name FROM table_city WHERE id=%d", idNewCity))
	if err = row.Scan(&testCity.Name); err != nil {
		t.Fatalf("scan get new city name: %v", err)
	}
	if row.Err() != nil {
		t.Fatalf("check scan err: %v", row.Err())
	}

	if testCity.Name != "testCityName" {
		t.Fatalf("new city name: got: %s, expected: %s", testCity.Name, "testCityName")
	}
}

func TestStorage_GetCity(t *testing.T) {
	s := tHelperInitEmptyDB(t)

	testCity := model.City{Name: "testCityName"}
	row := s.db.QueryRow(fmt.Sprintf("INSERT INTO table_city (name) VALUES('%s') RETURNING id", testCity.Name))
	err := row.Scan(&testCity.ID)
	if err != nil {
		t.Fatalf("scan new city id: %v", err)
	}
	if row.Err() != nil {
		t.Fatalf("check scan err: %v", row.Err())
	}

	testCity, err = s.GetCity(context.Background(), testCity.ID)
	if err != nil {
		t.Fatalf("get city: %v", err)
	}

	if testCity.Name != "testCityName" {
		t.Fatalf("get city name: got: %s, expected: %s", testCity.Name, "testCityName")
	}
}

func TestStorage_GetAllCities(t *testing.T) {
	s := tHelperInitEmptyDB(t)

	testInputCities := []model.City{
		{Name: "testCityName1"},
		{Name: "testCityName2"},
		{Name: "testCityName3"},
	}
	for i := range testInputCities {
		row := s.db.QueryRow(fmt.Sprintf("INSERT INTO table_city (name) VALUES('%s') RETURNING id", testInputCities[i].Name))
		if err := row.Scan(&testInputCities[i].ID); err != nil {
			t.Fatalf("scan new city id: %v", err)
		}
		if row.Err() != nil {
			t.Fatalf("check scan err: %v", row.Err())
		}
	}

	citiesFromDB, err := s.GetAllCities(context.Background())
	if err != nil {
		t.Fatalf("GetAllCities: %v", err)
	}
	if len(citiesFromDB) != len(testInputCities) {
		t.Fatalf("expected %d cities, got %d cities", len(citiesFromDB), len(testInputCities))
	}

	sort.Slice(citiesFromDB, func(i, j int) bool {
		return citiesFromDB[i].Name < citiesFromDB[j].Name
	})
	assert.Equal(t, testInputCities, citiesFromDB)
}
