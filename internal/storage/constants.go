package storage

import "time"

const (
	maxOpenConns = 20
	maxIdleConns = 20
	maxIdleTime  = time.Second * 30
	maxLifeTime  = time.Minute * 2
)
