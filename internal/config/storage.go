package config

import "os"

const (
	defaultStorageHost              = "127.0.0.1"
	defaultStoragePort              = "5432"
	defaultStorageUser              = "postgres"
	defaultStoragePassword          = "postgres"
	defaultUserDBName               = "postgres"
	defaultUserTableName            = "users"
	defaultCityTableName            = "city"
	defaultInterestTableName        = "interest"
	defaultUserPerInterestTableName = "user_per_interest"
)
const (
	envNameStorageHost              = "STORAGE_HOST"
	envNameStoragePort              = "STORAGE_PORT"
	envNameStorageUser              = "STORAGE_USER"
	envNameStoragePassword          = "STORAGE_PASSWORD"
	envNameUserDBName               = "USER_DB_NAME"
	envNameUserTableName            = "USER_TABLE_NAME"
	envNameCityTableName            = "CITY_TABLE_NAME"
	envNameInterestTableName        = "INTEREST_TABLE_NAME"
	envNameUserPerInterestTableName = "USER_PER_INTEREST_TABLE_NAME"
)

type StorageConfig struct {
	Host                     string
	Port                     string
	User                     string
	Password                 string
	UserDBName               string
	UserTableName            string
	CityTableName            string
	InterestTableName        string
	UserPerInterestTableName string
}

func newDefaultStorageConfig() StorageConfig {
	return StorageConfig{
		Host:                     defaultStorageHost,
		Port:                     defaultStoragePort,
		User:                     defaultStorageUser,
		Password:                 defaultStoragePassword,
		UserDBName:               defaultUserDBName,
		UserTableName:            defaultUserTableName,
		CityTableName:            defaultCityTableName,
		InterestTableName:        defaultInterestTableName,
		UserPerInterestTableName: defaultUserPerInterestTableName,
	}
}

func (c *StorageConfig) parseEnv() {
	envHost := os.Getenv(envNameStorageHost)
	if envHost != "" {
		c.Host = envHost
	}

	envPort := os.Getenv(envNameStoragePort)
	if envPort != "" {
		c.Port = envPort
	}

	envUser := os.Getenv(envNameStorageUser)
	if envUser != "" {
		c.User = envUser
	}

	envPassword := os.Getenv(envNameStoragePassword)
	if envPassword != "" {
		c.Password = envPassword
	}

	envUserDBName := os.Getenv(envNameUserDBName)
	if envUserDBName != "" {
		c.UserDBName = envUserDBName
	}

	envUserTableName := os.Getenv(envNameUserTableName)
	if envUserTableName != "" {
		c.UserTableName = envUserTableName
	}

	envCityTableName := os.Getenv(envNameCityTableName)
	if envCityTableName != "" {
		c.CityTableName = envCityTableName
	}

	envInterestTableName := os.Getenv(envNameInterestTableName)
	if envInterestTableName != "" {
		c.InterestTableName = envInterestTableName
	}

	envUserPerInterestTableName := os.Getenv(envNameUserPerInterestTableName)
	if envUserPerInterestTableName != "" {
		c.UserPerInterestTableName = envUserPerInterestTableName
	}
}
