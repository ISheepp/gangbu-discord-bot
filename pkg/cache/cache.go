package cache

import (
	"context"
	"time"
)

type Cache interface {
	GetString(ctx context.Context, key string) (string, error)
	SetString(ctx context.Context, key string, value string, expiration time.Duration) error
	DeleteString(ctx context.Context, key string) error
}
