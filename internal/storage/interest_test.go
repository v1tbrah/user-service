//go:build with_db

package storage

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/pet-pr-social-network/user-service/internal/model"
)

func TestStorage_CreateInterest(t *testing.T) {
	s := initEmptyDB(t)

	testInterest := model.Interest{Name: "testInterestName"}
	idNewInterest, err := s.CreateInterest(context.Background(), testInterest)
	if err != nil {
		t.Fatalf("s.CreateInterest: %v", err)
	}

	row := s.dbConn.QueryRow(fmt.Sprintf("SELECT name FROM %s WHERE id=%d", s.cfg.InterestTableName, idNewInterest))
	if err = row.Scan(&testInterest.Name); err != nil {
		t.Fatalf("scan get new interest name: %v", err)
	}
	if row.Err() != nil {
		t.Fatalf("check scan get new interest name: %v", row.Err())
	}

	if testInterest.Name != "testInterestName" {
		t.Fatalf("new interest name: got: %s, expected: %s", testInterest.Name, "testInterestName")
	}
}

func TestStorage_GetInterest(t *testing.T) {
	s := initEmptyDB(t)

	testInterest := model.Interest{Name: "testInterestName"}
	row := s.dbConn.QueryRow(fmt.Sprintf("INSERT INTO %s (name) VALUES('%s') RETURNING id", s.cfg.InterestTableName, testInterest.Name))
	if err := row.Scan(&testInterest.ID); err != nil {
		t.Fatalf("scan new interest id: %v", err)
	}
	if row.Err() != nil {
		t.Fatalf("check scan new interest id: %v", row.Err())
	}

	testInterest, err := s.GetInterest(context.Background(), testInterest.ID)
	if err != nil {
		t.Fatalf("get interest: %v", err)
	}

	if testInterest.Name != "testInterestName" {
		t.Fatalf("get interest name: got: %s, expected: %s", testInterest.Name, "testInterestName")
	}
}

func TestStorage_GetAllInterests(t *testing.T) {
	s := initEmptyDB(t)

	testInputInterests := []model.Interest{
		{Name: "testInterestName1"},
		{Name: "testInterestName2"},
		{Name: "testInterestName3"},
	}
	for i := range testInputInterests {
		row := s.dbConn.QueryRow(fmt.Sprintf("INSERT INTO %s (name) VALUES('%s') RETURNING id", s.cfg.InterestTableName, testInputInterests[i].Name))
		if err := row.Scan(&testInputInterests[i].ID); err != nil {
			t.Fatalf("scan new interest id: %v", err)
		}
		if row.Err() != nil {
			t.Fatalf("check scan err: %v", row.Err())
		}
	}

	interestsFromDB, err := s.GetAllInterests(context.Background())
	if err != nil {
		t.Fatalf("GetAllCities: %v", err)
	}
	if len(interestsFromDB) != len(testInputInterests) {
		t.Fatalf("expected %d interests, got %d cities", len(interestsFromDB), len(testInputInterests))
	}

	sort.Slice(interestsFromDB, func(i, j int) bool {
		return interestsFromDB[i].Name < interestsFromDB[j].Name
	})
	assert.Equal(t, testInputInterests, interestsFromDB)
}