package cache

import "errors"

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}

var ErrKeyNotFound = errors.New("key not found")
