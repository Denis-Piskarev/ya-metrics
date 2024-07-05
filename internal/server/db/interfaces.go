// Package to operate with database

package db

import (
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/db.go -package=mocks

// DBInterface interface of database
type DBInterface interface {
	Ping(context.Context) error
	Writer
	Reader
}

// Writer interface for writing metrics
type Writer interface {
	WriteGauge(context.Context, string, float64) (float64, error)
	WriteCounter(context.Context, string, int64) (int64, error)
}

// Reader interface for reading metrics
type Reader interface {
	GetMetrics(context.Context) (string, error)
	GetGauge(context.Context, string) (float64, error)
	GetCounter(context.Context, string) (int64, error)
}
