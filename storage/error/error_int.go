package error

import (
	_ "github.com/lib/pq"
)

type ErrorStorageI interface {
	SaveError(Error string)  error
}
