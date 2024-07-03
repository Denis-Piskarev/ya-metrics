package db

import "context"

//go:generate mockgen -source=interfaces.go -destination=mocks/db.go -package=mocks

type DbInterface interface {
	Ping(context.Context) error
}
