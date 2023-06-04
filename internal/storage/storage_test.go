//go:build with_db

package storage

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gitlab.com/pet-pr-social-network/user-service/config"
)

func TestStorage_Init(t *testing.T) {
	cfg := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(cfg.LogLvl)

	if err := cfg.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(cfg.LogLvl)

	s, err := Init(cfg.Storage)
	if err != nil {
		t.Fatalf("init storage: %v", err)
	}

	// CHECK EXISTENCE AFTER INITIALIZATION
	if _, err = s.db.Query("SELECT 1 FROM table_city"); err != nil {
		t.Fatalf("select from table city: %v", err)
	}

	if _, err = s.db.Query("SELECT 1 FROM table_interest"); err != nil {
		t.Fatalf("select from table interest: %v", err)
	}

	if _, err = s.db.Query("SELECT 1 FROM table_user"); err != nil {
		t.Fatalf("select from table user: %v", err)
	}

	if _, err = s.db.Query("SELECT 1 FROM table_user_per_interest"); err != nil {
		t.Fatalf("select from table user per interest: %v", err)
	}
}

func tHelperInitEmptyDB(t *testing.T) *Storage {
	cfg := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(cfg.LogLvl)

	if err := cfg.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(cfg.LogLvl)

	s, err := Init(cfg.Storage)
	if err != nil {
		t.Fatalf("init storage: %v", err)
	}

	// DELETE FROM TABLES FOR CLEAR TEST SPACE
	if _, err = s.db.Query("DELETE FROM table_user CASCADE"); err != nil {
		t.Fatalf("delete from table user: %v", err)
	}

	if _, err = s.db.Query("DELETE FROM table_city CASCADE"); err != nil {
		t.Fatalf("delete from table city: %v", err)
	}

	if _, err = s.db.Query("DELETE FROM table_interest CASCADE"); err != nil {
		t.Fatalf("delete from table interest: %v", err)
	}

	if _, err = s.db.Query("DELETE FROM table_user_per_interest CASCADE"); err != nil {
		t.Fatalf("delete from table user per interest: %v", err)
	}

	return s
}
