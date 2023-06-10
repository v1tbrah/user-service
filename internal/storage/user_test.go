//go:build with_db

package storage

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/v1tbrah/user-service/internal/model"
)

func TestStorage_CreateUser(t *testing.T) {
	s := tHelperInitEmptyDB(t)

	tests := []struct {
		name           string
		userName       string
		userSurname    string
		getInterestsID func() []int64
		getCityIdPtr   func() *int64
		wantErr        bool
	}{
		{
			name:        "simple test with city and interests",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				result := make([]int64, 0, 3)
				tempInterest := model.Interest{}
				for i := 0; i < 3; i++ {
					tempInterest.Name = strconv.Itoa(i)
					id, err := s.CreateInterest(context.Background(), tempInterest)
					if err != nil {
						t.Fatalf("s.CreateInterest: %v", err)
					}
					result = append(result, id)
				}
				return result
			},
			getCityIdPtr: func() *int64 {
				testCity := model.City{Name: "testCityName"}
				id, err := s.CreateCity(context.Background(), testCity)
				if err != nil {
					t.Fatalf("s.CreateCity: %v", err)
				}
				return &id
			},
		},
		{
			name:        "test without interests",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				return nil
			},
			getCityIdPtr: func() *int64 {
				testCity := model.City{Name: "testCityName"}
				id, err := s.CreateCity(context.Background(), testCity)
				if err != nil {
					t.Fatalf("s.CreateCity: %v", err)
				}
				return &id
			},
		},
		{
			name:        "test without city",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				result := make([]int64, 0, 3)
				tempInterest := model.Interest{}
				for i := 0; i < 3; i++ {
					tempInterest.Name = strconv.Itoa(i)
					id, err := s.CreateInterest(context.Background(), tempInterest)
					if err != nil {
						t.Fatalf("s.CreateInterest: %v", err)
					}
					result = append(result, id)
				}
				return result
			},
			getCityIdPtr: func() *int64 {
				return nil
			},
		},
		{
			name:        "test without city and interests",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				return nil
			},
			getCityIdPtr: func() *int64 {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = tHelperInitEmptyDB(t)

			user := model.User{
				Name:        tt.userName,
				Surname:     tt.userSurname,
				InterestsID: tt.getInterestsID(),
				CityID:      tt.getCityIdPtr(),
			}
			id, err := s.CreateUser(context.Background(), user)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			userFromDB := model.User{}
			row := s.db.QueryRow(fmt.Sprintf("SELECT name, surname, city_id FROM table_user WHERE id=%d", id))
			if err = row.Scan(&userFromDB.Name, &userFromDB.Surname, &userFromDB.CityID); err != nil {
				t.Fatalf("scan new user: %v", err)
			}
			if row.Err() != nil {
				t.Fatalf("check scan err: %v", err)
			}

			rows, err := s.db.Query(fmt.Sprintf("SELECT interest_id FROM table_user_per_interest WHERE user_id=%d", id))
			if err != nil {
				t.Fatalf("select interest_id list by user_id (%d)", id)
			}
			var tempInterestID int64
			for rows.Next() {
				if err = rows.Scan(&tempInterestID); err != nil {
					t.Fatalf("scan interest id: %v", err)
				}
				userFromDB.InterestsID = append(userFromDB.InterestsID, tempInterestID)
			}
			if rows.Err() != nil {
				t.Fatalf("check scan err: %v", rows.Err())
			}

			assert.Equal(t, user, userFromDB)
		})
	}
}

func TestStorage_GetUser(t *testing.T) {
	s := tHelperInitEmptyDB(t)

	tests := []struct {
		name           string
		userName       string
		userSurname    string
		getInterestsID func() []int64
		getCityIdPtr   func() *int64
		wantErr        bool
	}{
		{
			name:        "simple test with city and interests",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				result := make([]int64, 0, 3)
				tempInterest := model.Interest{}
				for i := 0; i < 3; i++ {
					tempInterest.Name = strconv.Itoa(i)
					id, err := s.CreateInterest(context.Background(), tempInterest)
					if err != nil {
						t.Fatalf("s.CreateInterest: %v", err)
					}
					result = append(result, id)
				}
				return result
			},
			getCityIdPtr: func() *int64 {
				testCity := model.City{Name: "testCityName"}
				id, err := s.CreateCity(context.Background(), testCity)
				if err != nil {
					t.Fatalf("s.CreateCity: %v", err)
				}
				return &id
			},
		},
		{
			name:        "test without interests",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				return nil
			},
			getCityIdPtr: func() *int64 {
				testCity := model.City{Name: "testCityName"}
				id, err := s.CreateCity(context.Background(), testCity)
				if err != nil {
					t.Fatalf("s.CreateCity: %v", err)
				}
				return &id
			},
		},
		{
			name:        "test without city",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				result := make([]int64, 0, 3)
				tempInterest := model.Interest{}
				for i := 0; i < 3; i++ {
					tempInterest.Name = strconv.Itoa(i)
					id, err := s.CreateInterest(context.Background(), tempInterest)
					if err != nil {
						t.Fatalf("s.CreateInterest: %v", err)
					}
					result = append(result, id)
				}
				return result
			},
			getCityIdPtr: func() *int64 {
				return nil
			},
		},
		{
			name:        "test without city and interests",
			userName:    "testUserName",
			userSurname: "testUserSurname",
			getInterestsID: func() []int64 {
				return nil
			},
			getCityIdPtr: func() *int64 {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = tHelperInitEmptyDB(t)

			user := model.User{
				Name:        tt.userName,
				Surname:     tt.userSurname,
				InterestsID: tt.getInterestsID(),
				CityID:      tt.getCityIdPtr(),
			}
			id, err := s.CreateUser(context.Background(), user)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			user.ID = id

			userFromDB, err := s.GetUser(context.Background(), id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, user, userFromDB)
		})
	}
}
