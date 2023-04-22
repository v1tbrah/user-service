//go:build with_db

package storage

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/pet-pr-social-network/user-service/internal/config"
)

func TestStorage_Init(t *testing.T) {
	cfg := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(cfg.LogLvl)

	if err := cfg.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(cfg.LogLvl)

	s := initEmptyDB(t)

	// DROP TABLES TO CHECK THEIR EXISTENCE AFTER REINITIALIZATION
	if _, err := s.dbConn.Query(fmt.Sprintf("DROP TABLE %s CASCADE", cfg.StorageConfig.CityTableName)); err != nil {
		t.Fatalf("drop table city: %s", err)
	}

	if _, err := s.dbConn.Query(fmt.Sprintf("DROP TABLE %s CASCADE", cfg.StorageConfig.InterestTableName)); err != nil {
		t.Fatalf("drop table interest: %s", err)
	}

	if _, err := s.dbConn.Query(fmt.Sprintf("DROP TABLE %s CASCADE", cfg.StorageConfig.UserTableName)); err != nil {
		t.Fatalf("drop table users: %s", err)
	}

	if _, err := s.dbConn.Query(fmt.Sprintf("DROP TABLE %s CASCADE", cfg.StorageConfig.UserPerInterestTableName)); err != nil {
		t.Fatalf("drop table user_per_interest: %s", err)
	}
}

func initEmptyDB(t *testing.T) *Storage {
	cfg := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(cfg.LogLvl)

	if err := cfg.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(cfg.LogLvl)

	s, err := Init(cfg.StorageConfig)
	if err != nil {
		t.Fatalf("init storage: %v", err)
	}

	// DROP TABLES IF THEY ALREADY EXIST
	if _, err = s.dbConn.Query(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", s.cfg.CityTableName)); err != nil {
		t.Fatalf("drop table city: %s", err)
	}

	if _, err = s.dbConn.Query(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", s.cfg.InterestTableName)); err != nil {
		t.Fatalf("drop table interest: %s", err)
	}

	if _, err = s.dbConn.Query(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", s.cfg.UserTableName)); err != nil {
		t.Fatalf("drop table users: %s", err)
	}

	if _, err = s.dbConn.Query(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", s.cfg.UserPerInterestTableName)); err != nil {
		t.Fatalf("drop table user_per_interest: %s", err)
	}

	// REINIT
	if s, err = Init(cfg.StorageConfig); err != nil {
		t.Fatalf("init storage after drop tables: %v", err)
	}

	return s
}
